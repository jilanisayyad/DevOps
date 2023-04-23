package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func sendEvents(ws *websocket.Conn) {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	w, err := clientset.CoreV1().Pods("").Watch(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error watching pods: %s", err.Error())
	}
	defer w.Stop()
	for event := range w.ResultChan() {
		cus_event := map[string]interface{}{
			"event_type": event.Type,
			"object": map[string]interface{}{
				"resource":  event.Object.(*v1.Pod).Name,
				"namespace": event.Object.(*v1.Pod).Namespace,
			},
		}
		log.Printf("Event: %s %s %s", cus_event["event_type"], cus_event["object"].(map[string]interface{})["namespace"], cus_event["object"].(map[string]interface{})["resource"])
		if err := ws.WriteJSON(cus_event); err != nil {
			log.Printf("Error writing JSON to websocket: %s", err.Error())
			break
		}
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("Error upgrading websocket: %s", err.Error())
			return
		}
		defer ws.Close()
		sendEvents(ws)
	})
	log.Println("Server starting on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %s", err.Error())
	}
}
