package query

import(
    "github.com/cbsinteractive/bitmovin-api-sdk-go/common"
)

type UdpMulticastInputListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
    Name string `query:"name"`
}

func (q *UdpMulticastInputListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
