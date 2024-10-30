
# Image Digest Test

This project demonstrates how to retrieve the digest of a Docker image using the Go programming language and the `go-containerregistry` package.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Local Testing](#local-testing)
- [Kubernetes Environment Testing](#kubernetes-environment-testing)
- [Dockerfile Explanation](#dockerfile-explanation)

## Prerequisites

Make sure you have the following installed on your machine:

- [Go](https://golang.org/dl/) (version 1.22.2 or later)
- [Docker](https://www.docker.com/get-started)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/)

## Local Testing

To run the application locally, follow these steps:

1. Clone the repository or download the code.

   ```bash
   git clone <repository-url>
   cd image-digest-test
   ```

2. Build and run the Go application:

   ```bash
   go run main.go
   ```

   This command will execute the `main.go` file and output the image digest for the specified images.

## Kubernetes Environment Testing

To test the image digest retrieval in a Kubernetes environment, follow these steps:

1. **Build the Docker Image:**

   Build the Docker image using the provided Dockerfile.

   ```bash
   docker build -t neajmorshad/digest-tester:latest .
   ```

2. **Create a Kind Cluster:**

   If you haven't already, create a Kind cluster:

   ```bash
   kind create cluster
   ```

3. **Deploy the Pod:**

   Create a Kubernetes Pod using the provided configuration. Save the following YAML as `digest-tester-pod.yaml`:

   ```yaml
   apiVersion: v1
   kind: Pod
   metadata:
     name: digest-tester
   spec:
     containers:
     - name: digest-tester
       image: neajmorshad/digest-tester:latest
       imagePullPolicy: Always
       command: ["/bin/sh", "-c", "/digest-tester && tail -f /dev/null"]
   ```

   Apply the Pod configuration:

   ```bash
   kubectl apply -f digest-tester-pod.yaml
   ```

4. **Check Logs:**

   To view the logs from the Pod, run:

   ```bash
   kubectl logs -f digest-tester
   ```

   This command will display the output from the Go application, including the image digests.

5. **Exec into the Pod (optional):**

   If you want to interact with the Pod directly, you can exec into it:

   ```bash
   kubectl exec -it digest-tester -- sh
   ```

   This will give you a shell inside the Pod.