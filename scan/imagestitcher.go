package scan

import "goscanner/utils"

type ImageStitcher struct {
	file1path string
	file2path string
	outpath   string
	MaxWidth  float64
}

func ImageStitcherInit(f1, f2, out string, MaxWidth float64) *ImageStitcher {
	return &ImageStitcher{file1path: f1, file2path: f2, outpath: out, MaxWidth: MaxWidth}
}

func (i *ImageStitcher) Stitch() error {
	err := utils.Stitching(i.file1path, i.file2path, i.outpath, i.MaxWidth)
	return err
}
