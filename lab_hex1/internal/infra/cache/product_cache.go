package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type Product struct {
	ID    int
	Name  string
	Price float64
}

type ProductCache struct {
	cache       *cache.Cache
	refreshFunc func() ([]Product, error)
}

func NewProductCache(refreshFunc func() ([]Product, error), refreshMinutes int) *ProductCache {
	c := cache.New(time.Duration(refreshMinutes)*time.Minute, time.Duration(refreshMinutes)*time.Minute)
	pc := &ProductCache{cache: c, refreshFunc: refreshFunc}
	pc.refresh()
	go pc.autoRefresh(time.Duration(refreshMinutes) * time.Minute)
	return pc
}

func (p *ProductCache) refresh() {
	products, err := p.refreshFunc()
	if err == nil {
		p.cache.Flush()
		for _, prod := range products {
			p.cache.SetDefault(prod.Name, prod)
		}
	}
}

func (p *ProductCache) autoRefresh(interval time.Duration) {
	for {
		time.Sleep(interval)
		p.refresh()
	}
}

func (p *ProductCache) Get(name string) (Product, bool) {
	item, found := p.cache.Get(name)
	if !found {
		return Product{}, false
	}
	return item.(Product), true
}
