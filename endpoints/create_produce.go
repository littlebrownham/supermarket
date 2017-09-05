package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/littlebrownham/supermarket/shared"
	"github.com/littlebrownham/supermarket/util/validator"
)

type CreateProduceRequest struct {
	Name        string  `json:"name"`
	ProduceCode string  `json:"produce_code"`
	Price       float64 `json:"price"`
}

type dbInserter interface {
	Insert(key string, value interface{}, c chan error)
}

type CreateProduce struct {
	db dbInserter
}

func NewCreateProduce(db dbInserter) *CreateProduce {
	return &CreateProduce{
		db: db,
	}
}

func (c *CreateProduce) CreateProduce(w http.ResponseWriter, req *http.Request) {
	createProduceReq := &CreateProduceRequest{}
	if err := json.NewDecoder(req.Body).Decode(createProduceReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	chanErr := make(chan error)
	produce := shared.Product{
		Name:  createProduceReq.Name,
		Price: createProduceReq.Price,
	}

	go c.db.Insert(createProduceReq.ProduceCode, produce, chanErr)
	err := <-chanErr
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func validate(req CreateProduceRequest) error {
	validateCode := 

}
