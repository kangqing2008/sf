package dm

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"kangqing2008/gotools/tools/file"
	"fmt"
	"time"
	"strings"
	"os"
	"kangqing2008/sf/dti"
)

const(
	USERNAME 	= "root"
	PASSWORD 	= "powerdata"
	DBNAME		= "stocks"
	PROTOCOL	= "tcp"
	HOST		= "localhost"
	PORT		= 3306
	DRIVER		= "mysql"


	MA5 	= 5
	MA10 	= 10
	MA20	= 20
	MA30	= 30
	MA60	= 60
	MA120	= 120
	MA250	= 250

)

type DLine struct{
	Market		string
	Code		string
	Intcode		int64
	Day			time.Time
	Open		float64
	Close		float64
	Max			float64
	Min			float64
	Vol			int64
	Tno			float64
	Uad			float64
	Aiad		float64
	Am			float64
	Tr			float64
	Ma5			float64
	Ma10		float64
	Ma20		float64
	Ma30		float64
	Ma60		float64
	Ma120		float64
	Ma250		float64
}


// 解析文件并将数据解析后导入数据库
func ImportFile(filename string)(int64,error){

	// 通过解析文件全路径，得到文件名称，股票市场和股票代码
	name,_ := file.GetFileName(filename)
	market := file.Substr(name,0,2)
	code := file.SubstrToEnd(name,2)

	lines,err := readFileContent(filename)
	if err != nil{
		panic("读取文件出错：" + filename)
	}
	datas := parseDLines(market,code,lines)
	//fmt.Println(len(datas))
	//return -1,nil
	return insertDLines(datas)

}

func readFileContent(filename string) ([]string,error){
	// 判断文件是否存在
	if exists,finfo := file.Exists(filename); !exists || finfo == nil{
		panic("文件名不存在:" + filename)
	}

	// 读取文件内容到内存中，每行一个字符串
	return file.ReadAllLines(filename)
}

// 保存日线数据到数据库中
func insertDLines(datas []DLine)(int64,error){
	db := OpenDatabase()
	defer db.Close()
	tx , err := db.Begin()
	if err != nil{
		panic("启动事务时出错!")
	}
	defer tx.Rollback()
	stmt,err := tx.Prepare(` INSERT INTO D(MARKET,CODE,INTCODE,DAY,OPEN,CLOSE,MAX,MIN,VOL,TNO,UAD,AIAD,AM,TR)VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?) `)
	if err != nil{
		panic("解析插入日线数据的sql时出错!")
	}
	defer stmt.Close()
	var count int64 = 0
	for _,d := range datas{
		//fmt.Println("插入第",i,"条数据：",d)
		res,err := stmt.Exec(d.Market,d.Code,d.Intcode,d.Day,d.Open,d.Close,d.Max,d.Min,d.Vol,d.Tno,d.Uad,d.Aiad,d.Am,d.Tr)
		if err != nil{
			return -1,err
		}else{
			ra,_ := res.RowsAffected()
			count += ra
		}
	}
	tx.Commit()
	return count,nil
}

// 将数据文件中的一行解析成一个日线数据
func parseDLines(market,code string ,lines []string)[]DLine{
	// 用于暂存数据的
	var datas []DLine
	// 从第一行开始解析
	for _,line := range lines[1:]{
		//fmt.Println("第",i,"行:",line)
		var data = DLine{}
		data.Market 	= market
		data.Code 		= code
		data.Intcode,_	= strconv.ParseInt(code,10,64)
		array := strings.Split(line,",")
		data.Day,_ 		= time.Parse("2006-01-02",array[0])
		data.Open,_		= strconv.ParseFloat(array[1],64)
		data.Close,_	= strconv.ParseFloat(array[2],64)
		data.Max,_		= strconv.ParseFloat(array[3],64)
		data.Min,_		= strconv.ParseFloat(array[4],64)
		data.Vol,_		= strconv.ParseInt(array[5],10,64)
		data.Tno,_		= strconv.ParseFloat(array[6],64)
		data.Uad,_		= strconv.ParseFloat(array[7],64)
		data.Aiad,_		= strconv.ParseFloat(array[8],64)
		data.Am,_		= strconv.ParseFloat(array[9],64)
		data.Tr,_		= strconv.ParseFloat(array[10],64)
		datas = append(datas,data)
		//fmt.Println("data[",i,"]:",data)
	}
	return datas
}

