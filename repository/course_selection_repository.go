package repository

import (
	"github.com/gin-backend-application/model/business_object"
	"github.com/gin-backend-application/model/entity"
)

type CourseSelectionRepository interface {
	findAllByCourseId(courseId string) []business_object.CourseSelectionBO
	findAllByStudentId(studentId string) []business_object.CourseSelectionBO
	save(courseSelection business_object.CourseSelectionBO) business_object.CourseSelectionBO
}

type courseSelectionRepository struct {
	courseSelections []entity.CourseSelectionEntity
}

func (c courseSelectionRepository) findAllByCourseId(courseId string) []business_object.CourseSelectionBO {
	var all []business_object.CourseSelectionBO
	for _, obj := range c.courseSelections{
		if obj.CourseId == courseId{
			all = append(all, obj.ToBO())
		}
	}
	return all
}

func (c courseSelectionRepository) findAllByStudentId(studentId string) []business_object.CourseSelectionBO {
	var all []business_object.CourseSelectionBO
	for _, obj := range c.courseSelections{
		if obj.StudentId == studentId{
			all = append(all, obj.ToBO())
		}
	}
	return all
}

func (c *courseSelectionRepository) save(courseSelection business_object.CourseSelectionBO) business_object.CourseSelectionBO {
	c.courseSelections = append(c.courseSelections, courseSelection.ToEntity())
	return courseSelection
}

