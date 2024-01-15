//js:package native/ios/open-gles
package open_gles

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface OpenGLES.EAGLSharegroup : objc.NSObject
*/

type EAGLSharegroup struct {
  *objc.NSObject

}

func (r *EAGLSharegroup) DebugLabel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *EAGLSharegroup) SetDebugLabel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

