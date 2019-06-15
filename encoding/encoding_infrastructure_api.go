package encoding
import (
    "github.com/cbsinteractive/bitmovin-api-sdk-go/common"
    
)

type EncodingInfrastructureApi struct {
    apiClient *common.ApiClient
    Kubernetes *EncodingInfrastructureKubernetesApi
    Aws *EncodingInfrastructureAwsApi
}

func NewEncodingInfrastructureApi(configs ...func(*common.ApiClient)) (*EncodingInfrastructureApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInfrastructureApi{apiClient: apiClient}

    kubernetesApi, err := NewEncodingInfrastructureKubernetesApi(configs...)
    api.Kubernetes = kubernetesApi
    awsApi, err := NewEncodingInfrastructureAwsApi(configs...)
    api.Aws = awsApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

