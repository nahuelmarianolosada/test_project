package main

import "fmt"

func main() {

	for i := 0; i < 4; i++ {
		defer showMeTheVariable(i)
	}
}

func showMeTheVariable(i int) {
	fmt.Print(i)
}

//vars := []interface{}{"brand"}
/*json := map[string]interface{}{"tags": nil}

result, err := gval.Evaluate(`tags  null`, json)
if err != nil {
	fmt.Println(err)
}

v, ok := result.(bool)

if !ok {
	fmt.Errorf("No se pudo parser el resultado de gval")
}
fmt.Print(v)*/

/*v := interface{}(nil)
json.Unmarshal([]byte(`{
	"foo": {
		"bar": {
			"status":null
		}
	}
}`), &v)

value, err := gval.Evaluate(`foo.bar != null && foo.bar.status == null`, v)
if err != nil {
	fmt.Println(err)
}

fmt.Print(value)*/

/* v := interface{}(nil)
json.Unmarshal([]byte(`{
	"question":{
		"answer":{
			"status":"closed",
			"text":""
		}
	}
}`), &v)
//value2, err2 := gval.Evaluate(`question.status in ["closed", "open", "pending"]`,v, jsonpath.Language())
value2, err2 := gval.Evaluate(`question.answer != null && (question.answer.status in ["closed", "open", "pending"]) && question.answer.text == ""`, v, jsonpath.Language())
if err2 != nil {
	fmt.Println(err2)
}
fmt.Print(value2) */
//}
