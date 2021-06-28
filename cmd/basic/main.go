package main

import (
	"encoding/json"
	"fmt"

	findprocess "github.com/binRick/go-find-process"
	"github.com/k0kubun/pp"
)

func main() {
	//test_pids()
	test_envkey(`CONTAINER_HOSTNAME`)

}
func test_envkey(k string) {
	pid_wgcs_names := map[int64]string{}
	pids, vals, dur, err := findprocess.EnvKey(k)
	if err != nil {
		panic(err)
	}
	for i, val := range vals {
		pid_wgcs_names[pids[i]] = val.(string)
	}
	if false {
		fmt.Println(
			fmt.Sprintf(`%d pids`, len(pids)),
			fmt.Sprintf(`%d vals`, len(vals)),
			fmt.Sprintf(`%s`, dur),
			pp.Sprintf(`%s`, vals),
			pp.Sprintf(`%s`, pid_wgcs_names),
		)
	}
	pid_wgcs_names_bytes, err := json.Marshal(pid_wgcs_names)
	if err != nil {
		panic(err)
	}
	pid_wgcs_names_str := string(pid_wgcs_names_bytes)

	fmt.Println(pid_wgcs_names_str)
}

func test_pids() {
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
		if false {
			fmt.Println(
				pp.Sprintf(`%s`, V),
				pp.Sprintf(`V qty: %d`, len(V)),
			)
		}
	}

	msg := fmt.Sprintf(`Acquired %d Pids, %d Values in %s`, len(pids), len(vals), dur)
	pp.Println(msg)
}
