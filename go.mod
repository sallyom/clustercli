module github.com/sallyom/clustercli

go 1.13

require (
	github.com/openshift/api v0.0.0-20201012140924-16436fa6166b
	github.com/openshift/client-go v0.0.0-20200827190008-3062137373b5
	github.com/openshift/clustercli v0.0.0-00010101000000-000000000000
	github.com/openshift/library-go v0.0.0-20201013192036-5bd7c282e3e7
	github.com/prometheus/client_golang v1.7.1
	github.com/spf13/cobra v1.1.1
	k8s.io/apimachinery v0.19.2
	k8s.io/klog/v2 v2.3.0
)

replace (
	github.com/openshift/api => /home/somalley/code/gowork/src/github.com/openshift/api
	github.com/openshift/client-go => /home/somalley/code/gowork/src/github.com/openshift/client-go
	github.com/openshift/clustercli => /home/somalley/code/gowork/src/github.com/openshift/clustercli
)
