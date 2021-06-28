package main

import (
	"encoding/json"
	"fmt"

	findprocess "github.com/binRick/go-find-process"
	"github.com/k0kubun/pp"
)

func main() {
	pids, vals, dur, err := findprocess.Pids()
	if err != nil {
		panic(err)
	}

	for _, val := range vals {
		V := map[string]interface{}{}
		val_err := json.Unmarshal([]byte(val.(string)), &V)
		if val_err != nil {
			panic(val_err)
		}

		fmt.Println(
			pp.Sprintf(`%s`, V),
		)
	}
	msg := fmt.Sprintf(`Acquired %d Pids, %d Values in %s`, len(pids), len(vals), dur)
	pp.Println(msg)
}
