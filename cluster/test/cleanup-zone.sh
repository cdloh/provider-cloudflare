#!/usr/bin/env bash
echo "Cleanup zone"

# Cleanup the zones once all the others are gone
kubectl wait managed --selector='!cloudflare-zone' --for=delete
kubectl delete zone.zone.cloudflare.upbound.io/example --wait --ignore-not-found
