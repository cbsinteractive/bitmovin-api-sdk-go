package model
import (
	"time"
)

type ConvertSccCaption struct {
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
	// The input location to get the scc file from
	Input *InputPath `json:"input,omitempty"`
	Outputs []EncodingOutput `json:"outputs,omitempty"`
	// Name of the captions file
	FileName string `json:"fileName,omitempty"`
	OutputFormat StreamCaptionOutputFormat `json:"outputFormat,omitempty"`
	// Optional settings when converting SCC to WebVTT
	WebVttSettings *ConvertSccCaptionWebVttSettings `json:"webVttSettings,omitempty"`
}

