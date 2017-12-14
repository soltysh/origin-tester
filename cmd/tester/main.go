package main

import (
	"flag"
	"os"
	"path/filepath"
	"time"

	"github.com/golang/glog"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	config.UserAgent = "tester/v1.0.0"

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	i := 0
	for {
		glog.Infof("Starting %d batch...", i)
		setupSecretWatches(clientset, 200)
		i = i + 1
	}

}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func setupSecretWatches(clientset *kubernetes.Clientset, limit int) {
	for i := 0; i < limit; i++ {
		go func() {
			watch, err := clientset.CoreV1().Secrets("").Watch(metav1.ListOptions{})
			if err != nil {
				glog.Errorf("Error setting watch: %v", err)
			}
			select {
			case event := <-watch.ResultChan():
				glog.V(3).Infof("Got %v", event)
			case <-time.After(5. * time.Minute):
				watch.Stop()
			}
		}()
	}
}
