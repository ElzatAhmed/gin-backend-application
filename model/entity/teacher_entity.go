package entity

import (
	"github.com/gin-backend-application/model"
	"github.com/gin-backend-application/model/business_object"
	"time"
)

type TeacherEntity struct {
	Id 				uint64				`json:"id"`
	TeacherId		string				`json:"teacher_id"`
	Password		string				`json:"password"`
	Name			string 				`json:"name"`
	Gender 			model.GenderType	`json:"gender"`
	Age    			uint8				`json:"age"`
	EnrollTime		time.Time			`json:"enroll_time"`
}

func (entity *TeacherEntity) ToBO() business_object.TeacherBo {
	return business_object.TeacherBo{
		Id:         entity.Id,
		TeacherId:  entity.TeacherId,
		Password:   entity.Password,
		Name:       entity.Name,
		Gender:     entity.Gender,
		Age:        entity.Age,
		EnrollTime: entity.EnrollTime,
	}
}
