package main
import "fmt"

func even_odd(i int) int {
	if i%2==0 {
		return 0
	} else {
		return 1
	}
}

func main(){
	z := even_odd(3)
	if z==1 {
		fmt.Println("ODD")
	} else 	{
		fmt.Println("EVEN")
	}
}
