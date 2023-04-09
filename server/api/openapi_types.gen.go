// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.5-0.20230403173426-fd06f5aed350 DO NOT EDIT.
package api

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

// Defines values for ArtistRole.
const (
	ArtistRoleAlbumArtist ArtistRole = "albumArtist"
	ArtistRoleArtist      ArtistRole = "artist"
)

// Defines values for ResourceType.
const (
	ResourceTypeAlbum  ResourceType = "album"
	ResourceTypeArtist ResourceType = "artist"
	ResourceTypeTrack  ResourceType = "track"
)

// Defines values for ServerInfoFeatures.
const (
	Albums  ServerInfoFeatures = "albums"
	Artists ServerInfoFeatures = "artists"
	Images  ServerInfoFeatures = "images"
)

// Album defines model for Album.
type Album struct {
	Attributes *AlbumAttributes `json:"attributes,omitempty"`

	// Id The unique identifier for the resource
	Id            string `json:"id"`
	Relationships *struct {
		Artists []AlbumArtistRelationship `json:"artists"`
		Tracks  []AlbumTrackRelationship  `json:"tracks"`
	} `json:"relationships,omitempty"`

	// Type The type of the resource
	Type ResourceType `json:"type"`
}

// AlbumArtistRelationship defines model for AlbumArtistRelationship.
type AlbumArtistRelationship struct {
	Data ResourceObject `json:"data"`
	Meta interface{}    `json:"meta"`
}

// AlbumAttributes defines model for AlbumAttributes.
type AlbumAttributes struct {
	// Genre The genre of the album
	Genre string `json:"genre"`

	// ReleaseDate The release date of the album
	ReleaseDate openapi_types.Date `json:"releaseDate"`

	// Title The title of the album
	Title string `json:"title"`
}

// AlbumTrackRelationship defines model for AlbumTrackRelationship.
type AlbumTrackRelationship struct {
	Data ResourceObject `json:"data"`
}

// Artist defines model for Artist.
type Artist struct {
	Attributes *ArtistAttributes `json:"attributes,omitempty"`

	// Id The unique identifier for the resource
	Id            string `json:"id"`
	Relationships *struct {
		Albums *struct {
			Data []ArtistAlbumRelationship `json:"data"`
		} `json:"albums,omitempty"`
		Tracks *struct {
			Data []ArtistTrackRelationship `json:"data"`
		} `json:"tracks,omitempty"`
	} `json:"relationships,omitempty"`

	// Type The type of the resource
	Type ResourceType `json:"type"`
}

// ArtistAlbumRelationship defines model for ArtistAlbumRelationship.
type ArtistAlbumRelationship struct {
	Data ResourceObject   `json:"data"`
	Meta ArtistMetaObject `json:"meta"`
}

// ArtistAttributes defines model for ArtistAttributes.
type ArtistAttributes struct {
	// Bio A short biography of the artist
	Bio *string `json:"bio,omitempty"`

	// Name The name of the artist
	Name *string `json:"name,omitempty"`
}

// ArtistMetaObject defines model for ArtistMetaObject.
type ArtistMetaObject struct {
	// Role The role of an artist in a track or album
	Role ArtistRole `json:"role"`
}

// ArtistRole The role of an artist in a track or album
type ArtistRole string

// ArtistTrackRelationship defines model for ArtistTrackRelationship.
type ArtistTrackRelationship struct {
	Data ResourceObject `json:"data"`
	Meta struct {
		// Role The role of an artist in a track or album
		Role ArtistRole `json:"role"`
	} `json:"meta"`
}

// ErrorList defines model for ErrorList.
type ErrorList struct {
	Errors []ErrorObject `json:"errors"`
}

// ErrorObject defines model for ErrorObject.
type ErrorObject struct {
	Detail *string `json:"detail,omitempty"`
	Id     *string `json:"id,omitempty"`
	Status *string `json:"status,omitempty"`
	Title  *string `json:"title,omitempty"`
}

// PaginationLinks defines model for PaginationLinks.
type PaginationLinks struct {
	First *string `json:"first,omitempty"`
	Last  *string `json:"last,omitempty"`
	Next  *string `json:"next,omitempty"`
	Prev  *string `json:"prev,omitempty"`
}

