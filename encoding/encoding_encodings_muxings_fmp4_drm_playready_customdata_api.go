package encoding
import (
    "github.com/cbsinteractive/bitmovin-api-sdk-go/common"
    
    "github.com/cbsinteractive/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsMuxingsFmp4DrmPlayreadyCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsMuxingsFmp4DrmPlayreadyCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsFmp4DrmPlayreadyCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsFmp4DrmPlayreadyCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsFmp4DrmPlayreadyCustomdataApi) Get(encodingId string, muxingId string, drmId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["muxing_id"] = muxingId
        params.PathParams["drm_id"] = drmId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings/fmp4/{muxing_id}/drm/playready/{drm_id}/customData", &responseModel, reqParams)
    return responseModel, err
}

