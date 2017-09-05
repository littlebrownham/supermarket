package endpoints

import (
	"net/http"
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
