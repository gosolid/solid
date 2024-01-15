//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSURLConnectionDelegate
*/

type NSURLConnectionDelegate struct {

}

func (r *NSURLConnectionDelegate) Connection() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLConnectionDelegate) ConnectionShouldUseCredentialStorage() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

