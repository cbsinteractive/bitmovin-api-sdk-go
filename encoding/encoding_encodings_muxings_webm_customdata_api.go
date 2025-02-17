package encoding
import (
    "github.com/cbsinteractive/bitmovin-api-sdk-go/common"
    
    "github.com/cbsinteractive/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsMuxingsWebmCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsMuxingsWebmCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsWebmCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsWebmCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsWebmCustomdataApi) Get(encodingId string, muxingId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/webm/{muxing_id}/customData", &responseModel, reqParams)
    return responseModel, err
}

