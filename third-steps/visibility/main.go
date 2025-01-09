package visibility

// Student is a struct that is exported and all properties are exported
type Student struct {
	Name  string
	Email string
}

// Teacher is a struct that is exported and all properties are exported
type Teacher struct {
	Name  string
	Email string
}

// HiddenStudent is a struct that is exportaded but its properties are not
type HiddenStudent struct {
	email string
}
