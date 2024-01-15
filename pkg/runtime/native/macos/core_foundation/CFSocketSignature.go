//js:package native/macos/core-foundation
package core_foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreFoundation.CFSocketSignature
*/

type CFSocketSignature struct {
  ProtocolFamily int `v8:"protocolFamily"`
  SocketType int `v8:"socketType"`
  Protocol int `v8:"protocol"`
  Address *CFData `v8:"address"`
}
