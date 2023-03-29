package ospackage

import (
	"fmt"
	"os"
)

func GetTerminalArgs() {
	fmt.Printf("%#v\n", os.Args)
}
