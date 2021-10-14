package business_object

import (
	"github.com/gin-backend-application/model"
	"github.com/gin-backend-application/model/entity"
	"time"
)

type StudentBO struct {
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


func (bo *StudentBO) ToEntity() entity.StudentEntity{
	return entity.StudentEntity{
		Id:         bo.Id,
		StudentId:  bo.StudentId,
		Password:   bo.Password,
		Name:       bo.Name,
		Gender:     bo.Gender,
		Age:        bo.Age,
		Grade:      bo.Grade,
		EnrollTime: bo.EnrollTime,
		Graduated:  bo.Graduated,
	}
}