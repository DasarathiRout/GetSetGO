package main

import (
	"fmt"
	"runtime"

	"rsc.io/quote"
)

// MainGo Comments
// (1) Sequential
// (2) Iterative
// (3) Conditional

func main() {
	fmt.Println("Lets Get Set ... GO !!!")
	//HelloGo()
	//QuoteMe()
	//callSET1()
	//callSET2()
	//callSET3()
	//callSET4()
	//callSET5()
	//callSET6()
	callSET7()
}

/*
HelloGo Returns
Block Comments Function
*/
func HelloGo() string {
	return "Hi,Dasarathi"
}

/*
QuoteMe Returns Function
*/
func QuoteMe() string {
	return quote.Hello()
}

func callSET1() {
	byteSize, errorMsg := fmt.Println("GO", "LANG", 1.14, true)
	fmt.Println(byteSize, errorMsg)
	xVal := 87 + 11 // Sort Decelaration Operator  VAR := VALUE
	fmt.Println(xVal)
	xVal = 89
	fmt.Println(xVal)
}

var xSize = 111

//  size := 111
//  syntax error: non-declaration statement outside function body

var zSize int
var message string
var isTrue bool

func callSET2() {
	var nSize = 99
	fmt.Println(nSize, message, zSize)
	fmt.Printf("%T\n", zSize)
	fmt.Printf("%T\n", message)
	fmt.Printf("%T\n", isTrue)

}

type drout int // User Defined Type
var numD drout
var numX int

func callSET3() {
	numX = 123
	numD = 1987
	fmt.Println(numX)
	fmt.Printf("%T\n", numX)
	fmt.Println(numD)
	fmt.Printf("%T\n", numD)

	// numX = numD
	numX = int(numD)

}

func callSET4() {
	fmt.Println(runtime.GOOS)
	msg := "Dasarathi@1987"

	fmt.Println(msg)
	fmt.Printf("%T\n", msg)

	byMsg := []byte(msg)
	fmt.Println(byMsg)
	fmt.Printf("%T\n", byMsg)
	fmt.Println("\n====================================")
	for i := 0; i < len(msg); i++ {
		fmt.Printf("%#U : ", msg[i])
	}
	fmt.Println("\n====================================")
	for index, value := range msg {
		fmt.Println(index, value)
	}
	fmt.Println("\n====================================")
}

// CON1 Constant EG
const CON1 = `User: Dasarathi Rout
Mobile: 9876543210
Email : dasarathi@go.com`

// USER MOOBILE EMAIL UNTYPE STATUS & TYPED CONST
const (
	USER        = "Dasarathi Rout"
	MOBILE      = 9876543210
	EMAIL       = "dasarathi.go.com"
	STATUS bool = true
)

// N1 N2 N3 IOTA
const (
	N1 = iota
	N2 = iota
	N3 = iota
)

func callSET5() {

	fmt.Println(CON1)
	fmt.Printf("%T\n", CON1)

	fmt.Println(USER)
	fmt.Printf("%T\n", USER)

	fmt.Println(MOBILE)
	fmt.Printf("%T\n", MOBILE)

	fmt.Println(EMAIL)
	fmt.Printf("%T\n", EMAIL)

	fmt.Println(STATUS)
	fmt.Printf("%T\n", STATUS)

	fmt.Println(N1, N2, N3)
}

func callSET6() {
	X2 := 2
	fmt.Printf("DECIMAL=%d\t| BINARY=%b\n", X2, X2)
	X4 := X2 << 1
	fmt.Printf("DECIMAL=%d\t| BINARY=%b\n", X4, X4)
	X8 := X4 << 1
	fmt.Printf("DECIMAL=%d\t| BINARY=%b\n", X8, X8)
	X2 = X8 >> 2
	fmt.Printf("DECIMAL=%d\t| BINARY=%b\n", X2, X2)
}

// SIZE KB MB
const (
	_ = iota // THROW AWAY VALUED 0
	//KB = 1024
	//MB = 1024 * KB

	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
	TB = 1 << (iota * 10)
)

func callSET7() {
	fmt.Printf("DECIMAL=%d\t\t| BINARY=%b\n", KB, KB)
	fmt.Printf("DECIMAL=%d\t\t| BINARY=%b\n", MB, MB)
	fmt.Printf("DECIMAL=%d\t| BINARY=%b\n", GB, GB)
	fmt.Printf("DECIMAL=%d\t| BINARY=%b\n", TB, TB)
	var X = 16
	fmt.Printf("DECIMAL=%d | BINARY=%b | HEX=%#x\n", X, X, X)
}
