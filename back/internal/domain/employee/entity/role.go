package employee

// EmployeeRole represents the role of an employee
type EmployeeRole string

const (
	Regular EmployeeRole = "regular"
	Admin   EmployeeRole = "admin"
	Owner   EmployeeRole = "owner"
)

// IsValid returns true if the role is valid
func (r EmployeeRole) IsValid() bool {
	return r == Regular || r == Admin || r == Owner
}
