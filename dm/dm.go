package dm

import (
	"net/http"
	"fmt"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"kangqing2008/tools/errors"
	"math"
)

const(
	APP_CODE = "d3ab757c04764223ab5d8f1fd8de8a67"
	STOCK_API_SERVER = "https://ali-stock.showapi.com"
	STOCK_LIST = STOCK_API_SERVER + "/stocklist"


)

//Get Stock List Result
type GslResult struct{
	ResCode 		int64 			`json:"showapi_res_code"`	//状态码
	ResError		string 			`json:"showapi_res_error"`	//错误消息
	ResBody			GslResBody		`json:"showapi_res_body"`	//响应体

}

type GslResBody struct{
	AllPages		int64			`json:"allPages"`			//总页数
	Stocks			[]GslResStock	`json:"contentlist"`		//股票清单
	CurrentPage 	int64			`json:"currentPage"`		//当前页数
	AllNum			int64			`json:"allNum"`				//总数量
	MaxResult		int64			`json:"maxResult"`			//每页最多行数
}

type GslResStock struct{
	Market			string			`json:"market"`				//交易所代码
	Status			string			`json:"status"`				//状态
	Name			string			`json:"name"`				//名称
	CurrCapital		string			`json:"currcapital"`		//流通股本 万股
	ProfitFour		string			`json:"profit_four"`		//四季度净利润
	Code			string			`json:"code"`				//股票代码
	TotalCapital	string			`json:"totalcapital"`		//总股本 万股
	Mgjzc			string			`json:"mgjzc"`				//每股净资产 元
	Pinyin			string			`json:"pinyin"`				//拼音
}


type DTI struct{
	Market		string
	Code 		string
	Intcode		int64
	KDJ_K		float64
	KDJ_D		float64
	KDJ_J		float64
	EMA12		float64
	EMA26		float64
	DIF			float64
	DEA			float64
	MACD		float64
	RSI			float64
}



//market 证券交易所:sh 上海  sz 深圳 hk 港交所
//page   第几页数据
func GetStockList(market string,page int64)string{

	client := &http.Client{}
	req,err := http.NewRequest("GET",STOCK_LIST + "?market=" + market + "&page=" + strconv.Itoa(int(page)),nil)
	if err != nil{
		panic("无法获得请求对象!")
	}
	req.Header.Set("Authorization", "APPCODE " + APP_CODE)

	resp,err := client.Do(req)
	if err != nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	result := string(body)
	fmt.Println(result)
	return result
}

func ParseStockList(str string)(GslResult,error){
	var result = GslResult{}
	err := json.Unmarshal([]byte(str),&result)
	if err != nil {
		fmt.Println(err)
		//panic("无法解析字符串："+str)
		return result,err
	}
	return result,nil
}

func GetStockListAndParse(market string,page int64)(GslResult,error){
	return ParseStockList(GetStockList(market,page))
}


func UpdateAllStockInfo(){
	markets := []string{"sz","sh","hk"}
	for _,market := range markets{
		fmt.Println("开始更新股票市场：",market,"的股票信息！")
		r1,err := GetStockListAndParse(market,1)
		if err != nil{
			fmt.Println("下载股票代码清单出错：",market,1,err)
			panic(err)
		}
		ra1,err := SaveStockList(r1.ResBody.Stocks)
		if err != nil{
			fmt.Println("更新股票信息出错：",market,1,err)
			panic(err)
		}else{
			fmt.Println("更新股票信息成功，行数：",ra1)
			for j := int64(2) ; j <= r1.ResBody.AllPages ; j++{
				r2,err := GetStockListAndParse(market,j)
				if err != nil{
					fmt.Println("下载股票代码清单出错：",market,j,err)
					panic(err)
				}
				ra2,err := SaveStockList(r2.ResBody.Stocks)
				if err != nil{
					fmt.Println("更新股票信息出错：",market,j,err)
					panic(err)
				}
				fmt.Println("更新股票信息成功，行数：",ra2)
			}
		}
	}
}

