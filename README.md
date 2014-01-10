### Irc

Robust and flexible IRC library for Go.

### Usage

```go
irc := NewIrcClient()
irc.Nick = "nick"
irc.Host = "irc.freenode.net"
irc.CallBack = handle
irc.Channel = "#testGoLangIrc"
irc.Join()

func handle(irc *IRCClient, message string) {
    //
    // Simple echo
    //
    irc.SendMessage(message)
}
```

SSL example:

```go
irc := NewIrcClient()
irc.Nick = "nick"
irc.Host = "irc.freenode.net"
irc.CallBack = handle
irc.Channel = "#testGoLangIrc"
irc.SSL = true
irc.SslCert = "cert.pem"
irc.SslKey  = "key.pem"
irc.Join()

func handle(irc *IRCClient, message string) {
    //
    // Simple echo
    //
    irc.SendMessage(message)
}
```

### API

[irc at godoc](http://godoc.org/github.com/Bullet-Chat/irc)

### TODO

 1. Add error handling
 3. Think about concurrency
 4. Add real name

### Contribution

 1. Fork [irc](https://github.com/Bullet-Chat/irc);
 2. Make changes;
 3. Send pull request;
 4. Thank you!

### Author

[@0xAX](https://twitter.com/0xAX) and [@_voidPirate](https://twitter.com/_voidPirate)
