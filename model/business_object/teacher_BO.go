package business_object

import (
	"github.com/gin-backend-application/model"
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
