package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/osobotu/school_mgs/db/sqlc"
)

type createSessionRequest struct {
	Session   string `json:"session" binding:"required"`
	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`
}

func (server *Server) createSession(ctx *gin.Context) {
	var req createSessionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	startDate, err := time.Parse(time.RFC3339, req.StartDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var sd sql.NullTime
	sd.Scan(startDate)

	endDate, err := time.Parse(time.RFC3339, req.EndDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var ed sql.NullTime
	ed.Scan(endDate)

	arg := db.CreateSessionParams{
		Session:   req.Session,
		StartDate: sd,
		EndDate:   ed,
	}

	session, err := server.store.CreateSession(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, session)
}
