package marketdb

import (
	"fmt"

	"golang.org/x/sync/syncmap"
)

type MarketDB struct {
	// thread safe map; also could have wrapped Go's default map around a RW mutex
	db syncmap.Map
}

type Product struct {
	Name  string
	Price float32
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
func (m *MarketDB) GetAll() map[string]Product {
	seen := make(map[string]Product)
	m.db.Range(func(key, value interface{}) bool {
		k, ok := key.(string)
		if !ok {
			return false
		}
		v, ok := value.(Product)
		seen[k] = v
		return true
	})

	return seen
}
