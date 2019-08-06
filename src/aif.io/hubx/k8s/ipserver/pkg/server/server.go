package server

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/spf13/cobra"
	"aif.io/hubx/pkg/tools"
	"aif.io/hubx/pkg/root"
	"log"
)
var(
	stop        = make(chan struct{})
	serverPort string
	ipCmd      = &cobra.Command{
		Use:   "ipserver",
		Short: "ipserver service",
		RunE: func(c *cobra.Command, args []string) error {

			tools.PrintFlags(c.Flags())

			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "当前 IP：",strings.Split(r.RemoteAddr,":")[0],"  来自于：中国 四川 成都  电信")
			})

			if err:=http.ListenAndServe("0.0.0.0:"+serverPort, nil) ; err != nil {
				log.Println(err)
				return err
			}
			<-stop
			return nil
		},
	}
)


func init(){
	ipCmd.PersistentFlags().StringVar(&serverPort, "port", "","ip server's port")
	root.AddCmd(ipCmd)
}