// Copyright (C) 2016 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package irc

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	ircevent "github.com/thoj/go-ircevent"

	"github.com/nlamirault/k8s-bot/version"
)

var (
	messages = map[string]string{
		"Help": `You can use the following commands:
        !help     : Display help message
        !version  : Show bot version
        `,
	}
)

type Message struct {
	Room         string
	FromUserName string
	ToUserName   string
	Message      string
}

type Provider struct {
	channels []string
	Conn     *ircevent.Connection
	in       chan Message
	out      chan Message
	err      error
}

func NewProvider(server string, nick string, channels string) (*Provider, error) {
	log.Printf("[DEBUG] IRC Provider using: %s %s %s", server, nick, channels)
	provider := &Provider{
		channels: strings.Split(channels, ","),
		in:       make(chan Message),
		out:      make(chan Message),
	}
	ircConn := ircevent.IRC(nick, nick)
	ircConn.AddCallback(
		"001",
		func(e *ircevent.Event) {
			for _, channel := range provider.channels {
				ircConn.Join(channel)
			}
		},
	)

	ircConn.AddCallback("PRIVMSG", func(e *ircevent.Event) {
		log.Printf("[DEBUG] IRC Message: %s", e)
		msg := Message{
			Room:         e.Arguments[0],
			FromUserName: e.Nick,
			Message:      e.Message(),
		}
		provider.in <- msg
	})

	if err := ircConn.Connect(server); err != nil {
		return nil, err
	}
	provider.Conn = ircConn
	return provider, nil
}

func (p *Provider) Receiver() {
	for inMsg := range p.in {
		log.Printf("[INFO] Manage incomming message: %s", inMsg)
		var buffer bytes.Buffer
		if strings.Contains(inMsg.Message, "!version") {
			buffer.WriteString(fmt.Sprintf("v%s", version.Version))
		} else if strings.Contains(inMsg.Message, "!help") {
			buffer.WriteString(messages["Help"])
		} else {
			buffer.WriteString("Sorry, unknown command.")
		}
		outMsg := Message{
			Room:       inMsg.Room,
			ToUserName: inMsg.FromUserName,
			Message:    buffer.String(),
		}
		p.out <- outMsg
	}
}

func (p *Provider) Sender() {
	for msg := range p.out {
		channel := msg.Room
		if p.Conn.GetNick() == msg.Room {
			channel = msg.FromUserName
		}

		var finalMsg bytes.Buffer
		finalMsg.WriteString(msg.Message)
		msgs := strings.Split(finalMsg.String(), "\n")
		for _, m := range msgs {
			p.Conn.Privmsg(channel, m)
		}
	}
}
