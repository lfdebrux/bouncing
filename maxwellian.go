package bouncing

import "bitbucket.org/lfdebrux/gsl/interp"

const K_B = 1.3806e-23

var Q *interp.Interp

var Qx []float64 = []float64{
	0.,
	0.2,
	0.4,
	0.6000000000000001,
	0.8,
	1.,
	1.2000000000000002,
	1.4000000000000001,
	1.6,
	1.8,
	2.,
	2.2,
	2.4000000000000004,
	2.6,
	2.8000000000000003,
	3.,
	3.2,
	3.4000000000000004,
	3.6,
	3.8000000000000003,
	4.,
	4.2,
	4.4,
	4.6000000000000005,
	4.800000000000001,
	5.,
	5.2,
	5.4,
	5.6000000000000005,
	5.800000000000001,
	6.,
	6.2,
	6.4,
	6.6000000000000005,
	6.800000000000001,
	7.,
	7.2,
	7.4,
	7.6000000000000005,
	7.800000000000001,
	8.,
}
var Qy []float64 = []float64{
	0.,
	0.002102341288023736,
	0.016227370977989722,
	0.0516242410296933,
	0.11278271841483418,
	0.1987480430987991,
	0.30381412759687243,
	0.4192497777523722,
	0.5354547456266262,
	0.6439187918909294,
	0.7385358700508894,
	0.8161048964495845,
	0.8761111827355618,
	0.9200461839373973,
	0.9505632104744564,
	0.9707091134651116,
	0.9833675596347932,
	0.9909470511230297,
	0.9952760889228657,
	0.9976365072761346,
	0.9988660157102146,
	0.9994781867328134,
	0.999769681164832,
	0.9999024792515926,
	0.9999603848724221,
	0.9999845595017088,
	0.9999942252347149,
	0.9999979274083107,
	0.9999992861113233,
	0.9999997639994297,
	0.9999999251162304,
	0.9999999771924479,
	0.9999999933318036,
	0.9999999981284665,
	0.9999999994957272,
	0.9999999998695542,
	0.9999999999676028,
	0.9999999999922746,
	0.9999999999982313,
	0.9999999999996111,
	1.0,
}

func InitMaxwellian() {
	Q = interp.New(Qy,Qx,interp.Akima)
}

func FreeMaxwellian() {
	Q.Free()
	Q = nil
}
