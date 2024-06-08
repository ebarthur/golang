package data

type Student struct {
	Id        int
	FirstName string
	LastName  string
	GPA       float64
	DeansList bool
}

// Factory

// NewStudent is a factory of the Student struct that returns Student name and id
func NewStudent(id int, firstName, lastName string) Student {
	return Student{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
	}
}
