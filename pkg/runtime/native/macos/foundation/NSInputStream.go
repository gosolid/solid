//js:package native/macos/foundation
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

func (r *NSInputStream) GetBuffer() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSInputStream) InitWithData() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInputStream) InitWithURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSInputStream) HasBytesAvailable() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSInputStream) Read() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

