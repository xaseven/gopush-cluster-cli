package main

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

var request = gorequest.New()

//----------------------客户端获取订阅节点-----------------------------------------
//ex: $./gopush-client clicmd gn -key aaa -p 2
func printgetnoderes(resp gorequest.Response, body string, errs []error) {
	//fmt.Println(resp.Status)
	//fmt.Println(resp.Body)
	fmt.Println(body)
}

//获取订阅节点
//param: url,
//       key:订阅key
//       p:int(i订阅协议 1-websocket 2-tcp)
func Clientgetnode(url string, key string, proto string) {
	tmpstr := url + "?k=" + key + "&p=" + proto
	request.Get(tmpstr).End(printgetnoderes)
}

//----------------------客户端获取离线消息-----------------------------------------
//ex: $ ./gopush-client clicmd gom -key aaa -mid 1
//获取离线消息
//param: url,
//       key:订阅key
//       mid:int64(最新接收的私有消息ID)
func Clientofflinemsg(url string, key string, mid string) {
	tmpstr := url + "?k=" + key + "&m=" + mid
	request.Get(tmpstr).End(func(resp gorequest.Response, body string, errs []error) {
		fmt.Println(body)
	})
}

//---------------------客户端获取初始化消息ID-------------------------------------
//ex: $./gopush-client clicmd gom -key aaa -mid 143015

//客户端获取初始化消息ID
//param url,
func Clientgetinitmsgid(url string) {
	request.Get(url).End(func(resp gorequest.Response, body string, errs []error) {
		fmt.Println(body)
	})
}

//===----------------------管理员推送单个信息----------------------------------------------
//ex : $  ./gopush-client admin pspm -key=seven -expire 600 -msg {\"test\":1}
//管理员推送单个信息
//
func Adminpostmsg(url, key, expire string, msg string) {
	tmpurl := url + "?key=" + key + "&expire=" + expire
	request.Post(tmpurl).Send(msg).End(func(resp gorequest.Response, body string, errs []error) {
		fmt.Println(body)
	})
}

//-------------------------管理员给多个key推送同一条消息信息---------------------------------------------------
//ex: $ ./gopush-client admin pmpm -expire 600 -msg "{\"m\":\"{\\\"test\\\":1}\",\"k\":\"t1,t2,t3\"}"
//管理员给多个key推送同一条消息信息
//
func Adminpostmmsg(url, expire string, msg string) {
	tmpurl := url + "?expire=" + expire
	request.Post(tmpurl).Send(msg).End(func(resp gorequest.Response, body string, errs []error) {
		fmt.Println(body)
	})
}

//-----------------------管理员删除消息-------------------------------------------------------
//
func Admindelmsg(url string, key string, msg string) {
	tmpurl := url + "?key=" + key
	request.Post(tmpurl).Send(msg).End(func(resp gorequest.Response, body string, errs []error) {
		fmt.Println(body)
	})
}
