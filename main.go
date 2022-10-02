package main

import (
	"fmt"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		command := "./cmd.sh"
		cmd := exec.Command("/bin/bash", "-c", command)
		output, err := cmd.Output()
		if err != nil {
			fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
			return
		}
		fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
	})

	r.GET("/ping", func(c *gin.Context) {
		fmt.Println(c.Request)
		c.String(200, "pong")
	})

	r.Run()
}
