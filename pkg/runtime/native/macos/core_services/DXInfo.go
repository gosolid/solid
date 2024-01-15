//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
struct CoreServices.DXInfo
*/

type DXInfo struct {
  FrScroll objc.Point `v8:"frScroll"`
  FrOpenChain int `v8:"frOpenChain"`
  FrScript byte `v8:"frScript"`
  FrXFlags byte `v8:"frXFlags"`
  FrComment int16 `v8:"frComment"`
  FrPutAway int `v8:"frPutAway"`
}
