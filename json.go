package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	simpleJson "github.com/bitly/go-simplejson" // jsonをパースするライブラリ
)

func main() {
	url := "http://api.openweathermap.org/data/2.5/weather?q=API_KEY"
  
	resp, _ := http.Get(url)
	// defer 文を使うことで、それを呼び出した関数が終了する際に処理を実行できる
	defer resp.Body.Close()
  
	json, _ := ioutil.ReadAll(resp.Body)
	
	js, err := simpleJson.NewJson([]byte(json))

	if err != nil{
		panic(err)
	}
	weather, err := js.Get("weather").GetIndex(0).Get("main").String()
	
	if err != nil{
		panic(err)
	}
	weatherCountry, err := js.Get("sys").Get("country").String()
	
	if err != nil{
		panic(err)
	}
	
	fmt.Println("CurrentWeather:",weather)
	fmt.Println("CountryOfCurrentWeather:",weatherCountry)
    // result => CurrentWeather: Clouds
	// CountryOfCurrentWeather: JP

	// fmt.Println(string(byteArray)) // jsonをstringで取得
}