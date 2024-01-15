//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_ACL_ENTRY_PROTOTYPE_PTR
*/

type CSSMACLENTRYPROTOTYPEPTR struct {
  TypedSubject CSSMLISTPTR `v8:"typedSubject"`
  Delegate int `v8:"delegate"`
  Authorization CSSMAUTHORIZATIONGROUPPTR `v8:"authorization"`
  TimeRange CSSMACLVALIDITYPERIODPTR `v8:"timeRange"`
  EntryTag [68]byte `v8:"entryTag"`
}
