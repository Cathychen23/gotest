package main

import (
	"fmt"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	n := Add(1, 2)
	if n == 3 {
		t.Log("test.Add is success")
	} else {
		t.Error("test.Add is fail")
	}
}

func TestSwap(t *testing.T) {
	firstname := "Cathy"
	lastname := "Chen"
	x, y := Swap(firstname, lastname)
	if x == lastname && y == firstname {
		t.Log("test.Swap is success")
	} else {
		t.Error("test.Swap is fail")
	}
}

func TestSplit(t *testing.T) {
	sum := 17
	x, y := Split(sum)
	if x == sum*4/9 && y == sum-x {
		t.Log("test.Split is success")
	} else {
		t.Error("test.Split is fail")
	}
}

func TestBasic13(t *testing.T) {
	z := Basic13(3, 4)
	var ZZ uint = 5
	if ZZ == z {
		t.Log("test.Basic13 is success")
	} else {
		t.Error("test.Basic13 is fail")
	}
}

func TestBasic16(t *testing.T) {
	if needInt(Small) == 21 {
		if needFloat(Small) == 0.2 {
			if needFloat(Big) == Big*0.1 {
				t.Log("test.Basic16 is success")
			}
		}
	} else {
		t.Error("test.Basic16 is fail")
	}
}

func TestFlowcontrol3(t *testing.T) {
	sum, test := 1, 1
	for sum < 1000 {
		sum += sum
	}

	x := Flowcontrol3(test)

	if x == sum {
		t.Log("test.Flowcontrol3 is success")
	} else {
		t.Error("test.Flowcontrol3 is fail")
	}
}

func TestFlowcontrol5(t *testing.T) {
	var x float64 = 2
	var y float64 = -4
	var i, j string
	v := math.Sqrt(x)
	k := math.Sqrt(-y)
	i = fmt.Sprint(v)
	j = fmt.Sprint(k) + "i"

	if sqrt(x) == i {
		t.Log("test.Flowcontrol5 is success")
	} else {
		t.Error("test.Flowcontrol5 is fail")
	}
	if sqrt(y) == j {
		t.Log("test.Flowcontrol5 is success")
	} else {
		t.Error("test.Flowcontrol5 is fail")
	}
}

func test2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
func TestFlowcontrol7(t *testing.T) {

	if test2(3, 2, 10) == pow(3, 2, 10) {
		if test2(3, 3, 20) == pow(3, 3, 20) {
			t.Log("test.Flowcontrol7 is success")
		}
	} else {
		t.Error("test.Flowcontrol7 is fail")
	}
}
