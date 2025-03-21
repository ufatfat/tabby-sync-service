package github

import (
	"math/rand"
	"time"
)

const toks = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var tokBytes = []byte(toks)

func genRandString(min, max int) string {
	if min > max {
		return ""
	}
	l := 0
	if min == max {
		l = min
	} else {
		l = rand.Intn(max-min) + min // 长度[10,25]
	}

	rst := ""
	for i := 0; i < l; i++ {
		rst += string(toks[rand.Intn(len(toks))])
	}
	return rst
}

func (c *cache) setState(state string) {
	c.Lock()
	defer c.Unlock()
	c.data[state] = &cacheData{
		expire: time.Now().Add(stateExpireInterval),
	}
}
func (c *cache) setData(state string, data *userInfo) {
	c.Lock()
	defer c.Unlock()
	c.data[state].userInfo = data
}
func (c *cache) stateAvailable(state string) bool {
	c.RLock()
	defer c.RUnlock()
	return c.data[state].expire.After(time.Now())
}
func (c *cache) clearState(state string) {
	c.Lock()
	defer c.Unlock()
	delete(c.data, state)
}

// randomly choose 10 states to check if they're expired
func (c *cache) cleaner() {
	c.Lock()
	defer c.Unlock()
	counter := 0
	now := time.Now()
	for k := range c.data {
		if counter >= 10 {
			break
		}
		if now.After(c.data[k].expire) {
			delete(c.data, k)
		}
		counter++
	}
	time.Sleep(cleanerInterval)
}
