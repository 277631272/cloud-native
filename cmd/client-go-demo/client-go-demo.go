package main

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	"net"
)

func main() {
	{
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
		klog.Infof("podlist-v1: %v", podList)
	}

	{
		rootCAFile := "./ca.crt"
		tlsClientConfig := rest.TLSClientConfig{}
		tlsClientConfig.CAFile = rootCAFile
		host := "xxxxx"
		port := "xxx"
		tokenFile := "./token"
		config := &rest.Config{
			Host:            "https://" + net.JoinHostPort(host, port),
			TLSClientConfig: tlsClientConfig,
			BearerTokenFile: tokenFile,
		}
		clientset, err := clientset.NewForConfig(config)
		if err != nil {
			klog.Fatalf("NewForConfig failed, err: %s", err)
		}
		podList, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			klog.Fatalf("Pods List failed, err: %s", err)
		}
		klog.Infof("podlist-v2: %v", podList)
	}
	klog.Infof(v1.NamespaceDefault)
}
