package main
import "fmt"
import "regexp"
import "os"
func main() {
	r, err := regexp.Compile(`hello`)
	if err != nil {
		fmt.Println("regex failed!")
		os.Exit(255)
	}
	if r.MatchString("hello world") == true {
		fmt.Println("match")
	}else {
		fmt.Println("no match")
	}
}
