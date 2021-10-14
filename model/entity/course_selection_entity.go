package entity

import "github.com/gin-backend-application/model/business_object"

type CourseSelectionEntity struct {
	Id 			uint64		`json:"id"`
	CourseId	string		`json:"course_id"`
	StudentId	string		`json:"student_id"`
}


func (c *CourseSelectionEntity) ToBO() business_object.CourseSelectionBO{
	return business_object.CourseSelectionBO{
		Id:        c.Id,
		CourseId:  c.CourseId,
		StudentId: c.StudentId,
	}
}