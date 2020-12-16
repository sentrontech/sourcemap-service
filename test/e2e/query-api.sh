#!/bin/bash -e

ENDPOINT="https://sourcemaps.sentron-int.com"
DIRECTORY=`dirname $0`
FAIL=""

if [[ $API_KEY == "" ]]; then
  echo "Please specify API_KEY"
  exit 1
fi

function test-api {
    EXT=$1
    VARIANT=$2
    TEST_DIR="$DIRECTORY/$EXT-$VARIANT"
    echo "$EXT ($VARIANT)"
    rm -f $TEST_DIR/actual.json
    curl -sS -XPOST \
        -H "x-api-key: $API_KEY" \
        -H "Content-Type: application/json" \
        -d @"$TEST_DIR/input.json" \
        "$ENDPOINT/$EXT" | jq > $TEST_DIR/actual.json
    git --no-pager diff -b --no-index $TEST_DIR/actual.json $TEST_DIR/expected.json
    if [[ $? == 0 ]]; then echo "PASS"; else echo "FAIL"; exit 1; fi
    echo ""
}

echo ""
test-api "generate-source-extract" "jquery"
test-api "locate-sourcemap" "guess"
test-api "locate-sourcemap" "relative"
test-api "locate-sourcemap" "error"
