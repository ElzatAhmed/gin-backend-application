package business_object

import (
	"github.com/gin-backend-application/model"
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
