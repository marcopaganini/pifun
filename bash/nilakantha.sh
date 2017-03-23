#!/bin/bash

function nilakantha() {
  precision="$1"
  iterations="$2"

  echo "$precision k 3"

  for ((i=1; i<=$iterations; i++)) {
    x=$((i * 2))
    num=4
    den=$((x * (x+1) * (x+2)))
    signal="-"
    if [[ "$((i % 2))" -eq 1 ]]; then
      signal="+"
    fi
    echo "$num $den / $signal"
  }
  echo "p"
}

nilakantha 99991 100000 | dc
