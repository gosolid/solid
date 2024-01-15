//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSURLProtocolClient
*/

type NSURLProtocolClient struct {

}

func (r *NSURLProtocolClient) URLProtocol() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLProtocolClient) URLProtocolDidFinishLoading() error {
  return fmt.Errorf("unimplemented")
}

