//Написать программу, которая перемножает, делит, складывает, вычитает 2 числовых переменных a,b, значение которые > 2^20.
package main

import "fmt"
import "math/big"

func main(){
	var a,b float64
	a = 19023901239125891259815
	b = 1902349234958583274123
	x:= big.NewInt(123912949012401204)
	y:= big.NewInt(12312312412595959)
	fmt.Println(a)
	fmt.Println(x)
	z:= z.Mul(a,b)
	fmt.Println(z)
	fmt.Println(a / b )
	fmt.Println(a * b )
	fmt.Printf("%f\n", a * b ) 
	fmt.Printf("%g\n", a / b)
}