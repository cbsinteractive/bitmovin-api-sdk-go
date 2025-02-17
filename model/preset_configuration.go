package model
type PresetConfiguration string

// List of PresetConfiguration
const (
	PresetConfiguration_LIVE_HIGH_QUALITY PresetConfiguration = "LIVE_HIGH_QUALITY"
	PresetConfiguration_LIVE_LOW_LATENCY PresetConfiguration = "LIVE_LOW_LATENCY"
	PresetConfiguration_VOD_HIGH_QUALITY PresetConfiguration = "VOD_HIGH_QUALITY"
	PresetConfiguration_VOD_HIGH_SPEED PresetConfiguration = "VOD_HIGH_SPEED"
	PresetConfiguration_VOD_SPEED PresetConfiguration = "VOD_SPEED"
	PresetConfiguration_VOD_STANDARD PresetConfiguration = "VOD_STANDARD"
	PresetConfiguration_VOD_EXTRAHIGH_SPEED PresetConfiguration = "VOD_EXTRAHIGH_SPEED"
	PresetConfiguration_VOD_VERYHIGH_SPEED PresetConfiguration = "VOD_VERYHIGH_SPEED"
	PresetConfiguration_VOD_SUPERHIGH_SPEED PresetConfiguration = "VOD_SUPERHIGH_SPEED"
	PresetConfiguration_VOD_ULTRAHIGH_SPEED PresetConfiguration = "VOD_ULTRAHIGH_SPEED"
)