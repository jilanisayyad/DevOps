# Creating Our Dashboard for Kubernetes

This is an example of how to create a simple web dashboard for Kubernetes using Go programming language, Gin framework, and Kubernetes client-go library. The code provides REST API endpoints to retrieve information about deployments, pods, replica sets, and services in a Kubernetes cluster.

## Dependencies

To run this code, you will need the following dependencies:

- Go programming language
- Kubernetes cluster with `kubectl` configured
- `github.com/gin-gonic/gin` and `k8s.io/client-go` Go packages

## Usage

1. Clone the repository and navigate to the project directory.
2. Build the Go application using `go build` command.
3. Run the compiled binary to start the web application.
4. The server will start running on `http://localhost:8080` or a different port specified in your code.

## API Endpoints

The following API endpoints are available:

- `GET /`: Returns the main HTML dashboard page.
- `GET /api/v1/deployments/namespace/:namespace`: Returns a list of deployments in the specified namespace.
- `GET /api/v1/deployments/`: Returns a list of deployments in all namespaces.
- `GET /api/v1/namespaces/`: Returns a list of namespaces in the cluster.
- `GET /api/v1/pods/namespace/:namespace`: Returns a list of pods in the specified namespace.
- `GET /api/v1/pods/`: Returns a list of pods in all namespaces.
- `GET /api/v1/replicasets/namespace/:namespace`: Returns a list of replica sets in the specified namespace.
- `GET /api/v1/replicasets/`: Returns a list of replica sets in all namespaces.
- `GET /api/v1/services/namespace/:namespace`: Returns a list of services in the specified namespace.
- `GET /api/v1/services/`: Returns a list of services in all namespaces.

## Handlers

The code provides the following handlers to handle the API endpoints:

### `GetDeployments`

This handler retrieves a list of deployments in the specified namespace or in all namespaces using the Kubernetes client-go library for Go programming language. It returns the deployment name, ready replicas, up-to-date replicas, available replicas, and namespace in JSON format.

### `GetNamespaces`

This handler retrieves a list of namespaces in the cluster using the Kubernetes client-go library for Go programming language. It returns the namespace names in JSON format.

### `GetPods`

This handler retrieves a list of pods in the specified namespace or in all namespaces using the Kubernetes client-go library for Go programming language. It returns the pod name, readiness status, status, restarts, namespace, and age in JSON format.

### `GetReplicaSets`

This handler retrieves a list of replica sets in the specified namespace or in all namespaces using the Kubernetes client-go library for Go programming language. It returns the replica set name, desired replicas, current replicas, ready replicas, age, and namespace in JSON format.

### `GetServices`

This handler retrieves a list of services in the specified namespace or in all namespaces using the Kubernetes client-go library for Go programming language. It returns the service name, type, cluster IP, port, external IP, and namespace in JSON format.

## Deployment

You can build your own Docker image of the web dashboard using a Dockerfile, and then deploy it as a pod in your Kubernetes cluster using a manifest.yaml file. You can customize the manifest.yaml file to specify the image, ports, and other configurations for your deployment.


## Contributing

Thank you for your interest in contributing to this project! You can contribute by:

1. Forking the repository.
2. Creating a new branch for your changes.
3. Making your changes and committing them with descriptive commit messages.
4. Pushing your changes to your forked repository.
5. Creating a pull request with a detailed description of your changes.

Please make sure to follow the project's code of conduct and contribution guidelines. Contributions are greatly appreciated!

## Follow Me

You can follow me on my [Medium page](https://medium.com/@sayyedjilani8) for more articles and tutorials related to Kubernetes, Go programming language, and other technologies.

If you have any questions or suggestions, feel free to reach out to me. Thank you for your support!



