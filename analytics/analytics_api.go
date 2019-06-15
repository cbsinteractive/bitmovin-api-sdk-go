package analytics
import (
    "github.com/cbsinteractive/bitmovin-api-sdk-go/common"
    
)

type AnalyticsApi struct {
    apiClient *common.ApiClient
    Exports *AnalyticsExportsApi
    Impressions *AnalyticsImpressionsApi
    Queries *AnalyticsQueriesApi
    Licenses *AnalyticsLicensesApi
    Outputs *AnalyticsOutputsApi
}

func NewAnalyticsApi(configs ...func(*common.ApiClient)) (*AnalyticsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsApi{apiClient: apiClient}

    exportsApi, err := NewAnalyticsExportsApi(configs...)
    api.Exports = exportsApi
    impressionsApi, err := NewAnalyticsImpressionsApi(configs...)
    api.Impressions = impressionsApi
    queriesApi, err := NewAnalyticsQueriesApi(configs...)
    api.Queries = queriesApi
    licensesApi, err := NewAnalyticsLicensesApi(configs...)
    api.Licenses = licensesApi
    outputsApi, err := NewAnalyticsOutputsApi(configs...)
    api.Outputs = outputsApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

