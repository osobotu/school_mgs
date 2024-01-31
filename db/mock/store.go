// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/osobotu/school_mgs/db/sqlc (interfaces: Store)
//
// Generated by this command:
//
//	mockgen -package mockdb -destination db/mock/store.go github.com/osobotu/school_mgs/db/sqlc Store
//

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/osobotu/school_mgs/db/sqlc"
	gomock "go.uber.org/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateArm mocks base method.
func (m *MockStore) CreateArm(arg0 context.Context, arg1 string) (db.Arm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateArm", arg0, arg1)
	ret0, _ := ret[0].(db.Arm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateArm indicates an expected call of CreateArm.
func (mr *MockStoreMockRecorder) CreateArm(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateArm", reflect.TypeOf((*MockStore)(nil).CreateArm), arg0, arg1)
}

// CreateClass mocks base method.
func (m *MockStore) CreateClass(arg0 context.Context, arg1 string) (db.Class, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateClass", arg0, arg1)
	ret0, _ := ret[0].(db.Class)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateClass indicates an expected call of CreateClass.
func (mr *MockStoreMockRecorder) CreateClass(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateClass", reflect.TypeOf((*MockStore)(nil).CreateClass), arg0, arg1)
}

// CreateClassHasArms mocks base method.
func (m *MockStore) CreateClassHasArms(arg0 context.Context, arg1 db.CreateClassHasArmsParams) (db.ClassHasArm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateClassHasArms", arg0, arg1)
	ret0, _ := ret[0].(db.ClassHasArm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateClassHasArms indicates an expected call of CreateClassHasArms.
func (mr *MockStoreMockRecorder) CreateClassHasArms(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateClassHasArms", reflect.TypeOf((*MockStore)(nil).CreateClassHasArms), arg0, arg1)
}

// CreateDepartment mocks base method.
func (m *MockStore) CreateDepartment(arg0 context.Context, arg1 db.CreateDepartmentParams) (db.Department, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDepartment", arg0, arg1)
	ret0, _ := ret[0].(db.Department)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDepartment indicates an expected call of CreateDepartment.
func (mr *MockStoreMockRecorder) CreateDepartment(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDepartment", reflect.TypeOf((*MockStore)(nil).CreateDepartment), arg0, arg1)
}

// CreateDepartmentHasSubject mocks base method.
func (m *MockStore) CreateDepartmentHasSubject(arg0 context.Context, arg1 db.CreateDepartmentHasSubjectParams) (db.DepartmentHasSubject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDepartmentHasSubject", arg0, arg1)
	ret0, _ := ret[0].(db.DepartmentHasSubject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDepartmentHasSubject indicates an expected call of CreateDepartmentHasSubject.
func (mr *MockStoreMockRecorder) CreateDepartmentHasSubject(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDepartmentHasSubject", reflect.TypeOf((*MockStore)(nil).CreateDepartmentHasSubject), arg0, arg1)
}

// CreateFormMaster mocks base method.
func (m *MockStore) CreateFormMaster(arg0 context.Context, arg1 db.CreateFormMasterParams) (db.FormMaster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFormMaster", arg0, arg1)
	ret0, _ := ret[0].(db.FormMaster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFormMaster indicates an expected call of CreateFormMaster.
func (mr *MockStoreMockRecorder) CreateFormMaster(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFormMaster", reflect.TypeOf((*MockStore)(nil).CreateFormMaster), arg0, arg1)
}

// CreateScore mocks base method.
func (m *MockStore) CreateScore(arg0 context.Context, arg1 db.CreateScoreParams) (db.Score, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateScore", arg0, arg1)
	ret0, _ := ret[0].(db.Score)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateScore indicates an expected call of CreateScore.
func (mr *MockStoreMockRecorder) CreateScore(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateScore", reflect.TypeOf((*MockStore)(nil).CreateScore), arg0, arg1)
}

// CreateSession mocks base method.
func (m *MockStore) CreateSession(arg0 context.Context, arg1 db.CreateSessionParams) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockStoreMockRecorder) CreateSession(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockStore)(nil).CreateSession), arg0, arg1)
}

// CreateStudent mocks base method.
func (m *MockStore) CreateStudent(arg0 context.Context, arg1 db.CreateStudentParams) (db.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStudent", arg0, arg1)
	ret0, _ := ret[0].(db.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStudent indicates an expected call of CreateStudent.
func (mr *MockStoreMockRecorder) CreateStudent(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStudent", reflect.TypeOf((*MockStore)(nil).CreateStudent), arg0, arg1)
}

// CreateStudentOffersSubject mocks base method.
func (m *MockStore) CreateStudentOffersSubject(arg0 context.Context, arg1 db.CreateStudentOffersSubjectParams) (db.StudentOffersSubject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStudentOffersSubject", arg0, arg1)
	ret0, _ := ret[0].(db.StudentOffersSubject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStudentOffersSubject indicates an expected call of CreateStudentOffersSubject.
func (mr *MockStoreMockRecorder) CreateStudentOffersSubject(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStudentOffersSubject", reflect.TypeOf((*MockStore)(nil).CreateStudentOffersSubject), arg0, arg1)
}

// CreateSubject mocks base method.
func (m *MockStore) CreateSubject(arg0 context.Context, arg1 string) (db.Subject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubject", arg0, arg1)
	ret0, _ := ret[0].(db.Subject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSubject indicates an expected call of CreateSubject.
func (mr *MockStoreMockRecorder) CreateSubject(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubject", reflect.TypeOf((*MockStore)(nil).CreateSubject), arg0, arg1)
}

// CreateTeacher mocks base method.
func (m *MockStore) CreateTeacher(arg0 context.Context, arg1 db.CreateTeacherParams) (db.Teacher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTeacher", arg0, arg1)
	ret0, _ := ret[0].(db.Teacher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTeacher indicates an expected call of CreateTeacher.
func (mr *MockStoreMockRecorder) CreateTeacher(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTeacher", reflect.TypeOf((*MockStore)(nil).CreateTeacher), arg0, arg1)
}

// CreateTeacherTeachesClass mocks base method.
func (m *MockStore) CreateTeacherTeachesClass(arg0 context.Context, arg1 db.CreateTeacherTeachesClassParams) (db.TeacherTeachesClass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTeacherTeachesClass", arg0, arg1)
	ret0, _ := ret[0].(db.TeacherTeachesClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTeacherTeachesClass indicates an expected call of CreateTeacherTeachesClass.
func (mr *MockStoreMockRecorder) CreateTeacherTeachesClass(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTeacherTeachesClass", reflect.TypeOf((*MockStore)(nil).CreateTeacherTeachesClass), arg0, arg1)
}

// CreateTerm mocks base method.
func (m *MockStore) CreateTerm(arg0 context.Context, arg1 db.CreateTermParams) (db.Term, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTerm", arg0, arg1)
	ret0, _ := ret[0].(db.Term)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTerm indicates an expected call of CreateTerm.
func (mr *MockStoreMockRecorder) CreateTerm(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTerm", reflect.TypeOf((*MockStore)(nil).CreateTerm), arg0, arg1)
}

// CreateTermScore mocks base method.
func (m *MockStore) CreateTermScore(arg0 context.Context, arg1 db.CreateTermScoreParams) (db.TermScore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTermScore", arg0, arg1)
	ret0, _ := ret[0].(db.TermScore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTermScore indicates an expected call of CreateTermScore.
func (mr *MockStoreMockRecorder) CreateTermScore(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTermScore", reflect.TypeOf((*MockStore)(nil).CreateTermScore), arg0, arg1)
}

// DeleteArm mocks base method.
func (m *MockStore) DeleteArm(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteArm", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteArm indicates an expected call of DeleteArm.
func (mr *MockStoreMockRecorder) DeleteArm(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteArm", reflect.TypeOf((*MockStore)(nil).DeleteArm), arg0, arg1)
}

// DeleteClass mocks base method.
func (m *MockStore) DeleteClass(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteClass", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteClass indicates an expected call of DeleteClass.
func (mr *MockStoreMockRecorder) DeleteClass(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteClass", reflect.TypeOf((*MockStore)(nil).DeleteClass), arg0, arg1)
}

// DeleteClassHasArms mocks base method.
func (m *MockStore) DeleteClassHasArms(arg0 context.Context, arg1 db.DeleteClassHasArmsParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteClassHasArms", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteClassHasArms indicates an expected call of DeleteClassHasArms.
func (mr *MockStoreMockRecorder) DeleteClassHasArms(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteClassHasArms", reflect.TypeOf((*MockStore)(nil).DeleteClassHasArms), arg0, arg1)
}

// DeleteDepartment mocks base method.
func (m *MockStore) DeleteDepartment(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDepartment", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDepartment indicates an expected call of DeleteDepartment.
func (mr *MockStoreMockRecorder) DeleteDepartment(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDepartment", reflect.TypeOf((*MockStore)(nil).DeleteDepartment), arg0, arg1)
}

// DeleteDepartmentHasSubjects mocks base method.
func (m *MockStore) DeleteDepartmentHasSubjects(arg0 context.Context, arg1 db.DeleteDepartmentHasSubjectsParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDepartmentHasSubjects", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDepartmentHasSubjects indicates an expected call of DeleteDepartmentHasSubjects.
func (mr *MockStoreMockRecorder) DeleteDepartmentHasSubjects(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDepartmentHasSubjects", reflect.TypeOf((*MockStore)(nil).DeleteDepartmentHasSubjects), arg0, arg1)
}

// DeleteFormMasterByClassID mocks base method.
func (m *MockStore) DeleteFormMasterByClassID(arg0 context.Context, arg1 db.DeleteFormMasterByClassIDParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFormMasterByClassID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFormMasterByClassID indicates an expected call of DeleteFormMasterByClassID.
func (mr *MockStoreMockRecorder) DeleteFormMasterByClassID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFormMasterByClassID", reflect.TypeOf((*MockStore)(nil).DeleteFormMasterByClassID), arg0, arg1)
}

// DeleteScore mocks base method.
func (m *MockStore) DeleteScore(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteScore", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteScore indicates an expected call of DeleteScore.
func (mr *MockStoreMockRecorder) DeleteScore(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteScore", reflect.TypeOf((*MockStore)(nil).DeleteScore), arg0, arg1)
}

// DeleteSession mocks base method.
func (m *MockStore) DeleteSession(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockStoreMockRecorder) DeleteSession(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockStore)(nil).DeleteSession), arg0, arg1)
}

// DeleteStudent mocks base method.
func (m *MockStore) DeleteStudent(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStudent", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStudent indicates an expected call of DeleteStudent.
func (mr *MockStoreMockRecorder) DeleteStudent(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStudent", reflect.TypeOf((*MockStore)(nil).DeleteStudent), arg0, arg1)
}

// DeleteStudentOffersSubject mocks base method.
func (m *MockStore) DeleteStudentOffersSubject(arg0 context.Context, arg1 db.DeleteStudentOffersSubjectParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStudentOffersSubject", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStudentOffersSubject indicates an expected call of DeleteStudentOffersSubject.
func (mr *MockStoreMockRecorder) DeleteStudentOffersSubject(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStudentOffersSubject", reflect.TypeOf((*MockStore)(nil).DeleteStudentOffersSubject), arg0, arg1)
}

// DeleteSubject mocks base method.
func (m *MockStore) DeleteSubject(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSubject", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSubject indicates an expected call of DeleteSubject.
func (mr *MockStoreMockRecorder) DeleteSubject(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSubject", reflect.TypeOf((*MockStore)(nil).DeleteSubject), arg0, arg1)
}

// DeleteTeacher mocks base method.
func (m *MockStore) DeleteTeacher(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTeacher", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTeacher indicates an expected call of DeleteTeacher.
func (mr *MockStoreMockRecorder) DeleteTeacher(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTeacher", reflect.TypeOf((*MockStore)(nil).DeleteTeacher), arg0, arg1)
}

// DeleteTeacherTeachesClass mocks base method.
func (m *MockStore) DeleteTeacherTeachesClass(arg0 context.Context, arg1 db.DeleteTeacherTeachesClassParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTeacherTeachesClass", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTeacherTeachesClass indicates an expected call of DeleteTeacherTeachesClass.
func (mr *MockStoreMockRecorder) DeleteTeacherTeachesClass(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTeacherTeachesClass", reflect.TypeOf((*MockStore)(nil).DeleteTeacherTeachesClass), arg0, arg1)
}

// DeleteTerm mocks base method.
func (m *MockStore) DeleteTerm(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTerm", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTerm indicates an expected call of DeleteTerm.
func (mr *MockStoreMockRecorder) DeleteTerm(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTerm", reflect.TypeOf((*MockStore)(nil).DeleteTerm), arg0, arg1)
}

// DeleteTermScore mocks base method.
func (m *MockStore) DeleteTermScore(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTermScore", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTermScore indicates an expected call of DeleteTermScore.
func (mr *MockStoreMockRecorder) DeleteTermScore(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTermScore", reflect.TypeOf((*MockStore)(nil).DeleteTermScore), arg0, arg1)
}

// GetArmByID mocks base method.
func (m *MockStore) GetArmByID(arg0 context.Context, arg1 int32) (db.Arm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArmByID", arg0, arg1)
	ret0, _ := ret[0].(db.Arm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArmByID indicates an expected call of GetArmByID.
func (mr *MockStoreMockRecorder) GetArmByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArmByID", reflect.TypeOf((*MockStore)(nil).GetArmByID), arg0, arg1)
}

// GetClassByID mocks base method.
func (m *MockStore) GetClassByID(arg0 context.Context, arg1 int32) (db.Class, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClassByID", arg0, arg1)
	ret0, _ := ret[0].(db.Class)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClassByID indicates an expected call of GetClassByID.
func (mr *MockStoreMockRecorder) GetClassByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClassByID", reflect.TypeOf((*MockStore)(nil).GetClassByID), arg0, arg1)
}

// GetClassByName mocks base method.
func (m *MockStore) GetClassByName(arg0 context.Context, arg1 string) (db.Class, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClassByName", arg0, arg1)
	ret0, _ := ret[0].(db.Class)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClassByName indicates an expected call of GetClassByName.
func (mr *MockStoreMockRecorder) GetClassByName(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClassByName", reflect.TypeOf((*MockStore)(nil).GetClassByName), arg0, arg1)
}

// GetDepartmentByID mocks base method.
func (m *MockStore) GetDepartmentByID(arg0 context.Context, arg1 int32) (db.Department, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDepartmentByID", arg0, arg1)
	ret0, _ := ret[0].(db.Department)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDepartmentByID indicates an expected call of GetDepartmentByID.
func (mr *MockStoreMockRecorder) GetDepartmentByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDepartmentByID", reflect.TypeOf((*MockStore)(nil).GetDepartmentByID), arg0, arg1)
}

// GetFormMasterByID mocks base method.
func (m *MockStore) GetFormMasterByID(arg0 context.Context, arg1 int32) (db.FormMaster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFormMasterByID", arg0, arg1)
	ret0, _ := ret[0].(db.FormMaster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFormMasterByID indicates an expected call of GetFormMasterByID.
func (mr *MockStoreMockRecorder) GetFormMasterByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFormMasterByID", reflect.TypeOf((*MockStore)(nil).GetFormMasterByID), arg0, arg1)
}

// GetScoreByStudentID mocks base method.
func (m *MockStore) GetScoreByStudentID(arg0 context.Context, arg1 int32) (db.Score, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScoreByStudentID", arg0, arg1)
	ret0, _ := ret[0].(db.Score)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScoreByStudentID indicates an expected call of GetScoreByStudentID.
func (mr *MockStoreMockRecorder) GetScoreByStudentID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScoreByStudentID", reflect.TypeOf((*MockStore)(nil).GetScoreByStudentID), arg0, arg1)
}

// GetSessionByID mocks base method.
func (m *MockStore) GetSessionByID(arg0 context.Context, arg1 int32) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionByID", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionByID indicates an expected call of GetSessionByID.
func (mr *MockStoreMockRecorder) GetSessionByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionByID", reflect.TypeOf((*MockStore)(nil).GetSessionByID), arg0, arg1)
}

// GetStudentByID mocks base method.
func (m *MockStore) GetStudentByID(arg0 context.Context, arg1 int32) (db.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStudentByID", arg0, arg1)
	ret0, _ := ret[0].(db.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudentByID indicates an expected call of GetStudentByID.
func (mr *MockStoreMockRecorder) GetStudentByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudentByID", reflect.TypeOf((*MockStore)(nil).GetStudentByID), arg0, arg1)
}

// GetSubjectByID mocks base method.
func (m *MockStore) GetSubjectByID(arg0 context.Context, arg1 int32) (db.Subject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubjectByID", arg0, arg1)
	ret0, _ := ret[0].(db.Subject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubjectByID indicates an expected call of GetSubjectByID.
func (mr *MockStoreMockRecorder) GetSubjectByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubjectByID", reflect.TypeOf((*MockStore)(nil).GetSubjectByID), arg0, arg1)
}

// GetSubjectByName mocks base method.
func (m *MockStore) GetSubjectByName(arg0 context.Context, arg1 string) (db.Subject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubjectByName", arg0, arg1)
	ret0, _ := ret[0].(db.Subject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubjectByName indicates an expected call of GetSubjectByName.
func (mr *MockStoreMockRecorder) GetSubjectByName(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubjectByName", reflect.TypeOf((*MockStore)(nil).GetSubjectByName), arg0, arg1)
}

// GetTeacherByID mocks base method.
func (m *MockStore) GetTeacherByID(arg0 context.Context, arg1 int32) (db.Teacher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTeacherByID", arg0, arg1)
	ret0, _ := ret[0].(db.Teacher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTeacherByID indicates an expected call of GetTeacherByID.
func (mr *MockStoreMockRecorder) GetTeacherByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeacherByID", reflect.TypeOf((*MockStore)(nil).GetTeacherByID), arg0, arg1)
}

// GetTermByID mocks base method.
func (m *MockStore) GetTermByID(arg0 context.Context, arg1 int32) (db.Term, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTermByID", arg0, arg1)
	ret0, _ := ret[0].(db.Term)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTermByID indicates an expected call of GetTermByID.
func (mr *MockStoreMockRecorder) GetTermByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTermByID", reflect.TypeOf((*MockStore)(nil).GetTermByID), arg0, arg1)
}

// GetTermScoreByID mocks base method.
func (m *MockStore) GetTermScoreByID(arg0 context.Context, arg1 int32) (db.TermScore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTermScoreByID", arg0, arg1)
	ret0, _ := ret[0].(db.TermScore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTermScoreByID indicates an expected call of GetTermScoreByID.
func (mr *MockStoreMockRecorder) GetTermScoreByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTermScoreByID", reflect.TypeOf((*MockStore)(nil).GetTermScoreByID), arg0, arg1)
}

// ListAllDepartments mocks base method.
func (m *MockStore) ListAllDepartments(arg0 context.Context) ([]db.Department, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllDepartments", arg0)
	ret0, _ := ret[0].([]db.Department)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllDepartments indicates an expected call of ListAllDepartments.
func (mr *MockStoreMockRecorder) ListAllDepartments(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllDepartments", reflect.TypeOf((*MockStore)(nil).ListAllDepartments), arg0)
}

// ListArmsInClass mocks base method.
func (m *MockStore) ListArmsInClass(arg0 context.Context, arg1 int32) ([]db.ClassHasArm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListArmsInClass", arg0, arg1)
	ret0, _ := ret[0].([]db.ClassHasArm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListArmsInClass indicates an expected call of ListArmsInClass.
func (mr *MockStoreMockRecorder) ListArmsInClass(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListArmsInClass", reflect.TypeOf((*MockStore)(nil).ListArmsInClass), arg0, arg1)
}

// ListClasses mocks base method.
func (m *MockStore) ListClasses(arg0 context.Context, arg1 db.ListClassesParams) ([]db.Class, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListClasses", arg0, arg1)
	ret0, _ := ret[0].([]db.Class)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClasses indicates an expected call of ListClasses.
func (mr *MockStoreMockRecorder) ListClasses(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClasses", reflect.TypeOf((*MockStore)(nil).ListClasses), arg0, arg1)
}

// ListStudents mocks base method.
func (m *MockStore) ListStudents(arg0 context.Context, arg1 db.ListStudentsParams) ([]db.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStudents", arg0, arg1)
	ret0, _ := ret[0].([]db.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStudents indicates an expected call of ListStudents.
func (mr *MockStoreMockRecorder) ListStudents(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStudents", reflect.TypeOf((*MockStore)(nil).ListStudents), arg0, arg1)
}

// ListSubjects mocks base method.
func (m *MockStore) ListSubjects(arg0 context.Context, arg1 db.ListSubjectsParams) ([]db.Subject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSubjects", arg0, arg1)
	ret0, _ := ret[0].([]db.Subject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSubjects indicates an expected call of ListSubjects.
func (mr *MockStoreMockRecorder) ListSubjects(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSubjects", reflect.TypeOf((*MockStore)(nil).ListSubjects), arg0, arg1)
}

// ListSubjectsByDepartmentID mocks base method.
func (m *MockStore) ListSubjectsByDepartmentID(arg0 context.Context, arg1 int32) ([]db.DepartmentHasSubject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSubjectsByDepartmentID", arg0, arg1)
	ret0, _ := ret[0].([]db.DepartmentHasSubject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSubjectsByDepartmentID indicates an expected call of ListSubjectsByDepartmentID.
func (mr *MockStoreMockRecorder) ListSubjectsByDepartmentID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSubjectsByDepartmentID", reflect.TypeOf((*MockStore)(nil).ListSubjectsByDepartmentID), arg0, arg1)
}

// ListSubjectsOfferedByStudentID mocks base method.
func (m *MockStore) ListSubjectsOfferedByStudentID(arg0 context.Context, arg1 int32) ([]db.StudentOffersSubject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSubjectsOfferedByStudentID", arg0, arg1)
	ret0, _ := ret[0].([]db.StudentOffersSubject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSubjectsOfferedByStudentID indicates an expected call of ListSubjectsOfferedByStudentID.
func (mr *MockStoreMockRecorder) ListSubjectsOfferedByStudentID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSubjectsOfferedByStudentID", reflect.TypeOf((*MockStore)(nil).ListSubjectsOfferedByStudentID), arg0, arg1)
}

// ListTeacherTeachesClassByTeacherID mocks base method.
func (m *MockStore) ListTeacherTeachesClassByTeacherID(arg0 context.Context, arg1 int32) ([]db.TeacherTeachesClass, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTeacherTeachesClassByTeacherID", arg0, arg1)
	ret0, _ := ret[0].([]db.TeacherTeachesClass)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTeacherTeachesClassByTeacherID indicates an expected call of ListTeacherTeachesClassByTeacherID.
func (mr *MockStoreMockRecorder) ListTeacherTeachesClassByTeacherID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTeacherTeachesClassByTeacherID", reflect.TypeOf((*MockStore)(nil).ListTeacherTeachesClassByTeacherID), arg0, arg1)
}

// ListTeachers mocks base method.
func (m *MockStore) ListTeachers(arg0 context.Context, arg1 db.ListTeachersParams) ([]db.Teacher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTeachers", arg0, arg1)
	ret0, _ := ret[0].([]db.Teacher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTeachers indicates an expected call of ListTeachers.
func (mr *MockStoreMockRecorder) ListTeachers(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTeachers", reflect.TypeOf((*MockStore)(nil).ListTeachers), arg0, arg1)
}

// ListTermScoresForSubjectAndClass mocks base method.
func (m *MockStore) ListTermScoresForSubjectAndClass(arg0 context.Context, arg1 db.ListTermScoresForSubjectAndClassParams) ([]db.TermScore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTermScoresForSubjectAndClass", arg0, arg1)
	ret0, _ := ret[0].([]db.TermScore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTermScoresForSubjectAndClass indicates an expected call of ListTermScoresForSubjectAndClass.
func (mr *MockStoreMockRecorder) ListTermScoresForSubjectAndClass(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTermScoresForSubjectAndClass", reflect.TypeOf((*MockStore)(nil).ListTermScoresForSubjectAndClass), arg0, arg1)
}

// UpdateArm mocks base method.
func (m *MockStore) UpdateArm(arg0 context.Context, arg1 db.UpdateArmParams) (db.Arm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateArm", arg0, arg1)
	ret0, _ := ret[0].(db.Arm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateArm indicates an expected call of UpdateArm.
func (mr *MockStoreMockRecorder) UpdateArm(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateArm", reflect.TypeOf((*MockStore)(nil).UpdateArm), arg0, arg1)
}

// UpdateFormMaster mocks base method.
func (m *MockStore) UpdateFormMaster(arg0 context.Context, arg1 db.UpdateFormMasterParams) (db.FormMaster, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFormMaster", arg0, arg1)
	ret0, _ := ret[0].(db.FormMaster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFormMaster indicates an expected call of UpdateFormMaster.
func (mr *MockStoreMockRecorder) UpdateFormMaster(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFormMaster", reflect.TypeOf((*MockStore)(nil).UpdateFormMaster), arg0, arg1)
}

// UpdateStudent mocks base method.
func (m *MockStore) UpdateStudent(arg0 context.Context, arg1 db.UpdateStudentParams) (db.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStudent", arg0, arg1)
	ret0, _ := ret[0].(db.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStudent indicates an expected call of UpdateStudent.
func (mr *MockStoreMockRecorder) UpdateStudent(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStudent", reflect.TypeOf((*MockStore)(nil).UpdateStudent), arg0, arg1)
}

// UpdateTeacher mocks base method.
func (m *MockStore) UpdateTeacher(arg0 context.Context, arg1 db.UpdateTeacherParams) (db.Teacher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTeacher", arg0, arg1)
	ret0, _ := ret[0].(db.Teacher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTeacher indicates an expected call of UpdateTeacher.
func (mr *MockStoreMockRecorder) UpdateTeacher(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTeacher", reflect.TypeOf((*MockStore)(nil).UpdateTeacher), arg0, arg1)
}

// UpdateTermScoreByID mocks base method.
func (m *MockStore) UpdateTermScoreByID(arg0 context.Context, arg1 db.UpdateTermScoreByIDParams) (db.TermScore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTermScoreByID", arg0, arg1)
	ret0, _ := ret[0].(db.TermScore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTermScoreByID indicates an expected call of UpdateTermScoreByID.
func (mr *MockStoreMockRecorder) UpdateTermScoreByID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTermScoreByID", reflect.TypeOf((*MockStore)(nil).UpdateTermScoreByID), arg0, arg1)
}
