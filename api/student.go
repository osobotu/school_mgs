package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/osobotu/school_mgs/db/sqlc"
)

type createStudentRequest struct {
	FirstName    string  `json:"first_name" binding:"required"`
	LastName     string  `json:"last_name" binding:"required"`
	MiddleName   *string `json:"middle_name"`
	ClassID      int32   `json:"class_id"  binding:"required"`
	DepartmentID int32   `json:"department_id"  binding:"required"`
}

func (server *Server) createStudent(ctx *gin.Context) {
	var req createStudentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var mn sql.NullString
	if req.MiddleName != nil {
		mn.Scan(*req.MiddleName)
	}

	var classID sql.NullInt32
	classID.Scan(req.ClassID)

	var departmentID sql.NullInt32
	classID.Scan(req.DepartmentID)

	arg := db.CreateStudentParams{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		MiddleName:   mn,
		ClassID:      classID,
		DepartmentID: departmentID,
	}

	student, err := server.store.CreateStudent(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, student)
}

func (server *Server) getStudentByID(ctx *gin.Context) {
	var req idRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	student, err := server.store.GetStudentById(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, student)
}

func (server *Server) deleteStudent(ctx *gin.Context) {
	var req idRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteStudent(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, deleteMessage)
}

func (server *Server) listStudents(ctx *gin.Context) {
	var req listRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListStudentsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	students, err := server.store.ListStudents(ctx, arg)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, students)
}

type updateStudentRequest struct {
	FirstName    *string `json:"first_name"`
	LastName     *string `json:"last_name"`
	MiddleName   *string `json:"middle_name"`
	ClassID      *int32  `json:"class_id"`
	DepartmentID *int32  `json:"department_id"`
}

func (server *Server) updateStudent(ctx *gin.Context) {
	var reqID idRequest
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var reqData updateStudentRequest
	if err := ctx.ShouldBindJSON(&reqData); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	student, err := server.store.GetStudentById(ctx, reqID.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.UpdateStudentParams{
		ID:         student.ID,
		FirstName:  student.FirstName,
		LastName:   student.LastName,
		MiddleName: student.MiddleName,

		ClassID:      student.ClassID,
		DepartmentID: student.DepartmentID,
	}

	if reqData.FirstName != nil {
		arg.FirstName = *reqData.FirstName
	}

	if reqData.LastName != nil {
		arg.LastName = *reqData.LastName
	}

	if reqData.MiddleName != nil {
		var mn sql.NullString
		mn.Scan(*reqData.MiddleName)
		arg.MiddleName = mn
	}

	if reqData.ClassID != nil {
		var cid sql.NullInt32
		cid.Scan(*reqData.ClassID)
		arg.ClassID = cid
	}

	if reqData.DepartmentID != nil {
		var did sql.NullInt32
		did.Scan(*reqData.DepartmentID)
		arg.DepartmentID = did
	}

	student, err = server.store.UpdateStudent(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, student)
}
