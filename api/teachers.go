package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/osobotu/school_mgs/db/sqlc"
)

type createTeacherRequest struct {
	FirstName  string   `json:"first_name" binding:"required"`
	LastName   string   `json:"last_name" binding:"required"`
	MiddleName *string  `json:"middle_name"`
	SubjectID  *int32   `json:"subject_id"`
	Classes    *[]int32 `json:"classes"`
}

func (server *Server) createTeacher(ctx *gin.Context) {
	var req createTeacherRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var middleName sql.NullString
	if req.MiddleName != nil {
		middleName.Scan(*req.MiddleName)
	}
	var subjectID sql.NullInt32
	if req.SubjectID != nil {
		subjectID.Scan(*req.SubjectID)
	}

	arg := db.CreateTeacherParams{
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		MiddleName: middleName,
		SubjectID:  subjectID,
		Classes:    *req.Classes,
	}

	teacher, err := server.store.CreateTeacher(ctx, arg)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, teacher)

}

type getTeacherByIDRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getTeacherByID(ctx *gin.Context) {
	var req getTeacherByIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	teacher, err := server.store.GetTeacherById(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, teacher)
}

type deleteTeacherByIDRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteTeacherByID(ctx *gin.Context) {
	var req deleteTeacherByIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteTeacher(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "Deleted successfully")
}

type listTeachersRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listTeachers(ctx *gin.Context) {
	var req listTeachersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListTeachersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	teachers, err := server.store.ListTeachers(ctx, arg)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, teachers)
}

type updateTeacherID struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}
type updateTeacherByIDRequest struct {
	FirstName  *string  `json:"first_name"`
	LastName   *string  `json:"last_name"`
	MiddleName *string  `json:"middle_name"`
	SubjectID  *int32   `json:"subject_id"`
	Classes    *[]int32 `json:"classes"`
}

func (server *Server) updateTeacherByID(ctx *gin.Context) {

	var reqID updateTeacherID
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var req updateTeacherByIDRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	teacher, err := server.store.GetTeacherById(ctx, reqID.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.UpdateTeacherParams{
		ID:         teacher.ID,
		FirstName:  teacher.FirstName,
		LastName:   teacher.LastName,
		MiddleName: teacher.MiddleName,
		SubjectID:  teacher.SubjectID,
		Classes:    teacher.Classes,
	}

	if req.FirstName != nil {
		arg.FirstName = *req.FirstName
	}
	if req.LastName != nil {
		arg.LastName = *req.LastName
	}
	if req.MiddleName != nil {
		var mdn sql.NullString
		mdn.Scan(*req.MiddleName)
		arg.MiddleName = mdn
	}
	if req.SubjectID != nil {
		var id sql.NullInt32
		id.Scan(*req.SubjectID)
		arg.SubjectID = id
	}
	if req.Classes != nil {
		arg.Classes = *req.Classes
	}

	teacher, err = server.store.UpdateTeacher(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, teacher)
}
