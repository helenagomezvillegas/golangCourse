# golangCourse
## Install go from https://go.dev/doc/install
1. Go Download
Download Go package for (Mac, Linux or Windows)
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



