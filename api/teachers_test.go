package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	mockdb "github.com/osobotu/school_mgs/db/mock"
	db "github.com/osobotu/school_mgs/db/sqlc"
	"github.com/osobotu/school_mgs/token"
	"github.com/osobotu/school_mgs/utils"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCreateTeacher(t *testing.T) {
	user, _ := randomUser()
	teacher := randomTeacher()
	fmt.Print(teacher)

	testCases := []struct {
		name          string
		body          gin.H
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(T *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"first_name":    teacher.FirstName,
				"last_name":     teacher.LastName,
				"middle_name":   teacher.MiddleName,
				"subject_id":    teacher.SubjectID,
				"department_id": teacher.DepartmentID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				params := db.CreateTeacherParams{

					FirstName:    teacher.FirstName,
					LastName:     teacher.LastName,
					MiddleName:   teacher.MiddleName,
					SubjectID:    teacher.SubjectID,
					DepartmentID: teacher.DepartmentID,
				}

				store.EXPECT().
					CreateTeacher(gomock.Any(), gomock.Eq(params)).
					Times(1).
					Return(teacher, nil)

			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				compareTeacherResponse(t, recorder.Body, teacher)
			},
		},
		{
			name: "Missing required fields",
			body: gin.H{
				"middle_name":   teacher.MiddleName,
				"subject_id":    teacher.SubjectID,
				"department_id": teacher.DepartmentID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateTeacher(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Internal Server Error",
			body: gin.H{
				"first_name":    teacher.FirstName,
				"last_name":     teacher.LastName,
				"middle_name":   teacher.MiddleName,
				"subject_id":    teacher.SubjectID,
				"department_id": teacher.DepartmentID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				params := db.CreateTeacherParams{

					FirstName:    teacher.FirstName,
					LastName:     teacher.LastName,
					MiddleName:   teacher.MiddleName,
					SubjectID:    teacher.SubjectID,
					DepartmentID: teacher.DepartmentID,
				}

				store.EXPECT().
					CreateTeacher(gomock.Any(), gomock.Eq(params)).
					Times(1).
					Return(db.Teacher{}, sql.ErrConnDone)

			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
				// todo: figure out how to unmarshal sql.NullTypes to JSON
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(
			tc.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()

				store := mockdb.NewMockStore(ctrl)

				tc.buildStub(store)

				server := newTestServer(t, store)
				recorder := httptest.NewRecorder()

				data, err := json.Marshal(tc.body)
				require.NoError(t, err)

				url := "/v1/teachers"
				request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
				require.NoError(t, err)

				tc.setupAuth(t, request, server.tokenMaker)

				server.router.ServeHTTP(recorder, request)

				tc.checkResponse(t, recorder)
			},
		)
	}
}

func TestGetTeacherByID(t *testing.T) {
	user, _ := randomUser()
	teacher := randomTeacher()

	testCases := []struct {
		name          string
		teacherID     int32
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(T *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			teacherID: teacher.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetTeacherByID(gomock.Any(), gomock.Eq(teacher.ID)).
					Times(1).
					Return(teacher, nil)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)

			},
		},
		{
			name:      "Not Found",
			teacherID: 1001,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetTeacherByID(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Teacher{}, sql.ErrNoRows)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)

			},
		},
		{
			name:      "Invalid ID",
			teacherID: -1,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetTeacherByID(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)

			},
		},
		{
			name:      "Internal Server Error",
			teacherID: teacher.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetTeacherByID(gomock.Any(), gomock.Eq(teacher.ID)).
					Times(1).
					Return(db.Teacher{}, sql.ErrConnDone)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)

			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(
			tc.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()

				store := mockdb.NewMockStore(ctrl)
				tc.buildStub(store)

				server := newTestServer(t, store)
				recorder := httptest.NewRecorder()

				url := fmt.Sprintf("/v1/teachers/%d", tc.teacherID)
				request, err := http.NewRequest(http.MethodGet, url, nil)
				require.NoError(t, err)

				tc.setupAuth(t, request, server.tokenMaker)
				server.router.ServeHTTP(recorder, request)

				tc.checkResponse(t, recorder)
			},
		)
	}
}
func TestDeleteTeacher(t *testing.T) {
	user, _ := randomUser()
	teacher := randomTeacher()

	testCases := []struct {
		name          string
		teacherID     int32
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(T *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			teacherID: teacher.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteTeacher(gomock.Any(), gomock.Eq(teacher.ID)).
					Times(1)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)

			},
		},
		{
			name:      "Not Found",
			teacherID: teacher.ID + 1,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteTeacher(gomock.Any(), gomock.Any()).
					Times(1).
					Return(sql.ErrNoRows)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)

			},
		},
		{
			name:      "Invalid ID",
			teacherID: -1,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteTeacher(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)

			},
		},
		{
			name:      "Internal Server Error",
			teacherID: teacher.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteTeacher(gomock.Any(), gomock.Eq(teacher.ID)).
					Times(1).
					Return(sql.ErrConnDone)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)

			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(
			tc.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)
				defer ctrl.Finish()

				store := mockdb.NewMockStore(ctrl)
				tc.buildStub(store)

				server := newTestServer(t, store)
				recorder := httptest.NewRecorder()

				url := fmt.Sprintf("/v1/teachers/%d", tc.teacherID)
				request, err := http.NewRequest(http.MethodDelete, url, nil)
				require.NoError(t, err)

				tc.setupAuth(t, request, server.tokenMaker)

				server.router.ServeHTTP(recorder, request)

				tc.checkResponse(t, recorder)
			},
		)
	}
}

