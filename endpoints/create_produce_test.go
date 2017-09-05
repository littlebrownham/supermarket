package endpoints

import (
	// "errors"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeDBInserter struct {
	calledKey   string
	calledValue interface{}

	outputErr error
}

func (f *fakeDBInserter) Insert(key string, value interface{}, c chan error) {
	c <- f.outputErr
}

func TestCreateProduce(t *testing.T) {
	cases := []struct {
		name    string
		payload string
		err     error

		expectedStatusCode int
		expectedBody       string
	}{
		{
			name:               "success",
			payload:            `{"name":"test product","price": 11.22, "produce_code": "abcd-1234-abcd-1234"}`,
			expectedStatusCode: http.StatusCreated,
			expectedBody:       "abcd-1234-abcd-1234",
		},
		{
			name:               "bad request unmarshalling",
			payload:            `{"name":"test product","unit_price": "invalid", "prvalid": "abcd-1234-abcd-1234"}`,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "error writing to db duplicate entry",
			payload:            `{"name":"test product","price": 11.22, "produce_code": "abcd-1234-abcd-1234"}`,
			err:                errors.New("some error"),
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       "some error",
		},
	}

	for _, c := range cases {
		fakeDBInserter := &fakeDBInserter{}
		fakeDBInserter.outputErr = c.err
		fakeCreateProduce := NewCreateProduce(fakeDBInserter)
		assert.NotNil(t, fakeCreateProduce)

		r, _ := http.NewRequest("POST", "/createproduce", strings.NewReader(c.payload))
		w := httptest.NewRecorder()

		fakeCreateProduce.CreateProduce(w, r)
		assert.Equal(t, c.expectedStatusCode, w.Code, c.name)
		assert.Equal(t, c.expectedBody, w.Body.String(), c.name)
	}
}
