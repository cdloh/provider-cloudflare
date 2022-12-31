#!/usr/bin/env bash
echo "Cleanup resources that need a $1 priror to deleting the $1"

# Cleanup the resources that require a Cloudflare parent resource priror to deleting the parent
${KUBECTL} delete managed --selector="testing.cloudflare.upbound.io/require-$1=true" --wait
