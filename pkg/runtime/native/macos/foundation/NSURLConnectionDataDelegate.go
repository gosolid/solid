//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSURLConnectionDataDelegate
*/

type NSURLConnectionDataDelegate struct {

}

func (r *NSURLConnectionDataDelegate) ConnectionDidFinishLoading() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLConnectionDataDelegate) Connection() (*NSURLRequest, error) {
  return nil, fmt.Errorf("unimplemented")
}

