package repository

import (
	"github.com/gin-backend-application/model/business_object"
	"github.com/gin-backend-application/model/entity"
)

type StudentRepository interface {
	findAll() []business_object.StudentBO
	findByStudentId(studentId string) business_object.StudentBO
	findAllByName(name string) []business_object.StudentBO
	save(student business_object.StudentBO) business_object.StudentBO
}

type studentRepository struct {
	students []entity.StudentEntity
}

func NewStudentRepo() StudentRepository {
	return &studentRepository{}
}

func (s studentRepository) findAll() []business_object.StudentBO{
	var all []business_object.StudentBO
	for _, obj := range s.students {
		all = append(all, obj.ToBO())
	}
	return all
}

func (s studentRepository) findByStudentId(studentId string) business_object.StudentBO{
	var student business_object.StudentBO
	for _, obj := range s.students{
		if obj.StudentId == studentId {
			student = obj.ToBO()
		}
	}
	return student
}

func (s studentRepository) findAllByName(name string) []business_object.StudentBO{
	var all []business_object.StudentBO
	for _, obj := range s.students{
		if obj.Name == name{
			all = append(all, obj.ToBO())
		}
	}
	return all
}

func (s *studentRepository) save(student business_object.StudentBO) business_object.StudentBO{
	s.students = append(s.students, student.ToEntity())
	return student
}



