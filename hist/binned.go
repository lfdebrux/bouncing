package hist

import "fmt"

type B struct {
	Binmap []float64
	Binned []float64
	Start,Stop,Step float64
	Num int
}

func scale(val float64, in_min, in_max, out_min, out_max float64) float64 {
	return (val - in_min)*(out_max - out_min)/(in_max - in_min) + out_min
}

func Range(start,stop,step float64) []float64 {
	num := int((stop-start)/step) + 1
	s := make([]float64,num)
	for i := 0; i < num; i++ {
		x := scale(float64(i),0,float64(num-1),start,stop)
		s[i] = x
	}
	return s
}

func NewBins(start,stop,step float64) *B {
	num := int((stop-start)/step) + 1
	bm := Range(start,stop,step)
	bn := make([]float64,num)
	return &B{bm,bn,start,stop,step,num}
}

func (b *B) index(x float64) int {
	if x < b.Start || x > b.Stop {
		panic(fmt.Sprintf("plot: binned: value %f does not fit in range %f..%f",x,b.Start,b.Stop))
	}
	return int(x/b.Step)
}

func (b *B) Bin(xx ...float64) {
	for _,x := range xx {
		i := b.index(x)
		b.Binned[i]++
	}
	// TODO: return error?
}

func (b *B) Get(i int) float64 {
	return b.Binned[i]
}
