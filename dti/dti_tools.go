package dti

import (
	//"strconv"
	"math"
	"strconv"
)

const(
	CLOSE	= "CLOSE"
	OPEN	= "OPEN"
	HIGH	= "HIGH"
	LOW		= "LOW"
	VOL		= "VOL"

	RSV		= "RSV"
	KDJ_K	= "KDJ_K"
	KDJ_D	= "KDJ_D"
	KDJ_J	= "KDJ_J"

	EMA12	= "EMA12"
	EMA26	= "EMA26"
	DIF		= "DIF"
	DEA		= "DEA"
	MACD	= "MACD"

	BOLL    = "BOLL"
	UB      = "UB"
	LB      = "LB"
)

type DTITools struct{
	Data	[]LineData
	current int
}

type LineData struct{
	Market		string
	Code		string
	Intcode		int64
	Day			string
	OPEN		float64
	CLOSE		float64
	HIGH		float64
	LOW			float64
	VOL			float64

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

	RSV			float64
	KDJ_K		float64
	KDJ_D		float64
	KDJ_J		float64

	EMA12		float64
	EMA26		float64
	DIF			float64
	DEA			float64
	MACD		float64
	EXTRAS      map[string]float64
}

func (this *LineData)Get(X string)float64{
	if X == OPEN{
		return this.OPEN
	}else if X == CLOSE{
		return this.CLOSE
	}else if X == HIGH{
		return this.HIGH
	}else if X == LOW{
		return this.LOW
	}else if X == VOL{
		return this.VOL
	}else if X == RSV{
		return this.RSV
	}else if X == KDJ_K{
		return this.KDJ_K
	}else if X == KDJ_D{
		return this.KDJ_D
	}else if X == KDJ_J{
		return this.KDJ_J
	}else if X == EMA12{
		return this.EMA12
	}else if X == EMA26{
		return this.EMA26
	}else if X == DIF{
		return this.DIF
	}else if X == DEA{
		return this.DEA
	}else if X == MACD{
		return this.MACD
	}else{
		v,ok := this.EXTRAS[X]
		if ok {
			return v
		}else{
			panic("不支持的指标参数:" + X)
		}

	}
}

func NewLineData()LineData{
	return LineData{
		Market		: "",
		Code		: "",
		Intcode		: -1,
		Day			: "",
		OPEN		: math.NaN(),
		CLOSE       : math.NaN(),
		HIGH		: math.NaN(),
		LOW			: math.NaN(),
		VOL			:-1,

		Tno			: math.NaN(),
		Uad			: math.NaN(),
		Aiad		: math.NaN(),
		Am			: math.NaN(),
		Tr			: math.NaN(),

		Ma5			: math.NaN(),
		Ma10		: math.NaN(),
		Ma20		: math.NaN(),
		Ma30		: math.NaN(),
		Ma60		: math.NaN(),
		Ma120		: math.NaN(),
		Ma250		: math.NaN(),

		RSV			: math.NaN(),
		KDJ_K		: math.NaN(),
		KDJ_D		: math.NaN(),
		KDJ_J		: math.NaN(),

		EMA12		: math.NaN(),
		EMA26		: math.NaN(),
		DIF			: math.NaN(),
		DEA			: math.NaN(),
		MACD		: math.NaN(),
		EXTRAS      : make(map[string]float64)}
}

func (this *LineData)Set(X string,Val float64){
	if X == OPEN{
		this.OPEN = Val
	}else if X == CLOSE{
		this.CLOSE = Val
	}else if X == HIGH{
		this.HIGH = Val
	}else if X == LOW{
		this.LOW = Val
	}else if X == VOL{
		this.VOL = Val
	}else if X == RSV{
		this.RSV = Val
	}else if X == KDJ_K{
		this.KDJ_K = Val
	}else if X == KDJ_D{
		this.KDJ_D = Val
	}else if X == KDJ_J{
		this.KDJ_J = Val
	}else if X == EMA12{
		this.EMA12 = Val
	}else if X == EMA26{
		this.EMA26 = Val
	}else if X == DIF{
		this.DIF = Val
	}else if X == DEA{
		this.DEA = Val
	}else if X == MACD{
		this.MACD = Val
	}else{
		this.EXTRAS[X] = Val
		//panic("不支持的指标参数:" + X)
	}
}


func NewTools(data []LineData)*DTITools{
	if(data == nil || len(data) == 0){
		panic("数据为空，无法创建技术指标计算工具")
	}
	tools := DTITools{Data:data,current:len(data) -1}
	return &tools
}

func (this *DTITools)GetCurrent()int{
	return this.current;
}

func (this *DTITools)CurrentData()*LineData{
	return &this.Data[this.current];
}


//求X的N日移动平均值
func (this *DTITools)SetCurrent(crt int){
	length := len(this.Data)
	if (crt > (length -1)) || (crt < 0){
		panic("数组长度为：" + strconv.Itoa(length) +",实际设置位置：" + strconv.Itoa(crt))
	}
	this.current = crt
}

func (this *DTITools)Length()int{
	return len(this.Data)
}

//将current设置为0
//从前往后遍历所有数据
//执行完成后定位调用前的位置
func (this *DTITools)Each(run func(dti *DTITools)){
	temp := this.current
	for this.current = 0;this.current <this.Length();this.current++{
		run(this)
	}
	this.current = temp
}

//将current设置为length-1
//从后往前遍历所有数据
//执行完成后定位调用前的位置
func (this *DTITools)REach(run func(dti *DTITools)){
	temp := this.current
	for this.current = this.Length()-1;this.current >= 0;this.current--{
		run(this)
	}
	this.current = temp
}