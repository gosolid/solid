//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMProfileIterateData
*/

type CMProfileIterateData struct {
  DataVersion uint `v8:"dataVersion"`
  Header CM2Header `v8:"header"`
  Code int16 `v8:"code"`
  Name [256]byte `v8:"name"`
  Location CMProfileLocation `v8:"location"`
  UniCodeNameCount uint64 `v8:"uniCodeNameCount"`
  UniCodeName uint16 `v8:"uniCodeName"`
  AsciiName byte `v8:"asciiName"`
  MakeAndModel CMMakeAndModel `v8:"makeAndModel"`
  Digest [16]byte `v8:"digest"`
}
