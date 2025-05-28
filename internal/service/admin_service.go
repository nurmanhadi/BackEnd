package service

import (
	"liva/internal/entity"
	"liva/internal/model"
	"liva/pkg"
	"liva/pkg/exception"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// ==========================
// admin section
// ==========================
func (s *adminService) CreateAdmin(request model.AdminAddRequest) error {
	if err := s.validation.Struct(request); err != nil {
		s.log.WithError(err).Warn("failed validation")
		return err
	}
	countUsername, err := s.adminRepository.CountUsername(request.Username)
	if err != nil {
		s.log.WithError(err).Error("failed count username from database")
		return err
	}
	if countUsername > 0 {
		s.log.Warn("username already exists")
		return exception.NewError(409, "username already exists")
	}
	adminId := uuid.NewString()
	newPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.WithError(err).Error("failed hash password from bcrypt")
		return err
	}
	countIdentifier, err := s.userRepository.CountByIdentifier(request.Username)
	if err != nil {
		s.log.WithError(err).Error("failed count identifier from database")
		return err
	}
	if countIdentifier > 0 {
		s.log.Warn("identifier already exists")
		return exception.NewError(409, "identifier already exists")
	}
	countReference, err := s.userRepository.CountByReferenceId(adminId)
	if err != nil {
		s.log.WithError(err).Error("failed count reference_id from database")
		return err
	}
	if countReference > 0 {
		s.log.Warn("reference_id already exists")
		return exception.NewError(409, "reference_id already exists")
	}
	if err := s.userRepository.Save(&entity.User{
		Id:          uuid.NewString(),
		Identifier:  request.Username,
		Password:    string(newPassword),
		Role:        string(pkg.RoleAdmin),
		ReferenceId: adminId,
	}); err != nil {
		s.log.WithError(err).Error("failed save user to database")
		return err
	}
	admin := &entity.Admin{
		Id:       adminId,
		Username: request.Username,
	}
	err = s.adminRepository.Save(*admin)
	if err != nil {
		s.log.WithError(err).Error("failed save admin to database")
		return err
	}
	return nil
}
func (s *adminService) UpdateAdmin(adminId string, request model.AdminUpdateRequest) error {
	if err := s.validation.Struct(request); err != nil {
		s.log.WithError(err).Warn("failed validation")
		return err
	}
	countId, err := s.adminRepository.CountById(adminId)
	if err != nil {
		s.log.WithError(err).Error("failed count username from database")
		return err
	}
	if countId < 1 {
		s.log.Warn("admin not found")
		return exception.NewError(404, "admin not found")
	}

	err = s.adminRepository.Updates(adminId, &request)
	if err != nil {
		s.log.WithError(err).Error("failed update admin to database")
		return err
	}
	return nil
}
func (s *adminService) GetAdminById(adminId string) (*entity.Admin, error) {
	admin, err := s.adminRepository.FindById(adminId)
	if err != nil {
		s.log.WithError(err).Warn("admin not found")
		return nil, exception.NewError(404, "admin not found")
	}
	return admin, nil
}
func (s *adminService) GetAllAdmin() ([]entity.Admin, error) {
	admins, err := s.adminRepository.FindAll()
	if err != nil {
		s.log.WithError(err).Warn("admins not found")
		return nil, exception.NewError(404, "admins not found")
	}
	return admins, nil
}
func (s *adminService) DeleteAdmin(adminId string) error {
	countId, err := s.adminRepository.CountById(adminId)
	if err != nil {
		s.log.WithError(err).Error("failed count username from database")
		return err
	}
	if countId < 1 {
		s.log.Warn("admin not found")
		return exception.NewError(404, "admin not found")
	}
	err = s.adminRepository.Delete(adminId)
	if err != nil {
		s.log.WithError(err).Error("failed delete admin")
		return err
	}
	return nil
}

