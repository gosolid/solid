//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CMNativeDisplayInfo
*/

type CMNativeDisplayInfo struct {
  DataSize uint `v8:"dataSize"`
  RedPhosphor CMFixedXYColor `v8:"redPhosphor"`
  GreenPhosphor CMFixedXYColor `v8:"greenPhosphor"`
  BluePhosphor CMFixedXYColor `v8:"bluePhosphor"`
  WhitePoint CMFixedXYColor `v8:"whitePoint"`
  RedGammaValue int `v8:"redGammaValue"`
  GreenGammaValue int `v8:"greenGammaValue"`
  BlueGammaValue int `v8:"blueGammaValue"`
  GammaChannels uint16 `v8:"gammaChannels"`
  GammaEntryCount uint16 `v8:"gammaEntryCount"`
  GammaEntrySize uint16 `v8:"gammaEntrySize"`
  GammaData [1]byte `v8:"gammaData"`
}
