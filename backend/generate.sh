#!/bin/bash
set -e

echo "==> generating Ent code..."
(cd internal/ent && go generate)

echo "==> generating GraphQL code..."
(cd internal/graphql && gqlgen generate)
