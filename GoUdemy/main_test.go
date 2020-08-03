package main

import "testing"

/*
HelloGoTest Returns Block Comments
Must Have File_test.go & Start with TestFunction
*/
func TestHelloGo(tester *testing.T) {
	actualResult := HelloGo()
	var expectedResult = "Hi,Dasarathi"
	if actualResult != expectedResult {
		tester.Fatalf("Expected %s But Got %s", expectedResult, actualResult)
	}
}
