# ix-gpu-operator
The ix-gpu-operator is designed to help the user to provision and configure gpu components in the kubernetes cluster. 

## Motivation
There are many services about gpu services in the gpu cluster. To make it work, it requires different components to be provisioned and configured accordingly. It makes sense to have one operator to coordinate those relevant components in one place, instead of having them managed by different operators. 

## Features
The operator will automatically create the necessary resources to run the [ix-exporter](https://gitee.com/deep-spark/ix-exporter) and [ix-device-plugin](https://gitee.com/deep-spark/ix-device-plugin).

## Getting Started

You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster.
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

### Running on the cluster

1. Install Instances of Custom Resources:

```sh
kubectl apply -f config/samples/
```

2. Build and push your image to the location specified by `IMG`:

```sh
make docker-build docker-push IMG=<some-registry>/ix-gpu-operator:<version>
```

3. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=<some-registry>/ix-gpu-operator:tag
```

### Uninstall CRDs

To delete the CRDs from the cluster:

```sh
make uninstall
```

### Undeploy controller

UnDeploy the controller from the cluster:

```sh
make undeploy
```

### Test It Out

1. Install the CRDs into the cluster:

```sh
make install
```

2. Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):

```sh
make run
```

**NOTE:** You can also run this in one step by running: `make install run`

### Modifying the API definitions

If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright (c) 2024 Iluvatar CoreX. All rights reserved. This project has an Apache-2.0 license, as
found in the [LICENSE](LICENSE) file.
