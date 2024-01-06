package singleton

var cache map[string]string

func SetupCache() {
	if cache == nil {
		cache = make(map[string]string)
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
