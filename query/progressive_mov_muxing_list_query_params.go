package query

import(
    "github.com/cbsinteractive/bitmovin-api-sdk-go/common"
)

type ProgressiveMovMuxingListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

func (q *ProgressiveMovMuxingListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
