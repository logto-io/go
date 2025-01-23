#!/bin/bash

PKGNAME=logto
LOGOTO_ENDPOINT=https://default.logto.app
rm -rf ${PKGNAME}
openapi-generator generate -i ${LOGOTO_ENDPOINT}/api/swagger.json \
     -g go -o ./${PKGNAME} \
     --git-user-id logto-io --git-repo-id go/management-api/${PKGNAME} \
     --additional-properties=packageName=${PKGNAME},withGoMod=false
python3 fix_vars.py
go mod tidy
