//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMMInfo
*/

type CMMInfo struct {
  DataSize uint64 `v8:"dataSize"`
  CMMType uint `v8:"cMMType"`
  CMMMfr uint `v8:"cMMMfr"`
  CMMVersion uint `v8:"cMMVersion"`
  ASCIIName [32]byte `v8:"aSCIIName"`
  ASCIIDesc [256]byte `v8:"aSCIIDesc"`
  UniCodeNameCount uint64 `v8:"uniCodeNameCount"`
  UniCodeName [32]uint16 `v8:"uniCodeName"`
  UniCodeDescCount uint64 `v8:"uniCodeDescCount"`
  UniCodeDesc [256]uint16 `v8:"uniCodeDesc"`
}
