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
        serverIp string
	ipCmd      = &cobra.Command{
		Use:   "ipserver",
		Short: "ipserver service",
		RunE: func(c *cobra.Command, args []string) error {

			tools.PrintFlags(c.Flags())

			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				clientIp:=strings.Split(r.RemoteAddr,":")[0]
				fmt.Println(clientIp)
				fmt.Fprintln(w, "当前 IP：",clientIp,"  来自于：中国 四川 成都  电信")
			})

			if err:=http.ListenAndServe(serverIp+":"+serverPort, nil) ; err != nil {
				log.Println(err)
				return err
			}
			<-stop
			return nil
		},
	}
)


func init(){
	ipCmd.PersistentFlags().StringVar(&serverPort, "port", "8000","ip server's port")
	ipCmd.PersistentFlags().StringVar(&serverIp, "ip", "0.0.0.0","ip server's ip")
	root.AddCmd(ipCmd)
}
