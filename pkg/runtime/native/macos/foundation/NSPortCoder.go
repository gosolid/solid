//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSPortCoder : Foundation.NSCoder
*/

type NSPortCoder struct {
  *NSCoder

}

func (r *NSPortCoder) InitWithReceivePort() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPortCoder) Dispatch() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPortCoder) IsBycopy() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPortCoder) IsByref() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSPortCoder) EncodePortObject() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSPortCoder) DecodePortObject() (*NSPort, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPortCoder) Connection() (*NSConnection, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSPortCoder) PortCoderWithReceivePort() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

