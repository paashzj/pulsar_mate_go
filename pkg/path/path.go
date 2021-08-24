package path

import (
	"os"
	"path/filepath"
)

// pulsar
var (
	PulsarHome                   = os.Getenv("PULSAR_HOME")
	PulsarConfigDir              = filepath.FromSlash(PulsarHome + "/conf")
	PulsarConfig                 = filepath.FromSlash(PulsarConfigDir + "/broker.conf")
	PulsarOriginalConfig         = filepath.FromSlash(PulsarConfigDir + "/broker_original.conf")
	PulsarFunctionOriginalConfig = filepath.FromSlash(PulsarConfigDir + "/function_original.yaml")
)

// mate
var (
	PulsarMatePath              = filepath.FromSlash(PulsarHome + "/mate")
	PulsarScripts               = filepath.FromSlash(PulsarMatePath + "/scripts")
	PulsarInitScript            = filepath.FromSlash(PulsarScripts + "/init-pulsar.sh")
	PulsarStartScript           = filepath.FromSlash(PulsarScripts + "/start-pulsar.sh")
	PulsarStartStandaloneScript = filepath.FromSlash(PulsarScripts + "/start-pulsar-standalone.sh")
	PulsarStartFunctionScript   = filepath.FromSlash(PulsarScripts + "/start-pulsar-function.sh")
)
