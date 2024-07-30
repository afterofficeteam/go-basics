package model

// use first character capital letter to make access-modifier = public. else = private
type Student struct {
	id    int    // access-modifier: private
	Name  string // public
	Score int    // public
}

// in golang compiler, this func SetId() is equivalent with: func SetId(s *Student, id int) {}
// func (s *Student) SetId(id int) {
// 	s.id = id
// }

// func (s Student) GetId() int {
// 	return s.id
// }
