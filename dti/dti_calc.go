package dti

import (
	"strconv"
	//"math"
	"math"
	"fmt"
	"kangqing2008/gotools/tools/str"
)



//求X的N日移动平均值
func (this *DTITools)MA(X string,N int)float64{
	//其实位置为当前位置-N+1，应为要加上current位置的数据
	length := this.Length()
	start := this.current - N + 1
	end   := this.current + 1
	//如果长度不够，则有多长，算多长.
	if start < 0{
		start = 0
	}
	//如果计算出来的位置超出了数组长度，是不允许的
	if end > length{
		panic("数据长度：" + strconv.Itoa(length) + "，当前位置：" + strconv.Itoa(length) + "，截取的数组位置Array[" + strconv.Itoa(start) + ":" + strconv.Itoa(end) + "]")
	}
	temp := this.Data[start:end]
	var sum float64
	for _,line := range temp{
		sum += line.Get(X)
	}
	return str.RF64(sum/float64(end -start),2)
}


//求X的N日随机未成熟指标
func (this *DTITools)RSV(N int)float64{
	c := this.CurrentData().CLOSE
	h := this.HHV(HIGH,N)
	l := this.LLV(LOW,N)
	r := ((c-l)/(h-l))*100
	//fmt.Println(c,h,l,r)
	if math.IsNaN(r){
		return 0
	}else{
		return str.RF64(r,2)
	}
}

//获得往前Index各周期的数据
func (this *DTITools)REF(X string,N int)float64{
	index := this.current - N
	if (index < 0) || (index > (this.Length() -1)){
		panic("数据总长度是：" + strconv.Itoa(this.Length()) + ",不允许引用：" + strconv.Itoa(index))
	}
	d := this.Get(index)
	return d.Get(X)
}

//获得往前Index各周期的数据
func (this *DTITools)Move(to int){
	length := this.Length()
	moved := this.current + to
	if (moved < 0) || (length < (moved + 1)){
		panic("数据总长度是：" + strconv.Itoa(length) + ",当前位置:" + strconv.Itoa(this.current) + ",移动位置：" + strconv.Itoa(to))
	}
	this.current = moved
}

//获得第Index个周期的数据
func (this *DTITools)Get(Index int)LineData{
	length := this.Length()
	if (length <= Index){
		panic("数据总长度是：" + strconv.Itoa(length) + ",不允许定位到：" + strconv.Itoa(Index))
	}
	return this.Data[Index]
}

//获得第Index个周期的数据
func (this *DTITools)Last()LineData{
	length := len(this.Data)
	if (length < 1){
		panic("nil")
	}
	return this.Data[length - 1]
}


//求X的N日移动平均值
func (this *DTITools)EMA(X string,N int)float64{
	return this.ema_impl(X,N,0)
}

//EMA的递归实现
func (this *DTITools)ema_impl(X string,N int,Index int)float64{
	//d := this.Pre(Index)
	//val := d.Get(X)
	//return (2*val+(N-1)*this.ema_impl(X,N,Index+1))/(N+1)
	return 1
}



//DMA:求X的动态移动平均
//func (this *DTITools)DMA(X string,A float64)float64{
//	return this.dma_impl(X,A,0)
//}




//SMA函数实现,原始实现为:Y=(X*M+Y'*(N-M))/N
//增加了PRe参数，代表待计算指标的前一个值
func SMA(X float64,N,M int,PRE float64)float64{
	return (X*float64(M)+PRE*float64(N-M))/float64(N)
}


//求X的N日移动平均值
//func EMA(X float64,N int,PRE float64)float64{
//	return (2*X+(N-1)*PRE)/(N+1)
//}


//求X的N日移动平均值
func DMA(X float64,A float64,PRE float64)float64{
	return A*X+(1-A)*PRE
}

//SMA函数实现,原始实现为:Y=(X*M+Y'*(N-M))/N
//增加了PRE参数，用于代表Y'
//func SMA(X float64,N,M int,PRE float64)float64{
//	return (X*M+PRE*(N-M))/N
//}

//计算data数据中N天内X指标的最低值
func (this *DTITools)LLV(X string,N int) float64 {
	//从数据中取出最近N天内的数据
	start := this.current - N + 1
	end   := this.current
	//返回的结果
	var minValue float64 = 0
	if start < 0{
		start = 0
	}
	for i := start;i<=end;i++{
		d := this.Get(i)
		value :=  d.Get(X)
		if i == start{
			minValue = value
		}else{
			if value < minValue{
				minValue = value
			}
		}
	}
	return str.RF64(minValue,2)
}


//计算data数据中N天内X指标的最高值
func (this *DTITools)HHV(X string,N int)float64{
	//从数据中取出最近N天内的数据
	start := this.current - N + 1
	end   := this.current
	//返回的结果
	var maxValue float64 = 0

	if start < 0{
		start = 0
	}
	for i := start;i<=end;i++{
		d :=  this.Get(i)
		value := d.Get(X)
		if i == start{
			maxValue = value
		}else{
			if value > maxValue{
				maxValue = value
			}
		}
	}
	return str.RF64(maxValue,2)
}

