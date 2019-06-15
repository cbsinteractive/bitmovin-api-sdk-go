package encoding
import (
    "github.com/cbsinteractive/bitmovin-api-sdk-go/common"
    "github.com/cbsinteractive/bitmovin-api-sdk-go/query"
    "github.com/cbsinteractive/bitmovin-api-sdk-go/model"
    "github.com/cbsinteractive/bitmovin-api-sdk-go/pagination"
)

type EncodingManifestsSmoothApi struct {
    apiClient *common.ApiClient
    Default *EncodingManifestsSmoothDefaultApi
    Customdata *EncodingManifestsSmoothCustomdataApi
    Representations *EncodingManifestsSmoothRepresentationsApi
    Contentprotection *EncodingManifestsSmoothContentprotectionApi
}

func NewEncodingManifestsSmoothApi(configs ...func(*common.ApiClient)) (*EncodingManifestsSmoothApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingManifestsSmoothApi{apiClient: apiClient}

    defaultApi, err := NewEncodingManifestsSmoothDefaultApi(configs...)
    api.Default = defaultApi
    customdataApi, err := NewEncodingManifestsSmoothCustomdataApi(configs...)
    api.Customdata = customdataApi
    representationsApi, err := NewEncodingManifestsSmoothRepresentationsApi(configs...)
    api.Representations = representationsApi
    contentprotectionApi, err := NewEncodingManifestsSmoothContentprotectionApi(configs...)
    api.Contentprotection = contentprotectionApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingManifestsSmoothApi) Create(smoothStreamingManifest model.SmoothStreamingManifest) (*model.SmoothStreamingManifest, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.SmoothStreamingManifest
    err := api.apiClient.Post("/encoding/manifests/smooth", &smoothStreamingManifest, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsSmoothApi) Delete(manifestId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/encoding/manifests/smooth/{manifest_id}", &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsSmoothApi) Get(manifestId string) (*model.SmoothStreamingManifest, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel *model.SmoothStreamingManifest
    err := api.apiClient.Get("/encoding/manifests/smooth/{manifest_id}", &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsSmoothApi) List(queryParams ...func(*query.SmoothStreamingManifestListQueryParams)) (*pagination.SmoothStreamingManifestsListPagination, error) {
    queryParameters := &query.SmoothStreamingManifestListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.SmoothStreamingManifestsListPagination
    err := api.apiClient.Get("/encoding/manifests/smooth", &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsSmoothApi) Start(manifestId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Post("/encoding/manifests/smooth/{manifest_id}/start", struct{}{}, &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsSmoothApi) Status(manifestId string) (*model.ModelTask, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel *model.ModelTask
    err := api.apiClient.Get("/encoding/manifests/smooth/{manifest_id}/status", &responseModel, reqParams)
    return responseModel, err
}

func (api *EncodingManifestsSmoothApi) Stop(manifestId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["manifest_id"] = manifestId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Post("/encoding/manifests/smooth/{manifest_id}/stop", struct{}{}, &responseModel, reqParams)
    return responseModel, err
}

