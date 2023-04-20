package main

type Map struct {
	M   map[string]**string   `json:"m,select(test)"`
	T   map[string]**string   `json:"t,select(),omit(test)"`
	MP  *map[string]**string  `json:"mp,select(test)"`
	MPP **map[string]**string `json:"mpp,select(test)"`
}

func newTestMap() Map {

	str := "c++从研发到脱发"
	ptr := &str
	maps := make(map[string]**string)
	maps["test"] = &ptr
	mp := &maps
	mpp := &mp

	return Map{M: maps, T: maps, MP: mp, MPP: mpp}
}