// PaginationMeta defines model for PaginationMeta.
type PaginationMeta struct {
	// CurrentPage The current page in the collection
	CurrentPage *int32 `json:"currentPage,omitempty"`

	// TotalItems The total number of items in the collection
	TotalItems *int32 `json:"totalItems,omitempty"`

	// TotalPages The total numeber of pages in the collection
	TotalPages *int32 `json:"totalPages,omitempty"`
}

// ResourceObject defines model for ResourceObject.
type ResourceObject struct {
	// Id The unique identifier for the resource
	Id string `json:"id"`

	// Type The type of the resource
	Type ResourceType `json:"type"`
}

// ResourceType The type of the resource
type ResourceType string

// ServerInfo defines model for ServerInfo.
type ServerInfo struct {
	// AuthRequired Whether the user has access to the server.
	AuthRequired bool `json:"authRequired"`

	// Features A list of optional features the server supports.
	Features []ServerInfoFeatures `json:"features"`

	// Server The name of the server software.
	Server string `json:"server"`

	// ServerVersion The version number of the server.
	ServerVersion string `json:"serverVersion"`
}

// ServerInfoFeatures defines model for ServerInfo.Features.
type ServerInfoFeatures string

// Track defines model for Track.
type Track struct {
	Attributes *TrackAttributes `json:"attributes,omitempty"`

	// Id The unique identifier for the resource
	Id            string              `json:"id"`
	Relationships *TrackRelationships `json:"relationships,omitempty"`

	// Type The type of the resource
	Type ResourceType `json:"type"`
}

// TrackArtistRelationship defines model for TrackArtistRelationship.
type TrackArtistRelationship struct {
	Data ResourceObject `json:"data"`
	Meta interface{}    `json:"meta"`
}

// TrackAttributes defines model for TrackAttributes.
type TrackAttributes struct {
	// Album The name of the album the track belongs to
	Album string `json:"album"`

	// Albumartist The primary artist of the album the track belongs to.
	Albumartist string `json:"albumartist"`

	// Artist The name of the artist who performed the track
	Artist string `json:"artist"`

	// Bitrate The bitrate of the audio file in kilobits per second (kbps).
	Bitrate int `json:"bitrate"`

	// Bpm The beats per minute (BPM) of the track.
	Bpm *int `json:"bpm,omitempty"`

	// Channels The number of audio channels in the track.
	Channels int `json:"channels"`

	// Comments Any additional comments or notes about the track.
	Comments *string `json:"comments,omitempty"`

	// Disc The disc number within a multi-disc album.
	Disc *int `json:"disc,omitempty"`

	// Duration The duration of the track in seconds
	Duration float32 `json:"duration"`

	// Genre The genre of the track.
	Genre *string `json:"genre,omitempty"`

	// Mimetype The MIME type of the audio file.
	Mimetype string `json:"mimetype"`

	// RecordingMbid The MusicBrainz identifier for the recording of the track.
	RecordingMbid *string `json:"recording-mbid,omitempty"`

	// Size The size of the audio file in bytes.
	Size int `json:"size"`

	// Title The title of the track
	Title string `json:"title"`

	// Track The track number within the album.
	Track int `json:"track"`

	// TrackMbid The MusicBrainz identifier for the track.
	TrackMbid *string `json:"track-mbid,omitempty"`

	// Year The release year of the track or album.
	Year *int `json:"year,omitempty"`
}

// TrackRelationships defines model for TrackRelationships.
type TrackRelationships struct {
	Albums  *[]AlbumTrackRelationship `json:"albums,omitempty"`
	Artists []TrackArtistRelationship `json:"artists"`
}

// FilterContains defines model for filterContains.
type FilterContains = []string

// FilterEndsWith defines model for filterEndsWith.
type FilterEndsWith = []string

// FilterEquals defines model for filterEquals.
type FilterEquals = []string

// FilterGreaterOrEqual defines model for filterGreaterOrEqual.
type FilterGreaterOrEqual = []string

// FilterGreaterThan defines model for filterGreaterThan.
type FilterGreaterThan = []string

// FilterLessOrEqual defines model for filterLessOrEqual.
type FilterLessOrEqual = []string

// FilterLessThan defines model for filterLessThan.
type FilterLessThan = []string

// FilterStartsWith defines model for filterStartsWith.
type FilterStartsWith = []string

// Include defines model for include.
type Include = string

