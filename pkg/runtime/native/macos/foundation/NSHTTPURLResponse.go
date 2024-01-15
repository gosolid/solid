//js:package native/macos/foundation
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

func (r *NSHTTPURLResponse) InitWithURL() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPURLResponse) ValueForHTTPHeaderField() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPURLResponse) LocalizedStringForStatusCode() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSHTTPURLResponse) StatusCode() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSHTTPURLResponse) AllHeaderFields() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

