//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_ACL_ENTRY_INFO_PTR
*/

type CSSMACLENTRYINFOPTR struct {
  EntryPublicInfo CSSMACLENTRYPROTOTYPEPTR `v8:"entryPublicInfo"`
  EntryHandle int64 `v8:"entryHandle"`
}
