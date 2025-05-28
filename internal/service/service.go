package service

import (
	"liva/internal/entity"
	"liva/internal/model"
	"liva/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type AdminService interface {
	CreateAdmin(request model.AdminAddRequest) error
	UpdateAdmin(adminId string, request model.AdminUpdateRequest) error
	GetAdminById(adminId string) (*entity.Admin, error)
	GetAllAdmin() ([]entity.Admin, error)
	DeleteAdmin(adminId string) error

	CreateTeacher(request model.TeacherAddRequest) error
	GetTeacherById(teacherId string) (*model.TeacherResponse, error)
	GetAllTeacher() ([]entity.Teacher, error)
	UpdateTeacher(teacherId string, request model.TeacherUpdateRequest) error

	CreateStudent(request *model.StudentAddRequest) error
	GetStudentById(studentId string) (*model.StudentResponse, error)
	GetAllStudent() ([]entity.Student, error)
	UpdateStudent(studentId string, request *model.StudentUpdateRequest) error

	CreateClass(request model.ClassAddRequest) error
	UpdateClass(classId string, request model.ClassUpdateRequest) error
	GetClassById(classId string) (*entity.Class, error)
	GetAllClass() ([]entity.Class, error)

	CreateSubject(request model.SubjectAddRequest) error
	UpdateSubject(subjectId string, request model.SubjectUpdateRequest) error
	GetSubjectById(subjectId string) (*entity.Subject, error)
	GetAllSubject() ([]entity.Subject, error)

	CreateClassSubject(request model.ClassSubjectAddRequest) error
	DeleteClassSubject(classSubjectId string) error

	CreateSchedule(request model.ScheduleAddRequest) error
	UpdateSchedule(scheduleId string, request model.ScheduleUpdateRequest) error
	GetAllSchedule() ([]entity.Schedule, error)
	GetScheduleById(scheduleId string) (*entity.Schedule, error)
}
type adminService struct {
	adminRepository        repository.AdminRepository
	userRepository         repository.UserRepository
	teacherRepository      repository.TeacherRepository
	studentRepository      repository.StudentRepository
	classRepository        repository.ClassRepository
	subjectRepository      repository.SubjectRepository
	classSubjectRepository repository.ClassSubjectRepository
	scheduleRepository     repository.ScheduleRepository

	validation *validator.Validate
	log        *logrus.Logger
}

func NewAdminService(
	adminRepository repository.AdminRepository,
	userRepository repository.UserRepository,
	teacherRepository repository.TeacherRepository,
	studentRepository repository.StudentRepository,
	classRepository repository.ClassRepository,
	subjectRepository repository.SubjectRepository,
	classSubjectRepository repository.ClassSubjectRepository,
	scheduleRepository repository.ScheduleRepository,

	validation *validator.Validate,
	log *logrus.Logger,
) AdminService {
	return &adminService{
		adminRepository:        adminRepository,
		userRepository:         userRepository,
		teacherRepository:      teacherRepository,
		studentRepository:      studentRepository,
		classRepository:        classRepository,
		subjectRepository:      subjectRepository,
		classSubjectRepository: classSubjectRepository,
		scheduleRepository:     scheduleRepository,

		validation: validation,
		log:        log,
	}
}
