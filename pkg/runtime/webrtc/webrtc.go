//js:package webrtc
package webrtc

//go:generate go run github.com/grexie/isolates/codegen

import (
	"fmt"
	"io"
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/gosolid/solid/pkg/runtime/events"
	"github.com/pion/datachannel"
	"github.com/pion/ice/v2"
	webrtc "github.com/pion/webrtc/v3"
)

var PeerConnectionStateClosed = "closed"
var PeerConnectionStateOpen = "open"
var PeerConnectionStateConnecting = "connecting"

//js:alias string
type PeerConnectionState = string
type DataChannelOptions = webrtc.DataChannelInit

//js:class
type WebRTC struct {
	api *webrtc.API
}

func NewWebRTC() *WebRTC {
	var settingEngine = webrtc.SettingEngine{}

	settingEngine.DetachDataChannels()
	settingEngine.SetIncludeLoopbackCandidate(true)
	settingEngine.SetInterfaceFilter(func(iface string) bool {
		return iface == "lo0"
	})
	settingEngine.SetICEMulticastDNSMode(ice.MulticastDNSModeDisabled)

	var api = webrtc.NewAPI(webrtc.WithSettingEngine(settingEngine))

	return &WebRTC{api}
}

type PeerConnection struct {
	mutex sync.Mutex
	ID    ID
	conn  *webrtc.PeerConnection

	iceCandidates []webrtc.ICECandidateInit
	addCandidates bool

	closed       bool
	dataChannels map[string]*DataChannel

	onConnecting            events.EventHandler
	onConnected             events.EventHandler
	onClose                 events.EventHandler
	onICECandidate          events.EventHandler
	onDataChannel           events.EventHandler
	onConnectionStateChange events.EventHandler
}

type ID string

//js:class
type SessionDescription webrtc.SessionDescription

func NewID() ID {
	return ID(uuid.NewString())
}

//js:class
type ICECandidate webrtc.ICECandidateInit

type OfferMessage struct {
	ID                 ID
	SessionDescription *SessionDescription `json:"sessionDescription"`
}

type AnswerMessage struct {
	SessionDescription *SessionDescription `json:"sessionDescription"`
}

type AnswerResponse struct {
	ID ID
}

var rtc = NewWebRTC()

//js:constructor
func NewPeerConnection(id ID, iceServers []string) (*PeerConnection, error) {
	return rtc.NewPeerConnection(id, iceServers)
}

//js:method createPeerConnection
func (rtc *WebRTC) NewPeerConnection(id ID, iceServers []string) (*PeerConnection, error) {
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: iceServers,
			},
		},
	}

	peerConnection, err := rtc.api.NewPeerConnection(config)
	if err != nil {
		return nil, err
	}

	c := &PeerConnection{
		mutex:         sync.Mutex{},
		conn:          peerConnection,
		ID:            id,
		iceCandidates: []webrtc.ICECandidateInit{},
		addCandidates: false,
		dataChannels:  map[string]*DataChannel{},
	}

	peerConnection.OnConnectionStateChange(func(pcs webrtc.PeerConnectionState) {
		if pcs == webrtc.PeerConnectionStateClosed || pcs == webrtc.PeerConnectionStateDisconnected || pcs == webrtc.PeerConnectionStateFailed {
			pcs = webrtc.PeerConnectionStateClosed
			if !c.closed {
				c.closed = true
				c.onClose.Emit()
			}
		} else if pcs == webrtc.PeerConnectionStateConnected {
			c.mutex.Lock()
			if dataChannel, ok := c.dataChannels["signalling"]; ok && dataChannel.IsOpen() {
				c.mutex.Unlock()
				// c.onConnected.Emit()
			} else {
				c.mutex.Unlock()
				pcs = webrtc.PeerConnectionStateConnecting
				return
			}
		} else {
			pcs = webrtc.PeerConnectionStateConnecting
			c.onConnecting.Emit()
		}

		pcsp := &struct{ pcs webrtc.PeerConnectionState }{pcs}
		c.onConnectionStateChange.Emit(pcsp)
	})

	peerConnection.OnICECandidate(func(i *webrtc.ICECandidate) {
		if i == nil {
			return
		}

		candidate := i.ToJSON()

		go func() {
			c.onICECandidate.Emit((*ICECandidate)(&candidate))
		}()
	})

	peerConnection.OnDataChannel(func(dc *webrtc.DataChannel) {
		stream := NewDataChannel()

		go func() {
			if err := stream.setDataChannel(dc); err != nil {
				dc.Close()
			}

			c.mutex.Lock()
			c.dataChannels[dc.Label()] = stream
			c.mutex.Unlock()

			dc.OnClose(func() {
				c.mutex.Lock()
				delete(c.dataChannels, dc.Label())
				c.mutex.Unlock()
			})

			go c.onDataChannel.Emit(stream)
		}()
	})

	return c, nil
}

