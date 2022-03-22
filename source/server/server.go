package server

import (
	Bot "bot"
	"job"
	"job/message"
	"time"
)

func Serve() {
	//router := gin.Default()

	//router.GET("/", func(c *gin.Context) {
	//	c.Status(http.StatusOK)
	//})

	messageJob := message.NewConsole(time.Now(), "Ola'!!!")
	bot := Bot.New([]*job.Job{&messageJob})

	bot.Start()

	//router.Run(":8080")
}
