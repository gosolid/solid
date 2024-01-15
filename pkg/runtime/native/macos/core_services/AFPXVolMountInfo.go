//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.AFPXVolMountInfo
*/

type AFPXVolMountInfo struct {
  Length int16 `v8:"length"`
  Media uint `v8:"media"`
  Flags int16 `v8:"flags"`
  NbpInterval byte `v8:"nbpInterval"`
  NbpCount byte `v8:"nbpCount"`
  UamType int16 `v8:"uamType"`
  ZoneNameOffset int16 `v8:"zoneNameOffset"`
  ServerNameOffset int16 `v8:"serverNameOffset"`
  VolNameOffset int16 `v8:"volNameOffset"`
  UserNameOffset int16 `v8:"userNameOffset"`
  UserPasswordOffset int16 `v8:"userPasswordOffset"`
  VolPasswordOffset int16 `v8:"volPasswordOffset"`
  ExtendedFlags int16 `v8:"extendedFlags"`
  UamNameOffset int16 `v8:"uamNameOffset"`
  AlternateAddressOffset int16 `v8:"alternateAddressOffset"`
  AFPData [176]byte `v8:"aFPData"`
}
