package lastFmUser

import (
	"fmt"
	lastFm "github.com/Cellularhacker/last-fm-go"
	"github.com/goccy/go-json"
	"net/http"
	"net/url"
)

const (
	ValueMethodUserRecentTracks = "user.getRecentTracks"

	QueryKeyMethod   = "method"
	QueryKeyUser     = "user"
	QueryKeyLimit    = "limit"
	QueryKeyPage     = "page"
	QueryKeyFrom     = "from"
	QueryKeyTo       = "to"
	QueryKeyExtended = "extended"

	QueryValueExtendedFalse = "0"
	QueryValueExtendedTrue  = "1"
)

// GetRecentTracksRequest
// limit (Optional) : The number of results to fetch per page. Defaults to 50. Maximum is 200.
// user (Required) : The last.fm username to fetch the recent tracks of.
// page (Optional) : The page number to fetch. Defaults to first page.
// from (Optional) : Beginning timestamp of a range - only display scrobbles after this time, in UNIX timestamp format (integer number of seconds since 00:00:00, January 1st 1970 UTC). This must be in the UTC time zone.
// extended (0|1) (Optional) : Includes extended data in each artist, and whether or not the user has loved each track
// to (Optional) : End timestamp of a range - only display scrobbles before this time, in UNIX timestamp format (integer number of seconds since 00:00:00, January 1st 1970 UTC). This must be in the UTC time zone.
// api_key
type GetRecentTracksRequest struct {
	User     *string `json:"user,omitempty" form:"user,omitempty" param:"user" query:"user,omitempty"`
	Limit    *int64  `json:"limit,omitempty" form:"limit,omitempty" param:"limit" query:"limit,omitempty"`
	Page     *int64  `json:"page,omitempty" form:"page,omitempty" param:"page" query:"page,omitempty"`
	From     *int64  `json:"from,omitempty" form:"from,omitempty" param:"from" query:"from,omitempty"`
	To       *int64  `json:"to,omitempty" form:"to,omitempty" param:"to" query:"to,omitempty"`
	Extended *bool   `json:"extended,omitempty" form:"extended,omitempty" query:"extended,omitempty"`
}

