package scan

import "goscanner/utils"

type TarGziper struct {
	FileInPath  string
	FileOutPath string
}

type UnTarGziper struct {
	TarFileInPath  string
	TarFileOutPath string
}

func TarGziperInit(filein, fileout string) *TarGziper {
	return &TarGziper{filein, fileout}
}

func UnTarGziperInit(filein, fileout string) *UnTarGziper {
	return &UnTarGziper{filein, fileout}
}

// TarGziper tar the file
func (t *UnTarGziper) UnTar() error {
	return utils.UnTar(t.TarFileOutPath, t.TarFileInPath)
}

func (t *TarGziper) Tar() error {
	return utils.Tar(t.FileInPath, t.FileOutPath)
}
