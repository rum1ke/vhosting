package config_stream

import (
	"sync"

	"github.com/deepch/vdk/av"
)

type Config struct {
	StreamsMutex      sync.RWMutex
	Server            Server            `json:"server"`
	Streams           map[string]Stream `json:"streams"`
	ConcatStreamPaths map[string]*[]string
	LastError         error
	StreamsCount      int
	StreamDropped     bool
}

type Server struct {
	ICEUsernameMutex   sync.RWMutex
	ICECredentialMutex sync.RWMutex
	WebRTCPortMinMutex sync.RWMutex
	WebRTCPortMaxMutex sync.RWMutex
	ICEUsername        string `json:"iceUsername"`
	ICECredential      string `json:"iceCredential"`
	WebRTCPortMin      uint16 `json:"webrtcPortMin"`
	WebRTCPortMax      uint16 `json:"webrtcPortMax"`
}

type Stream struct {
	URL          string `json:"url"`
	Status       bool   `json:"status"`
	OnDemand     bool   `json:"onDemand"`
	DisableAudio bool   `json:"disableAudio"`
	Debug        bool   `json:"debug"`
	RunLock      bool   `json:"-"`
	Codecs       []av.CodecData
	ClientList   map[string]Viewer
	Working      bool
}

type Viewer struct {
	Cast chan av.Packet
}
