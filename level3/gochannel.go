package main
import "fmt"
func count(ch chan int) {
	var i int
	for {
		fmt.Scan(&i)
		ch <- i
	}
}
func main() {
	ch:= make(chan int)
	go count(ch)
	for {	
		j := <- ch
		if j < 1 {
			break
		}
		fmt.Println("value ",j)
	}
}
