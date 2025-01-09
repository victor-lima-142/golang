package main

import (
	"fmt"
	department "fourth-steps/departments"
	"fourth-steps/employees"
	person "fourth-steps/people"
	role "fourth-steps/roles"
)

func main() {
	var person person.Person
	var role role.Role
	var department department.Department

	println("Empty:")
	fmt.Println(person)
	fmt.Println(role)
	fmt.Println(department)

	println("\n\nMemory Address:")
	println(&person)
	println(&role)
	println(&department)

	person.SetFirstName("John")
	person.SetLastName("Doe")
	person.SetBirthday("1990-01-01")
	department.SetName("Development")
	role.SetName("Developer")
	role.SetDepartment(&department)

	var employee employees.Employee
	employee.Person = &person
	employee.Role = &role

	println("\n\nPopulated:")
	fmt.Println(person)
	fmt.Println(role)
	fmt.Println(department)
	departmentReturn := employee.Role.GetDepartment().GetName()
	firstNameReturn := employee.Person.GetFirstName()

	println("\n\nResults:")
	println(firstNameReturn, ":", departmentReturn)
}
