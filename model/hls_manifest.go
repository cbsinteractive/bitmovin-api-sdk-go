package model
import (
	"time"
)

type HlsManifest struct {
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
	Type ManifestType `json:"type,omitempty"`
	// The outputs to store the manifest
	Outputs []EncodingOutput `json:"outputs,omitempty"`
	// The filename of your manifest
	ManifestName string `json:"manifestName,omitempty"`
	// If this is set, the EXT-X-VERSION tags of the Media Playlists are set to the provided version
	HlsMediaPlaylistVersion HlsVersion `json:"hlsMediaPlaylistVersion,omitempty"`
	// If this is set, the EXT-X-VERSION tag of the Master Playlist is set to the provided version
	HlsMasterPlaylistVersion HlsVersion `json:"hlsMasterPlaylistVersion,omitempty"`
}

