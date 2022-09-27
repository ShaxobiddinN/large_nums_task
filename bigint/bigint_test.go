package bigint_test

import (
	"testing"

	"github.com/ShaxobiddinN/large_nums_task/bigint"
)

//----------------------------------test for NewIntTest function------------------------------------

func TestBigint(t *testing.T){
	// got,_:=bigint.NewInt("047")
	// want:=bigint.Bigint{Value: "47"}

	// if got!=want {
	// 	t.Error("its problem")
	// }
	tests := []struct{
        num      string
        expected string
      
    }{
        {num:"123" , expected: "123"},
		{num:"0123" , expected: "123"},
		{num:"00000000000000123" , expected: "123"},
		{num:"-123" , expected: "123"},
		{num:"+123" , expected: "123"},
    }
    for _,  tt := range tests{
        res,_:= bigint.NewInt(tt.num)
        if res.Value != tt.expected{
            t.Error("its problem")
        }
    }


}
	//---------------------------Add
	// type addTest struct{
	// 	arg1,arg2,expected string
	// }
	// var addTest =[]addTest{
	// 	addTest{"025","9","34"},
	// }

//----------------------------------test for Add function------------------------------------
	func TestAdd(t *testing.T)  {
		tests:= []struct{
			num1 bigint.Bigint
			num2 bigint.Bigint
			expected string
		} {
		{num1: bigint.Bigint{Value: "123"},num2: bigint.Bigint{Value: "01"},expected:"124"},
		{num1: bigint.Bigint{Value: "0123"},num2: bigint.Bigint{Value: "-01"},expected:"122"},
		{num1: bigint.Bigint{Value: "-123"},num2: bigint.Bigint{Value: "01"},expected:"-122"},
		{num1: bigint.Bigint{Value: "-123"},num2: bigint.Bigint{Value: "-01"},expected:"-124"},
		}
		for i:=0;i<len(tests);i++{
			res:=bigint.Add(tests[i].num1,tests[i].num2)
			if res.Value!=tests[i].expected{
				t.Error("Error!!! there is an error adding")
			}
		}
	}

//----------------------------------test for Substaction function----------------------------------
	func TestSub(t *testing.T)  {
		tests:= []struct{
			num1 bigint.Bigint
			num2 bigint.Bigint
			expected string
		} {
		{num1: bigint.Bigint{Value: "123"},num2: bigint.Bigint{Value: "01"},expected:"122"},
		{num1: bigint.Bigint{Value: "0123"},num2: bigint.Bigint{Value: "-01"},expected:"124"},
		{num1: bigint.Bigint{Value: "-123"},num2: bigint.Bigint{Value: "01"},expected:"-124"},
		{num1: bigint.Bigint{Value: "-123"},num2: bigint.Bigint{Value: "-01"},expected:"-122"},
		{num1: bigint.Bigint{Value: "1"},num2: bigint.Bigint{Value: "123"},expected:"-122"},
		}
		for _,tt:=range tests{
			res:=bigint.Sub(tt.num1,tt.num2)
			if res.Value!=tt.expected{
				t.Error("Error!!! there is an error subtraction")
			}
		}
	}

//----------------------------------test for Multiply function------------------------------------
	func TestMultiply(t *testing.T)  {
		tests:= []struct{
			num1 bigint.Bigint
			num2 bigint.Bigint
			expected string
		} {
		{num1: bigint.Bigint{Value: "123"},num2: bigint.Bigint{Value: "01"},expected:"123"},
		{num1: bigint.Bigint{Value: "0123"},num2: bigint.Bigint{Value: "-01"},expected:"-123"},
		{num1: bigint.Bigint{Value: "-123"},num2: bigint.Bigint{Value: "01"},expected:"-123"},
		{num1: bigint.Bigint{Value: "-123"},num2: bigint.Bigint{Value: "-01"},expected:"123"},
		}
		for _,tt:=range tests{
			res:=bigint.Multiply(tt.num1,tt.num2)
			if res.Value!=tt.expected{
				t.Error("Error!!! there is an error multiply")
			}
		}
	}

//----------------------------------test for Mod function------------------------------------
	func TestMod(t *testing.T)  {
		tests:= []struct{
			num1 bigint.Bigint
			num2 bigint.Bigint
			expected string
		} {
		{num1: bigint.Bigint{Value: "123"},num2: bigint.Bigint{Value: "01"},expected:"0"},
		{num1: bigint.Bigint{Value: "0123"},num2: bigint.Bigint{Value: "-01"},expected:"0"},
		{num1: bigint.Bigint{Value: "-123"},num2: bigint.Bigint{Value: "015"},expected:"-3"},
		{num1: bigint.Bigint{Value: "-09"},num2: bigint.Bigint{Value: "-04"},expected:"-1"},
		}
		for _,tt:=range tests{
			res:=bigint.Mod(tt.num1,tt.num2)
			if res.Value!=tt.expected{
				t.Error("Error!!! there is an error modulo")
			}
		}
	}

//----------------------------------test for Abs function------------------------------------
	func TestAbs(t *testing.T){
		tests := []struct{
			num bigint.Bigint
			expected string
		}{
			{num:bigint.Bigint{Value: "123"}, expected: "123"},
			{num:bigint.Bigint{Value: "-56"}, expected: "56"},
			{num:bigint.Bigint{Value: "+1256"}, expected: "1256"},
			{num:bigint.Bigint{Value: "+0"}, expected: "0"},

		}
		for _, tt := range tests{
			res := tt.num.Abs()
			if res.Value != tt.expected{
				t.Errorf("Error!!! there is an error ABS")
			}
		}
	}


