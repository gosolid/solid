module github.com/gosolid/solid

go 1.18

require (
	github.com/google/uuid v1.3.0
	github.com/grexie/isolates v0.1.0
	github.com/hashicorp/yamux v0.1.1
	github.com/keybase/go-keychain v0.0.0-20231219164618-57a3676c3af6
	github.com/pion/datachannel v1.5.5
	github.com/pion/ice/v2 v2.3.10
	github.com/pion/webrtc/v3 v3.2.16
	golang.org/x/crypto v0.17.0
	golang.org/x/sys v0.15.0
	golang.org/x/term v0.15.0
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/grexie/isolates v0.1.0 => ../grexie-isolates

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/grexie/refutils v0.1.1 // indirect
	github.com/pion/dtls/v2 v2.2.7 // indirect
	github.com/pion/interceptor v0.1.17 // indirect
	github.com/pion/logging v0.2.2 // indirect
	github.com/pion/mdns v0.0.7 // indirect
	github.com/pion/randutil v0.1.0 // indirect
	github.com/pion/rtcp v1.2.10 // indirect
	github.com/pion/rtp v1.8.1 // indirect
	github.com/pion/sctp v1.8.7 // indirect
	github.com/pion/sdp/v3 v3.0.6 // indirect
	github.com/pion/srtp/v2 v2.0.16 // indirect
	github.com/pion/stun v0.6.1 // indirect
	github.com/pion/transport/v2 v2.2.1 // indirect
	github.com/pion/turn/v2 v2.1.3 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	golang.org/x/net v0.15.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/iancoleman/orderedmap => github.com/tbehrsin/orderedmap v0.0.0-20231006180158-896cb0db4f33
