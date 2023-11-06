package utils

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func Informat(target string, str_array []string) bool {
	for _, v := range str_array {
		if v == target {
			return true
		}
	}
	return false
}

// 转化视频帧率
func FormatFrameRate(inputVideoPath, outputDir string, fps int) string {
	_formatArr := []string{"mp4", "flv"}
	_, _file := filepath.Split(inputVideoPath)
	_tmps := strings.Split(_file, ".")
	_ext := _tmps[len(_tmps)-1]
	if !Informat(_ext, _formatArr) {
		return "格式不支持！"
	}
	_name := uuid.New()
	_resultVideoPath := filepath.Join(outputDir, fmt.Sprintf("%s.%s", _name.String(), _ext))
	err := ffmpeg.Input(inputVideoPath).
		Output(_resultVideoPath, ffmpeg.KwArgs{"qscale": 0, "r": fps}).
		OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		return err.Error()
	}
	return _resultVideoPath
}
