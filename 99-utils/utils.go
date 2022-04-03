package utils

import (
	"encoding/json"
	"fmt"

	"github.com/TylerBrock/colorjson"
)

func PrettyPrint(i interface{}) {
	data, _ := json.MarshalIndent(i, "", "  ")
	fmt.Println(string(data))
}

func PrettyPrintColor(i interface{}) {
	f := colorjson.NewFormatter()
	f.Indent = 2
	s, _ := f.Marshal(i)
	fmt.Println(string(s))
}
