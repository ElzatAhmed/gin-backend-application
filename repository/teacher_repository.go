package repository

import (
	"github.com/gin-backend-application/model/business_object"
	"github.com/gin-backend-application/model/entity"
)

type TeacherRepository interface {
	findAll() []business_object.TeacherBo
	findByTeacherId(teacherId string) business_object.TeacherBo
	findAllByName(name string) []business_object.TeacherBo
	save(teacher business_object.TeacherBo) business_object.TeacherBo
}

type teacherRepository struct {
	teachers []entity.TeacherEntity
}

func NewTeacherRepo() TeacherRepository {
	return &teacherRepository{}
}

func (t teacherRepository) findAll() []business_object.TeacherBo {
	var all []business_object.TeacherBo
	for _, obj := range t.teachers{
		all = append(all, obj.ToBO())
	}
	return all
}

func (t teacherRepository) findByTeacherId(teacherId string) business_object.TeacherBo {
	var teacher business_object.TeacherBo
	for _, obj := range t.teachers{
		if obj.TeacherId == teacherId{
			teacher = obj.ToBO()
		}
	}
	return teacher
}

func (t teacherRepository) findAllByName(name string) []business_object.TeacherBo {
	var all []business_object.TeacherBo
	for _, obj := range t.teachers{
		if obj.Name == name{
			all = append(all, obj.ToBO())
		}
	}
	return all
}

func (t *teacherRepository) save(teacher business_object.TeacherBo) business_object.TeacherBo {
	t.teachers = append(t.teachers, teacher.ToEntity())
	return teacher
}

