package main

import (
    "example.com/cdk8s-demo/imports/k8s"
    "github.com/aws/constructs-go/constructs/v10"
    "github.com/aws/jsii-runtime-go"
)

type AppProps struct {
    Image         *string
    Replicas      *float64
    Port          *float64
    ContainerPort *float64
}

// NewApp creates both Service and Deployment for the Go Web application.
func NewApp(scope constructs.Construct, id *string, props *AppProps) constructs.Construct {
    construct := constructs.NewConstruct(scope, id)

    // Default values if none are provided
    replicas := props.Replicas
    if replicas == nil {
        replicas = jsii.Number(1)  // Default replicas to 1
    }

    port := props.Port
    if port == nil {
        port = jsii.Number(80)  // Default service port to 80
    }

    containerPort := props.ContainerPort
    if containerPort == nil {
        containerPort = jsii.Number(8080)  // Default container port to 8080
    }

    label := map[string]*string{
        "app": constructs.Node_Of(construct).Id(),
    }

    // Create a Service for the Go Web application
    k8s.NewKubeService(construct, jsii.String("go-web-app-service"), &k8s.KubeServiceProps{
        Metadata: &k8s.ObjectMeta{Labels: &label},
        Spec: &k8s.ServiceSpec{
            Type: jsii.String("NodePort"),  // Set service type to NodePort
            Ports: &[]*k8s.ServicePort{{
                Port:       port,
                TargetPort: k8s.IntOrString_FromNumber(containerPort),
                Protocol:   jsii.String("TCP"),
            }},
            Selector: &label,  // Selector to match the deployment
        },
    })

    // Create a Deployment for the Go Web application
    k8s.NewKubeDeployment(construct, jsii.String("go-web-app-deployment"), &k8s.KubeDeploymentProps{
        Metadata: &k8s.ObjectMeta{Labels: &label},
        Spec: &k8s.DeploymentSpec{
            Replicas: replicas,
            Selector: &k8s.LabelSelector{MatchLabels: &label},
            Template: &k8s.PodTemplateSpec{
                Metadata: &k8s.ObjectMeta{Labels: &label},
                Spec: &k8s.PodSpec{
                    Containers: &[]*k8s.Container{{
                        Name:  jsii.String("go-web-app"),
                        Image: props.Image,
                        Ports: &[]*k8s.ContainerPort{{ContainerPort: containerPort}},
                    }},
                },
            },
        },
    })

    return construct
}

func main() {
    app := cdk8s.NewApp(nil)

    // Define the application properties
    appProps := &AppProps{
        Image:         jsii.String("rushi/go-web-app"),  // Docker image to use
        Replicas:      jsii.Number(1),                    // Set replicas to 1
        Port:          jsii.Number(80),                   // Service port
        ContainerPort: jsii.Number(8080),                 // Container port
    }

    // Create the Go Web application with the defined properties
    NewApp(app, "go-web-demo", appProps)

    // Generate the Kubernetes YAML files
    app.Synth()
}

