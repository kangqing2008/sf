package dti

import (
	"strconv"
	"math"
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
)

type LineData struct{
	Market		string
	Code		string
	Intcode		int64
	Day			string
	OPEN		float64
	CLOSE		float64
	HIGH		float64
	LOW			float64
	VOL			int64

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
		panic("不支持的指标参数:" + X)
	}
}

func NewLineData()LineData{
	return &LineData{
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
		MACD		: math.NaN()}
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
		panic("不支持的指标参数:" + X)
	}
}

