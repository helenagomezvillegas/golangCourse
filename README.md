# golangCourse

# Table of contents

## Install go from https://go.dev/doc/install
1. Go Download
package for (Mac, Linux or Windows)
2. Go Install
* Open the package file you downloaded and follow the prompts to install Go.
The package installs the Go distribution to /usr/local/go. The package should put the /usr/local/go/bin directory in your PATH environment variable. You may need to restart any open Terminal sessions for the change to take effect.
* Verify that you've installed Go by opening a command prompt and typing the following command:
   > go version
* Confirm that the command prints the installed version of Go.

## Hello World!
```
package main
import "fmt"
func main() {

	fmt.Println("Hello World!")
}

```
run goland using this comand on terminal
> go run main.go

## Single and Multiline comment
```
package main
import "fmt"
// Single comment

/*Multiline
 comment
 */
func main() {

        fmt.Println("Hello World!")
}

```

## Types and Variables
Asigne a variable in go


```
package main
import "fmt"

func main() {

	title := "Sir"
	fmt.Println(title)
}
```
### Number
#### Integer
> int
* The numeric data type of int has several versions, which include:
int8
int16
int32
int64
uint8
uint16
uint32
uint64
The data types starting from int store signed integers while those starting with uint contain unsigned integers, and the numeric value that follows each data type represents the number of bytes that is stored.package main
```
import "fmt"

func main() {
   var num int
   num = 20
   fmt.Printf("Data type of %d is %T\n", num, num);

}
```
#### Float
> float
The FloatType identifier in Golang serves as a stand-in for either in-built float data type float32 or float64. Both the float32 and float64 data types represent numbers with decimal components.

The float32 data type occupies 32
32
 bits in memory, whereas the float64 data type occupies 64
64
 bits in memory. The default type for floating-point numbers is float64.

The command below shows how to declare a floating-point number in Golang:
> var a float32 = 2.25
> var b float64 = 2.50
> var c = 3.20
### String
string in Golang is a set of all strings that contain 8-bit bytes. By default, strings in Golang are UTF-8 encoded.

Variable of type string is enclosed between double-quotes. The type string variable’s value is immutable.
```
package main
import "fmt"

func main() {
   var sentence string
   sentence = "Hello From Everybody"
   fmt.Printf("Sentence : %s\n", sentence);
   fmt.Printf("Data type of sentence is %T\n", sentence);

}
```
### Boolean
Boolean values are those which can be assigned true or false and has the type bool with it.

```
package main

import (
    "fmt"
)

func main() {
    var bVal bool   // default is false
    fmt.Printf("bVal: %v\n", bVal)
}
```
## Print variarables with format
![Printf- Format specifiers ](/img/fmtprint.png "Table formats")
## Local variable Global variables
### Example  Local and Global variable
```
package main

import "fmt"
// This is a Global variable
var name string = "Lisa"
func main() {
   This is a Local variable
   var  lastname  string = "Patterson"
   fmt.Println(name);

}
```
## Zero Values
![Printf- Format specifiers ](/img/zeroValues.png "Table formats")

### Example:
```
package main

import "fmt"

func main() {
   var num0 int
   var num1 float64
   fmt.Printf("Data type of %d is %T\n", num0, num0);
   fmt.Printf("Data type of %.2f is %T\n", num1, num1);

}
```
> go run zeroValues.go
> Data type of 0 is int
> Data type of 0.00 is float64

see more in this table
![Printf- Format specifiers ](/img/values.png "Table formats")

##  User input
### Single Input
```
package main
import "fmt"

func main() {
   var name string
   // single input
   fmt.Print("Enter your name");
   fmt.Scanf("%s", &name);

}
```
### Multiple Input
```
package main
import "fmt"

func main() {
   var name string
   var is_muggle bool

   // single input
   fmt.Print("Enter your name and your muggle");
   fmt.Scanf("%s %t", &name, &is_muggle);
   fmt.Println(name, is_muggle);

}
```
## Find type of variable
Variable is a placeholder of the information which can be changed at runtime. And variables allow to Retrieve and Manipulate the stored information.
There are 3 ways to find the type of variables in Golang as follows:

* Using reflect.TypeOf Function
* Using reflect.ValueOf.Kind() Function
* Using %T with Printf
```
package main
import "fmt"
import "reflect"

func main() {
   var grades int = 42
   var message  string = "Hello"
   var checkin bool = true
   var amount float64 = 0.9

   fmt.Printf("variables grades = %v is of type %T \n", grades, grades)
   fmt.Printf("variables message = %v is of type %T \n", message, message)
   fmt.Printf("variables checkin = %v is of type %T \n", checkin, checkin)
   fmt.Printf("variables amount = %v is of type %T \n", amount, amount)

   // using reflect.Type
   fmt.Printf("Type: %v \n", reflect.TypeOf(1000))
   fmt.Printf("Type: %v \n", reflect.TypeOf("Maria"))

}
```

## Converting Between types










