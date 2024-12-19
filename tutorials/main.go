package main
import "fmt"
import "unicode/utf8"
const constant=111;
func main(){
	fmt.Println("hey this is my first go program!")	

	// values
	fmt.Println(true && false)
    fmt.Println(true || false)

	// var
	var a int  //unintialized, will print 0
	fmt.Println("the value of a is", a)
	
	var b=1
	fmt.Println("the value of b is", b)

	// const - declares a constant value
	fmt.Println("the value of constant is",constant);

	// string
	var my_string="Hello"
	var my_modified_string=`modified                string`

	fmt.Println(my_string)
	fmt.Println(my_modified_string)
	fmt.Println("length of modified string (lenght in form of bytes): ", len(my_modified_string))
	fmt.Println("length of  string-(Divya) (lenght in form of charaacters): ", utf8.RuneCountInString("Divya"))

	var itsBoolean bool = true
	fmt.Println("boolean--",itsBoolean)



	// Default values-
	// for all ints (uint/int) = 0
	// for string = ''
	// boolean = false


}
