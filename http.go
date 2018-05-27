package main

import (
  "net/http"
  "fmt"
  "io/ioutil"
)

func main() {
  url := "http://example.com"

  resp, _ := http.Get(url)
  // defer 文を使うことで、それを呼び出した関数が終了する際に処理を実行できる
  defer resp.Body.Close()

  byteArray, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(string(byteArray)) // htmlをstringで取得
}