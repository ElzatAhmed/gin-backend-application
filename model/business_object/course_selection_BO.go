package business_object

import "github.com/gin-backend-application/model/entity"

type CourseSelectionBO struct {
	Id 			uint64		`json:"id"`
	CourseId	string		`json:"course_id"`
	StudentId	string		`json:"student_id"`
}

func (c *CourseSelectionBO) ToEntity() entity.CourseSelectionEntity {
	return entity.CourseSelectionEntity{
		Id:        c.Id,
		CourseId:  c.CourseId,
		StudentId: c.StudentId,
	}
}
