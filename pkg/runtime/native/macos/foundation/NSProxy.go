//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSProxy
*/

type NSProxy struct {
  
  *objc.NSObject
}

func (r *NSProxy) Alloc() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProxy) Dealloc() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProxy) RespondsToSelector() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSProxy) Description() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProxy) AllowsWeakReference() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSProxy) RetainWeakReference() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSProxy) DebugDescription() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProxy) AllocWithZone() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProxy) Class() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProxy) ForwardInvocation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSProxy) MethodSignatureForSelector() (*NSMethodSignature, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProxy) Finalize() error {
  return fmt.Errorf("unimplemented")
}