// ==========================
// teacher section
// ==========================
func (s *adminService) CreateTeacher(request model.TeacherAddRequest) error {
	if err := s.validation.Struct(request); err != nil {
		s.log.WithError(err).Warn("failed validation")
		return err
	}
	countNip, err := s.teacherRepository.CountByNip(request.Nip)
	if err != nil {
		s.log.WithError(err).Error("failed count nip to database")
		return err
	}
	if countNip > 0 {
		s.log.Warn("nip already exists")
		return exception.NewError(409, "nip already exists")
	}
	teacherId := uuid.NewString()
	newPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.WithError(err).Error("failed generate hash password")
		return err
	}
	countIdentifier, err := s.userRepository.CountByIdentifier(request.Nip)
	if err != nil {
		s.log.WithError(err).Error("failed count identifier from database")
		return err
	}
	if countIdentifier > 0 {
		s.log.Warn("identifier already exists")
		return exception.NewError(409, "identifier already exists")
	}
	countReference, err := s.userRepository.CountByReferenceId(teacherId)
	if err != nil {
		s.log.WithError(err).Error("failed count reference_id from database")
		return err
	}
	if countReference > 0 {
		s.log.Warn("reference_id already exists")
		return exception.NewError(409, "reference_id already exists")
	}
	if err := s.userRepository.Save(&entity.User{
		Id:          uuid.NewString(),
		Identifier:  request.Nip,
		Password:    string(newPassword),
		Role:        string(pkg.RoleTeacher),
		ReferenceId: teacherId,
	}); err != nil {
		s.log.WithError(err).Error("failed save user to database")
		return err
	}
	teacher := &entity.Teacher{
		Id:   teacherId,
		Nip:  request.Nip,
		Name: strings.ToUpper(request.Name),
	}
	if err := s.teacherRepository.Save(*teacher); err != nil {
		s.log.WithError(err).Error("failed save teacher to database")
		return err
	}
	return nil
}
func (s *adminService) GetTeacherById(teacherId string) (*model.TeacherResponse, error) {
	teacher, err := s.teacherRepository.FindById(teacherId)
	if err != nil {
		s.log.WithError(err).Warn("teacher not found")
		return nil, exception.NewError(404, "teacher not found")
	}
	response := &model.TeacherResponse{
		Id:   teacher.Id,
		Nip:  teacher.Nip,
		Name: teacher.Name,
	}
	return response, nil
}
func (s *adminService) GetAllTeacher() ([]entity.Teacher, error) {
	teachers, err := s.teacherRepository.FindAll()
	if err != nil {
		s.log.WithError(err).Warn("teacher not found")
		return nil, err
	}
	return teachers, nil
}
func (s *adminService) UpdateTeacher(teacherId string, request model.TeacherUpdateRequest) error {
	if err := s.validation.Struct(request); err != nil {
		s.log.WithError(err).Warn("failed validation")
		return err
	}
	countId, err := s.teacherRepository.CountById(teacherId)
	if err != nil {
		s.log.WithError(err).Error("failed count id to database")
		return err
	}
	if countId < 1 {
		s.log.Warn("teacher not found")
		return exception.NewError(404, "techer not found")
	}
	if err := s.teacherRepository.Updates(teacherId, &request); err != nil {
		s.log.WithError(err).Error("failed update teacher to database")
		return err
	}
	return nil
}

