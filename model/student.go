package model

// Use first character capital letter to make it "Exported Names" (access-modifier = public).
// Else = private to this package
type Student struct {
	id    int     // access-modifier: private
	Name  string  // public
	Score float64 // public
}

// in golang, func SetId() is equivalent with: func SetId(s *Student, id int)
// func (s *Student) SetId(id int) {
// 	s.id = id
// }

// func (s Student) GetId() int {
// 	return s.id
// }

var SingeltonStudent = Student{
	id:    99,
	Name:  "AfterOffice",
	Score: 99.9,
}
