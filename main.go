package findprocess

import (
	"fmt"

	"github.com/shirou/gopsutil/process"
)

func Pids() {
	procs, err := process.Processes()
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf(`procs type: %T`, procs))
	fmt.Println(fmt.Sprintf(`procs qty: %d`, len(procs)))
}
