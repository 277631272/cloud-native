package main

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

func main() {
	kubeconfig := "./config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		klog.Fatalf("BuildConfigFromFlags failed, err: %s", err)
	}
	clientset := clientset.NewForConfigOrDie(config)
	podList, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		klog.Fatalf("Pods List failed, err: %s", err)
	}
	klog.Infof("podlist: %v", podList)
}
