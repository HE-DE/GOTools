package scan

import (
	"bufio"
	"fmt"
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
		fmt.Println("未找到文件")
		return err
	}
	defer f.Close()
	fInfo, _ := f.Stat()
	fmt.Println("待处理文件大小:", fInfo.Size())

	br := bufio.NewReader(f)
	ff, err := os.OpenFile(fe.FOutPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件写入错误")
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
			fmt.Println("解密错误")
			return err
		}

		buf := bufio.NewWriter(ff)
		buf.Write(getByte)
		buf.Flush()
	}
	fmt.Println("解密次数：", num)
	ffInfo, _ := ff.Stat()
	fmt.Printf("文件解密成功，生成文件名为：%s 文件大小为：%v Byte \n", ffInfo.Name(), ffInfo.Size())
	return nil
}
