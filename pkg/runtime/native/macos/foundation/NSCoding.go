//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSCoding
*/

type NSCoding struct {

}

func (r *NSCoding) EncodeWithCoder() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSCoding) InitWithCoder() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

