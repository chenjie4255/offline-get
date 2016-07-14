package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/chenjie4255/offline-get/downloader"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAddDownloadMission(t *testing.T) {
	Convey("Init test env", t, func() {
		dlMgr := downloader.NewDownloadMgr("/tmp/")
		engine := createMainEngine(dlMgr)
		var testURL = "https://ss0.bdstatic.com/5aV1bjqh_Q23odCf/static/superman/img/logo/bd_logo1_31bdc765.png"

		type result struct {
			Error      string
			Mission_Id int64 // json的key是mission_id
		}
		r := &result{}

		Convey("test add mission url", func() {
			v := url.Values{}
			v.Set("url", testURL)
			body := ioutil.NopCloser(strings.NewReader(v.Encode()))
			req, _ := http.NewRequest("POST", "http://localhost/mission", body)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			w := httptest.NewRecorder()
			engine.ServeHTTP(w, req)

			Convey("should return a mission id with no error", func() {
				So(w.Code, ShouldEqual, http.StatusOK)
				err := json.Unmarshal(w.Body.Bytes(), r)
				So(err, ShouldBeNil)
				So(r.Error, ShouldBeEmpty)
				So(r.Mission_Id, ShouldBeGreaterThan, 0)
			})
		})

		Convey("test add a invalid url mission", func() {
			v := url.Values{}
			v.Set("url", testURL+"sdfsdfsdf")
			body := ioutil.NopCloser(strings.NewReader(v.Encode()))
			req, _ := http.NewRequest("POST", "http://localhost/mission", body)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			w := httptest.NewRecorder()
			engine.ServeHTTP(w, req)

			Convey("should return a mission id with no error", func() {
				So(w.Code, ShouldEqual, http.StatusOK)
				err := json.Unmarshal(w.Body.Bytes(), r)
				So(err, ShouldBeNil)
				So(r.Error, ShouldBeEmpty)
				So(r.Mission_Id, ShouldBeGreaterThan, 0)
			})
		})
	})

}
