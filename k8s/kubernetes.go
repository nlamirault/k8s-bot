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

package k8s

// import (
// 	"log"
// )

// // Manager define a Kubernetes manager
// type Manager struct {
// 	Client  *Client
// 	Watcher *Watcher
// }

// func NewKubernetesManager(kubeconfigPath string, out chan messages.Message) (*Manager, error) {
// 	log.Printf("[INFO] Create the Kubernetes manager")
// 	client, err := newKubernetesClient(kubeconfigPath)
// 	if err != nil {
// 		return nil, err
// 	}

// 	watcher, err := newKubernetesWatcher(client.Clientset, out)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Manager{
// 		Client:  client,
// 		Watcher: watcher,
// 	}, nil
// }
