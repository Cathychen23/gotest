package main

import "testing"

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
	x := 2
	z := sqrt(2)
	if z == Flowcontrol5(x) {
		t.Log("test.Flowcontrol5 is success")
	} else {
		t.Error("test.Flowcontrol5 is fail")
	}
}

func TestFlowcontrol7(t *testing.T) {
	test := pow(3, 3, 20)
	if test == Flowcontrol7(3, 3, 20) {
		t.Log("test.Flowcontrol7 is success")
	} else {
		t.Error("test.Flowcontrol7 is fail")
	}
}
