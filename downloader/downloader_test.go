package downloader

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func checkFilesize(filepath string) int64 {
	file, err := os.Open(filepath)
	if err != nil {
		return -1
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return -1
	}
	return fi.Size()
}

func TestDownloadMission(t *testing.T) {
	Convey("Init downloader", t, func() {
		var testURL = "https://ss0.bdstatic.com/5aV1bjqh_Q23odCf/static/superman/img/logo/bd_logo1_31bdc765.png"
		var savePath = "/tmp/ccc"
		Convey("download from a valid url", func() {
			mission := newMission(testURL, savePath)
			err := mission.start()
			So(err, ShouldBeNil)

			Convey("check file size", func() {
				size := checkFilesize(savePath)
				So(size, ShouldBeGreaterThan, 0)
			})
		})

		Convey("redownload from a invalid url and save it to a same file", func() {
			var invalidURL = "https://pokemongocannotployinchina.com/xx.img"
			mission := newMission(invalidURL, savePath)
			err := mission.start()
			So(err, ShouldNotBeNil)

			Convey("check file size", func() {
				size := checkFilesize(savePath)
				So(size, ShouldBeGreaterThan, 0)
			})
		})

	})
}
