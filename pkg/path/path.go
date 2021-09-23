package path

import (
	"os"
	"path/filepath"
)

// pulsar
var (
	PulsarHome                     = os.Getenv("PULSAR_HOME")
	PulsarConfigDir                = filepath.FromSlash(PulsarHome + "/conf")
	PulsarConfig                   = filepath.FromSlash(PulsarConfigDir + "/broker.conf")
	PulsarOriginalConfig           = filepath.FromSlash(PulsarConfigDir + "/broker_original.conf")
	PulsarStandaloneConfig         = filepath.FromSlash(PulsarConfigDir + "/standalone.conf")
	PulsarStandaloneOriginalConfig = filepath.FromSlash(PulsarConfigDir + "/standalone_original.conf")
	PulsarFunctionOriginalConfig   = filepath.FromSlash(PulsarConfigDir + "/function_original.yaml")
	PulsarClientConfig             = filepath.FromSlash(PulsarConfigDir + "/client.conf")
	PulsarClientOriginalConfig     = filepath.FromSlash(PulsarConfigDir + "/client_original.conf")
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
