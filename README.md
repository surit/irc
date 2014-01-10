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

func handle(irc *IrcClient, message string) {
    //
    // Simple echo
    //
    irc.SendMessage(message)
}

```

### TODO

 1. Add error handling
 2  Add password checking
 3. Think about concurrency
 4. Add real name
 5. Add ssl

### Contribution

 1. Fork [irc](https://github.com/Bullet-Chat/irc);
 2. Make changes;
 3. Send pull request;
 4. Thank you!

### Author

[@0xAX](https://twitter.com/0xAX) and [@_voidPirate](https://twitter.com/_voidPirate)
