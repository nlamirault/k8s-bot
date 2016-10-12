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

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nlamirault/k8s-bot/irc"
	"github.com/nlamirault/k8s-bot/k8s"
	"github.com/nlamirault/k8s-bot/version"
)

func main() {
	var (
		showVersion = flag.Bool("version", false, "Print version information.")
		ircServer   = flag.String("irc-server", "", "IRC server to used")
		ircPort     = flag.Int("irc-port", 6667, "The IRC port")
		ircNick     = flag.String("irc-nick", "k8s-bot", "Username for the Bot")
		ircChannels = flag.String("irc-channels", "", "Username for the Bot")
		kubeconfig  = flag.String("kubeconfig", "./config", "Absolute path to the kubeconfig file")
	)
	flag.Parse()

	if *showVersion {
		fmt.Printf("Kubernetes Bot. v%s\n", version.Version)
		os.Exit(0)
	}

	k8sManager, err := k8s.NewKubernetesManager(*kubeconfig)
	if err != nil {
		log.Printf("[ERROR] Kubernetes failed: %s", err.Error())
		os.Exit(1)
	}

	ircprovider, err := irc.NewProvider(fmt.Sprintf("%s:%d", *ircServer, *ircPort), *ircNick, *ircChannels, k8sManager)
	if err != nil {
		log.Printf("[ERROR] Can't create IRC provider : %s", err)
		os.Exit(1)
	}

	go ircprovider.Receiver()
	go ircprovider.Sender()
	go ircprovider.Kubernetes.Watcher.Watch()
	ircprovider.Conn.Loop()
}
