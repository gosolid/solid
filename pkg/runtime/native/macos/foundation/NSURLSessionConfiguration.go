//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/macos/objc"
  "fmt"
)

/*
interface Foundation.NSURLSessionConfiguration : objc.NSObject
*/

type NSURLSessionConfiguration struct {
  *objc.NSObject
  *NSCopying
}

func (r *NSURLSessionConfiguration) HTTPShouldUsePipelining() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetHTTPShouldUsePipelining() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetHTTPShouldSetCookies() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetMultipathServiceType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) IsDiscretionary() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) WaitsForConnectivity() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) TLSMaximumSupportedProtocol() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetTLSMinimumSupportedProtocolVersion() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetHTTPCookieAcceptPolicy() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) HTTPCookieStorage() (*NSHTTPCookieStorage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) DefaultSessionConfiguration() (*NSURLSessionConfiguration, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetAllowsExpensiveNetworkAccess() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) NetworkServiceType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetWaitsForConnectivity() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) TLSMinimumSupportedProtocolVersion() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) URLCredentialStorage() (*NSURLCredentialStorage, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) ShouldUseExtendedBackgroundIdleMode() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) TLSMinimumSupportedProtocol() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) New() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) EphemeralSessionConfiguration() (*NSURLSessionConfiguration, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) AllowsExpensiveNetworkAccess() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetDiscretionary() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetHTTPCookieStorage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) BackgroundSessionConfigurationWithIdentifier() (*NSURLSessionConfiguration, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SessionSendsLaunchEvents() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) HTTPCookieAcceptPolicy() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetShouldUseExtendedBackgroundIdleMode() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) RequiresDNSSECValidation() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetNetworkServiceType() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetAllowsConstrainedNetworkAccess() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetHTTPAdditionalHeaders() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) HTTPMaximumConnectionsPerHost() (int64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) Init() (*any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetURLCache() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetTimeoutIntervalForResource() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetAllowsCellularAccess() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetTLSMinimumSupportedProtocol() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) HTTPAdditionalHeaders() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) ProtocolClasses() (*NSArray, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) Identifier() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetSharedContainerIdentifier() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) URLCache() (*NSURLCache, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) MultipathServiceType() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetRequiresDNSSECValidation() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SharedContainerIdentifier() (*NSString, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetConnectionProxyDictionary() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetProtocolClasses() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetTimeoutIntervalForRequest() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetSessionSendsLaunchEvents() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) ConnectionProxyDictionary() (*NSDictionary, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) TLSMaximumSupportedProtocolVersion() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetTLSMaximumSupportedProtocolVersion() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) RequestCachePolicy() (int, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetURLCredentialStorage() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) TimeoutIntervalForRequest() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) AllowsCellularAccess() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) HTTPShouldSetCookies() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetHTTPMaximumConnectionsPerHost() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) TimeoutIntervalForResource() (float64, error) {
  return 0, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) AllowsConstrainedNetworkAccess() (bool, error) {
  return false, fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetTLSMaximumSupportedProtocol() error {
  return fmt.Errorf("unimplemented")
}

func (r *NSURLSessionConfiguration) SetRequestCachePolicy() error {
  return fmt.Errorf("unimplemented")
}

