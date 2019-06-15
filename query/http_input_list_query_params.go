package query

import(
    "github.com/cbsinteractive/bitmovin-api-sdk-go/common"
)

type HttpInputListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

func (q *HttpInputListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