// ==========================
// student section
// ==========================
func (s *adminService) CreateStudent(request *model.StudentAddRequest) error {
	if err := s.validation.Struct(request); err != nil {
		s.log.WithError(err).Warn("failed validation")
		return err
	}
	countNis, err := s.studentRepository.CountByNis(request.Nis)
	if err != nil {
		s.log.WithError(err).Error("failed count nis from database")
		return err
	}
	if countNis > 0 {
		s.log.Warn("nis already exists")
		return exception.NewError(409, "nis already exists")
	}
	studentId := uuid.NewString()
	newName := strings.ToUpper(request.Name)
	newPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.WithError(err).Error("failed hash password from bcrypt")
		return err
	}
	countIdentifier, err := s.userRepository.CountByIdentifier(request.Nis)
	if err != nil {
		s.log.WithError(err).Error("failed count identifier from database")
		return err
	}
	if countIdentifier > 0 {
		s.log.Warn("identifier already exists")
		return exception.NewError(409, "identifier already exists")
	}
	countReference, err := s.userRepository.CountByReferenceId(studentId)
	if err != nil {
		s.log.WithError(err).Error("failed count reference_id from database")
		return err
	}
	if countReference > 0 {
		s.log.Warn("reference_id already exists")
		return exception.NewError(409, "reference_id already exists")
	}
	if err := s.userRepository.Save(&entity.User{
		Id:          uuid.NewString(),
		Identifier:  request.Nis,
		Password:    string(newPassword),
		Role:        string(pkg.RoleStudent),
		ReferenceId: studentId,
	}); err != nil {
		s.log.WithError(err).Error("failed save user to database")
		return err
	}
	student := &entity.Student{
		Id:      studentId,
		Nis:     request.Nis,
		ClassId: request.ClassId,
		Name:    newName,
		Status:  request.Status,
	}
	err = s.studentRepository.Save(*student)
	if err != nil {
		s.log.WithError(err).Error("failed save student to database")
		return err
	}
	return nil
}
func (s *adminService) GetStudentById(studentId string) (*model.StudentResponse, error) {
	student, err := s.studentRepository.FindById(studentId)
	if err != nil {
		s.log.WithError(err).Warn("student not found")
		return nil, exception.NewError(404, "student not found")
	}
	response := &model.StudentResponse{
		Id:      student.Id,
		Nis:     student.Nis,
		ClassId: student.ClassId,
		Name:    student.Name,
		Status:  student.Status,
	}
	return response, nil
}
func (s *adminService) GetAllStudent() ([]entity.Student, error) {
	students, err := s.studentRepository.FindAll()
	if err != nil {
		s.log.WithError(err).Warn("student not found")
		return nil, exception.NewError(404, "student not found")
	}
	return students, nil
}
func (s *adminService) UpdateStudent(studentId string, request *model.StudentUpdateRequest) error {
	if err := s.validation.Struct(request); err != nil {
		s.log.WithError(err).Warn("failed validation")
		return err
	}
	countId, err := s.studentRepository.CountById(studentId)
	if err != nil {
		s.log.WithError(err).Error("failed count id from database")
		return err
	}
	if countId < 1 {
		return exception.NewError(404, "student not found")
	}
	err = s.studentRepository.Updates(studentId, request)
	if err != nil {
		s.log.WithError(err).Error("failed update student to database")
		return err
	}
	return nil
}

// ==========================
// class section
// ==========================
func (s *adminService) CreateClass(request model.ClassAddRequest) error {
	if err := s.validation.Struct(request); err != nil {
		s.log.WithError(err).Warn("failed validation")
		return err
	}
	class := &entity.Class{
		Name: request.Name,
	}
	if err := s.classRepository.Save(*class); err != nil {
		s.log.WithError(err).Error("failed save to database")
		return err
	}
	return nil
}
func (s *adminService) UpdateClass(classId string, request model.ClassUpdateRequest) error {
	newClassId, err := strconv.ParseInt(classId, 10, 32)
	if err != nil {
		s.log.WithError(err).Error("failed parse to int")
		return err
	}
	if err := s.validation.Struct(request); err != nil {
		s.log.WithError(err).Warn("failed validation")
		return err
	}
	countClass, err := s.classRepository.CountById(int(newClassId))
	if err != nil {
		s.log.WithError(err).Error("failed count id to database")
		return err
	}
	if countClass < 1 {
		s.log.Warn("class not found")
		return exception.NewError(404, "class not found")
	}
	if err := s.classRepository.Updates(int(newClassId), request); err != nil {
		s.log.WithError(err).Error("failed update to database")
		return err
	}
	return nil
}
func (s *adminService) GetClassById(classId string) (*entity.Class, error) {
	newClassId, err := strconv.ParseInt(classId, 10, 32)
	if err != nil {
		s.log.WithError(err).Error("failed parse to int")
		return nil, err
	}
	class, err := s.classRepository.FindById(int(newClassId))
	if err != nil {
		s.log.WithError(err).Warn("class not found")
		return nil, exception.NewError(404, "class not found")
	}
	return class, nil
}
func (s *adminService) GetAllClass() ([]entity.Class, error) {
	classes, err := s.classRepository.FindAll()
	if err != nil {
		s.log.WithError(err).Warn("failed find all class to database")
		return nil, err
	}
	return classes, nil
}

