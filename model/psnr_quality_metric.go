package model

type PsnrQualityMetric struct {
	TimeSpan *TimeSpan `json:"timeSpan,omitempty"`
	// Peak signal-to-noise ratio
	Psnr *float64 `json:"psnr,omitempty"`
}

