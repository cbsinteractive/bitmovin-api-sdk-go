package encoding
import (
    "github.com/cbsinteractive/bitmovin-api-sdk-go/common"
    "github.com/cbsinteractive/bitmovin-api-sdk-go/query"
    "github.com/cbsinteractive/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMuxingsApi struct {
    apiClient *common.ApiClient
    Fmp4 *EncodingEncodingsMuxingsFmp4Api
    Cmaf *EncodingEncodingsMuxingsCmafApi
    SegmentedRaw *EncodingEncodingsMuxingsSegmentedRawApi
    Ts *EncodingEncodingsMuxingsTsApi
    Webm *EncodingEncodingsMuxingsWebmApi
    Mp3 *EncodingEncodingsMuxingsMp3Api
    Mp4 *EncodingEncodingsMuxingsMp4Api
    ProgressiveTs *EncodingEncodingsMuxingsProgressiveTsApi
    BroadcastTs *EncodingEncodingsMuxingsBroadcastTsApi
    ProgressiveWebm *EncodingEncodingsMuxingsProgressiveWebmApi
    ProgressiveMov *EncodingEncodingsMuxingsProgressiveMovApi
}

func NewEncodingEncodingsMuxingsApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMuxingsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMuxingsApi{apiClient: apiClient}

    fmp4Api, err := NewEncodingEncodingsMuxingsFmp4Api(configs...)
    api.Fmp4 = fmp4Api
    cmafApi, err := NewEncodingEncodingsMuxingsCmafApi(configs...)
    api.Cmaf = cmafApi
    segmentedRawApi, err := NewEncodingEncodingsMuxingsSegmentedRawApi(configs...)
    api.SegmentedRaw = segmentedRawApi
    tsApi, err := NewEncodingEncodingsMuxingsTsApi(configs...)
    api.Ts = tsApi
    webmApi, err := NewEncodingEncodingsMuxingsWebmApi(configs...)
    api.Webm = webmApi
    mp3Api, err := NewEncodingEncodingsMuxingsMp3Api(configs...)
    api.Mp3 = mp3Api
    mp4Api, err := NewEncodingEncodingsMuxingsMp4Api(configs...)
    api.Mp4 = mp4Api
    progressiveTsApi, err := NewEncodingEncodingsMuxingsProgressiveTsApi(configs...)
    api.ProgressiveTs = progressiveTsApi
    broadcastTsApi, err := NewEncodingEncodingsMuxingsBroadcastTsApi(configs...)
    api.BroadcastTs = broadcastTsApi
    progressiveWebmApi, err := NewEncodingEncodingsMuxingsProgressiveWebmApi(configs...)
    api.ProgressiveWebm = progressiveWebmApi
    progressiveMovApi, err := NewEncodingEncodingsMuxingsProgressiveMovApi(configs...)
    api.ProgressiveMov = progressiveMovApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMuxingsApi) List(encodingId string, queryParams ...func(*query.MuxingListQueryParams)) (*pagination.MuxingsListPagination, error) {
    queryParameters := &query.MuxingListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.MuxingsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/muxings", &responseModel, reqParams)
    return responseModel, err
}

