package hashtable

import "fmt"

func Set() {
	Data := make(map[string]bool)

	Data["apple"] = true
	Data["banana"] = true
	Data["apple"] = true
	Data["orange"] = true

	for val := range Data {
		fmt.Println(val)
	}
}