// ==========================
// subject section
// ==========================
func (s *adminService) CreateSubject(request model.SubjectAddRequest) error {
	if err := s.validation.Struct(request); err != nil {
		s.log.WithError(err).Warn("failed validation")
		return err
	}
	countTeacher, err := s.teacherRepository.CountById(request.TeacherId)
	if err != nil {
		s.log.WithError(err).Error("failed count id to database")
		return err
	}
	if countTeacher < 1 {
		s.log.Warn("teacher not found")
		return exception.NewError(404, "techer not found")
	}
	subject := &entity.Subject{
		TeacherId: request.TeacherId,
		Name:      request.Name,
	}
	if err := s.subjectRepository.Save(*subject); err != nil {
		s.log.WithError(err).Error("failed save subject to database")
	}
	return nil
}
func (s *adminService) UpdateSubject(subjectId string, request model.SubjectUpdateRequest) error {
	newSubjectId, err := strconv.ParseInt(subjectId, 10, 32)
	if err != nil {
		s.log.WithError(err).Error("failed parse string to int")
		return err
	}
	if err := s.validation.Struct(request); err != nil {
		s.log.WithError(err).Warn("failed validation")
		return err
	}
	countSubject, err := s.subjectRepository.CountById(int(newSubjectId))
	if err != nil {
		s.log.WithError(err).Error("failed count subject to database")
		return err
	}
	if countSubject < 1 {
		s.log.Warn("subject not found")
		return exception.NewError(404, "subject not found")
	}
	if request.TeacherId != "" {
		countTeacher, err := s.teacherRepository.CountById(request.TeacherId)
		if err != nil {
			s.log.WithError(err).Error("failed count teacher to database")
			return err
		}
		if countTeacher < 1 {
			s.log.Warn("teacher not found")
			return exception.NewError(404, "techer not found")
		}
	}
	if err := s.subjectRepository.Updates(int(newSubjectId), request); err != nil {
		s.log.WithError(err).Error("failed update subject to database")
	}
	return nil
}
func (s *adminService) GetSubjectById(subjectId string) (*entity.Subject, error) {
	newSubjectId, err := strconv.ParseInt(subjectId, 10, 32)
	if err != nil {
		s.log.WithError(err).Error("failed parse string to int")
		return nil, err
	}
	subject, err := s.subjectRepository.FindById(int(newSubjectId))
	if err != nil {
		s.log.WithError(err).Warn("subject not found")
		return nil, exception.NewError(404, "subject not found")
	}
	return subject, nil
}
func (s *adminService) GetAllSubject() ([]entity.Subject, error) {

	subjects, err := s.subjectRepository.FindAll()
	if err != nil {
		s.log.WithError(err).Warn("subject not found")
		return nil, exception.NewError(404, "subject not found")
	}
	return subjects, nil
}

// ==========================
// class-subject section
// ==========================
func (s *adminService) CreateClassSubject(request model.ClassSubjectAddRequest) error {
	if err := s.validation.Struct(request); err != nil {
		s.log.WithError(err).Warn("failed validation")
		return err
	}
	countClass, err := s.classRepository.CountById(request.ClassId)
	if err != nil {
		s.log.WithError(err).Error("failed count id to database")
		return err
	}
	countSubject, err := s.subjectRepository.CountById(request.SubjectId)
	if err != nil {
		s.log.WithError(err).Error("failed count id to database")
		return err
	}
	if countClass < 1 {
		s.log.Warn("class not found")
		return exception.NewError(404, "class not found")
	}
	if countSubject < 1 {
		s.log.Warn("subject not found")
		return exception.NewError(404, "subject not found")
	}
	classSubject := &entity.ClassSubject{
		ClassId:   request.ClassId,
		SubjectId: request.SubjectId,
	}
	if err := s.classSubjectRepository.Save(*classSubject); err != nil {
		s.log.WithError(err).Error("failed save class-subject to database")
	}
	return nil
}
func (s *adminService) DeleteClassSubject(classSubjectId string) error {
	newClassSubjectId, err := strconv.ParseInt(classSubjectId, 10, 32)
	if err != nil {
		s.log.WithError(err).Error("failed parse string to int")
		return err
	}
	countId, err := s.classSubjectRepository.CountById(int(newClassSubjectId))
	if err != nil {
		s.log.WithError(err).Error("failed count id from database")
		return err
	}
	if countId < 1 {
		s.log.Warn("class-subject not found")
		return exception.NewError(404, "class-subject not found")
	}
	err = s.classSubjectRepository.Delete(int(newClassSubjectId))
	if err != nil {
		s.log.WithError(err).Error("failed delete class-subject")
		return err
	}
	return nil
}

