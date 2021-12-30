package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stianeikeland/go-rpio/v4"
)

type Sender struct {
	Login string `json:"login"`
}

type Output struct {
	Title string `json:"title"`
}

type CheckRun struct {
	Status string `json:"status"`
	Output Output `json:"output"`
}

type CheckRunPayload struct {
	Action   string   `json:"action"`
	CheckRun CheckRun `json:"check_run"`
	Sender   Sender   `json:"sender"`
}

func buzz() {
	pin := rpio.Pin(17)
	pin.Output()

	sum := 1
	for sum < 500 {
		pin.Toggle()
		time.Sleep(time.Second)
		sum += sum
	}

	pin.High()

}

func main() {
	username := "drewalth"             // change to whatever GitHub username
	buildErrorTitle := "Build Errored" // change to what your CI titles failed builds

	err := rpio.Open()

	if err != nil {
		fmt.Println(err)
	}

	r := gin.Default()

	r.StaticFile("/", "./public/index.html")
	r.StaticFile("/pi.png", "./public/pi.png")

	r.GET("/test", func(c *gin.Context) {
		go buzz()

		c.JSON(200, gin.H{
			"message": "buzzing",
		})
	})

	r.POST("/check-run", func(c *gin.Context) {
		res := CheckRunPayload{}

		c.BindJSON(&res)

		userNameMatches := res.Sender.Login == username
		errorTitleMatches := res.CheckRun.Output.Title == buildErrorTitle
		statusComplete := res.CheckRun.Status == "completed"

		if userNameMatches && errorTitleMatches && statusComplete {
			go buzz()
			c.JSON(200, gin.H{
				"message": "buzzing",
			})
		} else {
			c.JSON(200, gin.H{
				"message": "safe for now",
			})
		}
	})
	r.Run(":5000")
}