func SaveStockList(stocks []GslResStock)(int64,error){
	db := OpenDatabase()
	defer db.Close()
	tx,err := db.Begin()
	if err != nil{
		fmt.Println("更新股票信息时，开启事务失败：",err)
		panic(err)
	}
	var result int64 = 0
	defer tx.Rollback()
	s1 ,err1 := tx.Prepare(" SELECT MARKET FROM S WHERE MARKET = ? AND CODE = ? ")
	s2 ,err2 := tx.Prepare(" INSERT INTO S (MARKET,CODE,INTCODE,NAME,STATUS,TC,CC,PY,MGJZC) VALUES(?,?,?,?,?,?,?,?,?) ")
	s3 ,err3 := tx.Prepare(" UPDATE S SET NAME = ?,STATUS = ?,TC = ?,CC = ?,PY = ?,MGJZC = ? WHERE MARKET = ? AND CODE = ? ")
	if err1 != nil || err2 != nil || err3 != nil{
		panic("编译sql语句时出错!")
	}
	defer s1.Close()
	defer s2.Close()
	defer s3.Close()
	for _,s := range stocks{
		rows,err4 := s1.Query(s.Market,s.Code)
		if err4 != nil{
			fmt.Println(err4)
			panic("查询股票信息是否存在时出错")
		}
		exists := rows.Next()
		rows.Close()
		if exists {
			fmt.Println("股票：",s.Market,s.Code,"已经存在，执行更新操作！")
			r5,err5 := s3.Exec(s.Name,s.Status,s.TotalCapital,s.CurrCapital,s.Pinyin,s.Mgjzc,s.Market,s.Code)
			if err5 != nil{
				fmt.Println("更新股票信息时失败:",err5)
				panic(err5)
			}
			ra5 , err7 := r5.RowsAffected()
			if err7 != nil{
				panic(err7)
			}
			result += ra5
			fmt.Println("更新股票：",s.Market,s.Code,"受影响行数：",ra5)
		}else{
			fmt.Println("股票：",s.Market,s.Code,"还不存在，执行插入操作！")
			intcode,_ := strconv.ParseInt(s.Code,10,64)
			r6,err6 := s2.Exec(s.Market,s.Code,intcode,s.Name,s.Status,s.TotalCapital,s.CurrCapital,s.Pinyin,s.Mgjzc)
			if err6 != nil{
				fmt.Println("插入股票信息时失败:",err6)
				panic(err6)
			}
			ra6 , err8 := r6.RowsAffected()
			if err8 != nil{
				panic(err8)
			}
			result += ra6
			fmt.Println("插入股票：",s.Market,s.Code,"受影响行数：",ra6)
		}

	}
	tx.Commit()
	return result,nil
}


func CalcDTIDetail(market,code string)[]DTI{
	lines := GetStockDLinesDetail(market,code)
	j := 0
	var dtis []DTI
	for i,l := range lines{
		j = i+1
		var dti DTI = DTI{}
		dti.Market 	= market
		dti.Code	= code
		intcode,err	:= strconv.ParseInt(code,10,64)
		dti.Intcode = intcode
		errors.PanicIfError(err)
		//kdj 使用9天做默认周期
		fmt.Println(l.Code,l.Close,l.Min,l.Max)
		dtis = append(dtis,dti)

		if j <= 2 {
			dti.KDJ_K = 100
			dti.KDJ_D = 100
			dti.KDJ_J = 100
		} else if j <= 9 {
			dti.KDJ_K,dti.KDJ_D,dti.KDJ_J = CalcKDJ(lines[0:j],dtis[0:j])
		} else {
			dti.KDJ_K,dti.KDJ_D,dti.KDJ_J  = CalcKDJ(lines[j-9:j],dtis[j-9:j])
		}

		//fmt.Println(dti.KDJ_K,dti.KDJ_D,dti.KDJ_J)
		// rsi 使用14天做默认周期
		if j >= 14{
			//dti.RSI = CalcRSI()
		}

		// macd 使用12,26作为默认的周期
		if j > 24{

		}

	}

	return dtis
}

func CalcKDJ(lines []DLine,dtis []DTI)(float64,float64,float64){
	cl := lines[len(lines) -1]

	max,min,_ := MaxMinAvg(lines,CLOSE)

	rsv := (cl.Close - min)/(max - min) * 100

	if math.IsNaN(rsv) {
		rsv = 50
	}
	fmt.Println("长度：",len(lines),"max:",max,"min:",min,"rsv:",rsv)
	var ti13 float64 = float64(1)/float64(3)
	var ti23 float64 = float64(2)/float64(3)
	var k,d,j float64
	//fmt.Println(min,max,rsv)
	k = ti23*PreK(dtis) + ti13*rsv
	d = ti23*PreD(dtis) + ti13*k
	j = 3*d -2*k
	fmt.Println("k:",k,"d:",d,"j:",j)
	return k,d,j
}

func PreK(dtis []DTI)float64{
	if len(dtis) > 1{
		//因为当前集合里面已经包含了需要计算的日期的数据
		//所以需要找倒数第二个数据
		return dtis[len(dtis) -2].KDJ_K
	}else{
		return float64(100)
	}
}

func PreD(dtis []DTI)float64{
	if len(dtis) > 1{
		//因为当前集合里面已经包含了需要计算的日期的数据
		//所以需要找倒数第二个数据
		return dtis[len(dtis) -2].KDJ_D
	}else{
		return float64(100)
	}
}


// 计算最大值，最小值和均值
func MaxMinAvg(lines []DLine,ti string)(float64,float64,float64){
	var max float64 = 0
	var min float64 = 0
	var sum float64 = 0
	for i,l := range lines{
		//如果是使用close指标
		if ti == DTI_CLOSE{
			//第一个值需要初始化
			if i == 0{
				max = l.Max
				min = l.Min
			}

			if l.Max > max {
				max = l.Max
			}
			if l.Min < min{
				min = l.Min
			}
			sum += l.Close
		}
	}
	return max,min,sum/float64(len(lines))
}