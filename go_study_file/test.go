package main 
import "fmt"

func main() {  
    var a1 = 123
    var a2 = "qwe"
    var a3 = "123%dqwe%s"
    var a4 = fmt.Sprintf(a3, a1, a2)
    fmt.Printf(a4)
    fmt.Printf("Hello, world!\n");

}