func (this *DTITools)KDJ(N,M1,M2 int){
	this.Each(func (t *DTITools){
		//fmt.Println("当前行号:",this.current + 1)
		p := t.CurrentData()
		p.RSV = t.RSV(N)
		p.KDJ_K = kdj_k(p.RSV,M1,t)
		p.KDJ_D = kdj_d(p.KDJ_K,M2,t)
		j := str.RF64(3*p.KDJ_K - 2*p.KDJ_D,2)
		if j < 0 {
			j = 0
		}
		if j > 100{
			j = 100
		}
		p.KDJ_J = j
	})
}



func kdj_k(RSV float64,M1 int,this *DTITools)float64{
	PRE := float64(50)
	//如果前一天的指标为0
	if this.current > 0{
		PRE = this.REF(KDJ_K,1)
	}
	r := str.RF64(SMA(RSV,M1,1,PRE),2)
	if r < 0 {
		r = 0
	}
	if r > 100 {
		r = 100
	}
	//fmt.Println("k-method,行号：",this.current,RSV,PRE,r)
	return r
}

func kdj_d(KDJ_K float64,M2 int,this *DTITools)float64{
	PRE := float64(50)
	//如果前一天的指标为0
	if this.current > 0{
		PRE = this.REF(KDJ_D,1)
	}
	r := str.RF64(SMA(KDJ_K,M2,1,PRE),2)
	if r < 0 {
		r = 0
	}
	if r > 100 {
		r = 100
	}
	return r
}


//计算当前数据中所有的MACD指标
func (this *DTITools)MACD(SHORT,LONG,MID int){
	//存储数据的Key名称
	ESHORT := "EMA" + strconv.Itoa(SHORT)
	ELONG  := "EMA" + strconv.Itoa(LONG)
	this.Each(func (t *DTITools){
		fmt.Println("当前行号:",this.current + 1)
		p := t.CurrentData()
		c := p.CLOSE
		pre_eshort := float64(-1)
		pre_elong  := float64(-1)
		pre_dea    := float64(0)
		if t.current == 0 {
			pre_eshort = c
			pre_elong  = c
			pre_dea    = float64(0)
		}else{
			pre_eshort = t.REF(ESHORT,1)
			pre_elong = t.REF(ELONG,1)
			pre_dea    = t.REF(DEA,1)

		}

		eshort := EMA(c,SHORT,pre_eshort)
		elong  := EMA(c,LONG,pre_elong)
		dif    := eshort - elong
		dea    := EMA(dif,MID,pre_dea)
		macd   := (dif - dea)*2
		p.DIF  = str.RF64(dif,2)
		p.DEA  = str.RF64(dea,2)
		p.MACD = str.RF64(macd,2)
		p.Set(ESHORT,eshort)
		p.Set(ELONG,elong)
		//p.EMA12 = eshort
		//p.EMA26 = elong
		//fmt.Println("前一个数据",p.Day,pre_eshort,pre_elong,pre_dea,eshort,elong,dea)
		//fmt.Println("内部",p.Day,"Close",p.CLOSE,"MA12",p.Get("EMA12"),"MA26",p.Get("EMA26"),"DIF",p.DIF,"DEA",p.DEA,"MACD",p.MACD)

	})
}

func EMA(C float64,N int,PRE float64)float64{
	a := float64(2)/float64(N+1)
	r := a * (C - PRE) + PRE
	return str.RF64(r,2)
}


func (this *DTITools)BOLL(M int){
	this.Each(func(t *DTITools){
		boll := this.MA(CLOSE,M)
		ub   := boll + 2*this.STD(CLOSE,M)
		lb   := boll - 2*this.STD(CLOSE,M)
		p := t.CurrentData()
		p.Set(BOLL,boll)
		p.Set(UB,str.RF64(ub,2))
		p.Set(LB,str.RF64(lb,2))
	})
}

//计算估算标准偏差
func (this *DTITools)STD(X string,M int)float64{
	length := M
	if length > (this.GetCurrent() + 1){
		length = this.GetCurrent() + 1
	}
	ma := this.MA(X,length)
	count := float64(0)
	for i:=0;i<length;i++{
		ref := this.REF(X,i)
		count += (ref-ma)*(ref-ma)
	}
	sqrt := count/float64(length)
	return math.Sqrt(sqrt)
}

//计算估算标准偏差
func (this *DTITools)STDP(X string,M int)float64{
	length := M
	if length > (this.GetCurrent() + 1){
		length = this.GetCurrent() + 1
	}
	ma := this.MA(X,length)
	count := float64(0)
	for i:=0;i<length;i++{
		ref := this.REF(X,i)
		count += (ref-ma)*(ref-ma)
	}
	sqrt := count/float64(length-1)
	return math.Sqrt(sqrt)
}
