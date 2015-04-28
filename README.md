# gopush-cluster-cli
gopush-cluster client for golang.
## 命令
#帮助
$ ./gopush-client -h
$ ./gopush-client h
$ ./gopush-client -help
$ ./gopush-client --help

#主命令帮助
$ ./gopush-client admin -h
$ ./gopush-client clicmd -h
$ ./gopush-client cli -h


#子命令帮助
$ ./gopush-client admin pspm -h
$ ./gopush-client admin pmpm -h
...
$ ./gopush-client clicmd gn -h
...

## Example
#客户端获取订阅节点
$./gopush-client clicmd gn -key seven -p 2 //获取节点，URL为代码中的URL key：seven 协议：2（Tcp）
$./gopush-client clicmd gn -url http://192.168.2.168:8090/1/server/get -key seven -p 2

#客户端获取离线消息
$./gopush-client clicmd gom -key seven -mid 1 //URL为代码中的URL key：seven 最新接收的私有消息ID：1
$./gopush-client clicmd gom -url http://192.168.2.168:8090/1/msg/get -key seven -mid 1

#客户端获取初始化消息ID
$./gopush-client clicmd gmi  
$./gopush-client clicmd gmi -url http://192.168.2.168:8090/1/time/get

#管理员推送单个信息
$./gopush-client admin pspm -key seven -expire 600 -msg {\"test\":1}
$./gopush-client admin pspm -url http://192.168.2.168:8091/1/admin/push/private -key seven -expire 600 -msg {\"test\":1}

#管理员给多个key推送同一条消息信息
$./gopush-client admin pmpm -expire 600 -msg "{\"m\":\"{\\\"test\\\":1}\",\"k\":\"seven,testkey1,testkey2\"}"
$./gopush-client admin pmpm -url http://192.168.2.168:8091/1/admin/push/mprivate -expire 600 -msg "{\"m\":\"{\\\"test\\\":1}\",\"k\":\"seven,testkey1,testkey2\"}"

#启动客户端
$./gopush-client cli -addr 192.168.2.168:6969 -key seven -heartbeat 30
$./gopush-client startclient -key seven -heartbeat 30
$./gopush-client cli 
$./gopush-client startclient
