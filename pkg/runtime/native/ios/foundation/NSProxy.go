//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSProxy
*/

type NSProxy struct {
  

}

func (r *NSProxy) Class() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProxy) ForwardInvocation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProxy) Dealloc() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProxy) Finalize() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProxy) Alloc() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProxy) MethodSignatureForSelector() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProxy) RespondsToSelector() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProxy) AllowsWeakReference() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProxy) RetainWeakReference() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProxy) Description() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProxy) DebugDescription() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSProxy) AllocWithZone() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

