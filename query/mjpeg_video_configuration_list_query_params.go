package query

import(
    "github.com/cbsinteractive/bitmovin-api-sdk-go/common"
)

type MjpegVideoConfigurationListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

func (q *MjpegVideoConfigurationListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
