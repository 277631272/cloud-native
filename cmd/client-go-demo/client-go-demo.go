package main

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	"net"
	"os"
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
		host, port := os.Getenv("k8s_service_host"), os.Getenv("k8s_service_port")
		if len(host) == 0 || len(port) == 0 {
			klog.Fatalf("invalid evn k8s_service_host, k8s_service_port")
		}
		klog.Infof("host|port: %s:%s", host, port)
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
}
