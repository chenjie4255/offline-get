package downloader

type DownloadMgr interface {
	AddMission(url string) (int64, error)

	CheckProgress(ID int64) (int, error)

	GetMissionFilename(ID int64) (string, error)
}
