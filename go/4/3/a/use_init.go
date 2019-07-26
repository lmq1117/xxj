package main

//fail
import "fmt"
import "./transs"

var TwoPi = 2 * transs.Pi

func main(){
	fmt.Printf("2 * Pi = %g\n",TwoPi)
}
