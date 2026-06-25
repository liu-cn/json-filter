package filter

import (
	"encoding/json"
	"testing"
)

var benchmarkMultiSceneFilter Filter

type multiSceneProfile struct {
	Age     int    `json:"age,select(profile.age|admin)"`
	Email   string `json:"email,select(profile.email|admin)"`
	Address string `json:"address,select(profile.address),omit(profile.address)"`
}

type multiSceneOrder struct {
	ID     int    `json:"id,select(orders.id|admin)"`
	Amount int    `json:"amount,select(orders.amount|admin)"`
	Status string `json:"status,select(orders.status),omit(orders.status)"`
}

type multiSceneUser struct {
	ID           int               `json:"id,select(id|public)"`
	Name         string            `json:"name,select(name|public)"`
	Password     string            `json:"password,select(password),omit(password)"`
	Email        string            `json:"email,select(member)"`
	InternalNote string            `json:"internal_note,select(admin)"`
	Profile      multiSceneProfile `json:"profile,select(profile.age|profile.email|profile.address|admin)"`
	Orders       []multiSceneOrder `json:"orders,select(orders.id|orders.amount|orders.status|admin)"`
}

func newMultiSceneUser() multiSceneUser {
	return multiSceneUser{
		ID:           1,
		Name:         "Ada",
		Password:     "secret",
		Email:        "ada@example.com",
		InternalNote: "admin-only",
		Profile: multiSceneProfile{
			Age:     28,
			Email:   "profile@example.com",
			Address: "Hangzhou",
		},
		Orders: []multiSceneOrder{
			{ID: 101, Amount: 500, Status: "paid"},
			{ID: 102, Amount: 300, Status: "pending"},
		},
	}
}

func assertJSONEqual(t *testing.T, got, want string) {
	t.Helper()

	equal, err := EqualJSON(got, want)
	if err != nil {
		t.Fatalf("compare json failed: %v", err)
	}
	if !equal {
		t.Fatalf("unexpected json:\n got: %s\nwant: %s", got, want)
	}
}

func TestSelectFilterWithMultipleScenesString(t *testing.T) {
	got := SelectFilter("id|name|profile.age|orders.id|orders.amount", newMultiSceneUser()).MustJSON()

	assertJSONEqual(t, got, `{
		"id": 1,
		"name": "Ada",
		"profile": {
			"age": 28
		},
		"orders": [
			{"id": 101, "amount": 500},
			{"id": 102, "amount": 300}
		]
	}`)
}

func TestSelectScenesFilterWithPermissionLevels(t *testing.T) {
	got := SelectScenesFilter(newMultiSceneUser(), "public", "member").MustJSON()

	assertJSONEqual(t, got, `{
		"id": 1,
		"name": "Ada",
		"email": "ada@example.com"
	}`)
}

func TestSelectScenesAPIWithPermissionLevels(t *testing.T) {
	data, err := json.Marshal(SelectScenes(newMultiSceneUser(), "public", "member", "admin"))
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	assertJSONEqual(t, string(data), `{
		"id": 1,
		"name": "Ada",
		"email": "ada@example.com",
		"internal_note": "admin-only",
		"profile": {
			"age": 28,
			"email": "profile@example.com"
		},
		"orders": [
			{"id": 101, "amount": 500},
			{"id": 102, "amount": 300}
		]
	}`)
}

func TestSelectScenesAPIWithDynamicScenes(t *testing.T) {
	fields := []string{"id", "name", "profile.age", "orders.id"}

	data, err := json.Marshal(SelectScenes(newMultiSceneUser(), fields...))
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	assertJSONEqual(t, string(data), `{
		"id": 1,
		"name": "Ada",
		"profile": {
			"age": 28
		},
		"orders": [
			{"id": 101},
			{"id": 102}
		]
	}`)
}

func TestSelectScenesAPIWithSlice(t *testing.T) {
	users := []multiSceneUser{
		{ID: 1, Name: "Ada", Email: "ada@example.com"},
		{ID: 2, Name: "Grace", Email: "grace@example.com"},
	}

	data, err := json.Marshal(SelectScenes(users, "public", "member"))
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	assertJSONEqual(t, string(data), `[
		{
			"id": 1,
			"name": "Ada",
			"email": "ada@example.com"
		},
		{
			"id": 2,
			"name": "Grace",
			"email": "grace@example.com"
		}
	]`)
}

func TestOmitFilterWithMultipleScenesString(t *testing.T) {
	got := OmitFilter("password|profile.address|orders.status", newMultiSceneUser()).MustJSON()

	assertJSONEqual(t, got, `{
		"id": 1,
		"name": "Ada",
		"email": "ada@example.com",
		"internal_note": "admin-only",
		"profile": {
			"age": 28,
			"email": "profile@example.com"
		},
		"orders": [
			{"id": 101, "amount": 500},
			{"id": 102, "amount": 300}
		]
	}`)
}

func TestOmitScenesFilter(t *testing.T) {
	got := OmitScenesFilter(newMultiSceneUser(), "password", "profile.address", "orders.status").MustJSON()

	assertJSONEqual(t, got, `{
		"id": 1,
		"name": "Ada",
		"email": "ada@example.com",
		"internal_note": "admin-only",
		"profile": {
			"age": 28,
			"email": "profile@example.com"
		},
		"orders": [
			{"id": 101, "amount": 500},
			{"id": 102, "amount": 300}
		]
	}`)
}

func TestOmitScenesAPI(t *testing.T) {
	data, err := json.Marshal(OmitScenes(newMultiSceneUser(), "password", "profile.address", "orders.status"))
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	assertJSONEqual(t, string(data), `{
		"id": 1,
		"name": "Ada",
		"email": "ada@example.com",
		"internal_note": "admin-only",
		"profile": {
			"age": 28,
			"email": "profile@example.com"
		},
		"orders": [
			{"id": 101, "amount": 500},
			{"id": 102, "amount": 300}
		]
	}`)
}

func BenchmarkSelectFilterSingleScene(b *testing.B) {
	user := newMultiSceneUser()
	for i := 0; i < b.N; i++ {
		benchmarkMultiSceneFilter = SelectFilter("public", user)
	}
}

func BenchmarkSelectFilterMultipleScenesString(b *testing.B) {
	user := newMultiSceneUser()
	for i := 0; i < b.N; i++ {
		benchmarkMultiSceneFilter = SelectFilter("public|member|admin", user)
	}
}

func BenchmarkSelectScenesFilterMultipleScenes(b *testing.B) {
	user := newMultiSceneUser()
	for i := 0; i < b.N; i++ {
		benchmarkMultiSceneFilter = SelectScenesFilter(user, "public", "member", "admin")
	}
}

func BenchmarkOmitScenesFilterMultipleScenes(b *testing.B) {
	user := newMultiSceneUser()
	for i := 0; i < b.N; i++ {
		benchmarkMultiSceneFilter = OmitScenesFilter(user, "password", "profile.address", "orders.status")
	}
}
