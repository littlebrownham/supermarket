package endpoints

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/littlebrownham/supermarket/adapter/marketdb"
	"github.com/stretchr/testify/assert"
)

type fakeDBGetter struct {
	expectedOutput []marketdb.GetProduce
}

func (f *fakeDBGetter) GetAll() []marketdb.GetProduce {
	return f.expectedOutput
}

func TestGetProduce(t *testing.T) {
	produce := []marketdb.GetProduce{marketdb.GetProduce{Name: "lettuce", Price: 12.12, ProduceCode: "1235-abcd"}}
	fakeDBGetter := &fakeDBGetter{}
	fakeDBGetter.expectedOutput = produce
	fakeGetProduce := NewGetProduce(fakeDBGetter)
	assert.NotNil(t, fakeGetProduce)

	r, _ := http.NewRequest("GET", "/getproduce", nil)
	w := httptest.NewRecorder()

	fakeGetProduce.GetProduce(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `[{"name":"lettuce","produce_code":"1235-abcd","price":12.12}]`, w.Body.String())
}
