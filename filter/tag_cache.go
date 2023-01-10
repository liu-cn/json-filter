package filter

var tagCache cache

var enableCache = true

func init() {
	tagCache.c = make(map[string]tag)
}

type cache struct {
	c map[string]tag
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

func (c *cache) getTag(tagStr string, scene string, pkgInfo string, isSelect bool, fieldName string, omitNotTag *bool) (tagEl tag, find bool) {

	key := c.key(pkgInfo, scene, fieldName, isSelect)

	v, ok := c.c[key]

	if !ok {
		if tagStr != "" {
			var tagE tag
			if isSelect {
				tagE = newSelectTag(tagStr, scene, fieldName)
			} else {

				if omitNotTag == nil {
					return tag{}, false
				}

				if *omitNotTag {
					tagE = newOmitNotTag(tagStr, fieldName)
				} else {
					tagE = newOmitTag(tagStr, scene, fieldName)
				}
			}
			c.c[key] = tagE
			return tagE, true
		} else {
			return tag{}, false
		}
	}
	return v, true
}
func (c *cache) getOmitTag(scene string, pkgInfo string, fieldName string) (tag, bool) {
	key := c.key(pkgInfo, scene, fieldName, false)

	v, ok := c.c[key]
	if !ok {
		return tag{}, false
	}
	return v, true
}
