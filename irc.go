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
	CallBack   func(*IRCClient, string, string)
	Channels   map[string]string
	Ssl        bool
	SSLCert    string
	SSLKey     string
}

func NewIrcClient() *IRCClient {
	return &IRCClient{Channels: make(map[string]string)}
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

func (i *IRCClient) SendMessage(message string, channel string) {
	i.Connection.Write([]byte("PRIVMSG " + channel + " " + message + " " + " \r\n"))
}

func connect(i *IRCClient) (conn net.Conn, e error) {
	var config tls.Config

	if i.Ssl == true {
		if i.SSLCert == "" {
			config = tls.Config{InsecureSkipVerify: true}
			return tls.Dial("tcp", i.Host+":"+strconv.Itoa(i.Port), &config)
		} else {
			cert, _ := tls.LoadX509KeyPair(i.SSLCert, i.SSLKey)
			config = tls.Config{Certificates: []tls.Certificate{cert}}
			return tls.Dial("tcp", i.Host+":"+strconv.Itoa(i.Port), &config)
		}
	} else {
		return net.Dial("tcp", i.Host+":"+strconv.Itoa(i.Port))
	}
}

func (i *IRCClient) Join(channel string, channel_key string) {
	_, val := i.Channels[channel]

	if val == false {
		i.Channels[channel] = channel_key
		conn, _ := connect(i)
		i.Connection = conn

		if i.Pass != "" {
			i.Connection.Write([]byte("PASS " + i.Pass + " \r\n"))
		}

		i.Connection.Write([]byte("NICK " + i.Nick + " \r\n"))
		i.Connection.Write([]byte("USER " + i.Nick + " nohost noserver :golang\r\n"))

		if channel_key != "" {
			i.Connection.Write([]byte("JOIN " + channel + " " + channel_key + " \r\n"))
		} else {
			i.Connection.Write([]byte("JOIN " + channel + " \r\n"))
		}

		start_connect(i, channel)
	}
}

func start_connect(client *IRCClient, channel string) {
	reader := bufio.NewReader(client.Connection)
	tp := textproto.NewReader(reader)

	defer client.Connection.Close()

	for {
		line, _ := tp.ReadLine()

		resp := strings.Split(line, " ")

		if resp[0] == "PING" {
			client.Connection.Write([]byte("PONG \r\n"))
			continue
		}

		if resp[1] == "PRIVMSG" {
			mess := strings.Split(line, channel)
			client.CallBack(client, channel, mess[len(mess)-1])
			continue
		}

		if resp[0] == "QUIT" {
			client.Connection.Close()
			return
		}
	}
}