func ImportDirectory(path string){
	files,err := file.ListFiles(path)
	if err != nil{
		fmt.Println("罗列所有文件时出错：",err)
		panic(err)
	}

	for i,file := range files{
		fullname := path + "\\" + file
		if strings.HasSuffix(file,".ok"){
			fmt.Println("解析第",i+1,"个文件:",file,"后缀为.ok,已经入库!")
		}else{
			count,er := ImportFile(fullname)
			if er == nil {
				fmt.Println("解析第",i+1,"个文件",file,",入库数据",count,"行!")
				os.Rename(fullname,fullname + ".ok")
			}else{
				fmt.Println("解析第",i+1,"个文件",file,",出错：",er)
			}
		}

	}
}

func GetStockStat()([]string,[]string,[]int64){
	db := OpenDatabase()
	defer db.Close()
	// WHERE CODE = '600000'
	//rows ,err := db.Query(" SELECT MARKET,CODE,COUNT(1) SIZE FROM D GROUP BY MARKET,CODE ORDER BY MARKET DESC,CODE ASC ")
	rows,err := db.Query(`SELECT MARKET, CODE,COUNT(1) SIZE,SUM(CASE WHEN MA5 > 0.001 THEN 1 ELSE 0 END) MA5C
				FROM D
				GROUP BY CODE
				HAVING (SIZE - MA5C) > 5
				ORDER BY CODE ASC`)

	if err != nil {
		fmt.Println("统计股票日线数据量的时候出错：",err)
		panic(err)
	}
	var markets,codes []string
	var sizes []int64
	defer rows.Close()
	for rows.Next(){
		var market ,code string
		var size,ma5c int64
		err := rows.Scan(&market,&code,&size,&ma5c)
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println(market,code,size,ma5c)
		markets = append(markets,market)
		codes	= append(codes,code)
		sizes	= append(sizes,size)
	}
	if rows.Err() != nil {
		fmt.Println(rows.Err())
	}
	return markets,codes,sizes
}

func CalcDayStat(){
	fmt.Println("开始统计每只股票的数据量!",time.Now())

	markets,codes,sizes := GetStockStat()
	fmt.Println("统计完毕，开始逐个股票分析")
	for i,market := range markets{
		fmt.Println("开始分析：",market,codes[i],sizes[i])
		dlines := GetStockDLines(market,codes[i])
		results := CalcDLineStat(dlines,true,true,true,true,true,true,true)
		fmt.Println("开始分析：",market,codes[i],sizes[i],"完毕!",len(results))
		//for j,line := range results{
		//	fmt.Println("line:",j+1,"day:",line.Day,"close:",line.Close,":ma5:",line.Ma5 ,":ma10:",line.Ma10,":ma20:",line.Ma20,":ma30:",line.Ma30,":ma60:",line.Ma60,":ma120:",line.Ma120,":ma250:",line.Ma250)
		//}
		fmt.Println("开始更新",codes[i])
		fmt.Println("更新",codes[i],"完成，更新行数：",UpdateDLineMa(results))
	}
	fmt.Println("全部完成!",time.Now())
}

func UpdateDLineMa(results []DLine)int64{
	db := OpenDatabase()
	defer db.Close()
	tx,err := db.Begin()
	if err != nil{
		fmt.Println("更新ma数据时发生错误：",err)
		panic(err)
	}
	defer tx.Rollback()
	stmt,err := tx.Prepare(" UPDATE D SET MA5 = ?,MA10 = ?, MA20 = ?,MA30 = ?,MA60 = ?,MA120 = ?,MA250 = ? WHERE CODE = ? AND DAY = ? ")
	if err != nil{
		fmt.Println("编译ma数据更新语句时发生错误：",err)
		panic(err)
	}
	defer stmt.Close()
	var count int64 = 0
	for _,m := range results{
		r,err := stmt.Exec(m.Ma5,m.Ma10,m.Ma20,m.Ma30,m.Ma60,m.Ma120,m.Ma250,m.Code,m.Day)
		if err != nil {
			fmt.Println("执行ma数据更新语句时发生错误：",err)
			panic(err)
		}
		ra,_ := r.RowsAffected()
		count += ra
	}
	tx.Commit()
	return count
}

