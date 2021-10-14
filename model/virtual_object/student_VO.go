package virtual_object

import (
	"time"
)

type StudentVo struct {
	StudentId	string				`json:"student_id"`
	Name		string 				`json:"name"`
	Gender 		string				`json:"gender"`
	Age    		uint8				`json:"age"`
	Grade 		string				`json:"grade"`
	EnrollTime	time.Time			`json:"enroll_time"`
	Graduated	bool				`json:"graduated"`
}
