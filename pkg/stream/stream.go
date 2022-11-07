package stream

import (
	"database/sql"
	"vhosting/pkg/user"

	"github.com/deepch/vdk/av"
	webrtc "github.com/deepch/vdk/format/webrtcv3"
	"github.com/gin-gonic/gin"
)

type Stream struct {
	Id           int    `json:"id" db:"id"`
	Stream       string `json:"stream" db:"Stream"`
	DateTime     string `json:"dateTime" db:"DateTime"`
	StatePublic  int    `json:"-" db:"StatePublic"`
	StatusPublic int    `json:"statusPublic" db:"StatusPublic"`
	StatusRecord int    `json:"-" db:"StatusRecord"`
	PathStream   string `json:"pathStream" db:"pathStream"`
}

type StreamGet struct {
	Id           int            `db:"id"`
	Stream       sql.NullString `db:"Stream"`
	DateTime     sql.NullString `db:"DateTime"`
	StatePublic  sql.NullInt16  `db:"StatePublic"`
	StatusPublic sql.NullInt16  `db:"StatusPublic"`
	StatusRecord sql.NullInt16  `db:"StatusRecord"`
	PathStream   sql.NullString `db:"pathStream"`
}

type JCodec struct {
	Type string
}

type Response struct {
	Tracks []string `json:"tracks"`
	Sdp64  string   `json:"sdp64"`
}

type StreamCommon interface {
	IsStreamExists(id int) (bool, error)
}

type StreamUseCase interface {
	StreamCommon

	GetStream(id int) (*Stream, error)
	GetAllStreams(urlparams *user.Pagin) (map[int]*Stream, error)

	AtoiRequestedId(ctx *gin.Context) (int, error)
	ParseURLParams(ctx *gin.Context) *user.Pagin

	ServeStreams()
	Exit(suuid string) bool
	RunIfNotRun(uuid string)
	CodecGet(suuid string) []av.CodecData
	GetICEServers() []string
	GetICEUsername() string
	GetICECredential() string
	GetWebRTCPortMin() uint16
	GetWebRTCPortMax() uint16
	WritePackets(url string, muxerWebRTC *webrtc.Muxer, audioOnly bool)
	CastListAdd(suuid string) (string, chan av.Packet)
	CastListDelete(suuid, cuuid string)
	List() (string, []string)
}

type StreamRepository interface {
	StreamCommon

	GetStream(id int) (*StreamGet, error)
	GetAllStreams(urlparams *user.Pagin) (map[int]*StreamGet, error)
	GetAllWorkingStreams() (*[]string, error)
}
