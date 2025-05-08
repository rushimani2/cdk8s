package main

import (
    "example.com/cdk8s-demo/imports/k8s"
    "github.com/aws/constructs-go/constructs/v10"
    "github.com/aws/jsii-runtime-go"
)

type AppProps struct {
    Image         *string
    Replicas      *string
    Port          *float64
    ContainerPort *float64
    Name          *string
}

func NewDeployment(scope constructs.Construct, id *string, props *AppProps) constructs.Construct {
    construct := constructs.NewConstruct(scope, id)

    replicas := props.Replicas
    if replicas == nil {
        replicas = jsii.Number(1)
    }

    appName := props.Name
    if appName == nil {
        appName = jsii.String("go-web-app")
    }

    label := map[string]*string{
        "app": appName,
    }

    k8s.NewKubeDeployment(construct, jsii.String("deployment"), &k8s.KubeDeploymentProps{
        Metadata: &k8s.ObjectMeta{
            Name:   appName,
            Labels: &label,
        },
        Spec: &k8s.DeploymentSpec{
            Replicas: replicas,
            Selector: &k8s.LabelSelector{MatchLabels: &label},
            Template: &k8s.PodTemplateSpec{
                Metadata: &k8s.ObjectMeta{Labels: &label},
                Spec: &k8s.PodSpec{
                    Containers: &[]*k8s.Container{{
                        Name:  appName,
                        Image: props.Image,
                        ImagePullPolicy: jsii.String("{{ .Values.image.pullPolicy }}"),
                        Ports: &[]*k8s.ContainerPort{{
                            ContainerPort: jsii.Number(8080),
                        }},
                    }},
                },
            },
        },
    })

    return construct
}

func NewService(scope constructs.Construct, id *string, props *AppProps) constructs.Construct {
    construct := constructs.NewConstruct(scope, id)

    port := props.Port
    if port == nil {
        port = jsii.Number(80)
    }

    appName := props.Name
    if appName == nil {
        appName = jsii.String("go-web-app")
    }

    label := map[string]*string{
        "app": appName,
    }

    k8s.NewKubeService(construct, jsii.String("service"), &k8s.KubeServiceProps{
        Metadata: &k8s.ObjectMeta{
            Name:   appName,
            Labels: &label,
        },
        Spec: &k8s.ServiceSpec{
            Type: jsii.String("LoadBalancer"),
            Ports: &[]*k8s.ServicePort{{
                Port:       port,
                TargetPort: k8s.IntOrString_FromNumber(jsii.Number(8080)),
                Protocol:   jsii.String("TCP"),
            }},
            Selector: &label,
        },
    })

    return construct
}
