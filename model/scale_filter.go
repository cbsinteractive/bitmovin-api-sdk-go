package model
import (
	"time"
)

type ScaleFilter struct {
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
	// The width of the output frame in pixel. If not set: codec configuration width will be used.
	Width *int32 `json:"width,omitempty"`
	// The height of the output frame in pixel. If not set: codec configuration height will be used.
	Height *int32 `json:"height,omitempty"`
	ScalingAlgorithm ScalingAlgorithm `json:"scalingAlgorithm,omitempty"`
	// The numerator of the sample aspect ratio (also known as pixel aspect ratio). Must be set if sampleAspectRatioDenominator is set.
	SampleAspectRatioNumerator *int32 `json:"sampleAspectRatioNumerator,omitempty"`
	// The denominator of the sample aspect ratio (also known as pixel aspect ratio). Must be set if sampleAspectRatioNumerator is set.
	SampleAspectRatioDenominator *int32 `json:"sampleAspectRatioDenominator,omitempty"`
}
func (o ScaleFilter) FilterType() FilterType {
    return FilterType_SCALE
}

