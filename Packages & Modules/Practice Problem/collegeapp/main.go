package main

import (
	"collegeapp/course"
	"collegeapp/student"
	"collegeapp/teacher"
)

func main()  {
	s := student.Student{
		Name: "Rahul",
		RollNo: 101,
	}
	t := teacher.Teacher{
		Name: "Dr.Sharma",
		Subject: "Go Programming",
	}
	c:=course.Course{
		Code: "CS201",
		Title: "Advanced Go",
	}

	s.Display()
	t.Display()
	c.Display()
}

// Student: Rahul (Roll: 101)Teacher: Dr.Sharma (Subject: Go Programming)Course: [CS201] Advanced Go

