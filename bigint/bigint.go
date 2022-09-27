package bigint

import (
	"errors"
	 "fmt"
	"strconv"
	"strings"
)

type Bigint struct {
	Value string
}

var(
	ErrorNumber = errors.New("input not a number")
	ErrorLength = errors.New("input not a int32 or ")

)
//----------------------validation------------------------------------------
func validation(num string)  (string , error) {
	
	allowed := "1234567890"
	var xato bool

	if strings.HasPrefix(num, "-") {
		num = strings.Replace(num, "-", "", 1)
	}
	// for strings.HasPrefix(num, "+") {
	// 	num = strings.TrimPrefix(num, "+")
	//   }
	if strings.HasPrefix(num, "+") {
		num = strings.Replace(num, "+", "", 1)
	}

	//sonlikka tekshirish
	arr := strings.Split(num, "")
	for _, v := range arr {
		if !strings.Contains(allowed, v) {
			xato = true
		}
	}
	if xato {
		return "", ErrorNumber
	}
		//check num int 32
		_, err := strconv.ParseInt(num, 10, 32)
	
		if err != nil {
			panic(ErrorLength)
		}

	return num , nil
}
//---------------------------------clean(tozalash)-------------------------
func clean(num string) string{
	sign := ""
	for strings.HasPrefix(num, "0") {
		num = strings.TrimPrefix(num, "0")
	}
	if strings.HasPrefix(num,"-"){
		sign = "-"
		num=num[1:]
	}
	
	// if strings.HasPrefix(num, "+"){
	// 	num=num[1:]
	// }
	
	return sign + num
}

//--------------new int----------------------------
func NewInt(num string) (Bigint,error){

	n, err:=validation(num)
	if err != nil{
		return Bigint{}, err
	}
	

	myint,err:= strconv.Atoi(num)
	if err==nil {
	  fmt.Println(myint)
	} else {
	  fmt.Println("Xatolik bor!",err)
	}
  
	return Bigint{
	  Value: clean(n),
	},nil
	
}

//set()...
// func (z *Bigint) Set(num string) error {
// 	err:=validation(num)
// 	if err != nil{
// 		return err
// 	}
  
// 	z.Value = clean(num) 

// 	return nil
// }

//-------------------Add two number(2 soni yig'indisini hisoblash)-------------------------
func Add(num1,num2 Bigint) Bigint{

	a,_ := strconv.ParseInt(num1.Value, 10, 32)
	// if err != nil {
	// 	panic(err)
	// }
	b, _ := strconv.ParseInt(num2.Value, 10, 32)
	// if  err != nil {
	// 	panic(err)
	// }

	c :=(a+b) 

	return Bigint{
		Value: strconv.Itoa(int(c)),
	}
	
}

//-------------------------sub two numbers(2 sonni ayirmasini hisoblash)----------------------
func Sub(num1,num2 Bigint) Bigint {
	a,_ := strconv.ParseInt(num1.Value, 10, 32)
	b, _ := strconv.ParseInt(num2.Value, 10, 32)

	d:=a-b
	
	return Bigint{
		Value: strconv.Itoa(int(d)),
	}
}

//--------------Multiply two numbers(2 sonni ko'paytmani hisoblash)---------------------------
func Multiply(num1,num2 Bigint) Bigint {
	a,_ := strconv.ParseInt(num1.Value, 10, 32)
	b, _ := strconv.ParseInt(num2.Value, 10, 32)

	e:=a*b
	return Bigint{
		Value: strconv.Itoa(int(e)),
	}
}

//--------------mod two numbers (2 sonni qoldig'ini olish)------------------------------
func Mod(num1,num2 Bigint) Bigint {
	a,_ := strconv.ParseInt(num1.Value, 10, 32)
	// if err != nil {
	// 	panic(err)
	// }
	b, _ := strconv.ParseInt(num2.Value, 10, 32)
	// if  err != nil {
	// 	panic(err)
	// }
	f:=a%b
	return Bigint{
		Value: strconv.Itoa(int(f)),
	}
}

//------------------------abs two numbers(sonni modulini olish)-------------------------------
func (z Bigint) Abs() Bigint {
	if z.Value[0] == '-'{
		return Bigint{
			Value: z.Value[1:],
		}
	}
	if z.Value[0] =='+' {
		return Bigint{
			Value: z.Value[1:],
		}
	}
// 	for  z.Value[0] =='0'{
// 		return Bigint{
// 			Value: z.Value[1:],
// 	}
// }
	
	return Bigint{
		Value: z.Value,
	}
}
