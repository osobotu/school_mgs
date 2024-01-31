package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/osobotu/school_mgs/db/sqlc"
)

type createTermScoreRequest struct {
	Assessment float64 `json:"assessment"`
	Exam       float64 `json:"exam"`
	SubjectID  int32   `json:"subject_id"`
	TermID     int32   `json:"term_id"`
	SessionID  int32   `json:"session_id"`
	ClassID    int32   `json:"class_id"`
}

func (server *Server) createTermScore(ctx *gin.Context) {
	var req createTermScoreRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateTermScoreParams{
		Assessment: req.Assessment,
		Exam:       req.Exam,
		SubjectID:  req.SubjectID,
		TermID:     req.TermID,
		SessionID:  req.SessionID,
		ClassID:    req.ClassID,
	}

	termScore, err := server.store.CreateTermScore(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, termScore)
}

func (server *Server) getTermScoreByID(ctx *gin.Context) {
	var req idRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	termScore, err := server.store.GetTermScoreByID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, termScore)
}

func (server *Server) deleteTermScoreByID(ctx *gin.Context) {
	var req idRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteTermScore(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, deleteMessage)
}

type listTermScoresForSubjectAndClassRequest struct {
	PageID    int32 `form:"page_id" binding:"required,min=1"`
	PageSize  int32 `form:"page_size" binding:"required,min=5,max=10"`
	SubjectID int32 `form:"subject_id" binding:"required,min=1"`
	ClassID   int32 `form:"class_id" binding:"required,min=1"`
}

func (server *Server) listTermScoresForSubjectAndClass(ctx *gin.Context) {

	var req listTermScoresForSubjectAndClassRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))

		return
	}

	arg := db.ListTermScoresForSubjectAndClassParams{
		Limit:     req.PageSize,
		Offset:    (req.PageID - 1) * req.PageSize,
		SubjectID: req.SubjectID,
		ClassID:   req.ClassID,
	}

	termScores, err := server.store.ListTermScoresForSubjectAndClass(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))

		return
	}
	ctx.JSON(http.StatusOK, termScores)
}

type updateTermScoreByIdRequest struct {
	Assessment *float64 `json:"assessment" binding:"required"`
	Exam       *float64 `json:"exam" binding:"required"`
}

func (server *Server) updateTermScoreByID(ctx *gin.Context) {
	var req idRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var reqData updateTermScoreByIdRequest
	if err := ctx.ShouldBindJSON(&reqData); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	termScore, err := server.store.GetTermScoreByID(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.UpdateTermScoreByIDParams{
		ID:         termScore.ID,
		Assessment: termScore.Assessment,
		Exam:       termScore.Exam,
		UpdatedAt:  time.Now().UTC(),
	}

	if reqData.Assessment != nil {
		arg.Assessment = *reqData.Assessment
	}
	if reqData.Exam != nil {
		arg.Exam = *reqData.Exam
	}

	termScore, err = server.store.UpdateTermScoreByID(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, termScore)
}
