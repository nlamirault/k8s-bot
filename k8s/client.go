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
	"k8s.io/client-go/1.4/pkg/api"
	"k8s.io/client-go/1.4/pkg/api/v1"
	"k8s.io/client-go/1.4/tools/clientcmd"
)

// Client define a Kubernetes client.
type Client struct {
	Clientset *kubernetes.Clientset
}

func NewKubernetesClient(kubeconfigPath string) (*Client, error) {
	// uses the current context in kubeconfig
	log.Printf("[DEBUG] Load Kubernetes configuration from %s", kubeconfigPath)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] Creates the Kubernetes clientset")
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return &Client{
		Clientset: clientset,
	}, nil
}

func (client *Client) GetServices() (*v1.ServiceList, error) {
	services, err := client.Clientset.Core().Services("").List(api.ListOptions{})
	if err != nil {
		return nil, err
	}
	return services, nil
}
