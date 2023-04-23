# Kube-Events using WebSockets

This application uses WebSockets to display Kubernetes events in real-time. It consists of a WebSocket server and client, both of which can be deployed using Kubernetes manifests.

## What is WebSocket?

WebSocket is a protocol that enables bidirectional communication between a client and a server over a single, long-lived connection. Unlike HTTP, which uses multiple connections to exchange data, WebSocket allows data to be transmitted in real-time using a single connection. This makes it ideal for applications that require real-time updates, such as chat apps, stock tickers, and game servers.

## Why WebSocket?

We chose to use WebSocket for this application because it provides real-time updates of Kubernetes events. This allows users to monitor their Kubernetes cluster and respond to events as they happen.

## How to use the application

To use the Kube-Events using WebSockets application, follow these steps:

1. Clone the repository using the following command:

```
git clone https://github.com/jilanisayyad/DevOps/tree/main/kube-events-websocket
```


2. Deploy the client and server manifests using the following commands:

```
kubectl apply -f kube-events-websocket/manifest.yaml
kubectl apply -f kube-events-websocket/wb-server/manifest.yaml
kubectl apply -f kube-events-websocket/wb-server-go/manifest.yaml
```


This will deploy both the client and server pods in your Kubernetes cluster.

3. Once the manifests have been deployed, you can access the client at `http://<node-ip>/` and start monitoring Kubernetes events in real-time.

## Architecture

The Kube-Events using WebSockets application consists of two pods:

- WebSocket client: The client is built using HTML and JavaScript and displays the events received from the server in real-time. The client is hosted at `/` path and is exposed on port `80`. The client Docker image can be pulled using `jilani1/kube-events`.
- WebSocket server: The server is built using Golang or Python and listens on port `8080` for incoming WebSocket connections. It receives Kubernetes events from the Kubernetes API server and broadcasts them to all connected clients. The server Docker images can be pulled using `jilani1/kube-connector-go` for Golang or `jilani1/kube-connector` for Python.

## Repository

The code for the Kube-Events using WebSockets application can be found in the following repository:

- https://github.com/jilanisayyad/DevOps/tree/main/kube-events-websocket

The repository also contains Kubernetes manifest files for deploying the client and server pods.

## Contributing

If you would like to contribute to the Kube-Events using WebSockets application, please fork the repository and submit a pull request. We welcome contributions of all kinds, including bug fixes, new features, and documentation improvements.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Follow and Subscribe

You can follow me and subscribe to my Medium blog for more articles and tutorials:

- https://medium.com/@sayyedjilani8



