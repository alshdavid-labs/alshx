#!/bin/bash

export YEAR="$(date -u +%y)"
export MONTH="$(date -u +%m)"
export DAY="$(date -u +%d)"
export HOUR="$(date -u +%H)"
export MINUTE="$(date -u +%M)"
export VERSION="${YEAR}-${MONTH}-${DAY}-${HOUR}-${MINUTE}"
