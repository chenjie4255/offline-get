package downloader

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDownload(t *testing.T) {
	Convey("Init downloadmgr env..", t, func() {
		var savePath = "/tmp"
		mgr := NewDownloadMgr(savePath)
		var testURL = "https://ss0.bdstatic.com/5aV1bjqh_Q23odCf/static/superman/img/logo/bd_logo1_31bdc765.png"

		Convey("test download..", func() {
			missionID, err := mgr.AddMission(testURL)

			Convey("mission should be added succeed", func() {
				So(err, ShouldBeNil)

				Convey("download should be finished in 10 seconds...", func() {
					for i := 0; ; i++ {
						percent, err := mgr.CheckProgress(missionID)
						So(err, ShouldBeNil)
						So(percent, ShouldBeBetweenOrEqual, 0, 100)
						if percent < 100 {
							time.Sleep(1 * time.Second)
							So(i, ShouldBeLessThan, 10)
						} else {
							break
						}
					}
				})

				filename, err := mgr.GetMissionFilename(missionID)
				So(err, ShouldBeNil)
				So(filename, ShouldEqual, "/tmp/bd_logo1_31bdc765.png")
			})
		})
	})

}
