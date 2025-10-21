#!/bin/bash
set -e

(cd internal/ent && go generate)
echo "==> ent code generated"

(cd internal/graphql && gqlgen generate)
echo "==> gqlgen code generated"
