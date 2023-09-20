module github.com/pluralsh/deployment-argocd-driver

go 1.19

require (
	github.com/pluralsh/deployment-api v0.0.0-20230919142746-6052930922e9
	github.com/pluralsh/deployment-operator v0.0.0-20230920071952-496df605c67f
	github.com/spf13/cobra v1.7.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.16.0
	google.golang.org/grpc v1.58.1
	k8s.io/klog v1.0.0
	k8s.io/klog/v2 v2.100.1
	github.com/argoproj/argo-cd/v2 v2.8.4
)

