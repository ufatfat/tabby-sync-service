package cache

var passCache map[string]uint64

func init() {
	passCache = make(map[string]uint64)
}

func GetUser(pass string) uint64 {
	return passCache[pass]
}

func SetUser(pass string, user uint64) {
	passCache[pass] = user
}
