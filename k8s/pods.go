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

import (
	"fmt"
	"log"

	"k8s.io/client-go/1.4/kubernetes"
	"k8s.io/client-go/1.4/pkg/api"
	"k8s.io/client-go/1.4/pkg/api/v1"
	"k8s.io/client-go/1.4/pkg/watch"

	"github.com/nlamirault/k8s-bot/messages"
)

func createPodsWatcher(clientset *kubernetes.Clientset) (watch.Interface, error) {
	watcher, err := clientset.Core().Pods("").Watch(api.ListOptions{})
	if err != nil {
		return nil, err
	}
	return watcher, nil
}

func managePodEvent(out chan messages.Message, eventType watch.EventType, pod *v1.Pod) {
	switch eventType {
	case watch.Added:
		log.Printf("[DEBUG] Add pod: %s\n", pod.Name)
		msg := makeMessage(fmt.Sprintf("Kubernetes: Pod added: %s", pod.Name))
		out <- msg
	case watch.Deleted:
		log.Printf("[DEBUG] Deleted pod: %s\n", pod.Name)
		msg := makeMessage(fmt.Sprintf("Kubernetes: Pod deleted: %s", pod.Name))
		out <- msg
	case watch.Modified:
		log.Printf("[DEBUG] Modified pod: %s\n", pod.Name)
		msg := makeMessage(fmt.Sprintf("Kubernetes: Pod modified: %s", pod.Name))
		out <- msg
	}
}
