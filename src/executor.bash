#!/bin/bash -e

mask_cmd() {
    (set -o pipefail && $@ | sed 's/:[^\/][^@]*@/:\*\*\*\*\*@/g')
}

/opt/helm > /tmp/run

echo ""
echo "Running the following script:"
echo "----------------------------"
mask_cmd cat /tmp/run
echo "----------------------------"
echo ""

if [ "$(basename $0)" = "executor" ]; then
    chmod +x /tmp/run
    mask_cmd /tmp/run
else
    . /tmp/run
fi