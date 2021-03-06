package marketdb

import (
	"errors"
	"strconv"
	"testing"

	"github.com/littlebrownham/supermarket/shared"
	"github.com/stretchr/testify/assert"
)

func TestNewMarketDB(t *testing.T) {
	db := NewMarketDB()
	assert.NotNil(t, db)
}

func TestMarketDBInsert(t *testing.T) {
	testCases := []struct {
		name        string
		inputKeys   []string
		inputValues []shared.Product
		inputChan   chan error

		expectedErr error
	}{
		{
			name:        "duplicate",
			inputKeys:   []string{"one"},
			inputValues: []shared.Product{shared.Product{Name: "first value", Price: float64(23.23)}},
			inputChan:   make(chan error),

			expectedErr: errors.New("Product one already exist"),
		},
	}

	for _, c := range testCases {
		db := NewMarketDB()

		for i := range c.inputKeys {
			go db.Insert(c.inputKeys[i], c.inputValues[i], c.inputChan)
			err := <-c.inputChan
			assert.NoError(t, err, c.name)
		}

		go db.Insert("one", "again", c.inputChan)
		expectedErr := <-c.inputChan
		assert.Equal(t, c.expectedErr.Error(), expectedErr.Error(), c.name)
	}
}

func TestMarketDBConcurrentDB(t *testing.T) {
	db := NewMarketDB()

	for i := 0; i < 10000; i++ {
		err := make(chan error)
		go db.Insert(strconv.Itoa(i), i, err)
		assert.NoError(t, <-err)
	}
}

func TestMarketDBGetAll(t *testing.T) {
	db := NewMarketDB()
	expectedEntries := 10

	for i := 0; i < expectedEntries; i++ {
		err := make(chan error)
		v := shared.Product{
			Name:  strconv.Itoa(i),
			Price: float64(i),
		}
		go db.Insert(strconv.Itoa(i), v, err)
		assert.NoError(t, <-err)
	}

	items := db.GetAll()
	assert.Equal(t, expectedEntries, len(items))
	for i := range items {
		assert.Contains(t, items, GetProduce{Name: strconv.Itoa(i), Price: float64(i), ProduceCode: strconv.Itoa(i)})
	}
}

func TestMarketDBDelete(t *testing.T) {
	db := NewMarketDB()
	for i := 0; i < 10; i++ {
		err := make(chan error)
		go db.Insert(strconv.Itoa(i), i, err)
		assert.NoError(t, <-err)
		go db.Delete(strconv.Itoa(i), err)
		assert.NoError(t, <-err)
	}

	items := db.GetAll()
	assert.Equal(t, 0, len(items))

	err := make(chan error)
	go db.Delete("non existent", err)
	assert.Error(t, <-err)
}
