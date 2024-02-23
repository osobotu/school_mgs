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

func TestGetSubjectByID(t *testing.T) {
	user, _ := randomUser()
	subject := randomSubject()

	testCases := []struct {
		name          string
		subjectID     int32
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(T *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			subjectID: subject.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				// build stubs
				store.EXPECT().
					GetSubjectByID(gomock.Any(), gomock.Eq(subject.ID)).
					Times(1).
					Return(subject, nil)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatch(t, recorder.Body, subject)
			},
		},
		{
			name:      "Not Found",
			subjectID: subject.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				// build stubs
				store.EXPECT().
					GetSubjectByID(gomock.Any(), gomock.Eq(subject.ID)).
					Times(1).
					Return(db.Subject{}, sql.ErrNoRows)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:      "Internal Server Error",
			subjectID: subject.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				// build stubs
				store.EXPECT().
					GetSubjectByID(gomock.Any(), gomock.Eq(subject.ID)).
					Times(1).
					Return(db.Subject{}, sql.ErrConnDone)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:      "Invalid ID",
			subjectID: 0,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				// build stubs
				store.EXPECT().
					GetSubjectByID(gomock.Any(), gomock.Any).
					Times(0)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
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

			// start test server and send request
			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/v1/subjects/%d", tc.subjectID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}

}

func TestCreateSubject(t *testing.T) {
	user, _ := randomUser()
	subject := randomSubject()

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
				"name": subject.Name,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateSubject(gomock.Any(), gomock.Eq(subject.Name)).
					Times(1).
					Return(subject, nil)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
				requireBodyMatch(t, recorder.Body, subject)
			},
		},
		{
			name: "Empty body",
			body: gin.H{},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateSubject(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Invalid body",
			body: gin.H{
				"subject_name": "test_invalid_name",
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateSubject(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Internal Server Error",
			body: gin.H{
				"name": subject.Name,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateSubject(gomock.Any(), gomock.Eq(subject.Name)).
					Times(1).
					Return(db.Subject{}, sql.ErrConnDone)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)

			},
		},
		// {
		// 	name: "Subject already exists",
		// 	body: gin.H{
		// 		"name": subject.Name,
		// 	},
		// 	buildStub: func(store *mockdb.MockStore) {
		// 		store.EXPECT().
		// 			CreateSubject(gomock.Any(), gomock.Eq(subject.Name)).
		// 			Times(1).
		// 			Return(db.Subject{}, ErrDuplicateValue{})
		// 	},
		// 	checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
		// 		require.Equal(t, http.StatusBadRequest, recorder.Code)

		// 	},
		// },
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStub(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			// marshal request to json
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/v1/subjects"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)

			server.router.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})
	}

}

func TestDeleteSubjectByID(t *testing.T) {
	user, _ := randomUser()
	subject := randomSubject()

	testCases := []struct {
		name          string
		subjectID     int32
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(T *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			subjectID: subject.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				// build stubs
				store.EXPECT().
					DeleteSubject(gomock.Any(), gomock.Eq(subject.ID)).
					Times(1)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)

			},
		},
		{
			name:      "Not found",
			subjectID: subject.ID + 1,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				// build stubs
				store.EXPECT().
					DeleteSubject(gomock.Any(), gomock.Any()).
					Times(1).
					Return(sql.ErrNoRows)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)

			},
		},
		{
			name:      "Invalid ID",
			subjectID: -1,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				// build stubs
				store.EXPECT().
					DeleteSubject(gomock.Any(), gomock.Any()).
					Times(0)

			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)

			},
		},
		{
			name:      "Internal Server Error",
			subjectID: subject.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				// build stubs
				store.EXPECT().
					DeleteSubject(gomock.Any(), gomock.Eq(subject.ID)).
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
		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStub(store)

			// start test server and send request
			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/v1/subjects/%d", tc.subjectID)
			request, err := http.NewRequest(http.MethodDelete, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestListSubjects(t *testing.T) {
	user, _ := randomUser()
	subjects := make([]db.Subject, 0)

	for i := 0; i < 5; i++ {
		subjects = append(subjects, randomSubject())
	}

	testCases := []struct {
		name          string
		pageID        int32
		pageSize      int32
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStub     func(store *mockdb.MockStore, params db.ListSubjectsParams)
		checkResponse func(T *testing.T, recorder *httptest.ResponseRecorder, params db.ListSubjectsParams)
	}{
		{
			name:     "OK",
			pageID:   1,
			pageSize: 5,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore, params db.ListSubjectsParams) {

				store.EXPECT().
					ListSubjects(gomock.Any(), gomock.Eq(params)).
					Times(1).
					Return(subjects, nil)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder, params db.ListSubjectsParams) {
				require.Equal(t, http.StatusOK, recorder.Code)

				// read body
				data, err := io.ReadAll(recorder.Body)
				require.NoError(t, err)

				// unmarshal JSON
				var gotSubjects []db.Subject
				err = json.Unmarshal(data, &gotSubjects)
				require.NoError(t, err)

				// compare length and values
				require.Len(t, gotSubjects, int(params.Limit))
				require.Equal(t, gotSubjects, subjects)

			},
		},
		{
			name:     "Invalid Page ID",
			pageID:   0,
			pageSize: 5,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore, params db.ListSubjectsParams) {

				store.EXPECT().
					ListSubjects(gomock.Any(), gomock.Eq(params)).
					Times(0)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder, params db.ListSubjectsParams) {
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
			buildStub: func(store *mockdb.MockStore, params db.ListSubjectsParams) {

				store.EXPECT().
					ListSubjects(gomock.Any(), gomock.Eq(params)).
					Times(0)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder, params db.ListSubjectsParams) {
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
			buildStub: func(store *mockdb.MockStore, params db.ListSubjectsParams) {

				store.EXPECT().
					ListSubjects(gomock.Any(), gomock.Eq(params)).
					Times(1).
					Return([]db.Subject{}, sql.ErrConnDone)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder, params db.ListSubjectsParams) {
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
			params := db.ListSubjectsParams{
				Limit:  tc.pageSize,
				Offset: (tc.pageID - 1) * tc.pageSize,
			}
			tc.buildStub(store, params)

			// start test server and send request
			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/v1/subjects?page_id=%d&page_size=%d", tc.pageID, tc.pageSize)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder, params)
		})
	}

}

type ErrDuplicateValue struct {
}

func (e ErrDuplicateValue) Error() string {
	return "duplicate key"
}

func randomSubject() db.Subject {
	return db.Subject{
		ID:        utils.RandomInt(1, 1000),
		Name:      utils.RandomString(7),
		CreatedAt: utils.RandomTime(),
		UpdatedAt: utils.RandomTime(),
	}
}

func requireBodyMatch(t *testing.T, body *bytes.Buffer, subject db.Subject) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotSubject db.Subject
	err = json.Unmarshal(data, &gotSubject)
	require.NoError(t, err)
	require.Equal(t, subject, gotSubject)
}
