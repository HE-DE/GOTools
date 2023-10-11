package scan

import (
	"database/sql"
	"goscanner/utils"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type Field struct {
	fieldName string
	fieldDesc string
	dataType  string
	isNull    string
	length    int
}

type Mysqlscanner struct {
	username string
	pwd      string
	dbname   string
	metadata []map[string][]Field
}

var DB *sql.DB

func MysqlscannerInit(username string, pwd string, dbname string) *Mysqlscanner {
	err := error(nil)
	DB, err = sql.Open("mysql", username+":"+pwd+"@tcp(127.0.0.1:3306)/"+dbname)
	if err != nil {
		panic(err)
	}
	return &Mysqlscanner{username, pwd, dbname, make([]map[string][]Field, 0)}
}

func (ms *Mysqlscanner) GetTables() []string {
	var tables []string
	rows, err := DB.Query("show tables")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		if err != nil {
			panic(err)
		}
		tables = append(tables, tableName)
	}
	rows.Close()
	return tables
}

func (ms *Mysqlscanner) Scan(tableName string) []Field {
	var fields []Field
	sqlStr := `SELECT COLUMN_NAME fName,column_comment fDesc,DATA_TYPE dataType,
						IS_NULLABLE isNull,IFNULL(CHARACTER_MAXIMUM_LENGTH,0) sLength
			FROM information_schema.columns 
			WHERE table_schema = ? AND table_name = ?`
	rows, err := DB.Query(sqlStr, ms.dbname, tableName)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var f Field
		err = rows.Scan(&f.fieldName, &f.fieldDesc, &f.dataType, &f.isNull, &f.length)
		if err != nil {
			panic(err)
		}
		fields = append(fields, f)
	}
	rows.Close()
	return fields
}

func (ms *Mysqlscanner) ScanAll() {
	var wg sync.WaitGroup
	tables := ms.GetTables()
	bar := utils.CreateBar(len(tables))
	for _, tableName := range tables {
		wg.Add(1)
		go func(tablename string) {
			defer wg.Done()
			bar.Add(1)
			ms.metadata = append(ms.metadata, map[string][]Field{tablename: ms.Scan(tablename)})
		}(tableName)
	}
	wg.Wait()
	defer DB.Close()
}

func (ms *Mysqlscanner) GetMetadata() []map[string][]Field {
	return ms.metadata
}
