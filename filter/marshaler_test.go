package filter

import "testing"

type customJSONValue struct {
	Value string
}

func (c customJSONValue) MarshalJSON() ([]byte, error) {
	return []byte(`{"value":"` + c.Value + `"}`), nil
}

type customJSONContainer struct {
	Data customJSONValue `json:"data,select(api)"`
}

type collectionChild struct {
	Name string `json:"name,select(api)"`
}

func TestJSONMarshalerField(t *testing.T) {
	got := SelectFilter("api", customJSONContainer{
		Data: customJSONValue{Value: "ok"},
	}).MustJSON()

	equal, err := EqualJSON(got, `{"data":{"value":"ok"}}`)
	if err != nil {
		t.Fatal(err)
	}
	if !equal {
		t.Fatalf("unexpected json: %s", got)
	}
}

func TestNilPointerElementsPreserved(t *testing.T) {
	child := &collectionChild{Name: "ok"}

	gotSlice := SelectFilter("api", []*collectionChild{nil, child}).MustJSON()
	if gotSlice != `[null,{"name":"ok"}]` {
		t.Fatalf("unexpected slice json: %s", gotSlice)
	}

	gotMap := SelectFilter("api", map[string]*collectionChild{
		"nil":   nil,
		"child": child,
	}).MustJSON()

	equal, err := EqualJSON(gotMap, `{"child":{"name":"ok"},"nil":null}`)
	if err != nil {
		t.Fatal(err)
	}
	if !equal {
		t.Fatalf("unexpected map json: %s", gotMap)
	}
}

func TestTopLevelNilAndScalar(t *testing.T) {
	var child *collectionChild

	if got := SelectFilter("api", child).MustJSON(); got != "null" {
		t.Fatalf("unexpected nil json: %s", got)
	}
	if got := SelectFilter("api", 12).MustJSON(); got != "12" {
		t.Fatalf("unexpected scalar json: %s", got)
	}
}

func TestBoolMapKeySupported(t *testing.T) {
	got := SelectFilter("api", map[bool]string{
		true:  "yes",
		false: "no",
	}).MustJSON()

	equal, err := EqualJSON(got, `{"true":"yes","false":"no"}`)
	if err != nil {
		t.Fatal(err)
	}
	if !equal {
		t.Fatalf("unexpected json: %s", got)
	}
}

func TestFilterBytesHelpers(t *testing.T) {
	filtered := SelectFilter("api", customJSONContainer{
		Data: customJSONValue{Value: "ok"},
	})

	got, err := filtered.Bytes()
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != `{"data":{"value":"ok"}}` {
		t.Fatalf("unexpected bytes: %s", string(got))
	}
	if string(filtered.MustBytes()) != string(got) {
		t.Fatalf("must bytes mismatch: %s", string(filtered.MustBytes()))
	}
}
