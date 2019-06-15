package query

import(
    "github.com/cbsinteractive/bitmovin-api-sdk-go/model"
    "github.com/cbsinteractive/bitmovin-api-sdk-go/common"
)

type MuxingListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    StreamMode model.StreamMode `query:"streamMode"`
}

func (q *MuxingListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
