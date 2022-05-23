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

## 实用库

### strings
- func FieldsFunc(s string, f func(rune) bool) []string 按方法f拆分字符串

### unicode
- IsLetter, IsNumber, IsSpace

### strconv
- Itoa int -> String
- Atoi String -> int
- ParseBool, ParseFloat, ParseInt, ParseUint  从String转