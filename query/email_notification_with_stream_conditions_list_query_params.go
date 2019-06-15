package query

import(
    "github.com/cbsinteractive/bitmovin-api-sdk-go/common"
)

type EmailNotificationWithStreamConditionsListQueryParams struct {
    Offset int32 `query:"offset"`
    Limit int32 `query:"limit"`
}

func (q *EmailNotificationWithStreamConditionsListQueryParams) Params() map[string]string {
    return common.GetParamsMap(q)
}
