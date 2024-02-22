package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/osobotu/school_mgs/db/sqlc"
	"github.com/osobotu/school_mgs/token"
	"github.com/osobotu/school_mgs/utils"
)

const deleteMessage = "Deleted Successfully"

type idRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type listRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

type Server struct {
	config     utils.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{config: config, store: store, tokenMaker: tokenMaker}

	server.setupRouter()

	return server, err
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) setupRouter() {
	router := gin.Default()

	// routes will be added here
	// if you pass in multiple routers, the last one should be the http router
	// while the others should be middlewares

	v1 := router.Group("/v1")
	{

		// ! users
		v1.POST("/users", server.createUser)
		v1.POST("/users/login", server.loginUser)
		v1.GET("/users/:id", server.getUserByID)

		// ! roles
		v1.POST("/roles", server.createRole)
		v1.GET("/roles/:id", server.getRoleByID)

		// ! subjects
		v1.POST("/subjects", server.createSubject)
		v1.GET("/subjects/:id", server.getSubjectByID)
		v1.GET("/subjects", server.listSubjects)
		v1.DELETE("/subjects/:id", server.deleteSubjectByID)

		// ! teachers
		v1.POST("/teachers", server.createTeacher)
		v1.GET("/teachers/:id", server.getTeacherByID)
		v1.GET("/teachers", server.listTeachers)
		v1.DELETE("/teachers/:id", server.deleteTeacherByID)
		v1.PATCH("/teachers/:id", server.updateTeacherByID)

		// ! classes
		v1.POST("/classes", server.createClass)
		v1.GET("/classes/:id", server.getClassByID)
		v1.GET("/classes", server.getClassByName)
		v1.GET("/all-classes", server.listClasses)
		v1.DELETE("/classes/:id", server.deleteClassByID)

		// ! sessions
		v1.POST("/sessions", server.createSession)
		v1.GET("/sessions/:id", server.getSessionByID)
		v1.DELETE("/sessions/:id", server.deleteSession)

		// ! terms
		v1.POST("/terms", server.createTerm)
		v1.GET("/terms/:id", server.getTermByID)
		v1.DELETE("/terms/:id", server.deleteTerm)

		// ! students
		v1.POST("/students", server.createStudent)
		v1.GET("/students/:id", server.getStudentByID)
		v1.PATCH("/students/:id", server.updateStudent)
		v1.GET("/students", server.listStudents)
		v1.DELETE("/students/:id", server.deleteStudent)

		// ! scores
		v1.POST("/scores", server.createScore)
		v1.GET("/scores/:id", server.getScoreByStudentID)
		v1.DELETE("/scores/:id", server.deleteScore)

		// ! term_scores
		v1.POST("/term_scores", server.createTermScore)
		v1.GET("/term_scores/:id", server.getTermScoreByID)
		v1.GET("/term_scores", server.listTermScoresForSubjectAndClass)
		v1.PATCH("/term_scores/:id", server.updateTermScoreByID)
		v1.DELETE("/term_scores/:id", server.deleteTermScoreByID)

	}

	server.router = router
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
