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