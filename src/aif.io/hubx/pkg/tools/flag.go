package tools

import "github.com/spf13/pflag"
import (
	"log"
	"strconv"
	"fmt"
)

func PrintFlags(flags *pflag.FlagSet) {
	flags.VisitAll(func(flag *pflag.Flag) {
		log.Printf("FLAG: --%s=%q", flag.Name, flag.Value)
	})
}

func ListenPortWrapper(port uint64) string {
	return fmt.Sprint(":",strconv.FormatUint(port,10))
}