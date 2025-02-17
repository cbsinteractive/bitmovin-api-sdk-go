package query

import(
    "github.com/cbsinteractive/bitmovin-api-sdk-go/common"
)

type Vp9VideoConfigurationListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

func (q *Vp9VideoConfigurationListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
