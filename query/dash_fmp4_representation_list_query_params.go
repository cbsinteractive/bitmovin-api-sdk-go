package query

import(
    "github.com/cbsinteractive/bitmovin-api-sdk-go/common"
)

type DashFmp4RepresentationListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

func (q *DashFmp4RepresentationListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
