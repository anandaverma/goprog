package main
import "fmt"
import "os"
import "log"
import "bufio"
func main() {
	file, err := os.Open("../level1/string.go") 
	if err != nil {
		log.Fatal("Error opening File", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if err:=scanner.Err(); err != nil {
		log.Fatal("Error Scanning File", scanner.Err())
	}
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
