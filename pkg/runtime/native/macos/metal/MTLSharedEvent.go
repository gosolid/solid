//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Metal.MTLSharedEvent
*/

type MTLSharedEvent struct {

}

func (r *MTLSharedEvent) NotifyListener() error {
  return fmt.Errorf("unimplemented")
}

func (r *MTLSharedEvent) NewSharedEventHandle() (*MTLSharedEventHandle, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *MTLSharedEvent) SignaledValue() (uint64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *MTLSharedEvent) SetSignaledValue() error {
  return fmt.Errorf("unimplemented")
}

