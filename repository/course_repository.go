package repository

import (
	"github.com/gin-backend-application/model/business_object"
	"github.com/gin-backend-application/model/entity"
)

type CourseRepository interface {
	findAll() []business_object.CourseBo
	findByCourseId(courseId string) business_object.CourseBo
	findAllByName(name string) []business_object.CourseBo
	save(course business_object.CourseBo) business_object.CourseBo
}

type courseRepository struct {
	courses []entity.CourseEntity
}

func NewCourseRepo() CourseRepository {
	return &courseRepository{}
}

func (c courseRepository) findAll() []business_object.CourseBo {
	var all []business_object.CourseBo
	for _, obj := range c.courses{
		all = append(all, obj.ToBO())
	}
	return all
}

func (c courseRepository) findByCourseId(courseId string) business_object.CourseBo {
	var course business_object.CourseBo
	for _, obj := range c.courses{
		if obj.CourseId == courseId{
			course = obj.ToBO()
		}
	}
	return course
}

func (c courseRepository) findAllByName(name string) []business_object.CourseBo {
	var all []business_object.CourseBo
	for _, obj := range c.courses{
		if obj.Name == name{
			all = append(all, obj.ToBO())
		}
	}
	return all
}

func (c *courseRepository) save(course business_object.CourseBo) business_object.CourseBo {
	c.courses = append(c.courses, course.ToEntity())
	return course
}

