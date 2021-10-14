package entity

import (
	"github.com/gin-backend-application/model/business_object"
	"time"
)

type CourseEntity struct {
	Id 			uint64		`json:"id"`
	CourseId	string		`json:"course_id"`
	Name 		string		`json:"name"`
	TeacherId	string		`json:"teacher_id"`
	Capacity 	uint8		`json:"capacity"`
	From		time.Time	`json:"from"`
	To			time.Time	`json:"to"`
}

func (c *CourseEntity) ToBO() business_object.CourseBo{
	return business_object.CourseBo{
		Id:        c.Id,
		CourseId:  c.CourseId,
		Name:      c.Name,
		TeacherId: c.TeacherId,
		Capacity:  c.Capacity,
		From:      c.From,
		To:        c.To,
	}
}
