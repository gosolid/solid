//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSHTTPURLResponse : Foundation.NSURLResponse
*/

type NSHTTPURLResponse struct {
  *NSURLResponse

}

func (r *NSHTTPURLResponse) LocalizedStringForStatusCode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPURLResponse) StatusCode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPURLResponse) AllHeaderFields() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPURLResponse) InitWithURL() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPURLResponse) ValueForHTTPHeaderField() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

