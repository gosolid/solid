//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Foundation.NSURLSessionConfiguration : objc.NSObject
*/

type NSURLSessionConfiguration struct {
  *objc.NSObject

}

func (r *NSURLSessionConfiguration) SetTimeoutIntervalForResource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SharedContainerIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetHTTPShouldSetCookies() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetURLCredentialStorage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) NetworkServiceType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) ConnectionProxyDictionary() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) TLSMinimumSupportedProtocol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) BackgroundSessionConfigurationWithIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) Identifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) WaitsForConnectivity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetTLSMinimumSupportedProtocolVersion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) HTTPMaximumConnectionsPerHost() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) HTTPCookieStorage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) DefaultSessionConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetSharedContainerIdentifier() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) AllowsConstrainedNetworkAccess() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) HTTPCookieAcceptPolicy() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetHTTPMaximumConnectionsPerHost() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) URLCache() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) TimeoutIntervalForRequest() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetRequiresDNSSECValidation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetTLSMinimumSupportedProtocol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) HTTPAdditionalHeaders() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetProtocolClasses() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) EphemeralSessionConfiguration() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetWaitsForConnectivity() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetSessionSendsLaunchEvents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) TLSMaximumSupportedProtocol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) TimeoutIntervalForResource() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetNetworkServiceType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SessionSendsLaunchEvents() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetMultipathServiceType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) RequiresDNSSECValidation() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) IsDiscretionary() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) TLSMinimumSupportedProtocolVersion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetRequestCachePolicy() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) AllowsCellularAccess() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) ProtocolClasses() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) MultipathServiceType() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetDiscretionary() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetTLSMaximumSupportedProtocolVersion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) HTTPShouldSetCookies() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetHTTPCookieAcceptPolicy() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetHTTPAdditionalHeaders() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) URLCredentialStorage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetAllowsCellularAccess() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) AllowsExpensiveNetworkAccess() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetAllowsExpensiveNetworkAccess() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetTLSMaximumSupportedProtocol() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) TLSMaximumSupportedProtocolVersion() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) HTTPShouldUsePipelining() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetHTTPShouldUsePipelining() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) ShouldUseExtendedBackgroundIdleMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) RequestCachePolicy() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetAllowsConstrainedNetworkAccess() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetTimeoutIntervalForRequest() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetConnectionProxyDictionary() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetHTTPCookieStorage() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetURLCache() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetShouldUseExtendedBackgroundIdleMode() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

