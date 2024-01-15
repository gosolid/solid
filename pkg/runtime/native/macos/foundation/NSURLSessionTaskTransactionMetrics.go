//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
)

/*
interface Foundation.NSURLSessionTaskTransactionMetrics : objc.NSObject
*/

type NSURLSessionTaskTransactionMetrics struct {
  *objc.NSObject

}

func (r *NSURLSessionTaskTransactionMetrics) NegotiatedTLSCipherSuite() (*NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) DomainLookupStartDate() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) DomainLookupEndDate() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) LocalAddress() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) IsProxyConnection() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) CountOfResponseBodyBytesReceived() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) CountOfResponseBodyBytesAfterDecoding() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) LocalPort() (*NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) RemotePort() (*NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) DomainResolutionProtocol() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) SecureConnectionStartDate() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) RequestEndDate() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) ResponseStartDate() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) CountOfRequestBodyBytesSent() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) CountOfResponseHeaderBytesReceived() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) RequestStartDate() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) ResourceFetchType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) CountOfRequestHeaderBytesSent() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) FetchStartDate() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) SecureConnectionEndDate() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) NetworkProtocolName() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) NegotiatedTLSProtocolVersion() (*NSNumber, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) IsCellular() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) Request() (*NSURLRequest, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) ResponseEndDate() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) IsReusedConnection() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) CountOfRequestBodyBytesBeforeEncoding() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) IsExpensive() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) IsConstrained() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) IsMultipath() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) Response() (*NSURLResponse, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) ConnectStartDate() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) ConnectEndDate() (*NSDate, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) RemoteAddress() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

