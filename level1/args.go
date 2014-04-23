package main
import "os"
import "fmt"
func main() {
	tot := len(os.Args)
	fmt.Printf("Total Args: %d\n", tot)
	for i:=0; i<tot; i++ {
		fmt.Println(i, os.Args[i])
	}
}
