package encoding
import (
    "github.com/cbsinteractive/bitmovin-api-sdk-go/common"
    
    "github.com/cbsinteractive/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsWebmDrmApi struct {
    apiClient *common.ApiClient
    Cenc *EncodingEncodingsMuxingsWebmDrmCencApi
}

func NewEncodingEncodingsMuxingsWebmDrmApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsWebmDrmApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsWebmDrmApi{apiClient: apiClient}

    cencApi, err := NewEncodingEncodingsMuxingsWebmDrmCencApi(configs...)
    api.Cenc = cencApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsWebmDrmApi) List(encodingId string, muxingId string) (*pagination.DrmsListPagination, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *pagination.DrmsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/drm", &responseModel, reqParams)
    return responseModel, err
}

