package marketdb

import (
	"fmt"

	"github.com/littlebrownham/supermarket/shared"
	"golang.org/x/sync/syncmap"
)

type MarketDB struct {
	// thread safe map; also could have wrapped Go's default map around a RW mutex
	db syncmap.Map
}

type GetProduce struct {
	Name        string  `json:"name"`
	ProduceCode string  `json:"produce_code"`
	Price       float32 `json:"price"`
}

// NewMarketDB returns a marketDB to handle concurrent inserts/deletes/gets
func NewMarketDB() *MarketDB {
	syncMap := syncmap.Map{}
	return &MarketDB{
		db: syncMap,
	}
}

// Insert adds a entry to syncmap
// returns err if entry already exist
func (m *MarketDB) Insert(key string, value interface{}, c chan error) {
	if _, ok := m.db.Load(key); !ok {
		m.db.Store(key, value)
		c <- nil
	}
	c <- fmt.Errorf("Product %s already exist", key)
}

// GetAll returns copy of the sync map
func (m *MarketDB) GetAll() []GetProduce {
	seen := make([]GetProduce, 0)
	m.db.Range(func(key, value interface{}) bool {
		k, ok := key.(string)
		if !ok {
			fmt.Println("not valid")
			return false
		}
		v, ok := value.(shared.Product)
		if !ok {
			fmt.Println("not valid")
			return false
		}
		getProduce := GetProduce{
			ProduceCode: k,
			Price:       v.Price,
			Name:        v.Name,
		}
		seen = append(seen, getProduce)
		return true
	})
	return seen
}
