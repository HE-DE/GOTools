package scan

import "goscanner/utils"

type VideoChanger struct {
	inputVideoPath string
	ouputDir       string
	fps            int
}

func InitVideoChanger(inputVideoPath string, ouputDir string, fps int) *VideoChanger {
	return &VideoChanger{inputVideoPath, ouputDir, fps}
}

func (vc *VideoChanger) Changefps() string {
	return utils.FormatFrameRate(vc.inputVideoPath, vc.ouputDir, vc.fps)
}
