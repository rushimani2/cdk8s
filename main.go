package main

import (
    "github.com/aws/constructs-go/constructs/v10"
    "github.com/aws/jsii-runtime-go"
    "github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

func main() {
    app := cdk8s.NewApp(nil)

    NewDeploymentChart(app, "deployment")
    NewServiceChart(app, "service")

    app.Synth()
}

func NewDeploymentChart(scope constructs.Construct, id string) cdk8s.Chart {
    chart := cdk8s.NewChart(scope, jsii.String(id), nil)
    NewDeployment(chart, jsii.String("go-web-app"), &AppProps{
        Image:         jsii.String("{{ .Values.image.repository }}:{{ .Values.image.tag }}"),
        Replicas:      jsii.Number(3),
        ContainerPort: jsii.Number(8080),
    })
    return chart
}

func NewServiceChart(scope constructs.Construct, id string) cdk8s.Chart {
    chart := cdk8s.NewChart(scope, jsii.String(id), nil)
    NewService(chart, jsii.String("go-web-app"), &AppProps{
        Port:          jsii.Number(8080),
        ContainerPort: jsii.Number(8080),
    })
    return chart
}
