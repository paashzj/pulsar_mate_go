#!/bin/bash

mkdir $PULSAR_HOME/logs
if [ $CLUSTER_INIT == "true" ]; then
    $PULSAR_HOME/mate/pulsar_mate 2>&1
else
    DIR="$( cd "$( dirname "$0"  )" && pwd  )"
    bash -x $DIR/start-daemon.sh
    tail -f /dev/null
fi

