package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/chenjie4255/offline-get/downloader"
)

func main() {
	engine := gin.Default()

	dlMgr := downloader.NewDownloadMgr("/tmp/")

	engine.POST("/mission", func(c *gin.Context) {
		c.Request.ParseForm()
		url := c.Request.PostForm.Get("url")
		ret := make(map[string]interface{})
		if url == "" {
			ret["error"] = "invalid parameter"
			ret["mission_id"] = 0
			c.JSON(http.StatusOK, ret)
			return
		}

		missionID, err := dlMgr.AddMission(url)
		if err != nil {
			ret["error"] = err.Error()
			ret["mission_id"] = 0
		} else {
			ret["error"] = ""
			ret["mission_id"] = missionID
		}

		c.JSON(http.StatusOK, ret)
		return
	})

	engine.Static("/offline_file/", "/tmp/")

	engine.Run(":5821")
}
