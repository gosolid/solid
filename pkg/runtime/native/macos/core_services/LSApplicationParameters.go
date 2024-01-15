//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
struct CoreServices.LSApplicationParameters
*/

type LSApplicationParameters struct {
  Version int64 `v8:"version"`
  Flags int `v8:"flags"`
  Application FSRef `v8:"application"`
  AsyncLaunchRefCon void `v8:"asyncLaunchRefCon"`
  Environment *core_foundation.CFDictionary `v8:"environment"`
  Argv *core_foundation.CFArray `v8:"argv"`
  InitialEvent AEDesc `v8:"initialEvent"`
}
