#!/bin/bash

if [ $CLUSTER_INIT == "true" ]; then
    $PULSAR_HOME/mate/pulsar_mate 2>&1
else
    nohup $PULSAR_HOME/mate/pulsar_mate >$PULSAR_HOME/pulsar_mate.stdout.log 2>$PULSAR_HOME/pulsar_mate.stderr.log
    tail -f /dev/null
fi