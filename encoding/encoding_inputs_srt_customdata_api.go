package encoding
import (
    "github.com/cbsinteractive/bitmovin-api-sdk-go/common"
    
    "github.com/cbsinteractive/bitmovin-api-sdk-go/model"
)

type EncodingInputsSrtCustomdataApi struct {
    apiClient *common.ApiClient
}

func NewEncodingInputsSrtCustomdataApi(configs ...func(*common.ApiClient)) (*EncodingInputsSrtCustomdataApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingInputsSrtCustomdataApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingInputsSrtCustomdataApi) Get(inputId string) (*model.CustomData, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["input_id"] = inputId
    }

    var responseModel *model.CustomData
    err := api.apiClient.Get("/encoding/inputs/srt/{input_id}/customData", &responseModel, reqParams)
    return responseModel, err
}

