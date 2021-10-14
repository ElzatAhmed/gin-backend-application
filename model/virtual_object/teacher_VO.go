package virtual_object

import (
	"time"
)

type TeacherVo struct {
	TeacherId		string				`json:"teacher_id"`
	Name			string 				`json:"name"`
	Gender 			string				`json:"gender"`
	Age    			uint8				`json:"age"`
	EnrollTime		time.Time			`json:"enroll_time"`
}
