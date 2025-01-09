package people

type Person struct {
	firstName string
	lastName  string
	birthday  string
}

func (p *Person) SetFirstName(newfirstName string) {
	p.firstName = newfirstName
}

func (p *Person) GetFirstName() string {
	return p.firstName
}

func (p *Person) SetLastName(newLastName string) {
	p.lastName = newLastName
}

func (p *Person) GetLastName() string {
	return p.lastName
}

func (p *Person) SetBirthday(newBirthday string) {
	p.birthday = newBirthday
}

func (p *Person) GetBirthday() string {
	return p.birthday
}
