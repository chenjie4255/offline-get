package downloader

import (
	"strings"
	"sync/atomic"
	"time"
)

type downloadInfo struct {
	filename  string
	dlMission *mission
}

type downloadMgr struct {
	savePath string
	fileMap  map[int64]*downloadInfo
	maxID    int64
}

func NewDownloadMgr(savePath string) DownloadMgr {
	mgr := &downloadMgr{savePath, make(map[int64]*downloadInfo), 0}
	return mgr
}

func (d *downloadMgr) AddMission(url string) (int64, error) {
	urlInfo := strings.Split(url, "/")
	filename := d.savePath + "/" + urlInfo[len(urlInfo)-1]
	mission := newMission(url, filename)
	mission.start()

	dlInfo := &downloadInfo{filename, mission}
	id := atomic.AddInt64(&d.maxID, 1)
	d.fileMap[id] = dlInfo
	d.maxID++

	return id, nil
}

func (d *downloadMgr) CheckProgress(ID int64) (int, error) {
	dlInfo, found := d.fileMap[ID]
	if !found {
		return 0, ErrInvalidMissionID
	}

	if dlInfo.dlMission.err != nil {
		return 0, dlInfo.dlMission.err
	}

	select {
	case <-dlInfo.dlMission.end:
		return 100, nil
	case <-time.After(1 * time.Microsecond):
		return 1, nil
	}
	return 1, nil
}

func (d *downloadMgr) GetMissionFilename(ID int64) (string, error) {
	dlInfo, found := d.fileMap[ID]
	if !found {
		return "", ErrInvalidMissionID
	}

	return dlInfo.filename, nil

}
