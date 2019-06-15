package query

import(
    "github.com/cbsinteractive/bitmovin-api-sdk-go/common"
)

type AudioMediaInfoListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

func (q *AudioMediaInfoListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
