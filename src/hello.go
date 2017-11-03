package main
import "fmt"

func main() {
    var input float64
    var x float64

    fmt.Print("Enter a number:\n")
    fmt.Scanf("%f", &input)
    x = input
    fmt.Printf("You entered this: %g\n", x)
}
