package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	mockdb "github.com/osobotu/school_mgs/db/mock"
	db "github.com/osobotu/school_mgs/db/sqlc"
	"github.com/osobotu/school_mgs/db/utils"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCreateRole(t *testing.T) {
	role := randomRole()

	testCases := []struct {
		name          string
		body          gin.H
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(T *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"name": role.Role,
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateRole(gomock.Any(), gomock.Eq(role.Role)).
					Times(1).
					Return(role, nil)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "Empty body",
			body: gin.H{},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateRole(gomock.Any(), gomock.Any()).
					Times(0)

			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Invalid Body",
			body: gin.H{
				"names": "First",
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateRole(gomock.Any(), gomock.Any()).
					Times(0)

			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Internal Server Error",
			body: gin.H{
				"name": role.Role,
			},
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateRole(gomock.Any(), gomock.Eq(role.Role)).
					Times(1).
					Return(db.Role{}, sql.ErrConnDone)

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

			// var params string
			// unmarshalParams(t, tc.body, &params)

			tc.buildStub(store)

			// start test server and send request
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			// marshal request to json
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/v1/roles"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})

	}
}

func TestGetRoleByID(t *testing.T) {
	role := randomRole()

	testCases := []struct {
		name          string
		RoleID        int32
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(T *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			RoleID: role.ID,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetRoleByID(gomock.Any(), gomock.Eq(role.ID)).
					Times(1).
					Return(role, nil)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:   "Not Found",
			RoleID: role.ID + 1,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetRoleByID(gomock.Any(), gomock.Eq(role.ID+1)).
					Times(1).
					Return(db.Role{}, sql.ErrNoRows)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:   "Invalid ID",
			RoleID: -1,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetRoleByID(gomock.Any(), gomock.Eq(-1)).
					Times(0)

			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:   "Internal server error",
			RoleID: role.ID,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetRoleByID(gomock.Any(), gomock.Eq(role.ID)).
					Times(1).
					Return(db.Role{}, sql.ErrConnDone)

			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		store := mockdb.NewMockStore(ctrl)
		tc.buildStub(store)

		// start test server and send request
		server := NewServer(store)
		recorder := httptest.NewRecorder()

		url := fmt.Sprintf("/v1/roles/%d", tc.RoleID)
		request, err := http.NewRequest(http.MethodGet, url, nil)
		require.NoError(t, err)

		server.router.ServeHTTP(recorder, request)
		tc.checkResponse(t, recorder)
	}
}

func randomRole() db.Role {
	return db.Role{
		ID:   utils.RandomInt(1, 1000),
		Role: utils.RandomString(5),
	}
}
