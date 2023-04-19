package filter

import "reflect"

func (t *fieldNodeTree) getSelectTag1(scene string, pkgInfo string, i int, typeOf reflect.Type) tagInfo {
	//selectTag := tagInfo{}
	selectTag := getFieldSelectTag(typeOf.Field(i), scene)

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
