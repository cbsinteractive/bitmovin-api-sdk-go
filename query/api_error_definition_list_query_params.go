package query

import(
    "github.com/cbsinteractive/bitmovin-api-sdk-go/common"
)

type ApiErrorDefinitionListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

func (q *ApiErrorDefinitionListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
