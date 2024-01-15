//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSScriptWhoseTest : objc.NSObject
*/

type NSScriptWhoseTest struct {
  *objc.NSObject
  *NSCoding
}

func (r *NSScriptWhoseTest) IsTrue() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSScriptWhoseTest) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSScriptWhoseTest) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

