//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSInputStream : Foundation.NSStream
*/

type NSInputStream struct {
  *NSStream

}

func (r *NSInputStream) InitWithURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInputStream) HasBytesAvailable() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInputStream) Read() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInputStream) GetBuffer() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInputStream) InitWithData() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