func CalcDLineStat(dlines []DLine,ma5,ma10,ma20,ma30,ma60,ma120,ma250 bool)[]DLine{
	//缓存计算结果,carray负责缓存前一个ma值计算时
	//cpre负责缓存前一个ma
	var precount map[int64]float64	= map[int64]float64{}
	results := make([]DLine,len(dlines))
	for i,line := range dlines{
		//当前数量
		j := i + 1
		if ma5 {
			line.Ma5,precount[MA5] = calcAndCache(j,MA5,precount[MA5],line,dlines)
		}
		if ma10 {
			line.Ma10,precount[MA10] = calcAndCache(j,MA10,precount[MA10],line,dlines)
		}
		if ma20 {
			line.Ma20,precount[MA20] = calcAndCache(j,MA20,precount[MA20],line,dlines)
		}
		if ma30 {
			line.Ma30,precount[MA30] = calcAndCache(j,MA30,precount[MA30],line,dlines)
		}
		if ma60 {
			line.Ma60,precount[MA60] = calcAndCache(j,MA60,precount[MA60],line,dlines)
		}
		if ma120 {
			line.Ma120,precount[MA120] = calcAndCache(j,MA120,precount[MA120],line,dlines)
		}
		if ma250 {
			line.Ma250,precount[MA250] = calcAndCache(j,MA250,precount[MA250],line,dlines)
		}
		results[i] = line
	}
	return results
}

func calcAndCache(j,ma int,precount float64,line DLine,lines []DLine)(float64,float64){
	var maresult,preresult float64
	if j > ma {
		maresult = (precount + line.Close)/float64(ma)
		preresult = precount + line.Close - lines[j-ma].Close
	}else if j == ma {
		maresult,preresult = CalcMa(lines[j-ma:j],ma)
	}
	return maresult,preresult
}

// 计算给定的日线列表dlines中，从后往前倒退malength个数据的均值
// 返回计算出来的均值和数据求和后减去第一个数据的值
// 比如dlines的长度为10,malength=5,本函数会计算dlines[5:]中Close值的均值作为第一个参数返回。
// 并将dlines[6:]中所有close的值求和后作为第二个参数返回
func CalcMa(dlines []DLine,malength int)(float64,float64){
	var count,precount float64
	for _,line := range dlines[len(dlines)-malength:]{
		count += line.Close
	}
	precount = count - dlines[len(dlines) - malength].Close
	return count/float64(malength),precount
}

func GetStockDLines(market,code string)([]DLine){
	db := OpenDatabase()
	defer db.Close()
	rows,err := db.Query(" SELECT DAY,OPEN,CLOSE FROM D WHERE CODE = ? ORDER BY DAY ASC",code)
	if err != nil {
		fmt.Println("查询",market,code,"日线数据时出错：",err)
		panic(err)
	}
	defer rows.Close()
	var results []DLine
	for rows.Next(){
		var open,close float64
		var day time.Time
		err = rows.Scan(&day,&open,&close)
		if err != nil{
			fmt.Println("从数据库提取日线时出错：",err)
			panic(err)
		}
		var line DLine = DLine{}
		line.Market = market
		line.Day 	= day
		line.Code 	= code
		line.Open	= open
		line.Close	= close
		results = append(results,line)
	}
	if rows.Err() != nil {
		fmt.Println("从数据库提取日线时出错：",err)
		panic(rows.Err())
	}
	return results
}