func TestListTeachers(t *testing.T) {
	user, _ := randomUser()
	teachers := make([]db.Teacher, 0)

	for i := 0; i < 5; i++ {
		teachers = append(teachers, randomTeacher())
	}

	testCases := []struct {
		name          string
		pageID        int32
		pageSize      int32
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStub     func(store *mockdb.MockStore, params db.ListTeachersParams)
		checkResponse func(T *testing.T, recorder *httptest.ResponseRecorder, params db.ListTeachersParams)
	}{
		{
			name:     "OK",
			pageID:   1,
			pageSize: 5,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore, params db.ListTeachersParams) {

				store.EXPECT().
					ListTeachers(gomock.Any(), gomock.Eq(params)).
					Times(1).
					Return(teachers, nil)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder, params db.ListTeachersParams) {
				require.Equal(t, http.StatusOK, recorder.Code)

				// read body
				data, err := io.ReadAll(recorder.Body)
				require.NoError(t, err)

				// unmarshal JSON
				var gotTeachers []db.Teacher
				err = json.Unmarshal(data, &gotTeachers)
				require.NoError(t, err)

				// compare length and values
				require.Len(t, gotTeachers, int(params.Limit))
				require.Equal(t, gotTeachers, teachers)

			},
		},
		{
			name:     "Invalid Page ID",
			pageID:   -1,
			pageSize: 5,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore, params db.ListTeachersParams) {

				store.EXPECT().
					ListTeachers(gomock.Any(), gomock.Eq(params)).
					Times(0)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder, params db.ListTeachersParams) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:     "Invalid Page Size",
			pageID:   1,
			pageSize: 15,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore, params db.ListTeachersParams) {

				store.EXPECT().
					ListTeachers(gomock.Any(), gomock.Eq(params)).
					Times(0)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder, params db.ListTeachersParams) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:     "Internal Server Error",
			pageID:   1,
			pageSize: 5,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore, params db.ListTeachersParams) {

				store.EXPECT().
					ListTeachers(gomock.Any(), gomock.Eq(params)).
					Times(1).
					Return([]db.Teacher{}, sql.ErrConnDone)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder, params db.ListTeachersParams) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			params := db.ListTeachersParams{
				Limit:  tc.pageSize,
				Offset: (tc.pageID - 1) * tc.pageSize,
			}
			tc.buildStub(store, params)

			// start test server and send request
			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/v1/teachers?page_id=%d&page_size=%d", tc.pageID, tc.pageSize)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder, params)
		})
	}

}

