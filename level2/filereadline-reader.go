package main
import "fmt"
import "os"
import "log"
import "bufio"
import "io"
func main (){
	file, err:= os.Open("../level1/string.go");
	if err != nil {
		log.Fatal("Error opening file", err)
	}	
	defer file.Close()
	bf := bufio.NewReader(file)
	for i := 0;;i++ {
		line, isPrefix, err:= bf.ReadLine()
		if err == io.EOF {
			break
		}
		if isPrefix {
			log.Fatal("Error: unexpected long line", file.Name)
		}	
		fmt.Println("line:",i," Text:", string(line))
	}
}
