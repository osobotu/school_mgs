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

func TestCreateUser(t *testing.T) {
	user := randomUser()

	testCases := []struct {
		name          string
		body          gin.H
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(T *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"email":    user.Email,
				"password": user.PasswordHash,
				"role_id":  user.RoleID,
			},
			buildStub: func(store *mockdb.MockStore) {
				params := db.CreateUserParams{
					Email:        user.Email,
					PasswordHash: user.PasswordHash,
					RoleID:       user.RoleID,
				}
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Eq(params)).
					Times(1).
					Return(user, nil)
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
					CreateUser(gomock.Any(), gomock.Any()).
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
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)

			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Internal Server Error",
			body: gin.H{
				"email":    user.Email,
				"password": user.PasswordHash,
				"role_id":  user.RoleID,
			},
			buildStub: func(store *mockdb.MockStore) {
				params := db.CreateUserParams{
					Email:        user.Email,
					PasswordHash: user.PasswordHash,
					RoleID:       user.RoleID,
				}
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Eq(params)).
					Times(1).
					Return(db.User{}, sql.ErrConnDone)

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

			url := "/v1/users"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})

	}
}

func TestGetUserByID(t *testing.T) {
	user := randomUser()

	testCases := []struct {
		name          string
		userID        int32
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(T *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			userID: user.ID,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserByID(gomock.Any(), gomock.Eq(user.ID)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:   "Not Found",
			userID: user.ID + 1,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserByID(gomock.Any(), gomock.Eq(user.ID+1)).
					Times(1).
					Return(db.User{}, sql.ErrNoRows)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:   "Invalid ID",
			userID: -1,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserByID(gomock.Any(), gomock.Eq(-1)).
					Times(0)

			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:   "Internal server error",
			userID: user.ID,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserByID(gomock.Any(), gomock.Eq(user.ID)).
					Times(1).
					Return(db.User{}, sql.ErrConnDone)

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

		url := fmt.Sprintf("/v1/users/%d", tc.userID)
		request, err := http.NewRequest(http.MethodGet, url, nil)
		require.NoError(t, err)

		server.router.ServeHTTP(recorder, request)
		tc.checkResponse(t, recorder)
	}
}

func randomUser() db.User {
	return db.User{
		ID:                utils.RandomInt(1, 1000),
		Email:             utils.RandomEmail(),
		PasswordHash:      "hash",
		RoleID:            utils.RandomInt(1, 3),
		PasswordChangedAt: utils.RandomTime(),
		CreatedAt:         utils.RandomTime(),
		UpdatedAt:         utils.RandomTime(),
	}
}
