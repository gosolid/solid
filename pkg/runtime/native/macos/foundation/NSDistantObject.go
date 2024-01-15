//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSDistantObject : Foundation.NSProxy
*/

type NSDistantObject struct {
  *NSProxy
  *NSCoding
}

func (r *NSDistantObject) InitWithTarget() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDistantObject) ProxyWithLocal() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDistantObject) InitWithLocal() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDistantObject) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDistantObject) SetProtocolForProxy() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSDistantObject) ConnectionForProxy() (*NSConnection, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSDistantObject) ProxyWithTarget() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

