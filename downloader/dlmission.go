package downloader

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/nu7hatch/gouuid"
)

type mission struct {
	url      string
	filePath string
}

func newMission(url, filePath string) *mission {
	return &mission{url, filePath}
}

func download(url, savePath string) error {
	cmd := exec.Command("wget", url, "-O", savePath)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	in := bufio.NewScanner(stdout)
	for in.Scan() {
		line := in.Text()
		fmt.Println(line)
	}
	if err := in.Err(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}

func (m *mission) start() error {
	id, _ := uuid.NewV4()
	tempPath := "/tmp/" + id.String()

	err := download(m.url, tempPath)
	if err != nil {
		return err
	}

	return moveFile(tempPath, m.filePath)
}

func moveFile(oldPath, newPath string) error {
	return os.Rename(oldPath, newPath)
}
