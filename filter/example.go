package filter

type User struct {
	Name     string
	Age      int
	Avatar   string
	Birthday int
	Password string
}

type Lang struct {
	Name string
	Arts []Art
}

type Art struct {
	Name    string
	Profile map[string]interface{}
	Values  []string
}
