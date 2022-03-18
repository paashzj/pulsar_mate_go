#!/bin/bash

mkdir -p $PULSAR_HOME/logs
nohup $PULSAR_HOME/mate/pulsar_mate >>$PULSAR_HOME/logs/pulsar_mate.stdout.log 2>>$PULSAR_HOME/logs/pulsar_mate.stderr.log &