// PageLimit defines model for pageLimit.
type PageLimit = int32

// PageOffset defines model for pageOffset.
type PageOffset = int32

// Sort defines model for sort.
type Sort = string

// GetAlbumsParams defines parameters for GetAlbums.
type GetAlbumsParams struct {
	// PageLimit The number of items per page
	PageLimit *PageLimit `form:"page[limit],omitempty" json:"page[limit],omitempty"`

	// PageOffset The offset for pagination
	PageOffset *PageOffset `form:"page[offset],omitempty" json:"page[offset],omitempty"`

	// FilterEquals Filter by any property with an exact match. Usage: filter[equals]=property:value
	FilterEquals *FilterEquals `form:"filter[equals],omitempty" json:"filter[equals],omitempty"`

	// FilterContains Filter by any property containing text. Usage: filter[contains]=property:value
	FilterContains *FilterContains `form:"filter[contains],omitempty" json:"filter[contains],omitempty"`

	// FilterLessThan Filter by any numeric property less than a value. Usage: filter[lessThan]=property:value
	FilterLessThan *FilterLessThan `form:"filter[lessThan],omitempty" json:"filter[lessThan],omitempty"`

	// FilterLessOrEqual Filter by any numeric property less than or equal to a value. Usage: filter[lessOrEqual]=property:value
	FilterLessOrEqual *FilterLessOrEqual `form:"filter[lessOrEqual],omitempty" json:"filter[lessOrEqual],omitempty"`

	// FilterGreaterThan Filter by any numeric property greater than a value. Usage: filter[greaterThan]=property:value
	FilterGreaterThan *FilterGreaterThan `form:"filter[greaterThan],omitempty" json:"filter[greaterThan],omitempty"`

	// FilterGreaterOrEqual Filter by any numeric property greater than or equal to a value. Usage: filter[greaterOrEqual]=property:value
	FilterGreaterOrEqual *FilterGreaterOrEqual `form:"filter[greaterOrEqual],omitempty" json:"filter[greaterOrEqual],omitempty"`

	// FilterStartsWith Filter by any property that starts with text. Usage: filter[startsWith]=property:value
	FilterStartsWith *FilterStartsWith `form:"filter[startsWith],omitempty" json:"filter[startsWith],omitempty"`

	// FilterEndsWith Filter by any property that ends with text. Usage: filter[endsWith]=property:value
	FilterEndsWith *FilterEndsWith `form:"filter[endsWith],omitempty" json:"filter[endsWith],omitempty"`

	// Sort Sort the results by one or more properties, separated by commas. Prefix the property with '-' for descending order.
	Sort *Sort `form:"sort,omitempty" json:"sort,omitempty"`

	// Include Related resources to include in the response, separated by commas
	Include *Include `form:"include,omitempty" json:"include,omitempty"`
}

// GetAlbumParams defines parameters for GetAlbum.
type GetAlbumParams struct {
	// Include Related resources to include in the response, separated by commas
	Include *Include `form:"include,omitempty" json:"include,omitempty"`
}

// GetArtistsParams defines parameters for GetArtists.
type GetArtistsParams struct {
	// PageLimit The number of items per page
	PageLimit *PageLimit `form:"page[limit],omitempty" json:"page[limit],omitempty"`

	// PageOffset The offset for pagination
	PageOffset *PageOffset `form:"page[offset],omitempty" json:"page[offset],omitempty"`

	// FilterEquals Filter by any property with an exact match. Usage: filter[equals]=property:value
	FilterEquals *FilterEquals `form:"filter[equals],omitempty" json:"filter[equals],omitempty"`

	// FilterContains Filter by any property containing text. Usage: filter[contains]=property:value
	FilterContains *FilterContains `form:"filter[contains],omitempty" json:"filter[contains],omitempty"`

	// FilterLessThan Filter by any numeric property less than a value. Usage: filter[lessThan]=property:value
	FilterLessThan *FilterLessThan `form:"filter[lessThan],omitempty" json:"filter[lessThan],omitempty"`

	// FilterLessOrEqual Filter by any numeric property less than or equal to a value. Usage: filter[lessOrEqual]=property:value
	FilterLessOrEqual *FilterLessOrEqual `form:"filter[lessOrEqual],omitempty" json:"filter[lessOrEqual],omitempty"`

	// FilterGreaterThan Filter by any numeric property greater than a value. Usage: filter[greaterThan]=property:value
	FilterGreaterThan *FilterGreaterThan `form:"filter[greaterThan],omitempty" json:"filter[greaterThan],omitempty"`

	// FilterGreaterOrEqual Filter by any numeric property greater than or equal to a value. Usage: filter[greaterOrEqual]=property:value
	FilterGreaterOrEqual *FilterGreaterOrEqual `form:"filter[greaterOrEqual],omitempty" json:"filter[greaterOrEqual],omitempty"`

	// FilterStartsWith Filter by any property that starts with text. Usage: filter[startsWith]=property:value
	FilterStartsWith *FilterStartsWith `form:"filter[startsWith],omitempty" json:"filter[startsWith],omitempty"`

	// FilterEndsWith Filter by any property that ends with text. Usage: filter[endsWith]=property:value
	FilterEndsWith *FilterEndsWith `form:"filter[endsWith],omitempty" json:"filter[endsWith],omitempty"`

	// Sort Sort the results by one or more properties, separated by commas. Prefix the property with '-' for descending order.
	Sort *Sort `form:"sort,omitempty" json:"sort,omitempty"`

	// Include Related resources to include in the response, separated by commas
	Include *Include `form:"include,omitempty" json:"include,omitempty"`
}

