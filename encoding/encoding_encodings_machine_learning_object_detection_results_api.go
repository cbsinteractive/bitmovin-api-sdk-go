package encoding
import (
    "github.com/cbsinteractive/bitmovin-api-sdk-go/common"
    "github.com/cbsinteractive/bitmovin-api-sdk-go/query"
    "github.com/cbsinteractive/bitmovin-api-sdk-go/pagination"
)

type EncodingEncodingsMachineLearningObjectDetectionResultsApi struct {
    apiClient *common.ApiClient
}

func NewEncodingEncodingsMachineLearningObjectDetectionResultsApi(configs ...func(*common.ApiClient)) (*EncodingEncodingsMachineLearningObjectDetectionResultsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &EncodingEncodingsMachineLearningObjectDetectionResultsApi{apiClient: apiClient}


	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *EncodingEncodingsMachineLearningObjectDetectionResultsApi) List(encodingId string, objectDetectionId string, queryParams ...func(*query.ObjectDetectionResultListQueryParams)) (*pagination.ObjectDetectionResultsListPagination, error) {
    queryParameters := &query.ObjectDetectionResultListQueryParams{}
	for _, queryParam := range queryParams {
		queryParam(queryParameters)
    }

    reqParams := func(params *common.RequestParams) {
        params.PathParams["encoding_id"] = encodingId
        params.PathParams["object_detection_id"] = objectDetectionId
        params.QueryParams = queryParameters
    }

    var responseModel *pagination.ObjectDetectionResultsListPagination
    err := api.apiClient.Get("/encoding/encodings/{encoding_id}/machine-learning/object-detection/{object_detection_id}/results", &responseModel, reqParams)
    return responseModel, err
}

