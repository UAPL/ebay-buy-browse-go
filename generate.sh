#!/usr/bin/env sh
openapi-generator generate --additional-properties=packageName=buybrowse -g go -i https://developer.ebay.com/api-docs/master/buy/browse/openapi/3/buy_browse_v1_oas3.yaml