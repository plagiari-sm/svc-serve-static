package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
	psmconfig "github.com/plagiari-sm/psm-config"
)

// APP ...
type APP struct {
	Router *gin.Engine
	Server *http.Server
}

// Initialize ...
func (a *APP) Initialize() {
	gin.SetMode(gin.ReleaseMode)

	cfg := psmconfig.Config
	// Assigin the router
	a.Router = gin.Default()
	// Inital router settings
	a.Router.RedirectTrailingSlash = true
	a.Router.RedirectFixedPath = true

	a.Router.Static("/"+cfg.HTMLPaths.Static, cfg.HTMLPaths.Serve+"/static")
	a.Router.LoadHTMLGlob(cfg.HTMLPaths.Serve + "/index.html")

	if cfg.Hash != "" {
		a.Router.Use(jwt.Auth(cfg.Hash))
	}

	a.Router.Use(static.Serve("/", static.LocalFile(cfg.HTMLPaths.Serve, true)))

	// This is usefull when you combine multiple microservices
	a.Router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// Assigin the http server
	a.Server = &http.Server{
		Addr:    ":8000",
		Handler: a.Router,
	}
}

// Serve ...
func (a *APP) Serve() {
	errChan := make(chan error, 10)

	go func() {
		errChan <- a.Server.ListenAndServe()
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.Fatal(err)
			}
		case s := <-signalChan:
			log.Println(fmt.Sprintf("Captured message %v. Exiting...", s))
			os.Exit(0)
		}
	}
}
