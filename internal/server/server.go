package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"os"
	"os/signal"
	"report_hn/internal/db"
	"report_hn/internal/logger"
	"syscall"
	"time"
)

func ApiServer() {
	r := gin.Default()

	psql := db.GetDB()
	r.POST("/login", func(c *gin.Context) { AuthLogin(psql, c) })

	protected := r.Group("/")
	protected.Use(AuthMiddleware)
	protected.Use(RateLimiter(rate.Every(time.Minute/30), 30))
	protected.POST("/reports", func(c *gin.Context) { CreateReport(psql, c) })
	protected.GET("/reports", func(c *gin.Context) { GetReports(psql, c) })

	srv := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Fatal("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		logger.Log.Println("timeout of 5 seconds.")
	}
	logger.Log.Println("Server exiting")
}
