#!/bin/bash

TOTALTESTS=$(cwl/run_test.sh RUNNER=yacle -l | tail -n 2 | head -n 1 | awk '{print $1}' | tr -dc '0123456789')

PASSED=$(cwl/run_test.sh RUNNER=yacle 2>&1 | grep "[0-9]* tests passed"  | awk '{ print $1 }')

if [ "All" = "${PASSED}" ]; then
  PASSED=$TOTALTESTS
fi

cat << EOS > .conformance.json
{
  "pass": $PASSED
}
EOS
