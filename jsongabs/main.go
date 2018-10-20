package main

import (
	"github.com/Jeffail/gabs"
	"fmt"
	//"reflect"
)

func main(){
	str := `
{
  "outter": {
    "inner": {
      "value": 10,
      "value2": 20
    },
    "inner2": {
      "value3": 30
    }
  }
}`

	jsonParsed, _ := gabs.ParseJSON([]byte(str))
	fmt.Println(jsonParsed.Path("outter.inner").String())

	//s := []string {"a","b"}

	//fmt.Println(reflect.TypeOf(s).String() == "[]string" )
	//jsonParsed.SetP(s,"outter.inner.amdhu")
	//fmt.Println(jsonParsed)

}