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

func TestCreateScore(t *testing.T) {
	score := randomScore()

	testCases := []struct {
		name          string
		body          gin.H
		buildStub     func(store *mockdb.MockStore, params db.CreateScoreParams)
		checkResponse func(T *testing.T, recorder *httptest.ResponseRecorder, params db.CreateScoreParams)
	}{
		{
			name: "OK",
			body: gin.H{
				"student_id":    score.StudentID,
				"term_score_id": score.TermScoreID,
			},
			buildStub: func(store *mockdb.MockStore, params db.CreateScoreParams) {
				store.EXPECT().
					CreateScore(gomock.Any(), gomock.Eq(params)).
					Times(1).
					Return(score, nil)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder, params db.CreateScoreParams) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "Empty body",
			body: gin.H{},
			buildStub: func(store *mockdb.MockStore, params db.CreateScoreParams) {
				store.EXPECT().
					CreateScore(gomock.Any(), gomock.Eq(params)).
					Times(0)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder, params db.CreateScoreParams) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Invalid body",
			body: gin.H{
				"student_id_id": "",
			},
			buildStub: func(store *mockdb.MockStore, params db.CreateScoreParams) {
				store.EXPECT().
					CreateScore(gomock.Any(), gomock.Eq(params)).
					Times(0)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder, params db.CreateScoreParams) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Internal Server Error",
			body: gin.H{
				"student_id":    score.StudentID,
				"term_score_id": score.TermScoreID,
			},
			buildStub: func(store *mockdb.MockStore, params db.CreateScoreParams) {
				store.EXPECT().
					CreateScore(gomock.Any(), gomock.Eq(params)).
					Times(1).
					Return(db.Score{}, sql.ErrConnDone)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder, params db.CreateScoreParams) {
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

			var params db.CreateScoreParams
			unmarshalParams(t, tc.body, &params)

			tc.buildStub(store, params)

			// start test server and send request
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			// marshal request to json
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/v1/scores"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder, params)
		})
	}

}
func TestGetScoreByStudentID(t *testing.T) {
	score := randomScore()
	termScore := randomTermScore()
	testCases := []struct {
		name          string
		studentID     int32
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(T *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			studentID: score.StudentID,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetScoreByStudentID(gomock.Any(), gomock.Eq(score.StudentID)).
					Times(1).
					Return(score, nil)

				store.EXPECT().
					GetTermScoreByID(gomock.Any(), gomock.Eq(score.TermScoreID)).
					Times(1).
					Return(termScore, nil)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:      "Not Found: student_id",
			studentID: score.StudentID + 1,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetScoreByStudentID(gomock.Any(), gomock.Eq(score.StudentID+1)).
					Times(1).
					Return(db.Score{}, sql.ErrNoRows)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:      "Not Found: term_score_id",
			studentID: score.StudentID,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetScoreByStudentID(gomock.Any(), gomock.Eq(score.StudentID)).
					Times(1).
					Return(score, nil)

				store.EXPECT().
					GetTermScoreByID(gomock.Any(), gomock.Eq(score.TermScoreID)).
					Times(1).
					Return(db.TermScore{}, sql.ErrNoRows)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:      "Invalid Student ID",
			studentID: -1,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetScoreByStudentID(gomock.Any(), gomock.Eq(-1)).
					Times(0)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:      "Internal Server Error: student_id",
			studentID: score.StudentID,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetScoreByStudentID(gomock.Any(), gomock.Eq(score.StudentID)).
					Times(1).
					Return(db.Score{}, sql.ErrConnDone)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:      "Internal Server Error: term_score_id",
			studentID: score.StudentID,
			buildStub: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetScoreByStudentID(gomock.Any(), gomock.Eq(score.StudentID)).
					Times(1).
					Return(score, nil)

				store.EXPECT().
					GetTermScoreByID(gomock.Any(), gomock.Eq(score.TermScoreID)).
					Times(1).
					Return(db.TermScore{}, sql.ErrConnDone)
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
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/v1/scores/%d", tc.studentID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}

}

func TestDeleteScoreByID(t *testing.T) {
	score := randomScore()

	testCases := []struct {
		name          string
		scoreID       int32
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(T *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:    "OK",
			scoreID: score.StudentID,
			buildStub: func(store *mockdb.MockStore) {
				// build stubs
				store.EXPECT().
					DeleteScore(gomock.Any(), gomock.Eq(score.StudentID)).
					Times(1)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)

			},
		},
		{
			name:    "Not found",
			scoreID: score.StudentID + 1,
			buildStub: func(store *mockdb.MockStore) {
				// build stubs
				store.EXPECT().
					DeleteScore(gomock.Any(), gomock.Any()).
					Times(1).
					Return(sql.ErrNoRows)
			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)

			},
		},
		{
			name:    "Invalid ID",
			scoreID: -1,
			buildStub: func(store *mockdb.MockStore) {
				// build stubs
				store.EXPECT().
					DeleteScore(gomock.Any(), gomock.Any()).
					Times(0)

			},
			checkResponse: func(T *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)

			},
		},
		{
			name:    "Internal Server Error",
			scoreID: score.StudentID,
			buildStub: func(store *mockdb.MockStore) {
				// build stubs
				store.EXPECT().
					DeleteScore(gomock.Any(), gomock.Eq(score.StudentID)).
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
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/v1/scores/%d", tc.scoreID)
			request, err := http.NewRequest(http.MethodDelete, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func unmarshalParams(t *testing.T, body gin.H, v interface{}) {
	jsonData, err := json.Marshal(body)
	require.NoError(t, err)

	err = json.Unmarshal(jsonData, v)
	require.NoError(t, err)
}

func randomScore() db.Score {
	return db.Score{
		StudentID:   utils.RandomInt(1, 100),
		TermScoreID: utils.RandomInt(1, 100),
		CreatedAt:   utils.RandomTime(),
		UpdatedAt:   utils.RandomTime(),
	}
}

func randomTermScore() db.TermScore {
	return db.TermScore{
		ID:         utils.RandomInt(1, 1000),
		Assessment: float64(utils.RandomInt(1, 40)),
		Exam:       float64(utils.RandomInt(1, 60)),
		SubjectID:  utils.RandomInt(1, 23),
		TermID:     utils.RandomInt(1, 3),
		SessionID:  utils.RandomInt(1, 10),
		ClassID:    utils.RandomInt(1, 6),
		ArmID:      utils.RandomInt(1, 6),
		CreatedAt:  utils.RandomTime(),
		UpdatedAt:  utils.RandomTime(),
	}
}
