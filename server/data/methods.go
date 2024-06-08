package data

import "fmt"

// Print returns the details of a student
func (i Student) Print() string {
	i.CheckDeanList()
	return fmt.Sprintf("Student ID: %d\nName: %s\nLast Name: %s\nGPA: %.2f\nDean's List: %t", i.Id, i.FirstName, i.LastName, i.GPA, i.DeansList)
}

// FullName returns the full name of a student
func (i Student) FullName() string {
	return fmt.Sprintf("%s %s", i.FirstName, i.LastName)
}

// CheckDeanList checks if a student is on the Dean's list
func (i Student) CheckDeanList() bool {
	if i.GPA >= 3.5 {
		i.DeansList = true
	}
	return i.DeansList
}

// GetGPA returns a student's GPA
func (i Student) GetGPA() string {
	return fmt.Sprintf("%s: %.2f", i.FullName(), i.GPA)
}
