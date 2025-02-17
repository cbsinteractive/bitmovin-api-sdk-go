package model
import (
	"time"
)

type BroadcastTsMuxing struct {
	// Name of the resource. Can be freely chosen by the user.
	Name string `json:"name,omitempty"`
	// Description of the resource. Can be freely chosen by the user.
	Description string `json:"description,omitempty"`
	// Creation timestamp formatted in UTC: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Modified timestamp formatted in UTC: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]map[string]interface{} `json:"customData,omitempty"`
	// Id of the resource
	Id string `json:"id,omitempty"`
	Streams []MuxingStream `json:"streams,omitempty"`
	Outputs []EncodingOutput `json:"outputs,omitempty"`
	// Average bitrate. Available after encoding finishes.
	AvgBitrate *int64 `json:"avgBitrate,omitempty"`
	// Min bitrate. Available after encoding finishes.
	MinBitrate *int64 `json:"minBitrate,omitempty"`
	// Max bitrate. Available after encoding finishes.
	MaxBitrate *int64 `json:"maxBitrate,omitempty"`
	// If this is set and contains objects, then this muxing has been ignored during the encoding process
	IgnoredBy []Ignoring `json:"ignoredBy,omitempty"`
	// Specifies how to handle streams that don't fulfill stream conditions
	StreamConditionsMode StreamConditionsMode `json:"streamConditionsMode,omitempty"`
	// Length of the segments in seconds.
	SegmentLength *float64 `json:"segmentLength,omitempty"`
	// Name of the new Video
	Filename string `json:"filename,omitempty"`
	Configuration *BroadcastTsMuxingConfiguration `json:"configuration,omitempty"`
	// Modifies the internal chunk length used for chunked encoding
	InternalChunkLength *InternalChunkLength `json:"internalChunkLength,omitempty"`
}
func (o BroadcastTsMuxing) MuxingType() MuxingType {
    return MuxingType_BROADCAST_TS
}

