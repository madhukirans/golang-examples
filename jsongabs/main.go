package main

import (
	"github.com/Jeffail/gabs"
	"fmt"
	//"reflect"
)

type AAA  struct {
	monitors []string
}

func main(){
	str := `
{"apiVersion":"sauron.oracledx.com/v1","items":[{"apiVersion":"sauron.oracledx.com/v1","kind":"SauronBackup","metadata":{"clusterName":"","creationTimestamp":"2018-10-31T19:19:47Z","name":"test-gbg-sauron-data","namespace":"sauron-1","resourceVersion":"27669482","selfLink":"/apis/sauron.oracledx.com/v1/namespaces/sauron-1/sauronbackups/test-gbg-sauron-data","uid":"ee016feb-dd41-11e8-85b5-0200170240e9"},"spec":{"bucket":"backupmadhu","envName":"sauron-operator","id":"test-gbg-sauron-data-test-gbg-backup","namespace":"odx-sre","sauronBackupPolicyName":"test-gbg-sauron-data","sauronName":"test-gbg-sauron-data","sauronNamespace":"sauron-1"}}],"kind":"SauronBackupList","metadata":{"continue":"","resourceVersion":"27669482","selfLink":"/apis/sauron.oracledx.com/v1/namespaces/sauron-1/sauronbackups"}}
`

	jsonParsed, _ := gabs.ParseJSON([]byte(str))
	a, _:= jsonParsed.Path("items").Children()
	fmt.Println(a[0])
		//jsonParsed.Path("spec.grafana").Children(),
		//jsonParsed.Path("spec.prometheus").Children())
	//
	//c , _ := jsonParsed.Path("outter.inner").Children()
	//fmt.Println(len(c))//.([]string))
	//fmt.Println(c[0].Path("value"))
	//m := new (AAA)
	//m.monitors = append(m.monitors, []string (jsonParsed.Path("outter.inner.value")))
	//s := []string {"a","b"}

	//fmt.Println(reflect.TypeOf(s).String() == "[]string" )
	//jsonParsed.SetP(s,"outter.inner.amdhu")
	//fmt.Println(jsonParsed)

}