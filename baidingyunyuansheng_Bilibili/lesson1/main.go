package main

import (
	"context"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}

	config.GroupVersion = &v1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs
	config.APIPath = "/api"
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}

	pod := v1.Pod{}
	err = restClient.Get().Namespace("default").Resource("pods").Name("nginx-deployment-5bf87f5f59-8tln7").Do(context.TODO()).Into(&pod)
	if err != nil {
		panic(err)
	} else {
		println(pod.Name)
	}

	//config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	//if err != nil {
	//	panic(err)
	//}
	//
	//clientset, err := kubernetes.NewForConfig(config)
	//if err != nil {
	//	panic(err)
	//}
	//
	//coreV1 := clientset.CoreV1()
	//pod, err := coreV1.Pods("default").Get(context.TODO(), "nginx-deployment-5bf87f5f59-8tln7", v1.GetOptions{})
	//if err != nil {
	//	panic(err)
	//} else {
	//	println(pod.Name)
	//}

}
