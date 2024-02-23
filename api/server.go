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

		authRoutes := v1.Group("/").Use(authMiddleware(server.tokenMaker))

		// ! roles
		authRoutes.POST("/roles", server.createRole)
		authRoutes.GET("/roles/:id", server.getRoleByID)

		// ! subjects
		authRoutes.POST("/subjects", server.createSubject)
		authRoutes.GET("/subjects/:id", server.getSubjectByID)
		authRoutes.GET("/subjects", server.listSubjects)
		authRoutes.DELETE("/subjects/:id", server.deleteSubjectByID)

		// ! teachers
		authRoutes.POST("/teachers", server.createTeacher)
		authRoutes.GET("/teachers/:id", server.getTeacherByID)
		authRoutes.GET("/teachers", server.listTeachers)
		authRoutes.DELETE("/teachers/:id", server.deleteTeacherByID)
		authRoutes.PATCH("/teachers/:id", server.updateTeacherByID)

		// ! classes
		authRoutes.POST("/classes", server.createClass)
		authRoutes.GET("/classes/:id", server.getClassByID)
		authRoutes.GET("/classes", server.getClassByName)
		authRoutes.GET("/all-classes", server.listClasses)
		authRoutes.DELETE("/classes/:id", server.deleteClassByID)

		// ! sessions
		authRoutes.POST("/sessions", server.createSession)
		authRoutes.GET("/sessions/:id", server.getSessionByID)
		authRoutes.DELETE("/sessions/:id", server.deleteSession)

		// ! terms
		authRoutes.POST("/terms", server.createTerm)
		authRoutes.GET("/terms/:id", server.getTermByID)
		authRoutes.DELETE("/terms/:id", server.deleteTerm)

		// ! students
		authRoutes.POST("/students", server.createStudent)
		authRoutes.GET("/students/:id", server.getStudentByID)
		authRoutes.PATCH("/students/:id", server.updateStudent)
		authRoutes.GET("/students", server.listStudents)
		authRoutes.DELETE("/students/:id", server.deleteStudent)

		// ! scores
		authRoutes.POST("/scores", server.createScore)
		authRoutes.GET("/scores/:id", server.getScoreByStudentID)
		authRoutes.DELETE("/scores/:id", server.deleteScore)

		// ! term_scores
		authRoutes.POST("/term_scores", server.createTermScore)
		authRoutes.GET("/term_scores/:id", server.getTermScoreByID)
		authRoutes.GET("/term_scores", server.listTermScoresForSubjectAndClass)
		authRoutes.PATCH("/term_scores/:id", server.updateTermScoreByID)
		authRoutes.DELETE("/term_scores/:id", server.deleteTermScoreByID)

	}

	server.router = router
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
