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
	"github.com/osobotu/school_mgs/utils"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCreateTerm(t *testing.T) {
	term := randomTerm()

	testCases := []struct {
		name          string
		body          gin.H
		buildStub     func(store *mockdb.MockStore, params db.CreateTermParams)
		checkResponse func(T *testing.T, recorder *httptest.ResponseRecorder, params db.CreateTermParams)
	}{
		{
			name: "OK",
			body: gin.H{
				"name":   "First",
				"number": 1,
			},
			buildStub: func(store *mockdb.MockStore, params db.CreateTermParams) {
				store.EXPECT().
					CreateTerm(gomock.Any(), gomock.Eq(params)).
					Times(1).
					Return(term, nil)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder, params db.CreateTermParams) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "Empty body",
			body: gin.H{},
			buildStub: func(store *mockdb.MockStore, params db.CreateTermParams) {
				store.EXPECT().
					CreateTerm(gomock.Any(), gomock.Eq(params)).
					Times(0)

			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder, params db.CreateTermParams) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Invalid Body",
			body: gin.H{
				"names":   "First",
				"numbers": "2",
			},
			buildStub: func(store *mockdb.MockStore, params db.CreateTermParams) {
				store.EXPECT().
					CreateTerm(gomock.Any(), gomock.Eq(params)).
					Times(0)

			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder, params db.CreateTermParams) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Internal Server Error",
			body: gin.H{
				"name":   "First",
				"number": 1,
			},
			buildStub: func(store *mockdb.MockStore, params db.CreateTermParams) {
				store.EXPECT().
					CreateTerm(gomock.Any(), gomock.Eq(params)).
					Times(1).
					Return(db.Term{}, sql.ErrConnDone)

			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder, params db.CreateTermParams) {
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

			var params db.CreateTermParams
			unmarshalParams(t, tc.body, &params)

			tc.buildStub(store, params)

			// start test server and send request
			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			// marshal request to json
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/v1/terms"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder, params)
		})

	}
}

func TestGetTermByID(t *testing.T) {
	term := randomTerm()

	testCases := []struct {
		name          string
		termID        int32
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(T *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			termID: term.ID,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetTermByID(gomock.Any(), gomock.Eq(term.ID)).
					Times(1).
					Return(term, nil)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:   "Not Found",
			termID: term.ID + 1,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetTermByID(gomock.Any(), gomock.Eq(term.ID+1)).
					Times(1).
					Return(db.Term{}, sql.ErrNoRows)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:   "Invalid ID",
			termID: -1,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetTermByID(gomock.Any(), gomock.Eq(-1)).
					Times(0)

			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:   "Internal server error",
			termID: term.ID,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetTermByID(gomock.Any(), gomock.Eq(term.ID)).
					Times(1).
					Return(db.Term{}, sql.ErrConnDone)

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

		url := fmt.Sprintf("/v1/terms/%d", tc.termID)
		request, err := http.NewRequest(http.MethodGet, url, nil)
		require.NoError(t, err)

		server.router.ServeHTTP(recorder, request)
		tc.checkResponse(t, recorder)
	}
}
func TestDeleteTerm(t *testing.T) {
	term := randomTerm()

	testCases := []struct {
		name          string
		termID        int32
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(T *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			termID: term.ID,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteTerm(gomock.Any(), gomock.Eq(term.ID)).
					Times(1).
					Return(nil)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:   "Not Found",
			termID: term.ID + 1,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteTerm(gomock.Any(), gomock.Eq(term.ID+1)).
					Times(1).
					Return(sql.ErrNoRows)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:   "Invalid ID",
			termID: -1,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteTerm(gomock.Any(), gomock.Eq(-1)).
					Times(0)

			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:   "Internal server error",
			termID: term.ID,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteTerm(gomock.Any(), gomock.Eq(term.ID)).
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

		url := fmt.Sprintf("/v1/terms/%d", tc.termID)
		request, err := http.NewRequest(http.MethodDelete, url, nil)
		require.NoError(t, err)

		server.router.ServeHTTP(recorder, request)
		tc.checkResponse(t, recorder)
	}
}

func randomTerm() db.Term {
	return db.Term{
		ID:        utils.RandomInt(1, 1000),
		Name:      utils.RandomString(3),
		Number:    utils.RandomInt(1, 3),
		CreatedAt: utils.RandomTime(),
		UpdatedAt: utils.RandomTime(),
	}
}
