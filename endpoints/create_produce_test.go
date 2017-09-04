package endpoints

import (
	// "errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeDB struct {
	calledKey   string
	calledValue interface{}

	outputErr error
}

func (f *fakeDB) Insert(key string, value interface{}, c chan error) {
	c <- f.outputErr
}

func TestCreateProduce(t *testing.T) {
	fakeDB := &fakeDB{}
	fakeCreateProduce := NewCreateProduce(fakeDB)
	assert.NotNil(t, fakeCreateProduce)
	payload := `{"name":"test product","price": 11.22, "produce_code": "abcd-1234"}`

	r, _ := http.NewRequest("POST", "/createProduce", strings.NewReader(payload))
	w := httptest.NewRecorder()

	fakeCreateProduce.CreateProduce(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}
