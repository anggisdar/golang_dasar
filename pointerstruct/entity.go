package pointerstruct

import (
	"fmt"
	"pertama/management"
)

// Print menerima pointer struct Employee sebagai paramater
func PrintEmployee(employee *management.Employee) {
	fmt.Printf("Hello, My Name: %s, Role: %s\n", employee.Name, employee.Role)
}

// Update mengubah nilai dari pointer struct
func UpdateEmployee(employee *management.Employee, newName, newRole string) {
	employee.Name = newName
	employee.Role = newRole
}
