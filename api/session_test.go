package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/osobotu/school_mgs/db/mock"
	db "github.com/osobotu/school_mgs/db/sqlc"
	"github.com/osobotu/school_mgs/utils"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

// func TestCreateSession(t *testing.T) {
// 	session := randomSession()
// 	sd := utils.RandomTime().UTC().String()

// 	testCases := []struct {
// 		name          string
// 		body          gin.H
// 		buildStub     func(store *mockdb.MockStore, params db.CreateSessionParams)
// 		checkResponse func(T *testing.T, recorder *httptest.ResponseRecorder, params db.CreateSessionParams)
// 	}{
// 		{
// 			name: "OK",
// 			body: gin.H{
// 				"session":    session.Session,
// 				"start_date": sd,
// 				"end_date":   sd,
// 			},
// 			buildStub: func(store *mockdb.MockStore, params db.CreateSessionParams) {

// 				store.EXPECT().
// 					CreateSession(gomock.Any(), gomock.Eq(params)).
// 					Times(1).
// 					Return(session, nil)
// 			},

// 			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder, params db.CreateSessionParams) {
// 				require.Equal(t, http.StatusOK, recorder.Code)
// 			},
// 		},
// 	}

// 	for i := range testCases {
// 		tc := testCases[i]

// 		ctrl := gomock.NewController(t)
// 		defer ctrl.Finish()

// 		store := mockdb.NewMockStore(ctrl)

// 		var params db.CreateSessionParams
// 		unmarshalParams(t, tc.body, &params)
// 		tc.buildStub(store, params)

// 		// start test server and send request
// 		server := newTestServer(t, store)
// 		recorder := httptest.NewRecorder()

// 		// marshal request to json
// 		data, err := json.Marshal(tc.body)
// 		require.NoError(t, err)

// 		url := "/v1/sessions"
// 		request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
// 		require.NoError(t, err)

// 		server.router.ServeHTTP(recorder, request)
// 		tc.checkResponse(t, recorder, params)
// 	}

// }

func TestGetSessionByID(t *testing.T) {
	session := randomSession()

	testCases := []struct {
		name          string
		sessionID     int32
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(T *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			sessionID: session.ID,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetSessionByID(gomock.Any(), gomock.Eq(session.ID)).
					Times(1).
					Return(session, nil)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:      "Not Found",
			sessionID: session.ID + 1,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetSessionByID(gomock.Any(), gomock.Eq(session.ID+1)).
					Times(1).
					Return(db.Session{}, sql.ErrNoRows)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:      "Invalid ID",
			sessionID: -1,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetSessionByID(gomock.Any(), gomock.Eq(-1)).
					Times(0)

			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:      "Internal server error",
			sessionID: session.ID,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetSessionByID(gomock.Any(), gomock.Eq(session.ID)).
					Times(1).
					Return(db.Session{}, sql.ErrConnDone)

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
		server := newTestServer(t, store)
		recorder := httptest.NewRecorder()

		url := fmt.Sprintf("/v1/sessions/%d", tc.sessionID)
		request, err := http.NewRequest(http.MethodGet, url, nil)
		require.NoError(t, err)

		server.router.ServeHTTP(recorder, request)
		tc.checkResponse(t, recorder)
	}
}
func TestDeleteSession(t *testing.T) {
	session := randomSession()

	testCases := []struct {
		name          string
		sessionID     int32
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(T *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			sessionID: session.ID,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteSession(gomock.Any(), gomock.Eq(session.ID)).
					Times(1).
					Return(nil)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:      "Not Found",
			sessionID: session.ID + 1,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteSession(gomock.Any(), gomock.Eq(session.ID+1)).
					Times(1).
					Return(sql.ErrNoRows)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:      "Invalid ID",
			sessionID: -1,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteSession(gomock.Any(), gomock.Eq(-1)).
					Times(0)

			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:      "Internal server error",
			sessionID: session.ID,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteSession(gomock.Any(), gomock.Eq(session.ID)).
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

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		store := mockdb.NewMockStore(ctrl)
		tc.buildStub(store)

		// start test server and send request
		server := newTestServer(t, store)
		recorder := httptest.NewRecorder()

		url := fmt.Sprintf("/v1/sessions/%d", tc.sessionID)
		request, err := http.NewRequest(http.MethodDelete, url, nil)
		require.NoError(t, err)

		server.router.ServeHTTP(recorder, request)
		tc.checkResponse(t, recorder)
	}
}

func randomSession() db.Session {
	return db.Session{
		ID:        utils.RandomInt(1, 1000),
		Session:   utils.RandomString(5),
		CreatedAt: utils.RandomTime(),
		UpdatedAt: utils.RandomTime(),
		StartDate: utils.RandomTime(),
		EndDate:   utils.RandomTime().AddDate(1, 0, 0),
	}
}
