package handlers

import (
	"log"

	"context"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Deployment struct {
	Name      string `json:"name"`
	Ready     int32  `json:"ready"`
	UpToDate  int32  `json:"up_to_date"`
	Available int32  `json:"available"`
	Namespace string `json:"namespace"`
}

type Pod struct {
	Name      string `json:"name"`
	Ready     bool   `json:"ready"`
	Status    string `json:"status"`
	Restarts  int32  `json:"restarts"`
	Namespace string `json:"namespace"`
	Age       int32  `json:"age"`
}

type ReplicaSet struct {
	Name      string `json:"name"`
	Desired   int32  `json:"desired"`
	Current   int32  `json:"current"`
	Ready     int32  `json:"ready"`
	Age       int32  `json:"age"`
	Namespace string `json:"namespace"`
}

type Services struct {
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	ClusterIP  string   `json:"cluster_ip"`
	Port       int32    `json:"port"`
	Namespace  string   `json:"namespace"`
	ExternalIP []string `json:"external_ip"`
}

func GetDeployments(c *gin.Context, namespace string) {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	var deployments []Deployment

	if namespace == "" {
		deploymentList, err := clientset.AppsV1().Deployments("").List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}

		for _, deployment := range deploymentList.Items {
			deployments = append(deployments, Deployment{
				Name:      deployment.Name,
				Ready:     deployment.Status.ReadyReplicas,
				UpToDate:  deployment.Status.UpdatedReplicas,
				Available: deployment.Status.AvailableReplicas,
				Namespace: deployment.Namespace,
			})
		}
	}

	if namespace != "" {
		deploymentList, err := clientset.AppsV1().Deployments(namespace).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}

		for _, deployment := range deploymentList.Items {
			deployments = append(deployments, Deployment{
				Name:      deployment.Name,
				Ready:     deployment.Status.ReadyReplicas,
				UpToDate:  deployment.Status.UpdatedReplicas,
				Available: deployment.Status.AvailableReplicas,
				Namespace: deployment.Namespace,
			})
		}
	}

	c.JSON(200, deployments)

}

func GetNamespaces(c *gin.Context) {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	namespaceList, err := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	var namespaces []string

	for _, namespace := range namespaceList.Items {
		namespaces = append(namespaces, namespace.Name)
	}

	c.JSON(200, namespaces)
}

// GetPods retrieves a list of Pod objects from the Kubernetes cluster
func GetPods(c *gin.Context, namespace string) {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	var pods []Pod

	if namespace == "" {
		podList, err := clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}

		for _, pod := range podList.Items {
			pods = append(pods, Pod{
				Name:      pod.Name,
				Ready:     pod.Status.ContainerStatuses[0].Ready,
				Status:    string(pod.Status.Phase),
				Restarts:  pod.Status.ContainerStatuses[0].RestartCount,
				Namespace: pod.Namespace,
				Age:       int32(pod.CreationTimestamp.UTC().Sub(pod.Status.StartTime.Time).Hours() / 24),
			})
		}
	}

	if namespace != "" {
		podList, err := clientset.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}

		for _, pod := range podList.Items {
			pods = append(pods, Pod{
				Name:      pod.Name,
				Ready:     pod.Status.ContainerStatuses[0].Ready,
				Status:    string(pod.Status.Phase),
				Restarts:  pod.Status.ContainerStatuses[0].RestartCount,
				Namespace: pod.Namespace,
				Age:       int32(pod.CreationTimestamp.UTC().Sub(pod.Status.StartTime.Time).Hours() / 24),
			})
		}
	}

	c.JSON(200, pods)
}

func GetReplicaSets(c *gin.Context, namespace string) {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	var replicaSets []ReplicaSet

	if namespace == "" {
		replicaSetList, err := clientset.AppsV1().ReplicaSets("").List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}

		for _, replicaSet := range replicaSetList.Items {
			replicaSets = append(replicaSets, ReplicaSet{
				Name:      replicaSet.Name,
				Desired:   replicaSet.Status.Replicas,
				Current:   replicaSet.Status.ReadyReplicas,
				Ready:     replicaSet.Status.AvailableReplicas,
				Namespace: replicaSet.Namespace,
			})
		}
	}

	if namespace != "" {
		replicaSetList, err := clientset.AppsV1().ReplicaSets(namespace).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}

		for _, replicaSet := range replicaSetList.Items {
			replicaSets = append(replicaSets, ReplicaSet{
				Name:      replicaSet.Name,
				Desired:   replicaSet.Status.Replicas,
				Current:   replicaSet.Status.ReadyReplicas,
				Ready:     replicaSet.Status.AvailableReplicas,
				Namespace: replicaSet.Namespace,
			})
		}
	}

	c.JSON(200, replicaSets)
}

func GetServices(c *gin.Context, namespace string) {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	var services []Services

	if namespace == "" {
		serviceList, err := clientset.CoreV1().Services("").List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}

		for _, service := range serviceList.Items {
			services = append(services, Services{
				Name:       service.Name,
				Type:       string(service.Spec.Type),
				ClusterIP:  service.Spec.ClusterIP,
				Namespace:  service.Namespace,
				ExternalIP: service.Spec.ExternalIPs,
				Port:       service.Spec.Ports[0].Port,
			})
		}
	}

	if namespace != "" {
		serviceList, err := clientset.CoreV1().Services(namespace).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}

		for _, service := range serviceList.Items {
			services = append(services, Services{
				Name:       service.Name,
				Type:       string(service.Spec.Type),
				ClusterIP:  service.Spec.ClusterIP,
				ExternalIP: service.Spec.ExternalIPs,
				Namespace:  service.Namespace,
				Port:       service.Spec.Ports[0].Port,
			})
		}
	}

	c.JSON(200, services)
}
