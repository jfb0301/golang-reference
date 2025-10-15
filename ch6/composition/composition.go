package main 

import "fmt"

type Employee struct {
	Name 	string 
	ID	 	string 
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.name, e.ID)
}

type Manager struct {
	Employee	
	Reports []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	// Business logic 
}

// Employee - Embedded field 

m := Manager{
	Employee : Employee{
		Name: "Robert Patison",
		ID: 	"1234"
	},
	Reports: []Employee
}

// Employee fields are "Promoted" to the containing struct 

fmt.Println(m.ID)
fmt.Println(m.Description())

