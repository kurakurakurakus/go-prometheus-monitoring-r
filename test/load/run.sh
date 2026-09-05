#!/bin/sh

set -e

TARGETS="/load/targets.txt"

TIMESTAMP=$(date +"%Y-%m-%d_%H%M%S")

RESULT="test/load/results/${TIMESTAMP}.bin"
JSON_RESULT="test/load/results/${TIMESTAMP}.json"

CONFIG_FILE="test/load/config.yml"

RATE=$(grep '^rate:' "$CONFIG_FILE" | sed 's/^rate:[[:space:]]*//')
DURATION=$(grep '^duration:' "$CONFIG_FILE" | sed 's/^duration:[[:space:]]*//')

mkdir -p test/load/results

echo "Load test"
echo "  Rate:     $RATE"
echo "  Duration: $DURATION"
echo "  Targets:  test/load/targets.txt"
echo ""

echo "Starting load test..."

MSYS_NO_PATHCONV=1 docker run --rm \
  -v "$(pwd)/test/load:/load:ro" \
  peterevans/vegeta:latest \
  vegeta attack \
  -targets="$TARGETS" \
  -rate="$RATE" \
  -duration="$DURATION" \
  > "$RESULT"

echo ""
echo "Results:"
echo ""

MSYS_NO_PATHCONV=1 docker run --rm \
  -v "$(pwd)/test/load:/load:ro" \
  peterevans/vegeta:latest \
  vegeta report \
  /load/results/latest.bin

echo ""
echo "Generating JSON report..."

MSYS_NO_PATHCONV=1 docker run --rm \
  -v "$(pwd)/test/load:/load:ro" \
  peterevans/vegeta:latest \
  vegeta report \
  -type=json \
  /load/results/latest.bin \
  > "$JSON_RESULT"

echo ""
echo "Saved:"
echo "  Raw:  $RESULT"
echo "  JSON: $JSON_RESULT"