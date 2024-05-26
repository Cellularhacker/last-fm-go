package lastFmUser

import (
	"fmt"
	"net/url"
)

const (
	ValueMethodUserGetLovedTracks = "user.getLovedTracks"
)

type GetLovedTracksRequest struct {
	User  *string `json:"user,omitempty" form:"user,omitempty" param:"user" query:"user,omitempty"`
	Limit *int64  `json:"limit,omitempty" form:"limit,omitempty" param:"limit" query:"limit,omitempty"`
	Page  *int64  `json:"page,omitempty" form:"page,omitempty" param:"page" query:"page,omitempty"`
}

func NewGetLovedTrackRequest() *GetLovedTracksRequest {
	return &GetLovedTracksRequest{}
}
func (gltr *GetLovedTracksRequest) SetUsername(username string) *GetLovedTracksRequest {
	gltr.User = &username

	return gltr
}
func (gltr *GetLovedTracksRequest) GetUsername() string {
	if gltr.User == nil {
		return ""
	}

	return *gltr.User
}
func (gltr *GetLovedTracksRequest) IsZeroUsername() bool {
	return len(gltr.GetUsername()) <= 0
}
func (gltr *GetLovedTracksRequest) SetLimit(limit int64) *GetLovedTracksRequest {
	gltr.Limit = &limit

	return gltr
}
func (gltr *GetLovedTracksRequest) GetLimit() int64 {
	if gltr.Limit == nil {
		return 0
	}

	return *gltr.Limit
}
func (gltr *GetLovedTracksRequest) StringLimit() string {
	return fmt.Sprintf("%d", gltr.GetLimit())
}
func (gltr *GetLovedTracksRequest) IsZeroLimit() bool {
	return gltr.GetLimit() <= 0
}
func (gltr *GetLovedTracksRequest) SetPage(page int64) *GetLovedTracksRequest {
	gltr.Page = &page

	return gltr
}
func (gltr *GetLovedTracksRequest) GetPage() int64 {
	if gltr.Page == nil {
		return 0
	}

	return *gltr.Page
}
func (gltr *GetLovedTracksRequest) StringPage() string {
	return fmt.Sprintf("%d", gltr.GetPage())
}
func (gltr *GetLovedTracksRequest) IsZeroPage() bool {
	return gltr.GetPage() <= 0
}
func (gltr *GetLovedTracksRequest) ToQuery() url.Values {
	p := url.Values{}
	p.Set(QueryKeyUser, ValueMethodUserGetLovedTracks)

	if !gltr.IsZeroUsername() {
		p.Set(QueryKeyUser, gltr.GetUsername())
	}
	if !gltr.IsZeroLimit() {
		p.Set(QueryKeyLimit, gltr.StringLimit())
	}
	if !gltr.IsZeroPage() {
		p.Set(QueryKeyPage, gltr.StringPage())
	}

	return p
}