func GetStockDayDetails(market,code string)([]dti.LineData){
	db := OpenDatabase()
	defer db.Close()
	rows,err := db.Query(` SELECT MARKET,CODE,INTCODE,DAY,OPEN,CLOSE,MAX,MIN,VOL,
 						   TNO,UAD,AIAD,AM,TR,MA5,MA10,MA20,MA30,MA60,MA120,MA250
 						   FROM D WHERE CODE = ? ORDER BY DAY ASC `,code)
	if err != nil {
		fmt.Println("查询",market,code,"日线数据时出错：",err)
		panic(err)
	}
	defer rows.Close()
	var results []dti.LineData
	for rows.Next(){
		var d dti.LineData = dti.NewLineData()
		//var open,close ,max,min,uad,aiad,am,tr,ma5,ma10,ma20,ma30,ma60,ma120,ma250 float64
		//var day time.Time
		//var intcode,vol,tno int64
		var day time.Time
		err = rows.Scan(&d.Market,&d.Code,&d.Intcode,&day,&d.OPEN,&d.CLOSE,&d.HIGH,&d.LOW,&d.VOL,
			&d.Tno,&d.Uad,&d.Aiad,&d.Am,&d.Tr,&d.Ma5,&d.Ma10,&d.Ma20,&d.Ma30,&d.Ma60,&d.Ma120,&d.Ma250)
		d.Day = day.Format("2006-01-02")
		if err != nil{
			fmt.Println("从数据库提取日线时出错：",err)
			panic(err)
		}
		results = append(results,d)
	}
	if rows.Err() != nil {
		fmt.Println("从数据库提取日线时出错：",err)
		panic(rows.Err())
	}
	return results
}

func GetStockDLinesDetail(market,code string)([]DLine){
	db := OpenDatabase()
	defer db.Close()
	rows,err := db.Query(` SELECT MARKET,CODE,INTCODE,DAY,OPEN,CLOSE,MAX,MIN,VOL,
 						   TNO,UAD,AIAD,AM,TR,MA5,MA10,MA20,MA30,MA60,MA120,MA250
 						   FROM D WHERE CODE = ? ORDER BY DAY ASC `,code)
	if err != nil {
		fmt.Println("查询",market,code,"日线数据时出错：",err)
		panic(err)
	}
	defer rows.Close()
	var results []DLine
	for rows.Next(){
		var d DLine = DLine{}
		//var open,close ,max,min,uad,aiad,am,tr,ma5,ma10,ma20,ma30,ma60,ma120,ma250 float64
		//var day time.Time
		//var intcode,vol,tno int64
		err = rows.Scan(&d.Market,&d.Code,&d.Intcode,&d.Day,&d.Open,&d.Close,&d.Max,&d.Min,&d.Vol,
			&d.Tno,&d.Uad,&d.Aiad,&d.Am,&d.Tr,&d.Ma5,&d.Ma10,&d.Ma20,&d.Ma30,&d.Ma60,&d.Ma120,&d.Ma250)
		if err != nil{
			fmt.Println("从数据库提取日线时出错：",err)
			panic(err)
		}
		results = append(results,d)
	}
	if rows.Err() != nil {
		fmt.Println("从数据库提取日线时出错：",err)
		panic(rows.Err())
	}
	return results
}


func OpenDatabase()*sql.DB{
	db,err := sql.Open(DRIVER,USERNAME + ":" + PASSWORD + "@" + PROTOCOL + "(" + HOST + ":" + strconv.Itoa(PORT) + ")/" + DBNAME + "?charset=utf8&parseTime=true")
	if err != nil{
		fmt.Println("连接数据库出错：",err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("ping数据库：",err)
	}
	return db
}

// 第一个验证方法的结果
func Alerm001(market,code string)[]time.Time{
	var results []time.Time
	dlines := GetStockDLinesDetail(market,code)
	for i,d := range dlines{
		j := i + 1
		if j >= 25{

			// ma5 > ma10 > ma20
			//c1 := d.Ma5 > d.Ma10 && d.Ma5 > d.Ma20
			c1 := d.Ma5 > d.Ma10 && d.Ma10 > d.Ma20
			// 当天成交量 > 前一天的成交量
			c2 := d.Vol > dlines[i-1].Vol
			// 当天成交量大于最近5天的平均成交量
			c3 := d.Vol > int64(avgVol(dlines[j-5:j]))
			// ma20大于5天前的ma20
			c4 := d.Ma20 > dlines[j-5].Ma20
			//fmt.Println(d.Day,c1,c2,c3,c4,d.Ma5,d.Ma10,d.Ma20,d.Vol)
			if c1 && c2 && c3 && c4{
				fmt.Println("符合条件:",market,code,d.Day.Format("2006-01-02"))
				results = append(results,d.Day)
			}
		}
	}
	return results
}


func avgVol(dlines []DLine)int64{
	var f,c int64
	c = int64(len(dlines))
	for _,d := range dlines{
		f += d.Vol
	}
	return f/c
}