package scan

import "goscanner/utils"

type Md5Encoder struct {
	filePath string
	md5Code  string
}

func Md5EncoderInit(filePath string) *Md5Encoder {
	return &Md5Encoder{filePath: filePath}
}

func (m *Md5Encoder) GetMd5Code() string {
	return m.md5Code
}

func (m *Md5Encoder) Md5Encode() error {
	m_result, err := utils.FileMD5(m.filePath)
	if err != nil {
		return err
	}
	m.md5Code = m_result
	return nil
}
