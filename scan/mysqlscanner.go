package scan

type Mysqlscanner struct {
	username string
	pwd      string
	dbname   string
	metadata []map[string]string
}

func MysqlscannerInit(username string, pwd string, dbname string) *Mysqlscanner {
	return &Mysqlscanner{username, pwd, dbname, make([]map[string]string, 0)}
}
