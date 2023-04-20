package main

type Slice struct {
	Slices   []**string   `json:"slices,select(test)"`
	Test     []**string   `json:"test,select(),omit(test)"`
	SliceP   *[]**string  `json:"slice_p,select(test)"`
	SlicesPP **[]**string `json:"slices_pp,select(test)"`
}

func newSliceTest() Slice {
	s := "值"
	p := &s

	slice := make([]**string, 0, 5)
	slice = append(slice, &p)
	pp := &slice
	ppp := &pp

	test := Slice{
		Slices:   slice,
		SliceP:   pp,
		SlicesPP: ppp,
		Test:     slice,
	}
	return test
}
