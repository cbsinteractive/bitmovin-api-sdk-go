package model
import (
	"time"
)

type Mp3AudioConfiguration struct {
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
	// Target bitrate for the encoded audio in bps
	Bitrate *int64 `json:"bitrate,omitempty"`
	// Audio sampling rate Hz
	Rate *float64 `json:"rate,omitempty"`
	// Channel layout of the audio codec configuration
	ChannelLayout ChannelLayout `json:"channelLayout,omitempty"`
}
func (o Mp3AudioConfiguration) CodecConfigType() CodecConfigType {
    return CodecConfigType_MP3
}

