package main

import (
	"github.com/codegangsta/cli"
	"os"
	"strconv"
)

func Definecommand() {
	app := cli.NewApp()
	app.Name = "gopush-cluster-cli"
	app.Version = "1.0.0"
	app.Authors = []cli.Author{cli.Author{Name: "oldseven", Email: "xaseven@qq.com"}, cli.Author{Name: "codeguest"}}
	app.Commands = []cli.Command{
		{
			Name:        "admin",
			Usage:       "admin commands",
			Description: "This is how we describe admin the function",
			Subcommands: []cli.Command{
				{ //管理员推送单个信息
					Name:        "postprivmsg",
					Aliases:     []string{"pspm"},
					Usage:       "admin post a private msg",
					Description: "admin post a private msg",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "url",
							Value: "http://192.168.2.168:8091/1/admin/push/private",
							Usage: "admin push private msg",
						},
						cli.StringFlag{
							Name:  "key",
							Value: "seven",
							Usage: "msg key",
						},
						cli.StringFlag{
							Name:  "expire",
							Value: "600",
							Usage: "msg expire",
						},
						cli.StringFlag{
							Name:  "msg",
							Value: "",
							Usage: "msg body",
						},
					},
					Action: func(c *cli.Context) {
						//println("Hello, ", c.String("url"), c.String("key"), c.String("expire"))
						Adminpostmsg(c.String("url"), c.String("key"), c.String("expire"), c.String("msg"))
					},
				}, { //管理员推送多个信息
					Name:    "postmprivmsg",
					Aliases: []string{"pmpm"},
					Usage:   "admin post private msgs",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "url",
							Value: "http://192.168.2.168:8091/1/admin/push/mprivate",
							Usage: "admin push private msgs",
						},

						cli.StringFlag{
							Name:  "expire",
							Value: "600",
							Usage: "msg expire",
						},
						cli.StringFlag{
							Name:  "msg",
							Value: "",
							Usage: "msg and key",
						},
					},
					Action: func(c *cli.Context) {
						//println("Hello, ", c.String("expire"))
						Adminpostmmsg(c.String("url"), c.String("expire"), c.String("msg"))
					},
				}, { //管理员删除消息
					Name:    "delmsg",
					Aliases: []string{"del"},
					Usage:   "admin delete msgs",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "url",
							Value: "http://192.168.2.168:8091/1/admin/msg/del",
							Usage: "admin delete private msgs",
						},

						cli.StringFlag{
							Name:  "key",
							Value: "seven",
							Usage: "msg key",
						},
						cli.StringFlag{
							Name:  "msg",
							Value: "",
							Usage: "msg",
						},
					},
					Action: func(c *cli.Context) {
						//println("Hello, ", c.String("key"))
						Admindelmsg(c.String("url"), c.String("key"), c.String("msg"))
					},
				},
			},
		}, {
			Name:  "clicmd",
			Usage: "client commands",
			Subcommands: []cli.Command{
				{ //客户端获取订阅节点
					Name:    "getnode",
					Aliases: []string{"gn"},
					Usage:   "client get node",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "url",
							Value: "http://192.168.2.168:8090/1/server/get",
							Usage: "",
						},
						cli.StringFlag{
							Name:  "key",
							Value: "seven",
							Usage: "msg key",
						},
						cli.StringFlag{
							Name:  "p",
							Value: "2",
							Usage: "proto",
						},
					},
					Action: func(c *cli.Context) {
						//println("Hello, ", c.String("url"), c.String("key"), c.String("p"))
						Clientgetnode(c.String("url"), c.String("key"), c.String("p"))
					},
				}, { //客户端获取离线消息
					Name:    "getofflinemsg",
					Aliases: []string{"gom"},
					Usage:   "client get offline msgs",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "url",
							Value: "http://192.168.2.168:8090/1/msg/get",
							Usage: "client get offline msgs point",
						},
						cli.StringFlag{
							Name:  "key",
							Value: "seven",
							Usage: "msg expire",
						},
						cli.StringFlag{
							Name:  "mid",
							Value: "1",
							Usage: "last msgid",
						},
					},
					Action: func(c *cli.Context) {
						//println("Hello, ", c.String("key"))
						Clientofflinemsg(c.String("url"), c.String("key"), c.String("mid"))
					},
				},
				{ //客户端获取初始化消息ID
					Name:    "getofflinemsg",
					Aliases: []string{"gmi"},
					Usage:   "client get first msgid",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "url",
							Value: "http://192.168.2.168:8090/1/time/get",
							Usage: "client get first msgid point",
						},
					},
					Action: func(c *cli.Context) {
						//println("Hello, ", c.String("url"))
						Clientgetinitmsgid(c.String("url"))
					},
				},
			},
		}, {
			Name:    "startclient",
			Aliases: []string{"cli"},
			Usage:   "comet client",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "addr",
					Value: "192.168.2.168:6969",
					Usage: "comet client",
				},
				cli.StringFlag{
					Name:  "key",
					Value: "seven",
					Usage: "comet client",
				},
				cli.StringFlag{
					Name:  "heartbeat",
					Value: "30",
					Usage: "comet client",
				},
			},
			Action: func(c *cli.Context) {
				hstr := c.String("heartbeat")
				iheart, _ := strconv.Atoi(hstr)
				Cometclient(c.String("addr"), c.String("key"), iheart)
			},
		}, {
			Name:  "bye",
			Usage: "says goodbye",
			Action: func(c *cli.Context) {
				println("bye")
			},
		},
	}

	app.Run(os.Args)
}

func main() {
	Definecommand()
}
