package gohttp

import (
	"github.com/gin-gonic/gin"
	"os"
)

// Default gin.New
func New() *gin.Engine {
	setProductionMode()

	return gin.New()
}

// Default gin.Default
func Default() *gin.Engine {
	setProductionMode()

	return gin.Default()
}

// setProductionMode production mode
func setProductionMode() {
	if os.Getenv("AppEnvironment") == "production" {
		SetMode(gin.ReleaseMode)
	}
}

// SetMode gin.SetMode
func SetMode(ginMode string) {
	gin.SetMode(ginMode)
}

// Use gin.Engine.Use
func Use(engine *gin.Engine, middleware ...gin.HandlerFunc) {
	engine.Use(middleware...)
}

// Run gin.Engine.Run
func Run(engine *gin.Engine, addr ...string) error {
	return engine.Run(addr...)
}

// RunTLS gin.Engine.RunTLS
func RunTLS(engine *gin.Engine, addr, certFile, keyFile string) error {
	return engine.RunTLS(addr, certFile, keyFile)
}

// RunUnix gin.Engine.RunUnix
func RunUnix(engine *gin.Engine, file string) error {
	return engine.RunUnix(file)
}
