#!/bin/bash

exec "$CLANG" -target "$TARGET" -isysroot "$SDK_PATH" "$@"