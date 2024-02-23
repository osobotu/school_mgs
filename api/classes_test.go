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

func TestCreateClass(t *testing.T) {
	user, _ := randomUser()
	class := randomClass()

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
				"name": class.Name,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateClass(gomock.Any(), gomock.Eq(class.Name)).
					Times(1).
					Return(class, nil)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
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
					CreateClass(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Invalid body",
			body: gin.H{
				"class_name": "test_invalid_name",
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateClass(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Internal Server Error",
			body: gin.H{
				"name": class.Name,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateClass(gomock.Any(), gomock.Eq(class.Name)).
					Times(1).
					Return(db.Class{}, sql.ErrConnDone)
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

			// marshal request to json
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/v1/classes"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestGetClassByID(t *testing.T) {
	user, _ := randomUser()
	class := randomClass()
	testCases := []struct {
		name          string
		classID       int32
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(T *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:    "OK",
			classID: class.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetClassByID(gomock.Any(), gomock.Eq(class.ID)).
					Times(1).
					Return(class, nil)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				compareClassResponse(t, recorder.Body, class)
			},
		},
		{
			name:    "Invalid ID",
			classID: -1,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetClassByID(gomock.Any(), gomock.Eq(-1)).
					Times(0)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:    "Not found",
			classID: class.ID + 1,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetClassByID(gomock.Any(), gomock.Eq(class.ID+1)).
					Times(1).
					Return(db.Class{}, sql.ErrNoRows)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:    "Internal Server Error",
			classID: class.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetClassByID(gomock.Any(), gomock.Eq(class.ID)).
					Times(1).
					Return(db.Class{}, sql.ErrConnDone)
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

			url := fmt.Sprintf("/v1/classes/%d", tc.classID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestGetClassByName(t *testing.T) {
	user, _ := randomUser()
	class := randomClass()
	testCases := []struct {
		name          string
		className     string
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(T *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			className: class.Name,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetClassByName(gomock.Any(), gomock.Eq(class.Name)).
					Times(1).
					Return(class, nil)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				compareClassResponse(t, recorder.Body, class)
			},
		},

		{
			name:      "Not found",
			className: "xxxxx",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetClassByName(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Class{}, sql.ErrNoRows)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:      "Internal Server Error",
			className: class.Name,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetClassByName(gomock.Any(), gomock.Eq(class.Name)).
					Times(1).
					Return(db.Class{}, sql.ErrConnDone)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:      "Empty name",
			className: "",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetClassByName(gomock.Any(), gomock.Eq("")).
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

			url := fmt.Sprintf("/v1/classes?name=%s", tc.className)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestDeleteClassByID(t *testing.T) {
	user, _ := randomUser()
	class := randomClass()

	testCases := []struct {
		name          string
		classID       int32
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(T *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:    "OK",
			classID: class.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				// build stubs
				store.EXPECT().
					DeleteClass(gomock.Any(), gomock.Eq(class.ID)).
					Times(1)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)

			},
		},
		{
			name:    "Not found",
			classID: class.ID + 1,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				// build stubs
				store.EXPECT().
					DeleteClass(gomock.Any(), gomock.Any()).
					Times(1).
					Return(sql.ErrNoRows)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)

			},
		},
		{
			name:    "Invalid ID",
			classID: -1,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				// build stubs
				store.EXPECT().
					DeleteClass(gomock.Any(), gomock.Any()).
					Times(0)

			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)

			},
		},
		{
			name:    "Internal Server Error",
			classID: class.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore) {
				// build stubs
				store.EXPECT().
					DeleteClass(gomock.Any(), gomock.Eq(class.ID)).
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

			url := fmt.Sprintf("/v1/classes/%d", tc.classID)
			request, err := http.NewRequest(http.MethodDelete, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestListClasses(t *testing.T) {
	user, _ := randomUser()
	classes := make([]db.Class, 0)

	for i := 0; i < 5; i++ {
		classes = append(classes, randomClass())
	}

	testCases := []struct {
		name          string
		pageID        int32
		pageSize      int32
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStub     func(store *mockdb.MockStore, params db.ListClassesParams)
		checkResponse func(T *testing.T, recorder *httptest.ResponseRecorder, params db.ListClassesParams)
	}{
		{
			name:     "OK",
			pageID:   1,
			pageSize: 5,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore, params db.ListClassesParams) {

				store.EXPECT().
					ListClasses(gomock.Any(), gomock.Eq(params)).
					Times(1).
					Return(classes, nil)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder, params db.ListClassesParams) {
				require.Equal(t, http.StatusOK, recorder.Code)

				// read body
				data, err := io.ReadAll(recorder.Body)
				require.NoError(t, err)

				// unmarshal JSON
				var gotClasses []db.Class
				err = json.Unmarshal(data, &gotClasses)
				require.NoError(t, err)

				// compare length and values
				require.Len(t, gotClasses, int(params.Limit))
				require.Equal(t, gotClasses, classes)

			},
		},
		{
			name:     "Invalid Page ID",
			pageID:   0,
			pageSize: 5,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Email, time.Minute)
			},
			buildStub: func(store *mockdb.MockStore, params db.ListClassesParams) {

				store.EXPECT().
					ListClasses(gomock.Any(), gomock.Eq(params)).
					Times(0)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder, params db.ListClassesParams) {
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
			buildStub: func(store *mockdb.MockStore, params db.ListClassesParams) {

				store.EXPECT().
					ListClasses(gomock.Any(), gomock.Eq(params)).
					Times(0)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder, params db.ListClassesParams) {
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
			buildStub: func(store *mockdb.MockStore, params db.ListClassesParams) {

				store.EXPECT().
					ListClasses(gomock.Any(), gomock.Eq(params)).
					Times(1).
					Return([]db.Class{}, sql.ErrConnDone)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder, params db.ListClassesParams) {
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
			params := db.ListClassesParams{
				Limit:  tc.pageSize,
				Offset: (tc.pageID - 1) * tc.pageSize,
			}
			tc.buildStub(store, params)

			// start test server and send request
			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/v1/all-classes?page_id=%d&page_size=%d", tc.pageID, tc.pageSize)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder, params)
		})
	}

}

func randomClass() db.Class {
	return db.Class{
		ID:        utils.RandomInt(1, 1000),
		Name:      utils.RandomString(4),
		CreatedAt: utils.RandomTime(),
		UpdatedAt: utils.RandomTime(),
	}
}

func compareClassResponse(t *testing.T, body *bytes.Buffer, class db.Class) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotClass db.Class
	err = json.Unmarshal(data, &gotClass)
	require.NoError(t, err)

	require.Equal(t, class, gotClass)

}
