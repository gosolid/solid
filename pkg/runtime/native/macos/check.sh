#!/bin/bash

ls -1R pkg/runtime/native/macos/ | grep .go | awk 'seen[$0]++' | while read FILE; do
  echo Duplicate $FILE:
  find pkg/runtime/native/macos -type f -name $FILE
  echo
done

