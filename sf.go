package main

import (
	//"kangqing2008/sf/dm"
	"kangqing2008/sf/dm"
	"fmt"
)

func main() {
	//dm.GetStockList("sh",1)
	//res := dm.ParseStockList(str)
	//fmt.Println(res.ResCode)
	//for i,stock := range res.ResBody.Stocks{
	//	fmt.Println(i,"-",stock.Code,":",stock.Name)
	//}
	//path := "E:\\stocks\\hk"
	//r,err := file.ReadLine(filename)
	//if err != nil{
	//	panic("读取文件:" + filename + " 失败!")
	//}
	//for i,line := range r{
	//	fmt.Println("第",i,"行:",line)
	//}
	//n,s := file.GetFileName(filename)
	//fmt.Println(n,".",s)
	//runes := []rune("sh000001")
	//fmt.Println(string(runes[0:2]),string(runes[2:]))
	//dm.ImportFile(filename)
	//dm.ImportDirectory(path)

	//testGetStockSize()
	dm.Alerm001("sh","600000")
	//str := []string{"0","1","2","3","4"}
	//fmt.Println(str[:])
	//fmt.Println(str[3:5])
	//fmt.Println(str[0:4])
	//fmt.Println(str[2:4])

}


func testGetStockSize(){
	fmt.Println("开始统计每只股票的数据量!")
	markets,codes,sizes := dm.GetStockStat()
	fmt.Println("统计完毕，开始展示")
	for i,market := range markets{
		fmt.Println(market,codes[i],sizes[i])
	}
}

