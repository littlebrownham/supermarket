package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/littlebrownham/supermarket/adapter/marketdb"
)

type dbGetter interface {
	GetAll() []marketdb.GetProduce
}

type GetProduce struct {
	db dbGetter
}

func NewGetProduce(db dbGetter) *GetProduce {
	return &GetProduce{
		db: db,
	}
}

func (c *GetProduce) GetProduce(w http.ResponseWriter, req *http.Request) {
	produce := c.db.GetAll()
	jsonString, err := json.Marshal(produce)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write(jsonString)
}
