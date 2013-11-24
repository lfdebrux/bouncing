package plot

import (
	"fmt"
	"reflect"
	"bitbucket.org/binet/go-gnuplot/pkg/gnuplot"
)

func scale(val float64, in_min, in_max, out_min, out_max float64) float64 {
	return (val - in_min)*(out_max - out_min)/(in_max - in_min) + out_min
}

func Xrange(start,stop,step float64) <-chan float64 {
	c := make(chan float64)
	go func() {
		for i := start; i <= stop; i += step {
			c <- i
		}
		close(c)
	} ()
	return c
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

type Plot struct {
	*gnuplot.Plotter
	plots []interface{}
	titles []string
}

func New(plots ...interface{}) *Plot {
	fname := ""
	persist := true
	debug := false

	gp,err := gnuplot.NewPlotter(fname, persist, debug)
	if err != nil {
			err_string := fmt.Sprintf("** err: %v\n", err)
			panic(err_string)
	}

	if len(plots) > 0 {
		t := reflect.TypeOf(plots[0])
		for _,f := range plots {
			if t != reflect.TypeOf(f) {
				panic("plot: New: all functions to plot must have same type")
			}
			t = reflect.TypeOf(f)
		}

		if t.NumOut() != 1 || t.Out(0).Kind() != reflect.Float64 {
			panic("plot: New: plot functions must return one and only one float64")
		}

		for i := 0; i < t.NumIn(); i++ {
			if t.In(i).Kind() != reflect.Float64 {
				panic("plot: New: arguments to plot functions must have type float64")
			}
		}
	}

	titles := make([]string,len(plots))
	for i := range titles {
		titles[i] = fmt.Sprintf("plot %d",i)
	}

	p := &Plot{gp,plots,titles}

	p.CheckedCmd("set datafile missing \"NaN\"")

	return p
}

func (p *Plot) Close() {
	p.CheckedCmd("q")
	p.Plotter.Close()
}

func (p *Plot) Vary(param int, data []float64, fixed ...float64) {
	if len(p.plots) < 1 {
		panic("plot: Vary: need functions to plot")
	}
	t := reflect.TypeOf(p.plots[0])
	params := make([]reflect.Value,len(fixed)+1)

	j := 0
	for i,x := range fixed {
		if i == param {
			params[j] = reflect.New(t.In(0))
			j++
		}
		params[j] = reflect.ValueOf(x)
		j++
	}

	if len(fixed)+1 != t.NumIn() {
		panic("plot: Vary: must have fixed points for all parameters except varying")
	}

	// TODO: This is very slow
	for num,f := range p.plots {
		curry := func(v float64) float64 {
			params[param] = reflect.ValueOf(v)
			result := reflect.ValueOf(f).Call(params)
			return result[0].Float()
		}
		p.PlotFunc(data,curry,p.titles[num])
	}
}

func (p *Plot) SetTitles(titles ...string) {
	for i := range titles {
		p.titles[i] = titles[i]
	}
}
