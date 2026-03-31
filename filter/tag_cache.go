package filter

import (
	"reflect"
	"sync"
)

var structFieldCache = cache{
	fields: make(map[reflect.Type][]fieldMeta),
	lock:   &sync.RWMutex{},
}

var enableCache = true

type fieldMeta struct {
	index      int
	fieldName  string
	anonymous  bool
	hasJSONTag bool
	ignored    bool
	spec       parsedTagSpec
}

type cache struct {
	lock   *sync.RWMutex
	fields map[reflect.Type][]fieldMeta
}

func (c *cache) GetField(key reflect.Type) ([]fieldMeta, bool) {
	c.lock.RLock()
	v, ok := c.fields[key]
	c.lock.RUnlock()
	return v, ok
}

func (c *cache) SetField(key reflect.Type, metas []fieldMeta) {
	c.lock.Lock()
	c.fields[key] = metas
	c.lock.Unlock()
}

func getFieldMetas(typeOf reflect.Type) []fieldMeta {
	if !enableCache {
		return buildFieldMetas(typeOf)
	}
	if metas, ok := structFieldCache.GetField(typeOf); ok {
		return metas
	}
	metas := buildFieldMetas(typeOf)
	structFieldCache.SetField(typeOf, metas)
	return metas
}

func buildFieldMetas(typeOf reflect.Type) []fieldMeta {
	metas := make([]fieldMeta, 0, typeOf.NumField())
	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		if !field.IsExported() {
			continue
		}

		meta := fieldMeta{
			index:     i,
			fieldName: field.Name,
			anonymous: field.Anonymous,
		}

		jsonTag, ok := field.Tag.Lookup("json")
		if !ok {
			metas = append(metas, meta)
			continue
		}

		meta.hasJSONTag = true
		if jsonTag == "-" {
			meta.ignored = true
			metas = append(metas, meta)
			continue
		}

		meta.spec = parseTagSpec(jsonTag, field.Name)
		metas = append(metas, meta)
	}
	return metas
}

func (m fieldMeta) tagInfo(scene string, isSelect bool) tagInfo {
	if isSelect {
		if !m.hasJSONTag || m.ignored {
			return tagInfo{omit: true}
		}
		return tagInfo{tag: m.spec.selectTag(scene, m.fieldName)}
	}

	if !m.hasJSONTag {
		return tagInfo{tag: newOmitNotTag(scene, m.fieldName)}
	}
	if m.ignored {
		return tagInfo{omit: true}
	}
	return tagInfo{tag: m.spec.omitTag(scene, m.fieldName)}
}
