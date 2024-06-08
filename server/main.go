package main

import (
	"fmt"
	d "github.com/ebarthur/golang/server/data"
	"time"
)

func main() {
	Kwarteng := d.Student{122, "Andrews", "Kwarteng", 3.54, false}
	Aninakwa := d.Student{154, "Kwaku", "Aninakwa", 3.45, false}

	Arthur := d.NewStudent(123, "Arthur", "Churchill")

	// Course
	linAlgrebra := d.Course{214, "Linear Algebra", "MTH 201", "Room 101", "Mr. Mensah"}

	// Lecturer
	Mensah := d.Lecturer{214, "Dr. Mensah", "Mathematics", 5}

	// Workshops
	Webdev := d.Workshop{Course: linAlgrebra, Date: time.Now()}
	AI := d.Workshop{Course: d.Course{Id: 354, Name: "Artificial Intelligence", Code: "CSC 301", Venue: "Room 312", Lecturer: "Mr. Agyare"}, Date: time.Date(2024, time.June, 24, 0, 0, 0, 0, time.UTC).Local()}

	// Embedding - Workshops; all fields and methods of Course are promoted to WorkshopEmbed
	DevOps := d.WorkshopEmbed{}

	// all fields of Course are promoted to WorkshopEmbed but cannot be accessed directly in the constructor
	DevOps.Id = 456
	DevOps.Name = "DevOps"
	DevOps.Code = "CSC 401"
	DevOps.Venue = "Room 421"
	DevOps.Lecturer = "Dr Annin"
	DevOps.Date = time.Date(2024, time.August, 6, 17, 15, 0, 0, time.UTC).Local()

	InfoSec := d.WorkshopEmbed{Course: d.Course{Id: 567, Name: "Information Security", Code: "CSC 501", Venue: "Room 501", Lecturer: "Mr. Osei"}, Date: time.Date(2024, time.June, 27, 0, 0, 0, 0, time.UTC).Local()}

	fmt.Println(Aninakwa.FullName())
	fmt.Println(Aninakwa.Print())
	fmt.Println(linAlgrebra)
	fmt.Println(Mensah)
	fmt.Println(Webdev)
	fmt.Println(AI)
	fmt.Println(DevOps)
	fmt.Println(InfoSec)

	fmt.Println(Arthur.FullName())

	// check if a student is on the Dean's list
	fmt.Println("Dean's List: ", Kwarteng.CheckDeanList()) // true

	// check a student's GPA
	fmt.Println(Kwarteng.GetGPA()) // Andrews Kwarteng: 3.54
}
