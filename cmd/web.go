package cmd

import (
	"../web"
	"fmt"
	"github.com/Unknwon/macaron"
	"github.com/codegangsta/cli"
	"net/http"
)

var CmdWeb = cli.Command{
	Name:        "web",
	Usage:       "start oct web service",
	Description: "oct is the module of handler docker repository and rkt image.",
	Action:      runWeb,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "address",
			Value: "0.0.0.0",
			Usage: "web service listen ip, default is 0.0.0.0; if listen with Unix Socket, the value is sock file path.",
		},
		cli.IntFlag{
			Name:  "port",
			Value: 80,
			Usage: "web service listen at port 80; if run with https will be 443.",
		},
	},
}

func runWeb(c *cli.Context) {
	m := macaron.New()

	web.SetWharfMacaron(m)

	listenaddr := fmt.Sprintf("%s:%d", c.String("address"), c.Int("port"))
	fmt.Println("Start to listen ", listenaddr)
	if err := http.ListenAndServe(listenaddr, m); err != nil {
		fmt.Printf("start oct http service error: %v", err.Error())
	}
}
