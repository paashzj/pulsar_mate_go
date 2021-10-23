#!/bin/bash

$PULSAR_HOME/bin/pulsar-daemon start standalone -nfw >>$PULSAR_HOME/logs/pulsar.stdout.log 2>>$PULSAR_HOME/logs/pulsar.stderr.log