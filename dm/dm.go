package dm

import (
	"net/http"
	"fmt"
	"strconv"
	"io/ioutil"
	"encoding/json"
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

//market 证券交易所:sh 上海  sz 深圳 hk 港交所
//page   第几页数据
func GetStockList(market string,page int)string{

	client := &http.Client{}
	req,err := http.NewRequest("GET",STOCK_LIST + "?market=" + market + "&page=" + strconv.Itoa(page),nil)
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

func ParseStockList(str string) GslResult{
	var result = GslResult{}
	err := json.Unmarshal([]byte(str),&result)
	if err != nil {
		fmt.Println(err)
		panic("无法解析字符串："+str)
	}
	return result
}