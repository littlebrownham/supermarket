package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/littlebrownham/supermarket/shared"
	"github.com/littlebrownham/supermarket/util"
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
	validationErr := validate(createProduceReq)

	if validationErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(string(validationErr.Error())))
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

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(createProduceReq.ProduceCode))
}

func validate(req *CreateProduceRequest) error {
	validateCodeErr := util.ValidateProduceCode(req.ProduceCode)
	if validateCodeErr != nil {
		return validateCodeErr
	}
	validatePriceErr := util.ValidatePrice(req.Price)
	if validatePriceErr != nil {
		return validatePriceErr
	}
	validateNameErr := util.ValidateName(req.Name)
	if validateNameErr != nil {
		return validateNameErr
	}

	return nil
}