func NewGetRecentTracksRequest() *GetRecentTracksRequest {
	return &GetRecentTracksRequest{}
}
func (grtr *GetRecentTracksRequest) SetUsername(username string) *GetRecentTracksRequest {
	grtr.User = &username

	return grtr
}
func (grtr *GetRecentTracksRequest) GetUsername() string {
	if grtr.User == nil {
		return ""
	}

	return *grtr.User
}
func (grtr *GetRecentTracksRequest) IsZeroUsername() bool {
	return grtr.User == nil || len(*grtr.User) <= 0
}
func (grtr *GetRecentTracksRequest) StringUsername() string {
	return grtr.GetUsername()
}
func (grtr *GetRecentTracksRequest) SetLimit(limit int64) *GetRecentTracksRequest {
	grtr.Limit = &limit

	return grtr
}
func (grtr *GetRecentTracksRequest) GetLimit() int64 {
	if grtr.Limit == nil {
		return 0
	}

	return *grtr.Limit
}
func (grtr *GetRecentTracksRequest) IsZeroLimit() bool {
	return grtr.Limit == nil
}
func (grtr *GetRecentTracksRequest) StringLimit() string {
	return fmt.Sprintf("%d", grtr.GetLimit())
}
func (grtr *GetRecentTracksRequest) SetPage(page int64) *GetRecentTracksRequest {
	grtr.Page = &page

	return grtr
}
func (grtr *GetRecentTracksRequest) GetPage() int64 {
	if grtr.Page == nil {
		return 0
	}

	return *grtr.Page
}
func (grtr *GetRecentTracksRequest) IsZeroPage() bool {
	return grtr.Page == nil
}
func (grtr *GetRecentTracksRequest) StringPage() string {
	return fmt.Sprintf("%d", grtr.GetPage())
}
func (grtr *GetRecentTracksRequest) SetFrom(from int64) *GetRecentTracksRequest {
	grtr.From = &from

	return grtr
}
func (grtr *GetRecentTracksRequest) GetFrom() int64 {
	if grtr.From == nil {
		return 0
	}

	return *grtr.From
}
func (grtr *GetRecentTracksRequest) IsZeroFrom() bool {
	return grtr.From == nil || *grtr.From <= 0
}
func (grtr *GetRecentTracksRequest) StringFrom() string {
	return fmt.Sprintf("%d", grtr.GetFrom())
}
func (grtr *GetRecentTracksRequest) SetTo(to int64) *GetRecentTracksRequest {
	grtr.To = &to

	return grtr
}
func (grtr *GetRecentTracksRequest) GetTo() int64 {
	if grtr.To == nil {
		return 0
	}

	return *grtr.To
}
func (grtr *GetRecentTracksRequest) IsZeroTo() bool {
	return grtr.To == nil || *grtr.To <= 0
}
func (grtr *GetRecentTracksRequest) StringTo() string {
	return fmt.Sprintf("%d", grtr.GetTo())
}
func (grtr *GetRecentTracksRequest) SetExtended(extended bool) *GetRecentTracksRequest {
	grtr.Extended = &extended

	return grtr
}
func (grtr *GetRecentTracksRequest) GetExtended() bool {
	if grtr.Extended == nil {
		return false
	}

	return *grtr.Extended
}
func (grtr *GetRecentTracksRequest) IsZeroExtended() bool {
	return grtr.Extended == nil || *grtr.Extended == false
}
func (grtr *GetRecentTracksRequest) StringExtended() string {
	if !grtr.GetExtended() {
		return QueryValueExtendedFalse
	}

	return QueryValueExtendedTrue
}
func (grtr *GetRecentTracksRequest) IsZero() bool {
	return grtr.IsZeroUsername() && grtr.IsZeroLimit() && grtr.IsZeroFrom() && grtr.IsZeroTo() && grtr.IsZeroExtended()
}
func (grtr *GetRecentTracksRequest) ToQuery() url.Values {
	p := url.Values{}
	p.Set(QueryKeyMethod, ValueMethodUserRecentTracks)

	if !grtr.IsZeroUsername() {
		p.Set(QueryKeyUser, grtr.GetUsername())
	}
	if !grtr.IsZeroLimit() {
		p.Set(QueryKeyLimit, grtr.StringLimit())
	}
	if !grtr.IsZeroPage() {
		p.Set(QueryKeyPage, grtr.StringPage())
	}
	if !grtr.IsZeroFrom() {
		p.Set(QueryKeyFrom, grtr.StringFrom())
	}
	if !grtr.IsZeroTo() {
		p.Set(QueryKeyTo, grtr.StringTo())
	}
	if !grtr.IsZeroExtended() {
		p.Set(QueryKeyExtended, grtr.StringExtended())
	}

	return p
}

type GetRecentTracksResponse struct {
	Recenttracks RecentTracks `json:"recenttracks"`
}
type RecentTracks struct {
	Track []RecentTrack `json:"track"`
	Attr  PagingInfo    `json:"@attr"`
}
type RecentTrack struct {
	Artist     MetaData        `json:"artist"`
	Streamable string          `json:"streamable"`
	Image      []Image         `json:"image"`
	Mbid       string          `json:"mbid"`
	Album      MetaData        `json:"album"`
	Name       string          `json:"name"`
	Attr       *AttrNowplaying `json:"@attr,omitempty"`
	Url        string          `json:"url"`
	Date       Date            `json:"date,omitempty"`
}
type PagingInfo struct {
	User       string `json:"user"`
	TotalPages string `json:"totalPages"`
	Page       string `json:"page"`
	PerPage    string `json:"perPage"`
	Total      string `json:"total"`
}
type MetaData struct {
	Url  string `json:"url,omitempty"`
	Mbid string `json:"mbid"`
	Text string `json:"#text"`
}
type Image struct {
	Size string `json:"size"`
	Text string `json:"#text"`
}
type AttrNowplaying struct {
	Nowplaying string `json:"nowplaying"`
}
type Date struct {
	UnixTimestamp string `json:"uts"`
	Text          string `json:"#text"`
}

func NewGetRecentTracksResponse() *GetRecentTracksResponse {
	return &GetRecentTracksResponse{}
}

func (a *AttrNowplaying) IsZero() bool {
	return a == nil || len(a.Nowplaying) <= 0 || a.Nowplaying != "true"
}

func GetRecentTracks(request *GetRecentTracksRequest) (*GetRecentTracksResponse, error) {
	resp, err := lastFm.MakeRequest(http.MethodGet, request.ToQuery(), nil)
	if err != nil {
		return nil, fmt.Errorf("lastFm.MakeRequest(http.MethodGet, request.ToQuery(), nil): %w", err)
	}

	respBody := NewGetRecentTracksResponse()
	err = json.Unmarshal(resp, respBody)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal(resp, respBody): %w", err)
	}

	return respBody, nil
}
