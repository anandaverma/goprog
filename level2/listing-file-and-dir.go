package main
import "fmt"
import "flag"
import "os"
import "path/filepath"
func fileVisit(path string, f os.FileInfo, _ error) error {
	if (f.IsDir()){
		fmt.Println("Directory: ", path)
	}else {
		fmt.Println("File: ", path)
	}
	return nil
}
func main() {
	flag.Parse()
	if (flag.NArg() != 1) {
		fmt.Println("go run " + os.Args[0] + " <file/directory path>")	
		os.Exit(255)
	}		
	dir := flag.Arg(0)
	fmt.Println(dir)
	filepath.Walk(dir, fileVisit)
}
