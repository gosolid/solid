//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSURLSessionTaskTransactionMetrics : objc.NSObject
*/

type NSURLSessionTaskTransactionMetrics struct {
  *objc.NSObject

}

func (r *NSURLSessionTaskTransactionMetrics) Request() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) SecureConnectionStartDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) RequestEndDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) CountOfRequestBodyBytesSent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) CountOfResponseHeaderBytesReceived() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) RemoteAddress() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) ConnectStartDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) RequestStartDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) LocalPort() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) IsConstrained() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) IsMultipath() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) ConnectEndDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) ResponseEndDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) NetworkProtocolName() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) CountOfResponseBodyBytesAfterDecoding() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) DomainLookupStartDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) DomainLookupEndDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) NegotiatedTLSCipherSuite() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) IsCellular() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) DomainResolutionProtocol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) Response() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) IsReusedConnection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) ResourceFetchType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) CountOfRequestHeaderBytesSent() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) CountOfResponseBodyBytesReceived() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) RemotePort() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) FetchStartDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) ResponseStartDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) LocalAddress() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) SecureConnectionEndDate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) IsProxyConnection() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) CountOfRequestBodyBytesBeforeEncoding() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) NegotiatedTLSProtocolVersion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionTaskTransactionMetrics) IsExpensive() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

