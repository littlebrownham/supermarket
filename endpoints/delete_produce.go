package endpoints

import (
	"net/http"

	"github.com/littlebrownham/supermarket/util"
)

type dbProducerDeleter interface {
	Delete(key string, c chan error)
}

type DeleteProduce struct {
	db dbProducerDeleter
}

func NewDeleteProduce(db dbProducerDeleter) *DeleteProduce {
	return &DeleteProduce{
		db: db,
	}
}

const ProduceCode = "produce_code"

func (d *DeleteProduce) DeleteProduce(w http.ResponseWriter, req *http.Request) {
	produce_code := req.URL.Query().Get(ProduceCode)
	if produce_code == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	if validationErr := validateDeleteReq(produce_code); validationErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(string(validationErr.Error())))
		return
	}

	chanErr := make(chan error)
	go d.db.Delete(produce_code, chanErr)
	err := <-chanErr
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}
func validateDeleteReq(produceCode string) error {
	return util.ValidateProduceCode(produceCode)
}
