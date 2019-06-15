package query

import(
    "github.com/cbsinteractive/bitmovin-api-sdk-go/common"
)

type SpriteListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

func (q *SpriteListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
