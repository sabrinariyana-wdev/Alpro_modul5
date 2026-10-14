package main
import "fmt"

func main(){
	var a,b,c int
	var sensor bool

	fmt.Scan(&a, &b, &c)
	sensor = (a == 1 && b == 1) || (a == 1 && c == 1) || (b == 1 && c ==1)

	fmt.Print(sensor)
}