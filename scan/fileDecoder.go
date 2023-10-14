package scan

import (
	"bufio"
	"goscanner/utils"
	"os"
)

type FileDecoder struct {
	FilePath string
	FOutPath string //加密后的文件路径
	Key      []byte
}

func FileDecoderInit(filePath string, key []byte, fout string) *FileDecoder {
	return &FileDecoder{
		FilePath: filePath,
		Key:      key,
		FOutPath: fout,
	}
}

func (fe *FileDecoder) Decrypt() error {
	f, err := os.Open(fe.FilePath)
	if err != nil {
		return err
	}
	defer f.Close()
	br := bufio.NewReader(f)
	ff, err := os.OpenFile(fe.FOutPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer ff.Close()
	num := 0
	//逐行读密文，逐行解密
	for {
		num = num + 1
		a, err := br.ReadString('\n')
		if err != nil {
			break
		}
		getByte, err := utils.DecryptByAes(a, fe.Key)
		if err != nil {
			return err
		}

		buf := bufio.NewWriter(ff)
		buf.Write(getByte)
		buf.Flush()
	}
	return nil
}
