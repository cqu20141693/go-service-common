### service common
#### 架构规范
1. event 能力必须由boot包下的启动方法进行开启和执行
2. 每个包需要利用event能力通过init函数注册对应事件的执行函数
3. boot包下的方法应改在main函数中调用，这样就能

#### test
``` 
go test -v -run TestLog cclog_test.go

```
#### 问题处理
1. 10054
``` 
git config --global http.sslVerify "false" 
```
2. 443
```
git config --global --unset https.proxy
```
