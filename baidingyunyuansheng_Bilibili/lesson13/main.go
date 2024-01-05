package main

import (
	"k8s.io/client-go/tools/clientcmd"
	"lesson13/pkg/generated/clientset/versioned"
	"log"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		log.Fatalln(err)
	}

	clientset, err := versioned.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	}

	clientset.CrdV1().Foos("default").List
}
