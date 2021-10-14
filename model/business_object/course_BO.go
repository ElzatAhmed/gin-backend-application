package business_object

import (
	"github.com/gin-backend-application/model/entity"
	"time"
)

type CourseBo struct {
	Id 			uint64		`json:"id"`
	CourseId	string		`json:"course_id"`
	Name 		string		`json:"name"`
	TeacherId	string		`json:"teacher_id"`
	Capacity 	uint8		`json:"capacity"`
	From		time.Time	`json:"from"`
	To			time.Time	`json:"to"`
}

func (c *CourseBo) ToEntity() entity.CourseEntity{
	return entity.CourseEntity{
		Id:        c.Id,
		CourseId:  c.CourseId,
		Name:      c.Name,
		TeacherId: c.TeacherId,
		Capacity:  c.Capacity,
		From:      c.From,
		To:        c.To,
	}
}