// ==========================
// schedule section
// ==========================
func (s *adminService) CreateSchedule(request model.ScheduleAddRequest) error {
	if err := s.validation.Struct(request); err != nil {
		s.log.WithError(err).Warn("failed validation")
		return err
	}
	countSubject, err := s.subjectRepository.CountById(request.SubjectId)
	if err != nil {
		s.log.WithError(err).Error("failed count subject to database")
		return err
	}
	if countSubject < 1 {
		s.log.Warn("subject not found")
		return exception.NewError(404, "subject not found")
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	newTimetable, err := time.ParseInLocation("2006-01-02 15:04:05", request.Timetable, loc)
	if err != nil {
		s.log.WithError(err).Error("failed parse string to date")
		return exception.NewError(400, "invalid timetable format, use 'YYYY-MM-DD HH:MM:SS'")

	}
	schedule := &entity.Schedule{
		SubjectId: request.SubjectId,
		Timetable: newTimetable,
	}
	if err := s.scheduleRepository.Save(*schedule); err != nil {
		s.log.WithError(err).Error("failed save schedule to database")
		return err
	}
	return nil
}
func (s *adminService) UpdateSchedule(scheduleId string, request model.ScheduleUpdateRequest) error {
	newScheduleId, err := strconv.ParseInt(scheduleId, 10, 64)
	if err != nil {
		s.log.WithError(err).Error("failed parse string to int")
		return err
	}
	if err := s.validation.Struct(request); err != nil {
		s.log.WithError(err).Warn("failed validation")
		return err
	}
	schedule := new(entity.Schedule)
	countSchedule, err := s.scheduleRepository.CountById(newScheduleId)
	if err != nil {
		s.log.WithError(err).Error("failed count class to database")
		return err
	}
	if countSchedule < 1 {
		s.log.Warn("schedule not found")
		return exception.NewError(404, "schedule not found")
	}
	if request.SubjectId != 0 {
		countSubject, err := s.subjectRepository.CountById(request.SubjectId)
		if err != nil {
			s.log.WithError(err).Error("failed count subject to database")
			return err
		}
		if countSubject < 1 {
			s.log.Warn("subject not found")
			return exception.NewError(404, "subject not found")
		}
		schedule.SubjectId = request.SubjectId
	}
	if request.Timetable != "" {
		loc, _ := time.LoadLocation("Asia/Jakarta")
		newTimetable, err := time.ParseInLocation("2006-01-02 15:04:05", request.Timetable, loc)
		if err != nil {
			s.log.WithError(err).Error("failed parse string to date")
			return exception.NewError(400, "invalid timetable format, use 'YYYY-MM-DD HH:MM:SS'")
		}
		schedule.Timetable = newTimetable
	}
	if err := s.scheduleRepository.Updates(newScheduleId, *schedule); err != nil {
		s.log.WithError(err).Error("failed save schedule to database")
		return err
	}
	return nil
}
func (s *adminService) GetAllSchedule() ([]entity.Schedule, error) {
	schedules, err := s.scheduleRepository.FindAll()
	if err != nil {
		s.log.WithError(err).Error("failed find schedules to database")
		return nil, err
	}
	return schedules, nil
}
func (s *adminService) GetScheduleById(scheduleId string) (*entity.Schedule, error) {
	newScheduleId, err := strconv.ParseInt(scheduleId, 10, 64)
	if err != nil {
		s.log.WithError(err).Error("failed parse string to int")
		return nil, err
	}
	schedule, err := s.scheduleRepository.FindById(newScheduleId)
	if err != nil {
		s.log.WithError(err).Error("failed find schedule to database")
		return nil, err
	}
	return schedule, nil
}
