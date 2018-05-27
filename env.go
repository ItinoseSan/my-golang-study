package main 
import "os"
import "fmt"

func main(){
	testEnvValue := os.Getenv("HOGE_ENV")
	testEnvValue2 := os.Getenv("PIYO_ENV")
	fmt.Println("HOGE_ENV:",testEnvValue)
	fmt.Println("PIYO_ENV:",testEnvValue2)
	// get env in array
   // fmt.Println(os.Environ());
}