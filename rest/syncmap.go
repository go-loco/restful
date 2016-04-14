package rest

import "sync"

// A synchronized map, using Read & Write Mutexes.
type syncMap struct {
	cache   map[string]interface{}
	rwMutex sync.RWMutex
}

func newSyncMap() *syncMap {
	return &syncMap{
		cache:   make(map[string]interface{}),
		rwMutex: sync.RWMutex{},
	}
}

// Get a value.
// It uses a Read Lock.
func (c *syncMap) get(key string) interface{} {

	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()

	return c.cache[key]
}

// Set a value for a Non Existent key.
// Return true if the key/value was set.
// It uses a Full Lock
func (c *syncMap) setNX(key string, value interface{}) bool {

	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()

	v := c.cache[key]
	if v == nil {
		c.cache[key] = value
		return true
	}

	return false
}
