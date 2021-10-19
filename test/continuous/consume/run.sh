#!/bin/bash

DIR="$( cd "$( dirname "$0"  )" && pwd  )"
export PULSAR_ALLOW_AUTO_TOPIC_CREATION=false
$PULSAR_HOME/mate/config_gen
$PULSAR_HOME/bin/pulsar standalone
sleep 30
go test $DIR/test_test.go