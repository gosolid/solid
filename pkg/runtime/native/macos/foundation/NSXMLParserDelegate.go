//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
protocol Foundation.NSXMLParserDelegate
*/

type NSXMLParserDelegate struct {

}

func (r *NSXMLParserDelegate) ParserDidStartDocument() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLParserDelegate) ParserDidEndDocument() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSXMLParserDelegate) Parser() error {
  return fmt.Errorf("unimplemented")
}

