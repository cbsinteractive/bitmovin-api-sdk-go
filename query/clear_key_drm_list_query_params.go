package query

import(
    "github.com/cbsinteractive/bitmovin-api-sdk-go/common"
)

type ClearKeyDrmListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

func (q *ClearKeyDrmListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
