package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/osobotu/school_mgs/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// routes will be added here
	// if you pass in multiple routers, the last one should be the http router
	// while the others should be middlewares

	v1 := router.Group("/v1")
	{
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
		v1.PATCH("/classes/:id", server.updateFormMaster)
		v1.GET("/classes", server.listClasses)
		v1.DELETE("/classes/:id", server.deleteClassByID)

		// ! sessions
		v1.POST("/sessions", server.createSession)
	}

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
