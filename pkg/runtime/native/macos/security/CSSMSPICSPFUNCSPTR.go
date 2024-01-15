//js:package native/macos/security
package security

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Security.CSSM_SPI_CSP_FUNCS_PTR
*/

type CSSMSPICSPFUNCSPTR struct {
  EventNotify func(...any) (any, error) `v8:"eventNotify"`
  QuerySize func(...any) (any, error) `v8:"querySize"`
  SignData func(...any) (any, error) `v8:"signData"`
  SignDataInit func(...any) (any, error) `v8:"signDataInit"`
  SignDataUpdate func(...any) (any, error) `v8:"signDataUpdate"`
  SignDataFinal func(...any) (any, error) `v8:"signDataFinal"`
  VerifyData func(...any) (any, error) `v8:"verifyData"`
  VerifyDataInit func(...any) (any, error) `v8:"verifyDataInit"`
  VerifyDataUpdate func(...any) (any, error) `v8:"verifyDataUpdate"`
  VerifyDataFinal func(...any) (any, error) `v8:"verifyDataFinal"`
  DigestData func(...any) (any, error) `v8:"digestData"`
  DigestDataInit func(...any) (any, error) `v8:"digestDataInit"`
  DigestDataUpdate func(...any) (any, error) `v8:"digestDataUpdate"`
  DigestDataClone func(...any) (any, error) `v8:"digestDataClone"`
  DigestDataFinal func(...any) (any, error) `v8:"digestDataFinal"`
  GenerateMac func(...any) (any, error) `v8:"generateMac"`
  GenerateMacInit func(...any) (any, error) `v8:"generateMacInit"`
  GenerateMacUpdate func(...any) (any, error) `v8:"generateMacUpdate"`
  GenerateMacFinal func(...any) (any, error) `v8:"generateMacFinal"`
  VerifyMac func(...any) (any, error) `v8:"verifyMac"`
  VerifyMacInit func(...any) (any, error) `v8:"verifyMacInit"`
  VerifyMacUpdate func(...any) (any, error) `v8:"verifyMacUpdate"`
  VerifyMacFinal func(...any) (any, error) `v8:"verifyMacFinal"`
  EncryptData func(...any) (any, error) `v8:"encryptData"`
  EncryptDataInit func(...any) (any, error) `v8:"encryptDataInit"`
  EncryptDataUpdate func(...any) (any, error) `v8:"encryptDataUpdate"`
  EncryptDataFinal func(...any) (any, error) `v8:"encryptDataFinal"`
  DecryptData func(...any) (any, error) `v8:"decryptData"`
  DecryptDataInit func(...any) (any, error) `v8:"decryptDataInit"`
  DecryptDataUpdate func(...any) (any, error) `v8:"decryptDataUpdate"`
  DecryptDataFinal func(...any) (any, error) `v8:"decryptDataFinal"`
  QueryKeySizeInBits func(...any) (any, error) `v8:"queryKeySizeInBits"`
  GenerateKey func(...any) (any, error) `v8:"generateKey"`
  GenerateKeyPair func(...any) (any, error) `v8:"generateKeyPair"`
  GenerateRandom func(...any) (any, error) `v8:"generateRandom"`
  GenerateAlgorithmParams func(...any) (any, error) `v8:"generateAlgorithmParams"`
  WrapKey func(...any) (any, error) `v8:"wrapKey"`
  UnwrapKey func(...any) (any, error) `v8:"unwrapKey"`
  DeriveKey func(...any) (any, error) `v8:"deriveKey"`
  FreeKey func(...any) (any, error) `v8:"freeKey"`
  PassThrough func(...any) (any, error) `v8:"passThrough"`
  Login func(...any) (any, error) `v8:"login"`
  Logout func(...any) (any, error) `v8:"logout"`
  ChangeLoginAcl func(...any) (any, error) `v8:"changeLoginAcl"`
  ObtainPrivateKeyFromPublicKey func(...any) (any, error) `v8:"obtainPrivateKeyFromPublicKey"`
  RetrieveUniqueId func(...any) (any, error) `v8:"retrieveUniqueId"`
  RetrieveCounter func(...any) (any, error) `v8:"retrieveCounter"`
  VerifyDevice func(...any) (any, error) `v8:"verifyDevice"`
  GetTimeValue func(...any) (any, error) `v8:"getTimeValue"`
  GetOperationalStatistics func(...any) (any, error) `v8:"getOperationalStatistics"`
  GetLoginAcl func(...any) (any, error) `v8:"getLoginAcl"`
  GetKeyAcl func(...any) (any, error) `v8:"getKeyAcl"`
  ChangeKeyAcl func(...any) (any, error) `v8:"changeKeyAcl"`
  GetKeyOwner func(...any) (any, error) `v8:"getKeyOwner"`
  ChangeKeyOwner func(...any) (any, error) `v8:"changeKeyOwner"`
  GetLoginOwner func(...any) (any, error) `v8:"getLoginOwner"`
  ChangeLoginOwner func(...any) (any, error) `v8:"changeLoginOwner"`
}
