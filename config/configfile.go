// Copyright 2022 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package config

import (
	"fmt"
	"os"

	"github.com/sourcenetwork/defradb/errors"
)

const (
	DefaultDefraDBConfigFileName = "config.yaml"
	configType                   = "yaml"
	defaultDirPerm               = 0o700
	defaultConfigFilePerm        = 0o644
)

func (cfg *Config) writeConfigFile(path string) error {
	buffer, err := cfg.toBytes()
	if err != nil {
		return err
	}
	if err := os.WriteFile(path, buffer, defaultConfigFilePerm); err != nil {
		return errors.Wrap("failed to write file", err)
	}
	return nil
}

// WriteConfigFile writes a config file in a given root directory.
func (cfg *Config) WriteConfigFileToRootDir(rootDir string) error {
	path := fmt.Sprintf("%v/%v", rootDir, DefaultDefraDBConfigFileName)
	return cfg.writeConfigFile(path)
}

// defaultConfigTemplate must reflect Config in content and configuration.
// All parameters must be represented here, to support Viper's automatic environment variable handling.
const defaultConfigTemplate = `# DefraDB configuration (YAML)

# NOTE: Paths below are relative to the DefraDB directory.
# By default, the DefraDB directory is "$HOME/.defradb", but
# can be changed via the $DEFRA_ROOT env variable or --rootdir CLI flag.

datastore:
  # Store can be badger | memory
    #   badger: fast pure Go key-value store optimized for SSDs (https://github.com/dgraph-io/badger)
    #   memory: in-memory version of badger
  store: {{ .Datastore.Store }}
  badger:
    path: {{ .Datastore.Badger.Path }}
    # Maximum file size of the value log files. The in-memory file size will be 2*valuelogfilesize.
    # Human friendly units can be used (ex: 500MB).
    valuelogfilesize: {{ .Datastore.Badger.ValueLogFileSize }}
  # memory:
  #    size: {{ .Datastore.Memory.Size }}

api:
  # Address of the HTTP API to listen on or connect to
  address: {{ .API.Address }}
  # Whether the API server should listen over HTTPS
  tls: {{ .API.TLS }}
  # The path to the public key file. Ignored if domains is set.
  pubkeypath: {{ .API.PubKeyPath }}
  # The path to the private key file. Ignored if domains is set.
  privkeypath: {{ .API.PrivKeyPath }}
  # Email address to let the CA (Let's Encrypt) send notifications via email when there are issues (optional).
  # email: {{ .API.Email }}

net:
  # Whether the P2P is disabled
  p2pdisabled: {{ .Net.P2PDisabled }}
  # Listening address of the P2P network
  p2paddress: {{ .Net.P2PAddress }}
  # Listening address of the RPC endpoint
  rpcaddress: {{ .Net.RPCAddress }}
  # gRPC server address
  tcpaddress: {{ .Net.TCPAddress }}
  # Time duration after which a RPC connection to a peer times out
  rpctimeout: {{ .Net.RPCTimeout }}
  # Whether the node has pubsub enabled or not
  pubsub: {{ .Net.PubSubEnabled }}
  # Enable libp2p's Circuit relay transport protocol https://docs.libp2p.io/concepts/circuit-relay/
  relay: {{ .Net.RelayEnabled }}
  # List of peers to boostrap with, specified as multiaddresses (https://docs.libp2p.io/concepts/addressing/)
  peers: {{ .Net.Peers }}
  # Amount of time after which an idle RPC connection would be closed
  RPCMaxConnectionIdle: {{ .Net.RPCMaxConnectionIdle }}

log:
  # Log level. Options are debug, info, error, fatal
  level: {{ .Log.Level }}
  # Include stacktrace in error and fatal logs
  stacktrace: {{ .Log.Stacktrace }}
  # Supported log formats are json, csv
  format: {{ .Log.Format }}
  # Where the log output is written to
  output: {{ .Log.Output }}
  # Disable colored log output
  nocolor: {{ .Log.NoColor }}
  # Caller location in log output
  caller: {{ .Log.Caller }}
`