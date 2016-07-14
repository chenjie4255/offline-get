package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/chenjie4255/offline-get/downloader"
)

func createMainEngine(mgr downloader.DownloadMgr) *gin.Engine {
	engine := gin.Default()
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

		missionID, err := mgr.AddMission(url)
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

	return engine
}

func main() {
	dlMgr := downloader.NewDownloadMgr("/tmp/")
	engine := createMainEngine(dlMgr)
	engine.Run(":5821")
}
