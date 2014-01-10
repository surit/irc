package irc

import (
	"bufio"
	"crypto/tls"
	"errors"
	"log"
	"net"
	"net/textproto"
	"strconv"
	"strings"
)

const (
	MaxMessageLength = 510
)

const (
	End = "\r\n"
)

type IRCClient struct {
	Nick       string
	Pass       string
	Host       string
	Port       int
	Connection net.Conn
	CallBack   func(*IRCClient, string)
	Channel    string
	Ssl        bool
	SslCert    string
	SslKey     string
}

func NewIrcClient() *IRCClient {
	return &IRCClient{}
}

func CheckPort(irc *IRCClient) *IRCClient {
	if irc.Port == 0 {
		irc.Port = 6667
		return irc
	} else {
		return irc
	}
}

func CheckHost(irc *IRCClient) (*IRCClient, error) {
	if irc.Host == "" {
		log.Fatal("[Error] Host can't be empty")
		return irc, errors.New("[Error] Host can't be empty")
	} else {
		return irc, nil
	}
}

func CheckChannel(irc *IRCClient) (*IRCClient, error) {
	if irc.Channel == "" {
		log.Fatal(("[Error] Channel can't be empty"))
		return irc, errors.New("[Error] Channel can't be empty")
	} else {
		return irc, nil
	}
}

func (i *IRCClient) SendMessage(message string) {
	i.Connection.Write([]byte("PRIVMSG " + i.Channel + " " + message + " " + " \r\n"))
}

func (i *IRCClient) Join() {
	var config tls.Config
	var conn net.Conn

	if i.Ssl == true {

		if i.SslCert == "" {
			config = tls.Config{InsecureSkipVerify: true}
		} else {
			cert, _ := tls.LoadX509KeyPair(i.SslCert, i.SslKey)
			config = tls.Config{Certificates: []tls.Certificate{cert}}
		}

		conn, _ = tls.Dial("tcp", i.Host+":"+strconv.Itoa(i.Port), &config)
	} else {
		conn, _ = net.Dial("tcp", i.Host+":"+strconv.Itoa(i.Port))
	}

	i.Connection = conn

	if i.Pass != "" {
		i.Connection.Write([]byte("PASS " + i.Pass + " \r\n"))
	}

	i.Connection.Write([]byte("NICK " + i.Nick + " \r\n"))
	i.Connection.Write([]byte("USER " + i.Nick + " nohost noserver :golang\r\n"))
	i.Connection.Write([]byte("JOIN " + i.Channel + " \r\n"))

	start_connect(i)
}

func start_connect(client *IRCClient) {
	reader := bufio.NewReader(client.Connection)
	tp := textproto.NewReader(reader)

	for {
		line, _ := tp.ReadLine()

		resp := strings.Split(line, " ")

		if resp[0] == "PING" {
			client.Connection.Write([]byte("PONG \r\n"))
			continue
		}

		if resp[1] == "PRIVMSG" {
			mess := strings.Split(line, client.Channel)
			client.CallBack(client, mess[len(mess)-1])
			continue
		}

		if resp[0] == "QUIT" {
			client.Connection.Close()
			return
		}
	}
}
