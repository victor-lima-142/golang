package employees

import (
	"fourth-steps/people"
	role "fourth-steps/roles"
)

type Employee struct {
	Person *people.Person
	Role   *role.Role
}
