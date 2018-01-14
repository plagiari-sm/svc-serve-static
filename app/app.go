package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/plagiari-sm/svc-serve-static/conf"
)

// APP ...
type APP struct {
	Router *gin.Engine
	Server *http.Server
}

// Initialize ...
func (a *APP) Initialize() {
	cfg := conf.Configuration
	// Assigin the router
	a.Router = gin.Default()
	// Inital router settings
	a.Router.RedirectTrailingSlash = true
	a.Router.RedirectFixedPath = true

	a.Router.Static("/static", cfg.StaticPath+"/static")
	a.Router.LoadHTMLGlob(cfg.StaticPath + "/index.html")
	a.Router.Use(static.Serve("/", static.LocalFile(cfg.StaticPath, false)))
	/*
		a.Router.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{})
		})
	*/
	// Forbid Access
	// This is usefull when you combine multiple microservices
	a.Router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// Assigin the http server
	a.Server = &http.Server{
		Addr:    cfg.Server.Host + ":" + strconv.Itoa(cfg.Server.Port),
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
