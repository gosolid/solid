//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSAppleScript : objc.NSObject
*/

type NSAppleScript struct {
  *objc.NSObject
  *NSCopying
}

func (r *NSAppleScript) InitWithContentsOfURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleScript) InitWithSource() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleScript) CompileAndReturnError() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSAppleScript) ExecuteAndReturnError() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleScript) ExecuteAppleEvent() (*NSAppleEventDescriptor, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleScript) Source() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSAppleScript) IsCompiled() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

