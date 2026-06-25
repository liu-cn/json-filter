package filter

import (
	"strings"
)

const (
	anySelect = "$any"
	//empty ="$empty"
)

type tag struct {
	SelectScene  string //执行的场景
	IsOmitField  bool   //该字段是否需要被忽略？
	IsSelect     bool   //是选中的情况,标识该字段是否需要被解析
	FieldName    string //字段名
	UseFieldName string //最终使用的字段名，没标签就用结构体字段名，tag里有标签就用标签名字
	IsAnonymous  bool   //IsAnonymous 标识该字段是否是匿名字段
	Omitempty    bool   //为空忽略
	Function     string //自定义处理函数
}

type parsedTagSpec struct {
	UseFieldName string
	IsAnonymous  bool
	Omitempty    bool
	Function     string
	SelectScenes map[string]struct{}
	OmitScenes   map[string]struct{}
}

func parseTagSpec(tagStr, fieldName string) parsedTagSpec {
	spec := parsedTagSpec{
		UseFieldName: fieldName,
	}

	parts := strings.Split(tagStr, ",")
	for i, raw := range parts {
		part := strings.TrimSpace(raw)
		if i == 0 {
			switch {
			case part == "":
				spec.IsAnonymous = true
				continue
			case isTagOption(part):
				// Leave the struct field name in place and parse as an option below.
			default:
				spec.UseFieldName = part
				continue
			}
		}

		switch {
		case part == "omitempty":
			spec.Omitempty = true
		case strings.HasPrefix(part, "select(") && strings.HasSuffix(part, ")"):
			spec.SelectScenes = addScenes(spec.SelectScenes, part[7:len(part)-1])
		case strings.HasPrefix(part, "omit(") && strings.HasSuffix(part, ")"):
			spec.OmitScenes = addScenes(spec.OmitScenes, part[5:len(part)-1])
		case strings.HasPrefix(part, "func(") && strings.HasSuffix(part, ")"):
			spec.Function = part[5 : len(part)-1]
		}
	}

	return spec
}

func isTagOption(part string) bool {
	return part == "omitempty" ||
		(strings.HasPrefix(part, "select(") && strings.HasSuffix(part, ")")) ||
		(strings.HasPrefix(part, "omit(") && strings.HasSuffix(part, ")")) ||
		(strings.HasPrefix(part, "func(") && strings.HasSuffix(part, ")"))
}

func addScenes(dst map[string]struct{}, scenes string) map[string]struct{} {
	if dst == nil {
		dst = make(map[string]struct{})
	}
	for _, scene := range strings.Split(scenes, "|") {
		scene = strings.TrimSpace(scene)
		if scene == "" {
			continue
		}
		dst[scene] = struct{}{}
	}
	return dst
}

func joinScenes(scenes ...string) string {
	return strings.Join(scenes, "|")
}

type sceneMatcher struct {
	raw    string
	single string
	scenes map[string]struct{}
}

func newSceneMatcher(scene string) sceneMatcher {
	matcher := sceneMatcher{
		raw: scene,
	}
	if !strings.Contains(scene, "|") {
		matcher.single = strings.TrimSpace(scene)
		return matcher
	}
	matcher.scenes = addScenes(nil, scene)
	return matcher
}

func newSceneMatcherFromScenes(scenes ...string) sceneMatcher {
	return newSceneMatcher(joinScenes(scenes...))
}

func hasScene(scenes map[string]struct{}, scene string) bool {
	if len(scenes) == 0 {
		return false
	}
	_, ok := scenes[scene]
	return ok
}

func (matcher sceneMatcher) matchAny(scenes map[string]struct{}) bool {
	if len(scenes) == 0 {
		return false
	}
	if hasScene(scenes, anySelect) {
		return true
	}
	if matcher.scenes == nil {
		return hasScene(scenes, matcher.single)
	}
	for requested := range matcher.scenes {
		if hasScene(scenes, requested) {
			return true
		}
	}
	return false
}

func (spec parsedTagSpec) selectTag(selectScene sceneMatcher, fieldName string) tag {
	tagEl := tag{
		FieldName:    fieldName,
		SelectScene:  selectScene.raw,
		IsOmitField:  true,
		UseFieldName: spec.UseFieldName,
		IsAnonymous:  spec.IsAnonymous,
		Omitempty:    spec.Omitempty,
		Function:     spec.Function,
	}

	if selectScene.matchAny(spec.SelectScenes) {
		tagEl.IsOmitField = false
		tagEl.IsSelect = true
	}
	return tagEl
}

func (spec parsedTagSpec) omitTag(omitScene sceneMatcher, fieldName string) tag {
	tagEl := tag{
		FieldName:    fieldName,
		SelectScene:  omitScene.raw,
		IsOmitField:  false,
		IsSelect:     true,
		UseFieldName: spec.UseFieldName,
		IsAnonymous:  spec.IsAnonymous,
		Omitempty:    spec.Omitempty,
		Function:     spec.Function,
	}

	if omitScene.matchAny(spec.OmitScenes) {
		tagEl.IsOmitField = true
		tagEl.IsSelect = false
	}
	return tagEl
}

func newSelectTag(tagStr, selectScene, fieldName string) tag {
	return parseTagSpec(tagStr, fieldName).selectTag(newSceneMatcher(selectScene), fieldName)
}

func newOmitTag(tagStr, omitScene, fieldName string) tag {
	return parseTagSpec(tagStr, fieldName).omitTag(newSceneMatcher(omitScene), fieldName)
}

func newOmitNotTag(omitScene sceneMatcher, fieldName string) tag {
	return tag{
		FieldName:    fieldName,
		IsSelect:     true,
		UseFieldName: fieldName,
		SelectScene:  omitScene.raw,
	}
}
