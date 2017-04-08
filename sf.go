package main

import (
	"fmt"
	"kangqing2008/sf/dm"
	"kangqing2008/sf/dti"
)

func main() {
	testMACD()
	//r := dti.EMA12 == ("EMA" + strconv.Itoa(12))
	//fmt.Println(r)
}

//测试内存中运行模式，计算MA
func testMA(){
	datas := dm.GetStockDayDetails("sh","002081")
	tools := dti.NewTools(datas)
	tools.Each(func(t *dti.DTITools){
		ma5 := t.MA(dti.CLOSE,5)
		ma10 := t.MA(dti.CLOSE,10)
		d := t.CurrentData()
		fmt.Println(d.Day,ma5,d.Ma5,ma10,d.Ma10,"RSV",t.RSV(9))
	})
}

func testKDJ(){
	datas := dm.GetStockDayDetails("sh","002081")
	tools := dti.NewTools(datas)
	tools.KDJ(9,3,3)
	tools.Each(func(this *dti.DTITools){
		p := this.CurrentData()
		fmt.Println(p.Day,p.RSV,p.KDJ_K,p.KDJ_D,p.KDJ_J)
	})
}

func testMACD(){
	datas := dm.GetStockDayDetails("sh","601881")
	tools := dti.NewTools(datas)
	tools.MACD(12,26,9)
	tools.Each(func(this *dti.DTITools){
		p := this.CurrentData()
		fmt.Println("外部",p.Day,"MA12",p.Get("EMA12"),"MA26",p.Get("EMA26"),"DIF",p.DIF,"DEA",p.DEA,"MACD",p.MACD)
	})
}
//array := []int{0,1,2,3,4,5,6}
//fmt.Println(array[0:7])
//testMA()
//s := fmt.Sprintf("%0.6f", 17.82671567890123456789987654324567898765432)
//f, _ := strconv.ParseFloat(s, 64)
//fmt.Println(s, f)

//func testGetStockSize(){
//	fmt.Println("开始统计每只股票的数据量!")
//	markets,codes,sizes := dm.GetStockStat()
//	fmt.Println("统计完毕，开始展示")
//	for i,market := range markets{
//		fmt.Println(market,codes[i],sizes[i])
//	}
//}

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
//dm.Alerm001("sh","600000")
//str := []string{"0","1","2","3","4"}
//fmt.Println(str[:])
//fmt.Println(str[3:5])
//fmt.Println(str[0:4])
//fmt.Println(str[2:4])
//dm.CalcDayStat()
//dm.UpdateAllStockInfo()
//dm.CalcDTIDetail("sh","601881")

