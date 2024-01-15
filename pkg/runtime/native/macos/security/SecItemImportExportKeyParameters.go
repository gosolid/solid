//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/core_foundation"
)

/*
struct Security.SecItemImportExportKeyParameters
*/

type SecItemImportExportKeyParameters struct {
  Version uint `v8:"version"`
  Flags int `v8:"flags"`
  Passphrase *void `v8:"passphrase"`
  AlertTitle *core_foundation.CFString `v8:"alertTitle"`
  AlertPrompt *core_foundation.CFString `v8:"alertPrompt"`
  AccessRef *SecAccess `v8:"accessRef"`
  KeyUsage *core_foundation.CFArray `v8:"keyUsage"`
  KeyAttributes *core_foundation.CFArray `v8:"keyAttributes"`
}
