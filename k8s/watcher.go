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
	"log"

	"k8s.io/client-go/1.4/kubernetes"
	// "k8s.io/client-go/1.4/pkg/api"
	"k8s.io/client-go/1.4/pkg/api/v1"
	"k8s.io/client-go/1.4/pkg/watch"
)

// Watcher define a Kubernetes watching resources
type Watcher struct {
	Pods      watch.Interface
	Services  watch.Interface
	Endpoints watch.Interface
}

func createWatchers(clientset *kubernetes.Clientset) (watch.Interface, watch.Interface, watch.Interface, error) {
	podsWatcher, err := createPodsWatcher(clientset)
	if err != nil {
		return nil, nil, nil, err
	}
	log.Printf("[DEBUG] Kubernetes pods watcher created.")
	svcsWatcher, err := createServicesWatcher(clientset)
	if err != nil {
		return nil, nil, nil, err
	}
	log.Printf("[DEBUG] Kubernetes services watcher created.")
	endpointWatcher, err := createEndpointsWatcher(clientset)
	if err != nil {
		return nil, nil, nil, err
	}
	log.Printf("[DEBUG] Kubernetes endpoints watcher created.")
	return podsWatcher, svcsWatcher, endpointWatcher, nil
}

func NewKubernetesWatcher(clientset *kubernetes.Clientset) (*Watcher, error) {
	podsWatcher, svcWatcher, endpointWatcher, err := createWatchers(clientset)
	if err != nil {
		return nil, err
	}
	return &Watcher{
		Pods:      podsWatcher,
		Services:  svcWatcher,
		Endpoints: endpointWatcher,
	}, nil
}

func (watcher *Watcher) listen() bool {
	select {
	case ev, ok := <-watcher.Pods.ResultChan():
		if !ok {
			return false
		}
		pod := ev.Object.(*v1.Pod)
		managePodEvent(ev.Type, pod)
	case ev, ok := <-watcher.Services.ResultChan():
		if !ok {
			return false
		}
		svc := ev.Object.(*v1.Service)
		manageServiceEvent(ev.Type, svc)
	case ev, ok := <-watcher.Endpoints.ResultChan():
		if !ok {
			return false
		}
		endpoints := ev.Object.(*v1.Endpoints)
		manageEndpointsEvent(ev.Type, endpoints)
	}
	return true

}

func (watcher *Watcher) Watch() {
	for {
		ok := watcher.listen()
		if !ok {
			log.Printf("[ERROR] Kubernetes watchers channels closed.")
		}
	}
}
