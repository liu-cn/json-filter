package filter

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
