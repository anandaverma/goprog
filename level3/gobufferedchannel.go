package main
import "fmt"
func count(ch chan int) {
	for i:=0; i<10; i++ {
		fmt.Println("pushing ", i)
		ch <- i
	}
}
func main() {
	ch := make(chan int, 2)
	go count(ch)
	for {
		j:= <- ch
		if j == 9 {
			break
		}
		fmt.Println("reading ", j)
	}	
}
