package main
import "fmt"
const (
	Big = 1<<100
	Small = Big>>99
	)

func main(){
	//var i int
	i:=Small
	i += 4
	for ;i<20;i++ {
		fmt.Println(i)
	}

	for i>0 {				//For is the while
		fmt.Println("Inside While",i)
		i--
	}
}
