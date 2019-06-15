package query

import(
    "github.com/cbsinteractive/bitmovin-api-sdk-go/common"
)

type StatisticsListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

func (q *StatisticsListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
