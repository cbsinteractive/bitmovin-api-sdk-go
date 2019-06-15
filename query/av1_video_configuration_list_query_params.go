package query

import(
    "github.com/cbsinteractive/bitmovin-api-sdk-go/common"
)

type Av1VideoConfigurationListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

func (q *Av1VideoConfigurationListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
