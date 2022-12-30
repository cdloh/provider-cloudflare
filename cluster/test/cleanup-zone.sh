#!/usr/bin/env bash
echo "Cleanup resources that need a zone priror to deleting the zone"

# Cleanup the resources that require a Cloudflare zone priror to deleting the zone
${KUBECTL} delete managed --selector='testing.cloudflare.upbound.io/require-zone=true' --wait