//js:event state
func (c *PeerConnection) OnConnectionStateChange(callback func(c PeerConnectionState) error) func() {
	return c.onConnectionStateChange.Add(func(args ...any) error {
		pcs := args[0].(*struct{ pcs webrtc.PeerConnectionState })
		return callback(PeerConnectionState(pcs.pcs))
	})
}

//js:async-method
func (c *PeerConnection) Wait() error {
	ch := make(chan error)
	done := false

	if c.State() == PeerConnectionStateOpen {
		return nil
	} else if c.State() != PeerConnectionStateConnecting {
		return fmt.Errorf("peer connection failed to establish")
	}

	remover := c.OnConnectionStateChange(func(c PeerConnectionState) error {
		if done {
			return nil
		}

		switch c {
		case PeerConnectionStateClosed:
			done = true
			ch <- fmt.Errorf("peer connection failed to establish")
			return nil
		case PeerConnectionStateOpen:
			done = true
			ch <- nil
			return nil
		case PeerConnectionStateConnecting:
			return nil
		}
		return nil
	})

	err := <-ch
	remover()
	return err
}

//js:get
func (c *PeerConnection) State() PeerConnectionState {
	pcs := c.conn.ConnectionState()

	if pcs == webrtc.PeerConnectionStateClosed || pcs == webrtc.PeerConnectionStateDisconnected || pcs == webrtc.PeerConnectionStateFailed {
		return PeerConnectionStateClosed
	} else if pcs == webrtc.PeerConnectionStateConnected {
		c.mutex.Lock()
		if dataChannel, ok := c.dataChannels["signalling"]; ok && dataChannel.IsOpen() {
			c.mutex.Unlock()
			return PeerConnectionStateOpen
		} else {
			c.mutex.Unlock()
			return PeerConnectionStateConnecting
		}

	} else {
		return PeerConnectionStateConnecting
	}
}

//js:method
func (c *PeerConnection) CreateDataChannel(options *DataChannelOptions) (*DataChannel, error) {
	if d, err := c.conn.CreateDataChannel(uuid.NewString(), options); err != nil {
		return nil, err
	} else {
		stream := NewDataChannel()
		stream.setDataChannel(d)
		return stream, nil
	}
}

//js:event channel
func (c *PeerConnection) OnRemoteDataChannel(callback func(*DataChannel) error) func() {
	return c.onDataChannel.Add(func(args ...any) error {
		d := args[0].(*DataChannel)
		return callback(d)
	})
}

//js:method
func (c *PeerConnection) Close() error {
	err := c.conn.Close()
	if !c.closed {
		c.closed = true
		go c.onClose.Emit()
	}
	return err
}

//js:event open
func (c *PeerConnection) OnConnected(callback func() error) func() {
	return c.onConnected.Add(func(...interface{}) error {
		return callback()
	})
}

//js:event connecting
func (c *PeerConnection) OnConnecting(callback func() error) func() {
	return c.onConnecting.Add(func(...interface{}) error {
		return callback()
	})
}

//js:event close
func (c *PeerConnection) OnClose(callback func() error) func() {
	return c.onClose.Add(func(...interface{}) error {
		return callback()
	})
}

//js:event candidate
func (c *PeerConnection) OnICECandidate(callback func(i *ICECandidate) error) func() {
	return c.onICECandidate.Add(func(args ...interface{}) error {
		i := args[0].(*ICECandidate)
		return callback(i)
	})
}

//js:method
func (c *PeerConnection) CreateOffer() (*SessionDescription, error) {
	if offer, err := c.conn.CreateOffer(nil); err != nil {
		return nil, err
	} else if err := c.conn.SetLocalDescription(offer); err != nil {
		return nil, err
	} else {
		return (*SessionDescription)(&offer), err
	}
}

