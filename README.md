# gopush-cluster-cli
gopush-cluster client for golang.
==============
[gopush-cluster服务器](https://github.com/Terry-Mao/gopush-cluster)
---------------------------------------
# 命令
##帮助
```sh
$ ./gopush-cluster-cli -h
$ ./gopush-cluster-cli h
$ ./gopush-cluster-cli -help
$ ./gopush-cluster-cli --help
```
##主命令帮助
```sh
$ ./gopush-cluster-cli admin -h
$ ./gopush-cluster-cli clicmd -h
$ ./gopush-cluster-cli cli -h
```

##子命令帮助
```sh
$ ./gopush-cluster-cli admin pspm -h
$ ./gopush-cluster-cli admin pmpm -h
...
$ ./gopush-cluster-cli clicmd gn -h
...
```

# Example
##客户端获取订阅节点
```sh
//获取节点，URL为代码中的URL key：seven 协议：2（Tcp）
$./gopush-cluster-cli clicmd gn -key seven -p 2 
$./gopush-cluster-cli clicmd gn -url http://192.168.2.168:8090/1/server/get -key seven -p 2
```
##客户端获取离线消息
```sh
//URL为代码中的URL key：seven 最新接收的私有消息ID：1
$./gopush-cluster-cli clicmd gom -key seven -mid 1 
$./gopush-cluster-cli clicmd gom -url http://192.168.2.168:8090/1/msg/get -key seven -mid 1
```
##客户端获取初始化消息ID
```sh
$./gopush-cluster-cli clicmd gmi 
$./gopush-cluster-cli clicmd gmi -url http://192.168.2.168:8090/1/time/get
```
##管理员推送单个信息
```sh
$./gopush-cluster-cli admin pspm -key seven -expire 600 -msg {\"test\":1}
$./gopush-cluster-cli admin pspm -url http://192.168.2.168:8091/1/admin/push/private -key seven -expire 600 -msg {\"test\":1}
```
##管理员给多个key推送同一条消息信息
```sh
$./gopush-cluster-cli admin pmpm -expire 600 -msg "{\"m\":\"{\\\"test\\\":1}\",\"k\":\"seven,testkey1,testkey2\"}"
$./gopush-cluster-cli admin pmpm -url http://192.168.2.168:8091/1/admin/push/mprivate -expire 600 -msg "{\"m\":\"{\\\"test\\\":1}\",\"k\":\"seven,testkey1,testkey2\"}"
```
##启动客户端
```sh
$./gopush-cluster-cli cli -addr 192.168.2.168:6969 -key seven -heartbeat 30
$./gopush-cluster-cli startclient -key seven -heartbeat 30
$./gopush-cluster-cli cli 
$./gopush-cluster-cli startclient
```
