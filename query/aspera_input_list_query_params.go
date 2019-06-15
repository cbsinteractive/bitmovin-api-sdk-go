package query

import(
    "github.com/cbsinteractive/bitmovin-api-sdk-go/common"
)

type AsperaInputListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

func (q *AsperaInputListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
