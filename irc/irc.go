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

	"github.com/nlamirault/k8s-bot/k8s"
	"github.com/nlamirault/k8s-bot/messages"
	"github.com/nlamirault/k8s-bot/version"
)

var (
	staticMessages = map[string]string{
		"Help": `You can use the following commands:
        !help     : Display help message
        !version  : Show bot version
        !k8s      : Display Kubernetes resouces
        `,
	}
)

type Provider struct {
	channels   []string
	Conn       *ircevent.Connection
	in         chan messages.Message
	Out        chan messages.Message
	Kubernetes *k8s.Client
}

func NewProvider(server string, nick string, channels string, k8sclient *k8s.Client) (*Provider, error) {
	log.Printf("[DEBUG] IRC Provider using: %s %s %s", server, nick, channels)
	provider := &Provider{
		channels:   strings.Split(channels, ","),
		in:         make(chan messages.Message),
		Out:        make(chan messages.Message),
		Kubernetes: k8sclient,
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
		msg := messages.Message{
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
		log.Printf("[INFO] IRC  message: %s", inMsg)
		var buffer bytes.Buffer
		if strings.Contains(inMsg.Message, "!version") {
			buffer.WriteString(fmt.Sprintf("v%s", version.Version))
		} else if strings.Contains(inMsg.Message, "!help") {
			buffer.WriteString(staticMessages["Help"])
		} else if strings.Contains(inMsg.Message, "!k8s") {
			buffer = p.dispatchKubernetesCommand(inMsg.Message)
		} else {
			buffer.WriteString("Sorry, unknown command.")
		}
		outMsg := messages.Message{
			Room:       inMsg.Room,
			ToUserName: inMsg.FromUserName,
			Message:    buffer.String(),
		}
		p.Out <- outMsg
	}
}

func (p *Provider) Sender() {
	for msg := range p.Out {
		log.Printf("[DEBUG] IRC Output message: %s", msg)
		channel := msg.Room
		if p.Conn.GetNick() == msg.Room {
			channel = msg.FromUserName
		}
		if channel == "" {
			channel = p.channels[0]
		}

		var finalMsg bytes.Buffer
		finalMsg.WriteString(msg.Message)
		msgs := strings.Split(finalMsg.String(), "\n")
		for _, m := range msgs {
			p.Conn.Privmsg(channel, m)
		}
	}
}

func (p *Provider) dispatchKubernetesCommand(msg string) bytes.Buffer {
	var buffer bytes.Buffer
	if strings.Contains(msg, ":services") {
		services, err := p.Kubernetes.GetServices()
		if err != nil {
			buffer.WriteString(fmt.Sprintf("Kubernetes error: %s", err.Error()))
		} else {
			buffer.WriteString("Services:\n")
			for _, service := range services.Items {
				buffer.WriteString(fmt.Sprintf("- %s\n", service.Name))
			}
		}

	} else if strings.Contains(msg, ":pods") {
		pods, err := p.Kubernetes.GetPods()
		if err != nil {
			buffer.WriteString(fmt.Sprintf("Kubernetes error: %s", err.Error()))
		} else {
			buffer.WriteString("Pods:\n")
			for _, pod := range pods.Items {
				buffer.WriteString(fmt.Sprintf("- %s\n", pod.Name))
			}
		}
	} else if strings.Contains(msg, ":nodes") {
		nodes, err := p.Kubernetes.GetNodes()
		if err != nil {
			buffer.WriteString(fmt.Sprintf("Kubernetes error: %s", err.Error()))
		} else {
			buffer.WriteString("Nodes:\n")
			for _, node := range nodes.Items {
				buffer.WriteString(fmt.Sprintf("- %s\n", node.Name))
			}
		}

	} else if strings.Contains(msg, ":namespaces") {
		namespaces, err := p.Kubernetes.GetPods()
		if err != nil {
			buffer.WriteString(fmt.Sprintf("Kubernetes error: %s", err.Error()))
		} else {
			buffer.WriteString("Namespaces:\n")
			for _, namespace := range namespaces.Items {
				buffer.WriteString(fmt.Sprintf("- %s\n", namespace.Name))
			}
		}
	}
	return buffer
}
