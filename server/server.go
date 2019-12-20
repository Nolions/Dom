package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"simpleChaceService/cache"
	"time"
)

var cacheHandler cache.Cache

func New(r *gin.Engine, addr string) *http.Server {
	log.Printf("Listening on http://localhost%s", addr)
	return &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func NewCacheHandler() {
	cacheHandler = cache.New()
}

func Run(s *http.Server) {
	if err := s.ListenAndServe(); err != nil {
		log.Println(err)
	}
}