package lastFmUser

import (
	"fmt"
	lastFm "github.com/Cellularhacker/last-fm-go"
	"github.com/goccy/go-json"
	"net/http"
	"net/url"
)

const (
	ValueMethodUserGetInfo = "user.getinfo"
)

type GetInfoRequest struct {
	User *string `json:"user,omitempty" form:"user,omitempty" param:"user" query:"user,omitempty"`
}

func NewGetInfoRequest() *GetInfoRequest {
	return &GetInfoRequest{}
}
func (gir *GetInfoRequest) SetUsername(username string) *GetInfoRequest {
	gir.User = &username

	return gir
}
func (gir *GetInfoRequest) GetUsername() string {
	if gir.User == nil {
		return ""
	}

	return *gir.User
}
func (gir *GetInfoRequest) IsZeroUsername() bool {
	return len(gir.GetUsername()) <= 0
}
func (gir *GetInfoRequest) IsZero() bool {
	return gir.IsZeroUsername()
}
func (gir *GetInfoRequest) ToQuery() url.Values {
	p := url.Values{}
	p.Set(QueryKeyMethod, ValueMethodUserGetInfo)
	if !gir.IsZeroUsername() {
		p.Set(QueryKeyUser, gir.GetUsername())
	}

	return p
}

type GetInfoResponse struct {
	User User `json:"user"`
}
type User struct {
	Name        string     `json:"name"`
	Age         string     `json:"age"`
	Subscriber  string     `json:"subscriber"`
	Realname    string     `json:"realname"`
	Bootstrap   string     `json:"bootstrap"`
	Playcount   string     `json:"playcount"`
	ArtistCount string     `json:"artist_count"`
	Playlists   string     `json:"playlists"`
	TrackCount  string     `json:"track_count"`
	AlbumCount  string     `json:"album_count"`
	Image       []Image    `json:"image"`
	Registered  Registered `json:"registered"`
	Country     string     `json:"country"`
	Gender      string     `json:"gender"`
	Url         string     `json:"url"`
	Type        string     `json:"type"`
}
type Registered struct {
	Unixtime string `json:"unixtime"`
	Text     int    `json:"#text"`
}

func NewGetInfoResponse() *GetInfoResponse {
	return &GetInfoResponse{}
}

func GetInfo(request *GetInfoRequest) (*GetInfoResponse, error) {
	resp, err := lastFm.MakeRequest(http.MethodGet, request.ToQuery(), nil)
	if err != nil {
		return nil, fmt.Errorf("lastFm.MakeRequest(http.MethodGet, request.ToQuery(), nil): %w", err)
	}

	respBody := NewGetInfoResponse()
	err = json.Unmarshal(resp, respBody)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal([]byte(resp), respBody): %w", err)
	}

	return respBody, nil
}
