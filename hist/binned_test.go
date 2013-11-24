package hist

import "testing"

func TestIndex(t *testing.T) {
	b := NewBins(0,1,0.1)
	test := map[float64]int{0.02:0,0.5:5,0.3:3,0.79:7}
	for v,i := range test {
		if j := b.index(v); j != i {
			t.Log(v,v/b.Step)
			t.Errorf("value %f -> %dth bin (%f); instead got -> %dth (%f)",v,i,b.Binmap[i],j,b.Binmap[j])
		}
	}
}
