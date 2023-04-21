package main

import (
    "fmt"
)

func main() {
    var num1, num2 float64
    var op string
    
    fmt.Println("Inserte el primer número:")
    fmt.Scanln(&num1)

    fmt.Println("Inserte la operación a realizar (+, -, *, /):")
    fmt.Scanln(&op)
    
    fmt.Println("Inserte el segundo número:")
    fmt.Scanln(&num2)
    
    switch op {
        case "+":
            fmt.Printf("La suma de %.2f y %.2f es: %.2f", num1, num2, num1+num2)
        case "-":
            fmt.Printf("La resta de %.2f y %.2f es: %.2f", num1, num2, num1-num2)
        case "*":
            fmt.Printf("La multiplicación de %.2f y %.2f es: %.2f", num1, num2, num1*num2)
        case "/":
            if num2 != 0 {
                fmt.Printf("La división de %.2f y %.2f es: %.2f", num1, num2, num1/num2)
            } else {
                fmt.Println("No se puede dividir entre cero")
            }
        default:
            fmt.Println("Operación no válida")
    }
}