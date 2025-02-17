package model

type BasicMediaInfo struct {
	// Id of the resource
	Id string `json:"id,omitempty"`
	// The value is a quoted-string which specifies the group to which the Rendition belongs.
	GroupId string `json:"groupId,omitempty"`
	// Primary language in the rendition.
	Language string `json:"language,omitempty"`
	// Identifies a language that is associated with the Rendition.
	AssocLanguage string `json:"assocLanguage,omitempty"`
	// Human readable description of the rendition.
	Name string `json:"name,omitempty"`
	// If set to true, the client SHOULD play this Rendition of the content in the absence of information from the user.
	IsDefault *bool `json:"isDefault,omitempty"`
	// If set to true, the client MAY choose to play this Rendition in the absence of explicit user preference.
	Autoselect *bool `json:"autoselect,omitempty"`
	// Contains Uniform Type Identifiers
	Characteristics []string `json:"characteristics,omitempty"`
}

