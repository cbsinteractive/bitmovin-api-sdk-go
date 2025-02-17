package encoding
import (
    "github.com/cbsinteractive/bitmovin-api-sdk-go/common"
    "github.com/cbsinteractive/bitmovin-api-sdk-go/query"
    "github.com/cbsinteractive/bitmovin-api-sdk-go/pagination"
)

type EncodingConfigurationsApi struct {
    apiClient *common.ApiClient
    Type *EncodingConfigurationsTypeApi
    Video *EncodingConfigurationsVideoApi
    Audio *EncodingConfigurationsAudioApi
}

func NewEncodingConfigurationsApi(configs ...func(*common.ApiClient)) (*EncodingConfigurationsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingConfigurationsApi{apiClient: apiClient}

    typeApi, err := NewEncodingConfigurationsTypeApi(configs...)
    api.Type = typeApi
    videoApi, err := NewEncodingConfigurationsVideoApi(configs...)
    api.Video = videoApi
    audioApi, err := NewEncodingConfigurationsAudioApi(configs...)
    api.Audio = audioApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingConfigurationsApi) List(queryParams ...func(*query.CodecConfigurationListQueryParams)) (*pagination.CodecConfigurationsListPagination, error) {
    queryParameters := &query.CodecConfigurationListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.CodecConfigurationsListPagination
    err := api.apiClient.Get("/encoding/configurations", &responseModel, reqParams)
    return responseModel, err
}

