package main

import (
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func main() {
	//read config from ~/.kube/config
	kubeconfig := homeDir() + "/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	//get all namespace
	namespaces, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	//List all namespaces
	for _, namespace := range namespaces.Items {
		fmt.Printf("Namespace : %s\n", namespace.GetName())
	}
	//List pods in the namespace default
	namespace := "default"
	//get info for only the namespace default
	clientset.CoreV1().Namespaces().Get(namespace, metav1.GetOptions{})
	//get all pods in the namespace default
	fmt.Println("Pods:")
	pods, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	for _, pod := range pods.Items {
		fmt.Printf("   %s , Created on : %v\n", pod.GetName(), pod.GetCreationTimestamp())
	}

}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
