package main

import (
	"context"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"time"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	NAMESPACE = os.Getenv("namespace")
	LABEL     = os.Getenv("label")
	DURATION  = os.Getenv("duration")
)

func main() {
	clientset, err := GetClient()
	if nil != err {
		panic(err)
	}
	pods, err := clientset.CoreV1().Pods(NAMESPACE).List(context.TODO(), metav1.ListOptions{
		LabelSelector: LABEL,
	})
	if err != nil {
		panic(err)
	}

	for _, pod := range pods.Items {
		createdAt := pod.CreationTimestamp.Time

		duration, err := strconv.Atoi(DURATION)

		if err != nil {
			panic(err)
		}

		allowedTime := time.Now().Add(-time.Duration(duration) * time.Hour)
		if createdAt.Before(allowedTime) {
			clientset.CoreV1().Pods(NAMESPACE).Delete(context.TODO(), pod.Name, metav1.DeleteOptions{})
		}
	}
}

func GetClient() (*kubernetes.Clientset, error) {
	config, err := getClientConfig()
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(config)
}

func getClientConfig() (*rest.Config, error) {
	config, err := rest.InClusterConfig()
	if err == nil {
		return config, nil
	}

	u, err := user.Current()
	if nil != err {
		return nil, err
	}

	return clientcmd.BuildConfigFromFlags("", filepath.Join(u.HomeDir, ".kube", "config"))
}
