package main

import (
	"fmt"

	"github.com/insionng/macross"
	"github.com/macross-contrib/cache"
	_ "github.com/macross-contrib/cache/redis"
)

func main() {

	v := macross.New()
	v.Use(cache.Cacher(cache.Options{Adapter: "redis", AdapterConfig: `{"Addr":"127.0.0.1:6379"}`, Section: "test", Interval: 5}))

	v.Get("/cache/put/", func(self *macross.Context) error {
		err := cache.Store(self).Set("name", "macross", 60)
		if err != nil {
			return err
		}

		return self.String("store okay")
	})

	v.Get("/cache/get/", func(self *macross.Context) error {
		var name string
		cache.Store(self).Get("name", &name)

		return self.String(fmt.Sprintf("get name %s", name))
	})

	v.Run(":7891")
}