//js:method
func (c *PeerConnection) CreateAnswer(offer *SessionDescription) (*SessionDescription, error) {
	if err := c.conn.SetRemoteDescription(webrtc.SessionDescription(*offer)); err != nil {
		return nil, err
	} else {
		if answer, err := c.conn.CreateAnswer(nil); err != nil {
			return nil, err
		} else if err := c.conn.SetLocalDescription(answer); err != nil {
			return nil, err
		} else {
			// c.mutex.Lock()
			c.addCandidates = true
			for _, iceCandidate := range c.iceCandidates {
				c.conn.AddICECandidate(iceCandidate)
			}
			c.iceCandidates = []webrtc.ICECandidateInit{}
			// c.mutex.Unlock()
			return (*SessionDescription)(&answer), nil
		}
	}
}

func (c *PeerConnection) OnAnswer(answer *SessionDescription) error {
	// <-time.After(1 * time.Second)

	if err := c.conn.SetRemoteDescription(webrtc.SessionDescription(*answer)); err != nil {
		return err
	}

	// c.mutex.Lock()
	// defer c.mutex.Unlock()
	c.addCandidates = true
	for _, iceCandidate := range c.iceCandidates {
		c.conn.AddICECandidate(iceCandidate)
	}
	c.iceCandidates = []webrtc.ICECandidateInit{}

	return nil
}

//js:method
func (c *PeerConnection) AddICECandidate(candidate *ICECandidate) error {
	// c.mutex.Lock()
	// defer c.mutex.Unlock()
	if c.addCandidates {
		return c.conn.AddICECandidate(webrtc.ICECandidateInit(*candidate))
	} else {
		c.iceCandidates = append(c.iceCandidates, webrtc.ICECandidateInit(*candidate))
	}

	return nil
}

//js:get channels
func (c *PeerConnection) DataChannels() []*DataChannel {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	channels := []*DataChannel{}
	for _, channel := range c.dataChannels {
		channels = append(channels, channel)
	}

	return channels
}

//js:class DataChannel
type DataChannel struct {
	mutex    sync.Mutex
	stream   datachannel.ReadWriteCloser
	label    string
	overflow []byte
	open     bool
}

func NewDataChannel() *DataChannel {
	c := &DataChannel{sync.Mutex{}, nil, "", []byte{}, false}
	c.mutex.Lock()

	return c
}

//js:get
func (d *DataChannel) Label() string {
	return d.label
}

//js:get
func (d *DataChannel) IsOpen() bool {
	return d.open
}

//js:method
func (d *DataChannel) Close() error {
	d.mutex.Lock()
	d.open = false
	d.mutex.Unlock()

	if d.stream == nil {
		return io.ErrClosedPipe
	}

	return d.stream.Close()
}

//js:method
func (d *DataChannel) Read(bytes []byte) (int, error) {
	d.mutex.Lock()
	d.mutex.Unlock()

	if d.stream == nil {
		return 0, io.ErrClosedPipe
	}

	n, err := d.stream.Read(bytes)
	return n, err
}

//js:method
func (d *DataChannel) Write(bytes []byte) (int, error) {
	d.mutex.Lock()
	d.mutex.Unlock()

	if d.stream == nil {
		return 0, io.ErrClosedPipe
	}

	n, err := d.stream.Write(bytes)
	return n, err
}

func (d *DataChannel) setDataChannel(channel *webrtc.DataChannel) error {
	if d.stream != nil {
		return fmt.Errorf("channel already acquired")
	}

	done := false
	defer func() {
		done = true
	}()
	errors := make(chan error)
	d.label = channel.Label()

	channel.OnOpen(func() {
		if stream, err := channel.Detach(); err != nil {
			d.mutex.Unlock()
			if !done {
				errors <- err
			}

			return
		} else {
			d.stream = stream
			d.mutex.Unlock()
			errors <- nil
		}
		d.open = true
	})

	channel.OnError(func(err error) {
		log.Println("DATACHANNEL ERROR", err)
		if !d.open {
			d.mutex.Unlock()
		}
		if !done {
			errors <- err
		}

		d.open = false
	})

	channel.OnClose(func() {
		log.Println("DATACHANNEL CLOSE")
		if !d.open {
			d.mutex.Unlock()
		}
		if !done {
			errors <- io.ErrClosedPipe
		}
		d.open = false
	})

	err := <-errors
	close(errors)
	return err
}
