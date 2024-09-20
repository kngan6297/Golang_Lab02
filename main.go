package main

import (
	"fmt"
)

type Student struct {
	ID     int
	Name   string
	Gender string 
}

type StudentMap struct {
	students map[int]Student
}

func NewStudentMap() *StudentMap {
	return &StudentMap{students: make(map[int]Student)}
}

func (sm *StudentMap) AddStudent(student Student) {
	sm.students[student.ID] = student
}

func (sm *StudentMap) UpdateStudent(id int, student Student) {
	if _, exists := sm.students[id]; exists {
		sm.students[id] = student
	} else {
		fmt.Println("Sinh viên không tồn tại.")
	}
}

func (sm *StudentMap) DeleteStudent(id int) {
	delete(sm.students, id)
}

func (sm *StudentMap) GetStudent(id int) (Student, bool) {
	student, exists := sm.students[id]
	return student, exists
}

func (sm *StudentMap) ListStudentsByGender(gender string) {
	for _, student := range sm.students {
		if student.Gender == gender {
			fmt.Println(student)
		}
	}
}

func main() {
	sm := NewStudentMap()

	sm.AddStudent(Student{ID: 1, Name: "Nguyễn Văn A", Gender: "Nam"})
	sm.AddStudent(Student{ID: 2, Name: "Trần Thị B", Gender: "Nữ"})
	sm.AddStudent(Student{ID: 3, Name: "Lê Văn C", Gender: "Nam"})
	sm.AddStudent(Student{ID: 4, Name: "Phạm Thị D", Gender: "Nữ"})
	sm.AddStudent(Student{ID: 5, Name: "Nguyễn Văn E", Gender: "Nam"})
	sm.AddStudent(Student{ID: 6, Name: "Nguyễn Thị F", Gender: "Nữ"})
	sm.AddStudent(Student{ID: 7, Name: "Vũ Văn G", Gender: "Nam"})
	sm.AddStudent(Student{ID: 8, Name: "Trần Thị H", Gender: "Nữ"})
	sm.AddStudent(Student{ID: 9, Name: "Hoàng Văn I", Gender: "Nam"})
	sm.AddStudent(Student{ID: 10, Name: "Nguyễn Thị J", Gender: "Nữ"})

	if student, exists := sm.GetStudent(1); exists {
		fmt.Println("Tìm thấy sinh viên:", student)
	} else {
		fmt.Println("Sinh viên không tồn tại.")
	}

	sm.UpdateStudent(1, Student{ID: 1, Name: "Nguyễn Văn A", Gender: "Nam"})

	sm.DeleteStudent(2)

	fmt.Println("Danh sách sinh viên là Nam:")
	sm.ListStudentsByGender("Nam")

	fmt.Println("Danh sách sinh viên là Nữ:")
	sm.ListStudentsByGender("Nữ")
}
