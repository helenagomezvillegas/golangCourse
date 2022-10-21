# golangCourse

# Table of contents

- [Install go from https://go.dev/doc/install](#install-go-from-httpsgodevdocinstall)
- [Hello World!](#hello-world)
- [Single and Multiline comment](#single-and-multiline-comment)
- [Types and Variables](#types-and-variables)
  - [Number](#number)
  - [String](#string)
  - [Boolean](#boolean)
  - [Print variarables with format](#print-variarables-with-format)
  - [Local and Global variables](#local-and-global-variables)

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

package main
import "fmt"
```
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
### Print variarables with format
![Printf- Format specifiers ](/img/fmtprint.png "Table formats")
### Local and Global variables







