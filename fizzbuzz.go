package main

import "fmt"


func main(){
	fmt.Println("FizzBuzz Start")
	for i := 0; i<100; i++{
		if i % 3 ==0 && i % 5 == 0{
			fmt.Println("FizzBuzz")
		}else if i % 3 == 0{
			fmt.Println("Fizz")
		}else if i % 5 == 0{
			fmt.Println("Buzz")
		}else{
			fmt.Println(i)
		}
	}
}
// インタラクティブなFizzBuzz(Qiitaより)
func CaseTypefizzbuzz() {

    i := 1
    for i < 101 {
        switch {
        case i%15 == 0:
            fmt.Println("FIZZ BUZZ!")
        case i%3 == 0:
            fmt.Println("FIZZ!")
        case i%5 == 0:
            fmt.Println("BUZZ!")
        default:
            fmt.Println(i)
        }

        i++
    }

}