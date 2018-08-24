package main

import (
	"k8s.io/client-go/discovery"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kubernetes/pkg/kubectl/cmd/util/openapi"
	"k8s.io/kubernetes/pkg/kubectl/cmd/util/openapi/validation"
)

type validator struct {
	validator *validation.SchemaValidation
}

func newValidator() (*validator, error) {
	config := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)
	clientConfig, err := config.ClientConfig()
	if err != nil {
		return nil, err
	}

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(clientConfig)
	if err != nil {
		return nil, err
	}

	getter := openapi.NewOpenAPIGetter(discoveryClient)
	resources, err := getter.Get()
	if err != nil {
		return nil, err
	}

	return &validator{validation.NewSchemaValidation(resources)}, nil
}

func (v *validator) validate(data []byte) error {
	return v.validator.ValidateBytes(data)
}
