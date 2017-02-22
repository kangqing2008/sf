package dm

import (
	"database/sql"
	"strconv"
	"kangqing2008/gotools/tools/file"
)

const(
	USERNAME 	= "root"
	PASSWORD 	= "power2000"
	DBNAME		= "stocks"
	PROTOCOL	= "tcp"
	HOST		= "localhost"
	PORT		= 3306
	DRIVER		= "mysql"
)

func ImportFile(filename string){
	if exists,finfo := file.Exists(filename); !exists || finfo == nil{
		panic("文件名不存在:" + filename)
	}

}

func OpenDatabase()*sql.DB{
	db,_ := sql.Open(DRIVER,USERNAME + ":" + PASSWORD + "@" + PROTOCOL + "(" + HOST + ":" + strconv.Itoa(PORT) + ")/" + DBNAME + "?charset=utf8")
	return db
}