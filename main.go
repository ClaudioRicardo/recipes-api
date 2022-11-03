package main

import (
	"time"

	"github.com/gin-conic/gin"
)

type Recipe struct {
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingradients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructiuons"`
	PublishedAt  time.Time `json:"publishedAt"`
}

func main() {
	router := gin.Deafult()
	router.Run()
}
