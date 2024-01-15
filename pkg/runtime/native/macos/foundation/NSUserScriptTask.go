//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSUserScriptTask : objc.NSObject
*/

type NSUserScriptTask struct {
  *objc.NSObject

}

func (r *NSUserScriptTask) InitWithURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSUserScriptTask) ExecuteWithCompletionHandler() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSUserScriptTask) ScriptURL() (*NSURL, error) {
  return nil, fmt.Errorf("unimplemented")
}

