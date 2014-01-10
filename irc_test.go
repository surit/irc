package irc

import (
	"testing"
)

func Test_NewIrcClient(t *testing.T) {
	irc := NewIrcClient()
	irc.Nick = "voidpirate"
	irc.Pass = ""

	if irc.Nick != "voidpirate" {
		t.Fatal("[Test_new_ircClient] nick matching failed")
	}

	if irc.Pass != "" {
		t.Fatal("[Test_new_ircClient] pass matching failed")
	}
}

func Test_CheckPort(t *testing.T) {
	irc := NewIrcClient()
	irc = CheckPort(irc)

	if irc.Port != 6667 {
		t.Fatal("[Test_checkPort] 6667 failed")
	}

	new_irc := NewIrcClient()
	new_irc.Port = 4000
	new_irc = CheckPort(new_irc)

	if new_irc.Port != 4000 {
		t.Fatal("[Test_checkPort] 4000 failed")
	}
}

func Test_Join(t *testing.T) {
	irc := NewIrcClient()
	irc.Host = "irc.freenode.net"
	irc.Nick = "weber222222"
	irc.Port = 6667
	irc.Ssl = false
	irc.CallBack = handle

	irc.Join("#testGoLa", "")
	irc.Join("#WeberMVC", "")
}

func handle(irc *IRCClient, channel string, message string) {
	irc.SendMessage(message, channel)
}
