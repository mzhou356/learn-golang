package os_package

import (
	"fmt"
	"os"
)

func GetTerminalArgs() {
	fmt.Printf("%#v\n", os.Args)
}
