package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Init() {
	fmt.Println("init done.")
}

func main() {
	fmt.Println("start..")
	// r := gin.New()
	//0.0.0.0:8080/
	r := gin.Default()
	r.Run()
}

