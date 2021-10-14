package entity

import (
	"github.com/gin-backend-application/model"
	"github.com/gin-backend-application/model/business_object"
	"time"
)

type StudentEntity struct {
	Id			uint64				`json:"id"`
	StudentId	string				`json:"student_id"`
	Password	string				`json:"password"`
	Name		string 				`json:"name"`
	Gender 		model.GenderType	`json:"gender"`
	Age    		uint8				`json:"age"`
	Grade 		model.GradeType		`json:"grade"`
	EnrollTime	time.Time			`json:"enroll_time"`
	Graduated	bool				`json:"graduated"`
}

func (entity *StudentEntity) ToBO() business_object.StudentBO{
	return business_object.StudentBO{
		Id:         entity.Id,
		StudentId:  entity.StudentId,
		Password:   entity.Password,
		Name:       entity.Name,
		Gender:     entity.Gender,
		Age:        entity.Age,
		Grade:      entity.Grade,
		EnrollTime: entity.EnrollTime,
		Graduated:  entity.Graduated,
	}
}