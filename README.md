# Docker2Kube

Go:

* _from_ a Docker image for a program that runs on a port
* _to_ the simplest possible Kubernetes config that runs that program

In more technical words:

Create a boilerplate Kubernetes deployment-service pair for a given container image.

## Installing

```
go build
sudo cp docker2kube /usr/local/bin/docker2kube
```

## Usage

```
$ cd myprogram
$ ls
Dockerfile main.go
$ docker build -t registry/myprogram:v1 .
[...]
$ docker push registry/myprogram:v1
$ cd ..
$ mkdir myprogram-config
$ git init .
$ docker2kube myprogram registry/myprogram:v1 80
$ git add *.yaml
$ git commit -am "Initial config"
```

## What next?

Once you've pushed Kubernetes config to a git repo, the practice known as _GitOps_ will help you get configs in version control reliably into a running application.

TODO: link to a public doc about _GitOps_.
