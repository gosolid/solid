//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.CM2Header
*/

type CM2Header struct {
  Size uint `v8:"size"`
  CMMType uint `v8:"cMMType"`
  ProfileVersion uint `v8:"profileVersion"`
  ProfileClass uint `v8:"profileClass"`
  DataColorSpace uint `v8:"dataColorSpace"`
  ProfileConnectionSpace uint `v8:"profileConnectionSpace"`
  DateTime CMDateTime `v8:"dateTime"`
  CS2profileSignature uint `v8:"cS2profileSignature"`
  Platform uint `v8:"platform"`
  Flags uint `v8:"flags"`
  DeviceManufacturer uint `v8:"deviceManufacturer"`
  DeviceModel uint `v8:"deviceModel"`
  DeviceAttributes [2]uint `v8:"deviceAttributes"`
  RenderingIntent uint `v8:"renderingIntent"`
  White CMFixedXYZColor `v8:"white"`
  Creator uint `v8:"creator"`
  Reserved [44]byte `v8:"reserved"`
}
