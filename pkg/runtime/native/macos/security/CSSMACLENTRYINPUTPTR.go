//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_ACL_ENTRY_INPUT_PTR
*/

type CSSMACLENTRYINPUTPTR struct {
  Prototype CSSMACLENTRYPROTOTYPEPTR `v8:"prototype"`
  Callback *func(...any) (any, error) `v8:"callback"`
  CallerContext void `v8:"callerContext"`
}
