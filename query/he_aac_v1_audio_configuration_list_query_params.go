package query

import(
    "github.com/cbsinteractive/bitmovin-api-sdk-go/common"
)

type HeAacV1AudioConfigurationListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

func (q *HeAacV1AudioConfigurationListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
