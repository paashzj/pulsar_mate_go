package path

import (
	"os"
	"path/filepath"
)

// zookeeper
var (
	PulsarHome           = os.Getenv("PULSAR_HOME")
	PulsarConfigDir      = filepath.FromSlash(PulsarHome + "/conf")
	PulsarConfig         = filepath.FromSlash(PulsarConfigDir + "/broker.conf")
	PulsarOriginalConfig = filepath.FromSlash(PulsarConfigDir + "/broker_original.conf")
)

// mate
var (
	PulsarMatePath              = filepath.FromSlash(PulsarHome + "/mate")
	PulsarScripts               = filepath.FromSlash(PulsarMatePath + "/scripts")
	PulsarInitScript            = filepath.FromSlash(PulsarScripts + "/init-pulsar.sh")
	PulsarStartScript           = filepath.FromSlash(PulsarScripts + "/start-pulsar.sh")
	PulsarStartStandaloneScript = filepath.FromSlash(PulsarScripts + "/start-pulsar-standalone.sh")
)
