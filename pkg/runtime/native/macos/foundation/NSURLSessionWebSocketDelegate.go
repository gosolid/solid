//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSURLSessionWebSocketDelegate
*/

type NSURLSessionWebSocketDelegate struct {

}

func (r *NSURLSessionWebSocketDelegate) URLSession() error {
  return fmt.Errorf("unimplemented")
}

