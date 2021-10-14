package business_object

import "time"

type CourseBo struct {
	Id 			uint64		`json:"id"`
	CourseId	string		`json:"course_id"`
	Name 		string		`json:"name"`
	TeacherId	string		`json:"teacher_id"`
	Capacity 	uint8		`json:"capacity"`
	From		time.Time	`json:"from"`
	To			time.Time	`json:"to"`
}
