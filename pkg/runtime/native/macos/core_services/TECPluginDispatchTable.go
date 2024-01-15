//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.TECPluginDispatchTable
*/

type TECPluginDispatchTable struct {
  Version uint `v8:"version"`
  CompatibleVersion uint `v8:"compatibleVersion"`
  PluginID uint `v8:"pluginID"`
  PluginNewEncodingConverter *func(...any) (any, error) `v8:"pluginNewEncodingConverter"`
  PluginClearContextInfo *func(...any) (any, error) `v8:"pluginClearContextInfo"`
  PluginConvertTextEncoding *func(...any) (any, error) `v8:"pluginConvertTextEncoding"`
  PluginFlushConversion *func(...any) (any, error) `v8:"pluginFlushConversion"`
  PluginDisposeEncodingConverter *func(...any) (any, error) `v8:"pluginDisposeEncodingConverter"`
  PluginNewEncodingSniffer *func(...any) (any, error) `v8:"pluginNewEncodingSniffer"`
  PluginClearSnifferContextInfo *func(...any) (any, error) `v8:"pluginClearSnifferContextInfo"`
  PluginSniffTextEncoding *func(...any) (any, error) `v8:"pluginSniffTextEncoding"`
  PluginDisposeEncodingSniffer *func(...any) (any, error) `v8:"pluginDisposeEncodingSniffer"`
  PluginGetCountAvailableTextEncodings *func(...any) (any, error) `v8:"pluginGetCountAvailableTextEncodings"`
  PluginGetCountAvailableTextEncodingPairs *func(...any) (any, error) `v8:"pluginGetCountAvailableTextEncodingPairs"`
  PluginGetCountDestinationTextEncodings *func(...any) (any, error) `v8:"pluginGetCountDestinationTextEncodings"`
  PluginGetCountSubTextEncodings *func(...any) (any, error) `v8:"pluginGetCountSubTextEncodings"`
  PluginGetCountAvailableSniffers *func(...any) (any, error) `v8:"pluginGetCountAvailableSniffers"`
  PluginGetCountWebTextEncodings *func(...any) (any, error) `v8:"pluginGetCountWebTextEncodings"`
  PluginGetCountMailTextEncodings *func(...any) (any, error) `v8:"pluginGetCountMailTextEncodings"`
  PluginGetTextEncodingInternetName *func(...any) (any, error) `v8:"pluginGetTextEncodingInternetName"`
  PluginGetTextEncodingFromInternetName *func(...any) (any, error) `v8:"pluginGetTextEncodingFromInternetName"`
}
