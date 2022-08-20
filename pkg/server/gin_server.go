package server

import (
	"calculator/internal/adapters/api/handler"
	"calculator/internal/core/ports"
	"calculator/internal/core/services"
	"calculator/pkg/router"
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type ginServer struct {
	log    ports.Logger
	router *router.Router
}

func newGinServer(l ports.Logger, r *router.Router) *ginServer {
	return &ginServer{
		log:    l,
		router: r,
	}
}

func (s *ginServer) setAppHandlers(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	service := services.NewDivisionService()
	handler := handler.NewDivsionHandler(service, s.log)

	v1.POST("/division", handler.Division)
}

func (s *ginServer) setupRouter() *gin.Engine {
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "test" {
		r := gin.New()
		s.setAppHandlers(r)
		return r
	}

	r := gin.New()
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())
	// setup cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))
	s.setAppHandlers(r)
	return r
}

func (s *ginServer) Run() {
	gin.Recovery()

	r := s.setupRouter()
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "5550"
	}

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		Addr:         fmt.Sprintf(":%v", port),
		Handler:      r,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		s.log.WithFields(ports.Fields{"port": port}).Infof("Starting HTTP Server")
		if err := server.ListenAndServe(); err != nil {
			s.log.WithError(err).Fatalln("Error starting HTTP server")
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		s.log.WithError(err).Fatalln("Server Shutdown Failed")
	}

	s.log.Infof("Service down")
}
