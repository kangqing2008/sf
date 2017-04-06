package dti

import "strconv"

type DTITools struct{
	Data	[]LineData
	current int
}

//求X的N日移动平均值
func (this *DTITools)SetCurrent(crt int){
	length := len(this.Data)
	if (crt > (length -1)) || (crt < 0){
		panic("数组长度为：" + strconv.Itoa(length) +",实际设置位置：" + strconv.Itoa(crt))
	}
	this.current = crt
}

//求X的N日移动平均值
func (this *DTITools)MA(X string,N int)float64{
	//其实位置为当前位置-N+1，应为要加上current位置的数据
	length := len(this.Data)
	start := this.current - N + 1
	end   := this.current
	if start < 0 || end > length{
		panic("计算MA时数组越界，length：" + strconv.Itoa(length) + "，current:" + strconv.Itoa(this.current))
	}

	temp := this.Data[start:end]
	var sum float64
	for _,line := range temp{
		sum += line.Get(X)
	}
	return sum/float64(N)
}

//获得往前Index各周期的数据
func (this *DTITools)Pre(Index int)LineData{
	length := len(this.Data)
	if (length < 1) || (length < (Index + 1)){
		panic("数据总长度是：" + strconv.Itoa(length) + ",不允许往前偏移：" + strconv.Itoa(Index))
	}
	return this.Data[len(this.Data) - 1 - Index]
}

//获得第Index个周期的数据
func (this *DTITools)Get(Index int)LineData{
	length := len(this.Data)
	if (length <= Index){
		panic("数据总长度是：" + strconv.Itoa(length) + ",不允许定位到：" + strconv.Itoa(Index))
	}
	return this.Data[Index]
}

//获得第Index个周期的数据
func (this *DTITools)Last()LineData{
	length := len(this.Data)
	if (length < 1){
		return nil
	}
	return this.Data[length - 1]
}


//求X的N日移动平均值
func (this *DTITools)EMA(X string,N int)float64{
	return this.ema_impl(X,N,0)
}

//EMA的递归实现
func (this *DTITools)ema_impl(X string,N int,Index int)float64{
	d := this.Pre(Index)
	val := d.Get(X)
	return (2*val+(N-1)*this.ema_impl(X,N,Index+1))/(N+1)
}


//DMA:求X的动态移动平均
func (this *DTITools)DMA(X string,A float64)float64{
	return this.dma_impl(X,A,0)
}

//DMA函数的递归实现
func (this *DTITools)dma_impl(X string,A float64,Index int)float64{
	if (A >= 1) || (A <= 0){
		panic("DMA函数的A值，必须大于0，而且小于1，实际值是：" + strconv.FormatFloat(A,'f',-1,64))
	}
	d := this.Pre(Index)
	val := d.Get(X)
	return A*val+(1-A)*this.dma_impl(X,A,Index +1)
}


//SMA函数实现,原始实现为:Y=(X*M+Y'*(N-M))/N
//增加了XD参数，用于代表待计算的指标
//比如在KDJ中的K值计算中,调用方法为：
//"KDJ_K" = SMA("RSV",3,1,"KDJ_K")
func (this *DTITools)SMA(X string,N,M int,XD string)float64{
	return this.sma_impl(X,N,M,0,XD)
}

//如果前一天的目标值已经计算则直接使用，如果没有计算，则全部计算出来
func (this *DTITools)sma_impl(X string,N,M,Index int,XD string)float64{
	//先获取到当天的用于计算的值
	d := this.Pre(Index)
	val := d.Get(X)
	//获取前一天的计算目标值
	xdVal := this.Pre(Index+1).Get(XD)
	//如果还没有计算，就先计算出来，并且存储进去
	//如果已经计算就直接执行公式
	if math.IsNaN(xdVal){
		xdVal = this.sma_impl(X,N,M,Index+1,XD)
		this.Pre(Index+1).Set(XD,xdVal)
	}
	return (val*M+xdVal*(N-M))/N
}

//求X的N日移动平均值
func EMA(X float64,N int,PRE float64)float64{
	return (2*X+(N-1)*PRE)/(N+1)
}


//求X的N日移动平均值
func DMA(X float64,A float64,PRE float64)float64{
	return A*X+(1-A)*PRE
}

//SMA函数实现,原始实现为:Y=(X*M+Y'*(N-M))/N
//增加了PRE参数，用于代表Y'
func SMA(X float64,N,M int,PRE float64)float64{
	return (X*M+PRE*(N-M))/N
}

//计算data数据中N天内X指标的最低值
func (this *DTITools)LLV(X string,N int) float64 {
	//从数据中取出最近N天内的数据
	temp := this.Data[len(this.Data) - N:]
	//返回的结果
	var minValue float64 = 0

	for i,line := range temp{
		if i == 0{
			minValue = line.Get(X)
		}else{
			if line.Get(X) < minValue{
				minValue = line.Get(X)
			}
		}
	}
	return minValue
}


//计算data数据中N天内X指标的最高值
func (this *DTITools)HHV(X string,N int)float64{
	//从数据中取出最近N天内的数据
	temp := this.Data[len(this.Data) - N:]
	//返回的结果
	var maxValue float64 = 0

	for i,line := range temp{
		if i == 0{
			maxValue = line.Get(X)
		}else{
			if line.Get(X) > maxValue{
				maxValue = line.Get(X)
			}
		}
	}
	return maxValue
}

func (this *DTITools)KDJ(k,d,j float64){
	length := len(this.Data)
	for i := length -1; i >= 0; i-- {

	}


	return this.Last().KDJ_K,this.Last().KDJ_D,this.Last().KDJ_J
}
