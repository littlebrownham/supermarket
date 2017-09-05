package endpoints

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeDBDeleter struct {
	expectedKey string

	outputErr error
}

func (f *fakeDBDeleter) Delete(key string, c chan error) {
	c <- f.outputErr
}

func TestDeleteProduce(t *testing.T) {
	url := "/deleteproduce?produce_code=1234-abcd-1234-abcd"
	cases := []struct {
		name string
		err  error
		url  string

		expectedStatusCode int
		expectedBody       string
	}{
		{
			name:               "success",
			url:                url,
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "invalid query param",
			url:                "/deleteproduce?something=something",
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       "invalid produce code",
		},
		{
			name: "not found",
			url:  url,
			err:  errors.New("item exist"),

			expectedStatusCode: http.StatusNotFound,
			expectedBody:       "item exist",
		},
	}

	for _, c := range cases {
		fakeDBDeleter := &fakeDBDeleter{}
		fakeDBDeleter.outputErr = c.err
		fakeDelete := NewDeleteProduce(fakeDBDeleter)
		assert.NotNil(t, fakeDelete)

		r, _ := http.NewRequest("POST", c.url, nil)

		w := httptest.NewRecorder()
		fakeDelete.DeleteProduce(w, r)
		assert.Equal(t, c.expectedStatusCode, w.Code)
		assert.Equal(t, c.expectedBody, w.Body.String())
	}

}
