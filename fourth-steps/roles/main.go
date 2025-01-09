package role

import (
	department "fourth-steps/departments"
)

type Role struct {
	name       string
	department *department.Department
}

func (r *Role) SetName(newName string) {
	r.name = newName
}

func (r *Role) GetName() string {
	return r.name
}

func (r *Role) SetDepartment(newDepartment *department.Department) {
	r.department = newDepartment
}

func (r *Role) GetDepartment() *department.Department {
	return r.department
}
