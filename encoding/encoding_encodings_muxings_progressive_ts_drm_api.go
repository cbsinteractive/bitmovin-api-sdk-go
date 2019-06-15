package encoding
import (
    "github.com/cbsinteractive/bitmovin-api-sdk-go/common"
    
    "github.com/cbsinteractive/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsProgressiveTsDrmApi struct {
    apiClient *common.ApiClient
    Fairplay *EncodingEncodingsMuxingsProgressiveTsDrmFairplayApi
}

func NewEncodingEncodingsMuxingsProgressiveTsDrmApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsProgressiveTsDrmApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsProgressiveTsDrmApi{apiClient: apiClient}

    fairplayApi, err := NewEncodingEncodingsMuxingsProgressiveTsDrmFairplayApi(configs...)
    api.Fairplay = fairplayApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsProgressiveTsDrmApi) List(encodingId string, muxingId string) (*pagination.DrmsListPagination, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *pagination.DrmsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/progressive-ts/{muxing_id}/drm", &responseModel, reqParams)
    return responseModel, err
}

