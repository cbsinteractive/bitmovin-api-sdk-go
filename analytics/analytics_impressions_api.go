package analytics
import (
    "github.com/cbsinteractive/bitmovin-api-sdk-go/common"
    
    "github.com/cbsinteractive/bitmovin-api-sdk-go/model"
)

type AnalyticsImpressionsApi struct {
    apiClient *common.ApiClient
}

func NewAnalyticsImpressionsApi(configs ...func(*common.ApiClient)) (*AnalyticsImpressionsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AnalyticsImpressionsApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AnalyticsImpressionsApi) Create(impressionId string, analyticsLicense model.AnalyticsLicense) (*model.AnalyticsImpressionDetails, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["impression_id"] = impressionId
    }

    var responseModel *model.AnalyticsImpressionDetails
    err := api.apiClient.Post("/analytics/impressions/{impression_id}", &analyticsLicense, &responseModel, reqParams)
    return responseModel, err
}

