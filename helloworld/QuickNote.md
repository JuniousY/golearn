## 基础语法


## 多线程

### rpc
调用方法的格式
```go
func (t *T) MethodName(argType T1, replyType *T2) error
```

client方法 
- Dial, DialHttp   建立连接，返回client
- Call, Go  同步、异步调用
- Close

server方法
- Accept 一般异步调用。参数为net.Listener
- Register, RegisterName 注册方法
- ServeCodec, ServeConn一般异步调用，阻塞等待连接，为client做处理