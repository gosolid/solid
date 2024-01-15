//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_ACL_PROCESS_SUBJECT_SELECTOR
*/

type CSSMACLPROCESSSUBJECTSELECTOR struct {
  Version uint16 `v8:"version"`
  Mask uint16 `v8:"mask"`
  Uid uint `v8:"uid"`
  Gid uint `v8:"gid"`
}
