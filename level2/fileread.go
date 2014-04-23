package main
import "fmt"
import "os"
import "log"
func main() {
	file, err := os.Open("../level1/string.go") 
	if err != nil {
		log.Fatal(err)
	}else {
		data:=make([]byte, 200)
		count, err := file.Read(data)
		if err !=nil {
			log.Fatal(err)
		}else {
			fmt.Printf("read %d bytes: %q\n", count, data[:count])
		}
	}
}
