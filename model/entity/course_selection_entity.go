package entity

type CourseSelectionEntity struct {
	Id 			uint64		`json:"id"`
	CourseId	string		`json:"course_id"`
	StudentId	string		`json:"student_id"`
}
