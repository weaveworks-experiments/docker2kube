package main

import (
	"fmt"
	"os"
	"text/template"
)

const SERVICE_TEMPLATE = `---
apiVersion: v1
kind: Service
metadata:
  name: {{.Name}}
  labels:
    name: {{.Name}}
spec:
  type: LoadBalancer
  ports:
  - port: {{.ContainerPort}}
  selector:
    name: {{.Name}}
`

const DEPLOYMENT_TEMPLATE = `---
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: {{.Name}}
spec:
  replicas: 1
  template:
    metadata:
      labels:
        name: {{.Name}}
    spec:
      containers:
      - name: {{.Name}}
        image: {{.Image}}
        ports:
        - containerPort: {{.ContainerPort}}
`

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: docker2kube <program-name> <docker-image:tag> <container-port>")
		return
	}

	programName := os.Args[1]
	dockerImageTag := os.Args[2]
	containerPort := os.Args[3]

	deploymentFilename := fmt.Sprintf("%s-deployment.yaml", programName)
	serviceFilename := fmt.Sprintf("%s-service.yaml", programName)

	if _, err := os.Stat(deploymentFilename); !os.IsNotExist(err) {
		fmt.Println("%s already exists. Rename it and re-run this tool to proceed.", deploymentFilename)
	}
	if _, err := os.Stat(serviceFilename); !os.IsNotExist(err) {
		fmt.Println("%s already exists. Rename it and re-run this tool to proceed.", serviceFilename)
	}

	type Inputs struct {
		Name          string
		Image         string
		ContainerPort string
	}

	input := Inputs{
		Name:          programName,
		Image:         dockerImageTag,
		ContainerPort: containerPort,
	}

	for filename, tmpl := range map[string]string{
		deploymentFilename: DEPLOYMENT_TEMPLATE,
		serviceFilename:    SERVICE_TEMPLATE,
	} {
		f, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		t, err := template.New(filename).Parse(tmpl)
		if err != nil {
			panic(err)
		}
		err = t.Execute(f, input)
		if err != nil {
			panic(err)
		}
		err = f.Close()
		if err != nil {
			panic(err)
		}
	}
}
