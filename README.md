# golangCourse
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
### Boolean
### String




