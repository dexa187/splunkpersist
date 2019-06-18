package splunkpersist

import (
	"encoding/json"
	"fmt"
)

func LogError(err error) {
	out := Response{Payload{Entry: []Entry{Entry{Content: fmt.Sprintf(err.Error())}}}}
	outString, _ := json.Marshal(out)
	fmt.Println("0")
	fmt.Printf("%v\n", len(string(outString)))
	fmt.Printf("%v", string(outString))
}
