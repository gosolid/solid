//js:package native/ios/open-gles
package open_gles

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface OpenGLES.EAGLContext : objc.NSObject
*/

type EAGLContext struct {
  *objc.NSObject

}

func (r *EAGLContext) SetDebugLabel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *EAGLContext) IsMultiThreaded() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *EAGLContext) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *EAGLContext) InitWithAPI() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *EAGLContext) SetCurrentContext() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *EAGLContext) API() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *EAGLContext) DebugLabel() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *EAGLContext) CurrentContext() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *EAGLContext) Sharegroup() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *EAGLContext) SetMultiThreaded() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

