package model
import (
	"time"
)

type SmoothStreamingRepresentation struct {
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
	// Id of the encoding
	EncodingId string `json:"encodingId,omitempty"`
	// Id of the muxing.
	MuxingId string `json:"muxingId,omitempty"`
	// The Smooth Streaming ismv or isma file that will be referenced in the manifest.
	MediaFile string `json:"mediaFile,omitempty"`
	// Language of the MP4 file
	Language string `json:"language,omitempty"`
	// Track where this MP4 shoudl be added
	TrackName string `json:"trackName,omitempty"`
}

