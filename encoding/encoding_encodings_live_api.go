package encoding
import (
    "github.com/cbsinteractive/bitmovin-api-sdk-go/common"
    
    "github.com/cbsinteractive/bitmovin-api-sdk-go/model"
)

type EncodingEncodingsLiveApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsLiveApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsLiveApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsLiveApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsLiveApi) Get(encodingId string) (*model.LiveEncoding, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.LiveEncoding
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/live", &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsLiveApi) GetStartRequest(encodingId string) (*model.StartLiveEncodingRequest, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.StartLiveEncodingRequest
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/live/start", &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsLiveApi) Restart(encodingId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/live/restart", struct{}{}, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsLiveApi) Start(encodingId string, startLiveEncodingRequest model.StartLiveEncodingRequest) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/live/start", &startLiveEncodingRequest, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingEncodingsLiveApi) Stop(encodingId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Post("/encoding/encodings/{encoding_id}/live/stop", struct{}{}, &responseModel, reqParams)
    return responseModel, err
}

