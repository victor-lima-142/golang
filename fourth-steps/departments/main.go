package department

type Department struct {
	name string
}

func (d *Department) SetName(newName string) {
	d.name = newName
}

func (d *Department) GetName() string {
	return d.name
}
