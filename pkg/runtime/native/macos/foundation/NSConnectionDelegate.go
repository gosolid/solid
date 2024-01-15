//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSConnectionDelegate
*/

type NSConnectionDelegate struct {

}

func (r *NSConnectionDelegate) MakeNewConnection() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSConnectionDelegate) Connection() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSConnectionDelegate) AuthenticationDataForComponents() (*NSData, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSConnectionDelegate) AuthenticateComponents() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSConnectionDelegate) CreateConversationForConnection() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

