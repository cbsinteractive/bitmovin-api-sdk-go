package pagination

import(
    "encoding/json"
    "github.com/cbsinteractive/bitmovin-api-sdk-go/serialization"
    "github.com/cbsinteractive/bitmovin-api-sdk-go/model"
)

type FtpInputsListPagination struct {
	TotalCount *int64           `json:"totalCount,omitempty"`
	Offset     *int32           `json:"offset,omitempty"`
	Limit      *int32           `json:"limit,omitempty"`
	Previous   string           `json:"previous,omitempty"`
	Next       string           `json:"next,omitempty"`
	Items      []model.FtpInput `json:"items,omitempty"`
}


  func (o *FtpInputsListPagination) UnmarshalJSON(b []byte) error {
    var items []model.FtpInput

    var pageResp model.PaginationResponse
    if err := json.Unmarshal(b, &pageResp); err != nil {
      return err
    }

    for _, i := range pageResp.Items {
      var v model.FtpInput
      serialization.Decode(i, &v)
      items = append(items, v)
    }

    o.TotalCount = pageResp.TotalCount
    o.Offset = pageResp.Offset
    o.Limit = pageResp.Limit
    o.Previous = pageResp.Previous
    o.Next = pageResp.Next
    o.Items = items
    return nil
  }
