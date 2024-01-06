package singleton

var cache map[string]string
var instanceCacheMap *map[string]string

func SetupCache() {
	if instanceCacheMap == nil {
		cache = make(map[string]string)
		instanceCacheMap = &cache
	}
}

func Add(key, value string) {
	SetupCache()
	cache[key] = value
}

func Get(key string) string {
	SetupCache()
	return cache[key]
}
