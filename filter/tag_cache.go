package filter

import "reflect"

var tagCache cache

var enableCache = true

func init() {
	tagCache.fields = make(map[string]tag)
}

type cache struct {
	fields map[string]tag
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
func getFieldOmitTag(field reflect.StructField, scene string) tagInfo {
	tagInfoEl := tagInfo{}
	//没开缓存就获取tag
	jsonTag, ok := field.Tag.Lookup("json")
	var tag tag
	if !ok {
		tag = newOmitNotTag(scene, field.Name)
	} else {
		if jsonTag == "-" {
			tagInfoEl.omit = true
			return tagInfoEl
		}
		tag = newOmitTag(jsonTag, scene, field.Name)
	}
	tagInfoEl.tag = tag
	return tagInfoEl
}
func getFieldSelectTag(field reflect.StructField, scene string) tagInfo {
	tagInfoEl := tagInfo{}
	//没开缓存就获取tag
	jsonTag, ok := field.Tag.Lookup("json")
	var t tag
	if !ok {
		tagInfoEl.omit = true
		return tagInfoEl
	} else {
		if jsonTag == "-" {
			tagInfoEl.omit = true
			return tagInfoEl
		}
		t = newSelectTag(jsonTag, scene, field.Name)
	}
	tagInfoEl.tag = t
	return tagInfoEl
}
func getOmitTag(scene string, pkgInfo string, i int, typeOf reflect.Type) tagInfo {
	omitTag := tagInfo{}
	if !enableCache { //没开缓存就获取tag
		omitTag = getFieldOmitTag(typeOf.Field(i), scene)
		return omitTag
	}
	fieldName := typeOf.Field(i).Name
	cacheKey := tagCache.key(pkgInfo, scene, fieldName, false)
	tagEl, exist := tagCache.fields[cacheKey]
	if !exist { //如果缓存里没取到
		omitTag = getFieldOmitTag(typeOf.Field(i), scene)
		tagCache.fields[cacheKey] = omitTag.tag
		return omitTag
	}
	omitTag.tag = tagEl

	return omitTag
}
func getSelectTag(scene string, pkgInfo string, i int, typeOf reflect.Type) tagInfo {
	selectTag := tagInfo{}
	//return getFieldSelectTag(typeOf.Field(i), scene)
	if !enableCache {
		return getFieldSelectTag(typeOf.Field(i), scene)
	}

	fieldName := typeOf.Field(i).Name
	cacheKey := tagCache.key(pkgInfo, scene, fieldName, true)
	tagEl, exist := tagCache.fields[cacheKey]
	if !exist { //如果缓存里没取到
		selectTag = getFieldSelectTag(typeOf.Field(i), scene)
		tagCache.fields[cacheKey] = selectTag.tag
		return selectTag
	}
	selectTag.tag = tagEl
	return selectTag
}