func TestUpdateTeacherByID(t *testing.T) {
	user, _ := randomUser()
	teacher := randomTeacher()

	testCases := []struct {
		name          string
		teacherID     int32
		body          gin.H
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(T *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			teacherID: teacher.ID,
			body: gin.H{
				"first_name":    "new first name",
				"last_name":     "new last name",
				"middle_name":   "new middle name",
				"subject_id":    1,
				"department_id": 13,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				params := db.UpdateTeacherParams{
					ID:           teacher.ID,
					FirstName:    "new first name",
					LastName:     "new last name",
					MiddleName:   "new middle name",
					SubjectID:    1,
					DepartmentID: 13,
				}
				store.EXPECT().
					GetTeacherByID(gomock.Any(), gomock.Eq(teacher.ID)).
					Times(1).
					Return(teacher, nil)

				updatedTeacher := db.Teacher{
					FirstName:    params.FirstName,
					LastName:     params.FirstName,
					MiddleName:   params.MiddleName,
					SubjectID:    params.SubjectID,
					DepartmentID: params.DepartmentID,
				}
				store.EXPECT().
					UpdateTeacher(gomock.Any(), gomock.Eq(params)).
					Times(1).
					Return(updatedTeacher, nil)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				data, err := io.ReadAll(recorder.Body)
				require.NoError(t, err)
				var gotTeacher db.Teacher
				err = json.Unmarshal(data, &gotTeacher)
				require.NoError(t, err)
				require.Equal(t, "new first name", gotTeacher.FirstName)
			},
		},
		{
			name:      "Invalid ID",
			teacherID: -1,
			body: gin.H{
				"first_name":    "new first name",
				"last_name":     "new last name",
				"middle_name":   "new middle name",
				"subject_id":    1,
				"department_id": 13,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				params := db.UpdateTeacherParams{
					ID:           teacher.ID,
					FirstName:    "new first name",
					LastName:     "new last name",
					MiddleName:   "new middle name",
					SubjectID:    1,
					DepartmentID: 13,
				}

				store.EXPECT().
					GetTeacherByID(gomock.Any(), gomock.Eq(teacher.ID)).
					Times(0)

				store.EXPECT().
					UpdateTeacher(gomock.Any(), gomock.Eq(params)).
					Times(0)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:      "Not Found",
			teacherID: teacher.ID + 1,
			body: gin.H{
				"first_name":    "new first name",
				"last_name":     "new last name",
				"middle_name":   "new middle name",
				"subject_id":    1,
				"department_id": 13,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				params := db.UpdateTeacherParams{
					ID:           teacher.ID,
					FirstName:    "new first name",
					LastName:     "new last name",
					MiddleName:   "new middle name",
					SubjectID:    1,
					DepartmentID: 13,
				}
				store.EXPECT().
					GetTeacherByID(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Teacher{}, sql.ErrNoRows)

				store.EXPECT().
					UpdateTeacher(gomock.Any(), gomock.Eq(params)).
					Times(0)

			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		// {
		// 	name:      "Invalid body [incorrect json key]",
		// 	teacherID: teacher.ID,
		// 	body: gin.H{
		// 		"f_names":       "new first name",
		// 		"l_name":        "new last name",
		// 		"middle_name":   "new middle name",
		// 		"subject_id":    1,
		// 		"department_id": 13,
		// 	},
		// 	buildStub: func(store *mockdb.MockStore) {
		// 		params := db.UpdateTeacherParams{
		// 			ID:           teacher.ID,
		// 			FirstName:    "new first name",
		// 			LastName:     "new last name",
		// 			MiddleName:   sql.NullString{String: "new middle name", Valid: true},
		// 			SubjectID:    sql.NullInt32{Int32: 1, Valid: true},
		// 			DepartmentID: sql.NullInt32{Int32: 13, Valid: true},
		// 		}
		// 		store.EXPECT().
		// 			GetTeacherByID(gomock.Any(), gomock.Eq(teacher.ID)).
		// 			Times(1)

		// 		store.EXPECT().
		// 			UpdateTeacher(gomock.Any(), gomock.Eq(params)).
		// 			Times(1)

		// 	},
		// 	checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
		// 		require.Equal(t, http.StatusBadRequest, recorder.Code)
		// 	},
		// },
		{
			name:      "Internal Server Error [UpdateTeacher]",
			teacherID: teacher.ID,
			body: gin.H{
				"first_name":    "new first name",
				"last_name":     "new last name",
				"middle_name":   "new middle name",
				"subject_id":    1,
				"department_id": 13,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				params := db.UpdateTeacherParams{
					ID:           teacher.ID,
					FirstName:    "new first name",
					LastName:     "new last name",
					MiddleName:   "new middle name",
					SubjectID:    1,
					DepartmentID: 13,
				}

				store.EXPECT().
					GetTeacherByID(gomock.Any(), gomock.Eq(teacher.ID)).
					Times(1).
					Return(teacher, nil)

				store.EXPECT().
					UpdateTeacher(gomock.Any(), gomock.Eq(params)).
					Times(1).
					Return(db.Teacher{}, sql.ErrConnDone)

			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:      "Internal Server Error [GetTeacher]",
			teacherID: teacher.ID,
			body: gin.H{
				"first_name":    "new first name",
				"last_name":     "new last name",
				"middle_name":   "new middle name",
				"subject_id":    1,
				"department_id": 13,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				params := db.UpdateTeacherParams{
					ID:           teacher.ID,
					FirstName:    "new first name",
					LastName:     "new last name",
					MiddleName:   "new middle name",
					SubjectID:    1,
					DepartmentID: 13,
				}

				store.EXPECT().
					GetTeacherByID(gomock.Any(), gomock.Eq(teacher.ID)).
					Times(1).
					Return(db.Teacher{}, sql.ErrConnDone)

				store.EXPECT().
					UpdateTeacher(gomock.Any(), gomock.Eq(params)).
					Times(0)

			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStub(store)

			// marshal request to json
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			// start test server and send request
			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/v1/teachers/%d", tc.teacherID)
			request, err := http.NewRequest(http.MethodPatch, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func compareTeacherResponse(t *testing.T, body *bytes.Buffer, teacher db.Teacher) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotTeacher db.Teacher
	err = json.Unmarshal(data, &gotTeacher)
	require.NoError(t, err)
	require.Equal(t, teacher, gotTeacher)
}
func randomTeacher() db.Teacher {
	return db.Teacher{
		ID:           utils.RandomInt(1, 1000),
		FirstName:    utils.RandomString(5),
		LastName:     utils.RandomString(7),
		MiddleName:   utils.RandomString(5),
		SubjectID:    utils.RandomInt(1, 23),
		DepartmentID: utils.RandomInt(1, 23),
		CreatedAt:    utils.RandomTime(),
		UpdatedAt:    utils.RandomTime(),
	}
}
