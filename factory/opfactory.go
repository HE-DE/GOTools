package factory

// ProcessorFactory 工厂函数创建相应的策略对象
func ProcessorFactory(command string) ProcessorStrategy {
	switch command {
	case "sc":
		return &ScanProcessor{}
	case "ps":
		return &PingScanProcessor{}
	case "ms":
		return &MysqlScanProcessor{}
	case "fe":
		return &FileEncryptProcessor{}
	case "fd":
		return &FileDecryptProcessor{}
	case "tg":
		return &TarGzipProcessor{}
	case "ug":
		return &UnTarGzipProcessor{}
	case "md5":
		return &Md5Processor{}
	case "sys":
		return &SystemInfoProcessor{}
	case "is":
		return &ImageStitcherProcessor{}
	case "vc":
		return &VideoFpsChangeProcessor{}
	default:
		return nil
	}
}
