package scan

import (
	"bufio"
	"fmt"
	"goscanner/utils"
	"os"
)

type FileEncoder struct {
	FilePath string
	FOutPath string //加密后的文件路径
	Key      []byte
}

func FileEncoderInit(filePath string, key []byte, fout string) *FileEncoder {
	return &FileEncoder{
		FilePath: filePath,
		Key:      key,
		FOutPath: fout,
	}
}

func (fe *FileEncoder) Encrypt() error {
	f, err := os.Open(fe.FilePath)
	if err != nil {
		return err
	}
	defer f.Close()
	fInfo, _ := f.Stat()
	fLen := fInfo.Size()
	fmt.Println("待处理文件大小:", fLen)
	maxLen := 1024 * 100 //100Kb  每 100Kb 加密一次
	var forNum int64 = 0
	getLen := fLen

	if fLen > int64(maxLen) {
		getLen = int64(maxLen)
		forNum = fLen / int64(maxLen)
		fmt.Println("需要加密次数：", forNum+1)
	}

	// 加密后的存储文件
	ff, err := os.OpenFile(fe.FOutPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("写入失败")
		return err
	}
	defer ff.Close()

	// 循环加密
	for i := 0; i < int(forNum+1); i++ {
		a := make([]byte, getLen)
		n, err := f.Read(a)
		if err != nil {
			fmt.Println("文件读取失败")
			return err
		}
		getByte, err := utils.EncryptByAes(a[:n], fe.Key)
		if err != nil {
			fmt.Println("加密错误")
			return err
		}
		//换行处理
		getBytes := append([]byte(getByte), []byte("\n")...)
		//写入
		buf := bufio.NewWriter(ff)
		buf.WriteString(string(getBytes[:]))
		buf.Flush()
	}
	ffInfo, _ := ff.Stat()
	fmt.Printf("文件加密成功，生成文件名为：%s 文件大小为：%v Byte \n", ffInfo.Name(), ffInfo.Size())
	return nil
}
