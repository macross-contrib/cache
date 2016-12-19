package cache

import (
	"github.com/insionng/macross"
)

const MacrossCacheStoreKey = "MacrossCacheStore"

func Store(value interface{}) Cache {
	var cacher Cache
	var okay bool
	switch v := value.(type) {
	case *macross.Context:
		if cacher, okay = v.Get(MacrossCacheStoreKey).(Cache); !okay {
			panic("Cacher not found, forget to Use Middleware ?")
		}
	default:

		panic("unknown Context")
	}

	if cacher == nil {
		panic("cache context not found")
	}

	return cacher
}

func Cacher(opt ...Options) macross.Handler {
	var option Options
	if len(opt) > 0 {
		option = opt[0]
	} else {
		option = Options{Adapter: "memory"}
	}
	return func(self *macross.Context) error {
		tagcache, err := New(option)
		if err != nil {
			return err
		}

		self.Set(MacrossCacheStoreKey, tagcache)

		if err = self.Next(); err != nil {
			return err
		}

		return nil
	}
}
