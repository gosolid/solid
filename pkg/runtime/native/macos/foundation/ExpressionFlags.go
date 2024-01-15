//js:package native/macos/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Foundation._expressionFlags
*/

type ExpressionFlags struct {
  EvaluationBlocked uint `v8:"evaluationBlocked"`
  UsesKVC uint `v8:"usesKVC"`
  ValidatedExpression uint `v8:"validatedExpression"`
  ValidatedKeys uint `v8:"validatedKeys"`
  ReservedExpressionFlags uint `v8:"reservedExpressionFlags"`
}