const str =`{
    "showapi_res_code": 0,
    "showapi_res_error": "",
    "showapi_res_body": {
        "allPages": 26,
        "contentlist": [
            {
                "market": "sh",
                "status": "0",
                "name": "数据港",
                "currcapital": "5265",
                "profit_four": "0.48",
                "code": "603881",
                "totalcapital": "21058.6508",
                "mgjzc": "2.496126",
                "pinyin": "sjg"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "安正时尚",
                "currcapital": "7126",
                "profit_four": "1.08",
                "code": "603839",
                "totalcapital": "28504",
                "mgjzc": "5.685118",
                "pinyin": "azss"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "海峡环保",
                "currcapital": "11250",
                "profit_four": "0.25333333333333",
                "code": "603817",
                "totalcapital": "45000",
                "mgjzc": "2.649781",
                "pinyin": "hxhb"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "奇精机械",
                "currcapital": "2000",
                "profit_four": "1.36",
                "code": "603677",
                "totalcapital": "8000",
                "mgjzc": "7.169465",
                "pinyin": "qjjx"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "镇海股份",
                "currcapital": "2557.63",
                "profit_four": "0.84",
                "code": "603637",
                "totalcapital": "10230.5089",
                "mgjzc": "4.836923",
                "pinyin": "zhgf"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "科森科技",
                "currcapital": "5266.67",
                "profit_four": "1.0557333333333",
                "code": "603626",
                "totalcapital": "21066.67",
                "mgjzc": "3.433535",
                "pinyin": "kskj"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "茶花股份",
                "currcapital": "6000",
                "profit_four": "0.568",
                "code": "603615",
                "totalcapital": "24000",
                "mgjzc": "4.419738",
                "pinyin": "chgf"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "*ST新集",
                "currcapital": "259054.18",
                "profit_four": "-1.2655",
                "code": "601918",
                "totalcapital": "259054.18",
                "mgjzc": "1.767933",
                "pinyin": "*stxj"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "*ST星湖",
                "currcapital": "55039.3465",
                "profit_four": "-0.5196",
                "code": "600866",
                "totalcapital": "64539.3465",
                "mgjzc": "1.647587",
                "pinyin": "*stxh"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "*ST宏盛",
                "currcapital": "15567.9074",
                "profit_four": "-0.0212",
                "code": "600817",
                "totalcapital": "16091.0082",
                "mgjzc": "0.583479",
                "pinyin": "*sths"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "*ST昆机",
                "currcapital": "39018.6291",
                "profit_four": "-3.1525",
                "code": "600806",
                "totalcapital": "53108.1103",
                "mgjzc": "1.367254",
                "pinyin": "*stkj"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "*ST黑豹",
                "currcapital": "34494.039",
                "profit_four": "-0.5705",
                "code": "600760",
                "totalcapital": "34494.039",
                "mgjzc": "1.090346",
                "pinyin": "*sthb"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "*ST新梅",
                "currcapital": "44638.308",
                "profit_four": "-1.2147",
                "code": "600732",
                "totalcapital": "44638.308",
                "mgjzc": "0.729704",
                "pinyin": "*stxm"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "*ST云维",
                "currcapital": "123247",
                "profit_four": "-13.0563",
                "code": "600725",
                "totalcapital": "123247",
                "mgjzc": "-4.603121",
                "pinyin": "*styw"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "*ST百花",
                "currcapital": "24852.4307",
                "profit_four": "-1.2507",
                "code": "600721",
                "totalcapital": "40038.6394",
                "mgjzc": "4.771563",
                "pinyin": "*stbh"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "*ST常林",
                "currcapital": "64028.4",
                "profit_four": "-2.9474",
                "code": "600710",
                "totalcapital": "130674.9434",
                "mgjzc": "1.651116",
                "pinyin": "*stcl"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "*ST中企",
                "currcapital": "186705.9398",
                "profit_four": "-8.7236",
                "code": "600675",
                "totalcapital": "186705.9398",
                "mgjzc": "1.557033",
                "pinyin": "*stzq"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "退市博元",
                "currcapital": "19032.8337",
                "profit_four": "-0.2434",
                "code": "600656",
                "totalcapital": "19034.3678",
                "mgjzc": "2.363317",
                "pinyin": "tsby"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "*ST兴业",
                "currcapital": "19464.192",
                "profit_four": "-2.1825",
                "code": "600603",
                "totalcapital": "52375.5844",
                "mgjzc": "-0.951964",
                "pinyin": "*stxy"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "*ST八钢",
                "currcapital": "76644.8935",
                "profit_four": "-13.3133",
                "code": "600581",
                "totalcapital": "76644.8935",
                "mgjzc": "-1.892123",
                "pinyin": "*stbg"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "ST慧球",
                "currcapital": "39479.3708",
                "profit_four": "-0.1893",
                "code": "600556",
                "totalcapital": "39479.3708",
                "mgjzc": "0.037641",
                "pinyin": "sthq"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "*ST山煤",
                "currcapital": "198245.614",
                "profit_four": "-5.9613",
                "code": "600546",
                "totalcapital": "198245.614",
                "mgjzc": "1.678770",
                "pinyin": "*stsm"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "ST狮头",
                "currcapital": "23000",
                "profit_four": "-0.2053",
                "code": "600539",
                "totalcapital": "23000",
                "mgjzc": "1.924925",
                "pinyin": "stst"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "*ST中发",
                "currcapital": "15389",
                "profit_four": "-0.0999",
                "code": "600520",
                "totalcapital": "15843",
                "mgjzc": "2.607852",
                "pinyin": "*stzf"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "*ST吉恩",
                "currcapital": "81112.1542",
                "profit_four": "-9.0278",
                "code": "600432",
                "totalcapital": "160372.3916",
                "mgjzc": "2.910907",
                "pinyin": "*stje"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "*ST金瑞",
                "currcapital": "45125.6401",
                "profit_four": "0.6895",
                "code": "600390",
                "totalcapital": "374838.7882",
                "mgjzc": "2.811502",
                "pinyin": "*stjr"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "*ST星马",
                "currcapital": "55574.0597",
                "profit_four": "-3.6114",
                "code": "600375",
                "totalcapital": "55574.0597",
                "mgjzc": "4.329025",
                "pinyin": "*stxm"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "*ST油工",
                "currcapital": "57815.4688",
                "profit_four": "-5.1573",
                "code": "600339",
                "totalcapital": "558314.7471",
                "mgjzc": "-0.637800",
                "pinyin": "*styg"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "*ST亚星",
                "currcapital": "31559.4",
                "profit_four": "0.2678",
                "code": "600319",
                "totalcapital": "31559.4",
                "mgjzc": "0.023287",
                "pinyin": "*styx"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "*ST商城",
                "currcapital": "17740.5726",
                "profit_four": "-1.1595",
                "code": "600306",
                "totalcapital": "17813.8918",
                "mgjzc": "-0.735738",
                "pinyin": "*stsc"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "*ST南化",
                "currcapital": "23514.814",
                "profit_four": "-0.6457",
                "code": "600301",
                "totalcapital": "23514.814",
                "mgjzc": "0.957970",
                "pinyin": "*stnh"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "ST景谷",
                "currcapital": "12980",
                "profit_four": "0.3371",
                "code": "600265",
                "totalcapital": "12980",
                "mgjzc": "0.156314",
                "pinyin": "stjg"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "ST成城",
                "currcapital": "33644.16",
                "profit_four": "-0.6892",
                "code": "600247",
                "totalcapital": "33644.16",
                "mgjzc": "0.112896",
                "pinyin": "stcc"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "*ST山水",
                "currcapital": "20244.588",
                "profit_four": "0.0418",
                "code": "600234",
                "totalcapital": "20244.588",
                "mgjzc": "0.235015",
                "pinyin": "*stss"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "*ST沧大",
                "currcapital": "29418.8216",
                "profit_four": "-1.4429",
                "code": "600230",
                "totalcapital": "29418.8216",
                "mgjzc": "4.143844",
                "pinyin": "*stcd"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "*ST江泉",
                "currcapital": "51169.7213",
                "profit_four": "0.4455",
                "code": "600212",
                "totalcapital": "51169.7213",
                "mgjzc": "1.449730",
                "pinyin": "*stjq"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "*ST新亿",
                "currcapital": "104407.1743",
                "profit_four": "0.2283",
                "code": "600145",
                "totalcapital": "149110.038",
                "mgjzc": "0.421424",
                "pinyin": "*stxy"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "ST明科",
                "currcapital": "33652.6",
                "profit_four": "2.2334",
                "code": "600091",
                "totalcapital": "43741.2524",
                "mgjzc": "2.115657",
                "pinyin": "stmk"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "博天环境",
                "currcapital": "4001",
                "profit_four": "0.28",
                "code": "603603",
                "totalcapital": "40001",
                "mgjzc": "2.228268",
                "pinyin": "bthj"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "上海天洋",
                "currcapital": "1500",
                "profit_four": "1.2",
                "code": "603330",
                "totalcapital": "6000",
                "mgjzc": "7.466511",
                "pinyin": "shty"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "白银有色",
                "currcapital": "69800",
                "profit_four": "0.048",
                "code": "601212",
                "totalcapital": "697296.5867",
                "mgjzc": "1.832463",
                "pinyin": "byys"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "工大高新",
                "currcapital": "49878.1936",
                "profit_four": "0.7658",
                "code": "600701",
                "totalcapital": "103473.5218",
                "mgjzc": "4.131187",
                "pinyin": "gdgx"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "江山欧派",
                "currcapital": "2021",
                "profit_four": "1.8",
                "code": "603208",
                "totalcapital": "8081.6061",
                "mgjzc": "6.419154",
                "pinyin": "jsop",
                "listing_date": "2017-02-10"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "新坐标",
                "currcapital": "1500",
                "profit_four": "0.98666666666667",
                "code": "603040",
                "totalcapital": "6000",
                "mgjzc": "5.723900",
                "pinyin": "xzb",
                "listing_date": "2017-02-09"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "德创环保",
                "currcapital": "5050",
                "profit_four": "0.28",
                "code": "603177",
                "totalcapital": "20200",
                "mgjzc": "2.266849",
                "pinyin": "dchb",
                "listing_date": "2017-02-07"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "百傲化学",
                "currcapital": "3334",
                "profit_four": "0.97333333333333",
                "code": "603360",
                "totalcapital": "13334",
                "mgjzc": "4.332523",
                "pinyin": "bahx",
                "listing_date": "2017-02-06"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "法兰泰克",
                "currcapital": "4000",
                "profit_four": "0.34",
                "code": "603966",
                "totalcapital": "16000",
                "mgjzc": "4.115552",
                "pinyin": "fltk",
                "listing_date": "2017-01-25"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "集友股份",
                "currcapital": "1700",
                "profit_four": "0.87093333333333",
                "code": "603429",
                "totalcapital": "6800",
                "mgjzc": "3.594939",
                "pinyin": "jygf",
                "listing_date": "2017-01-24"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "正裕工业",
                "currcapital": "2667",
                "profit_four": "1.2",
                "code": "603089",
                "totalcapital": "10667",
                "mgjzc": "4.563172",
                "pinyin": "zygy",
                "listing_date": "2017-01-26"
            },
            {
                "market": "sh",
                "status": "0",
                "name": "华达科技",
                "currcapital": "4000",
                "profit_four": "2.2266666666667",
                "code": "603358",
                "totalcapital": "16000",
                "mgjzc": "9.749437",
                "pinyin": "hdkj",
                "listing_date": "2017-01-25"
            }
        ],
        "currentPage": 1,
        "allNum": 1268,
        "maxResult": 50
    }
}`