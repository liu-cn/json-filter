package filter

import "sync"

var tagCache = cache{
	fields: make(map[string]tag),
	lock:   &sync.RWMutex{},
}

var enableCache = true

//func init() {
//	tagCache.fields = make(map[string]tag)
//}

type cache struct {
	lock   *sync.RWMutex
	fields map[string]tag
}

func (c *cache) GetField(key string) (tag, bool) {
	c.lock.RLock()
	v, ok := c.fields[key]
	c.lock.RUnlock()
	return v, ok
}

func (c *cache) SetField(key string, tagEl tag) {
	c.lock.Lock()
	c.fields[key] = tagEl
	c.lock.Unlock()
}

func (c *cache) key(pkgInfo string, scene string, fieldName string, isSelect bool) string {
	s := ""
	if !isSelect {
		s = "omit." + scene + "."
	} else {
		s = "select." + scene + "."
	}
	return s + pkgInfo + "." + fieldName
}
