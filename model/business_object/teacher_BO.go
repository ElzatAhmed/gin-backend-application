package business_object

import (
	"github.com/gin-backend-application/model"
	"github.com/gin-backend-application/model/entity"
	"time"
)

type TeacherBo struct {
	Id 				uint64				`json:"id"`
	TeacherId		string				`json:"teacher_id"`
	Password		string				`json:"password"`
	Name			string 				`json:"name"`
	Gender 			model.GenderType	`json:"gender"`
	Age    			uint8				`json:"age"`
	EnrollTime		time.Time			`json:"enroll_time"`
}

func (t *TeacherBo) ToEntity() entity.TeacherEntity{
	return entity.TeacherEntity{
		Id:         t.Id,
		TeacherId:  t.TeacherId,
		Password:   t.Password,
		Name:       t.Name,
		Gender:     t.Gender,
		Age:        t.Age,
		EnrollTime: t.EnrollTime,
	}
}
