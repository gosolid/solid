//js:package native/macos/app-kit
package app_kit

//go:generate go run github.com/grexie/isolates/codegen

/*
interface AppKit.NSSecureTextField : AppKit.NSTextField
*/

type NSSecureTextField struct {
  *NSTextField

}

