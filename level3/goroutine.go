package main
import "fmt"
import "time"
func ping() {
	for i:=0; i<20; i++{
		fmt.Println("ping")
		time.Sleep(1)
	}
}
func pong() {
	for i:=0; i<20; i++{
		fmt.Println("pong")
		time.Sleep(1)
	}
}
func main() {
	go ping()
	go pong()
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
