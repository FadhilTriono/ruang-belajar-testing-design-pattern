package employee

// Manager is a struct for manager
// It embeds Employee
// This means that Manager has all the fields and methods of Employee
type Manager struct {
	Employee 
	// TODO: Add the Employee struct
}

// NewManager creates a new manager
// It returns a pointer to the manager
// Creational method
func NewManager() *Manager {
	manager := &Manager{
		Employee: Employee{
			Name:   "Manager",
			Salary: 1000,
			Bonus:  0, // Assuming managers don't have a bonus by default
		},
	}
	return manager
}