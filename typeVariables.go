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