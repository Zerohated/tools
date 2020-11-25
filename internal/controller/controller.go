package controller

import (
	"time"

	conf "github.com/Zerohated/tools/config"
	"github.com/allegro/bigcache"
)

var (
	config = conf.Config
)

// Controller example
type Controller struct {
	Cache *bigcache.BigCache
}

func InitCache() (cache *bigcache.BigCache) {
	cache, _ = bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	return
}

// NewController example
func NewController() *Controller {
	return &Controller{
		Cache: InitCache(),
	}
}
