package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadFile(path string) any {
	var err error
	plan, _ := os.ReadFile(path)
	var data []map[string]interface{}
	err = json.Unmarshal(plan, &data)

	if err != nil {
		panic("ReadFile Error")
	}

	fmt.Printf("ReadFile = %T: %s\n", data, data)

	return data
}