// GetArtistParams defines parameters for GetArtist.
type GetArtistParams struct {
	// Include Related resources to include in the response, separated by commas
	Include *Include `form:"include,omitempty" json:"include,omitempty"`
}

// GetTracksParams defines parameters for GetTracks.
type GetTracksParams struct {
	// PageLimit The number of items per page
	PageLimit *PageLimit `form:"page[limit],omitempty" json:"page[limit],omitempty"`

	// PageOffset The offset for pagination
	PageOffset *PageOffset `form:"page[offset],omitempty" json:"page[offset],omitempty"`

	// FilterEquals Filter by any property with an exact match. Usage: filter[equals]=property:value
	FilterEquals *FilterEquals `form:"filter[equals],omitempty" json:"filter[equals],omitempty"`

	// FilterContains Filter by any property containing text. Usage: filter[contains]=property:value
	FilterContains *FilterContains `form:"filter[contains],omitempty" json:"filter[contains],omitempty"`

	// FilterLessThan Filter by any numeric property less than a value. Usage: filter[lessThan]=property:value
	FilterLessThan *FilterLessThan `form:"filter[lessThan],omitempty" json:"filter[lessThan],omitempty"`

	// FilterLessOrEqual Filter by any numeric property less than or equal to a value. Usage: filter[lessOrEqual]=property:value
	FilterLessOrEqual *FilterLessOrEqual `form:"filter[lessOrEqual],omitempty" json:"filter[lessOrEqual],omitempty"`

	// FilterGreaterThan Filter by any numeric property greater than a value. Usage: filter[greaterThan]=property:value
	FilterGreaterThan *FilterGreaterThan `form:"filter[greaterThan],omitempty" json:"filter[greaterThan],omitempty"`

	// FilterGreaterOrEqual Filter by any numeric property greater than or equal to a value. Usage: filter[greaterOrEqual]=property:value
	FilterGreaterOrEqual *FilterGreaterOrEqual `form:"filter[greaterOrEqual],omitempty" json:"filter[greaterOrEqual],omitempty"`

	// FilterStartsWith Filter by any property that starts with text. Usage: filter[startsWith]=property:value
	FilterStartsWith *FilterStartsWith `form:"filter[startsWith],omitempty" json:"filter[startsWith],omitempty"`

	// FilterEndsWith Filter by any property that ends with text. Usage: filter[endsWith]=property:value
	FilterEndsWith *FilterEndsWith `form:"filter[endsWith],omitempty" json:"filter[endsWith],omitempty"`

	// Sort Sort the results by one or more properties, separated by commas. Prefix the property with '-' for descending order.
	Sort *Sort `form:"sort,omitempty" json:"sort,omitempty"`

	// Include Related resources to include in the response, separated by commas
	Include *Include `form:"include,omitempty" json:"include,omitempty"`
}

// GetTrackParams defines parameters for GetTrack.
type GetTrackParams struct {
	// Include Related resources to include in the response, separated by commas
	Include *Include `form:"include,omitempty" json:"include,omitempty"`
}
