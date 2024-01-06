package singleton

var instance *Cache

type Cache struct {
	cache map[string]string
}

func NewCache() *Cache {
	if instance == nil {
		instance = &Cache{
			cache: make(map[string]string),
		}
	}
	return instance
}

func (c *Cache) Add(key, value string) {
	if c.cache == nil {
		c = NewCache()
	}
	c.cache[key] = value
}

func (c *Cache) Get(key string) string {
	if c.cache == nil {
		c = NewCache()
	}
	return c.cache[key]
}
