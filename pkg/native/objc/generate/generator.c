#include <stdlib.h>
#include <clang-c/Index.h>

int visitor(CXCursor, CXCursor, void *);

extern void visitChildren(const CXCursor cursor, const void *context)
{
  clang_visitChildren(cursor, (void *)visitor, (void *)context);
}

extern void ParseHeader(const char *target, const char *isysroot, const char *header, CXIndex *pIndex, CXTranslationUnit *pTranslationUnit, CXCursor *pCursor)
{
  const char *ARGUMENTS[] = {"-target", target, "-isysroot", isysroot, "-ObjC"};
  *pIndex = clang_createIndex(1, 0);
  *pTranslationUnit = clang_parseTranslationUnit(
      *pIndex, header, ARGUMENTS, 5,
      NULL, 0, CXTranslationUnit_SkipFunctionBodies);

  *pCursor = clang_getTranslationUnitCursor(*pTranslationUnit);
}