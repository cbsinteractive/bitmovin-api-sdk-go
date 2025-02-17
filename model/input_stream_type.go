package model
type InputStreamType string

// List of InputStreamType
const (
	InputStreamType_INGEST InputStreamType = "INGEST"
	InputStreamType_CONCATENATION InputStreamType = "CONCATENATION"
	InputStreamType_TRIMMING_TIME_BASED InputStreamType = "TRIMMING_TIME_BASED"
	InputStreamType_TRIMMING_TIME_CODE_TRACK InputStreamType = "TRIMMING_TIME_CODE_TRACK"
	InputStreamType_TRIMMING_H264_PICTURE_TIMING InputStreamType = "TRIMMING_H264_PICTURE_TIMING"
	InputStreamType_AUDIO_MIX InputStreamType = "AUDIO_MIX"
)