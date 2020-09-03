package common

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// Configuration defaults
var (
	// Max send and receive bytes for grpc clients and servers
	MaxRecvMsgSize = 100 * 1024 * 1024
	MaxSendMsgSize = 100 * 1024 * 1024
	// Default peer keepalive options
	DefaultKeepaliveOptions = &KeepaliveOptions{
		ClientInterval:    time.Duration(1) * time.Minute,  // 1 min
		ClientTimeout:     time.Duration(20) * time.Second, // 20 sec - gRPC default
		ServerInterval:    time.Duration(2) * time.Hour,    // 2 hours - gRPC default
		ServerTimeout:     time.Duration(20) * time.Second, // 20 sec - gRPC default
		ServerMinInterval: time.Duration(1) * time.Minute,  // match ClientInterval
	}
	// strong TLS cipher suites
	DefaultTLSCipherSuites = []uint16{
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
	}
	// default connection timeout
	DefaultConnectionTimeout = 5 * time.Second
)

// ServerConfig defines the parameters for configuring a GRPCServer instance
type ServerConfig struct {
	// ConnectionTimeout specifies the timeout for connection establishment
	// for all new connections
	ConnectionTimeout time.Duration
	// SecOpts defines the security parameters
	SecOpts *SecureOptions
	// KaOpts defines the keepalive parameters
	KaOpts *KeepaliveOptions
}

// ClientConfig defines the parameters for configuring a GRPCClient instance
type ClientConfig struct {
	// SecOpts defines the security parameters
	SecOpts *SecureOptions
	// KaOpts defines the keepalive parameters
	KaOpts *KeepaliveOptions
	// Timeout specifies how long the client will block when attempting to
	// establish a connection
	Timeout time.Duration
}

// SecureOptions defines the security parameters (e.g. TLS) for a
// GRPCServer or GRPCClient instance
type SecureOptions struct {
	// PEM-encoded X509 public key to be used for TLS communication
	Certificate []byte
	// PEM-encoded private key to be used for TLS communication
	Key []byte
	// Set of PEM-encoded X509 certificate authorities used by clients to
	// verify server certificates
	ServerRootCAs [][]byte
	// Set of PEM-encoded X509 certificate authorities used by servers to
	// verify client certificates
	ClientRootCAs [][]byte
	// Whether or not to use TLS for communication
	UseTLS bool
	// Whether or not TLS client must present certificates for authentication
	RequireClientCert bool
	// CipherSuites is a list of supported cipher suites for TLS
	CipherSuites []uint16
}

// KeepaliveOptions is used to set the gRPC keepalive settings for both
// clients and servers
type KeepaliveOptions struct {
	// ClientInterval is the duration after which if the client does not see
	// any activity from the server it pings the server to see if it is alive
	ClientInterval time.Duration
	// ClientTimeout is the duration the client waits for a response
	// from the server after sending a ping before closing the connection
	ClientTimeout time.Duration
	// ServerInterval is the duration after which if the server does not see
	// any activity from the client it pings the client to see if it is alive
	ServerInterval time.Duration
	// ServerTimeout is the duration the server waits for a response
	// from the client after sending a ping before closing the connection
	ServerTimeout time.Duration
	// ServerMinInterval is the minimum permitted time between client pings.
	// If clients send pings more frequently, the server will disconnect them
	ServerMinInterval time.Duration
}

// ServerKeepaliveOptions returns gRPC keepalive options for server.  If
// opts is nil, the default keepalive options are returned
func ServerKeepaliveOptions(ka *KeepaliveOptions) []grpc.ServerOption {
	// use default keepalive options if nil
	if ka == nil {
		ka = DefaultKeepaliveOptions
	}
	var serverOpts []grpc.ServerOption
	kap := keepalive.ServerParameters{
		Time:    ka.ServerInterval,
		Timeout: ka.ServerTimeout,
	}
	serverOpts = append(serverOpts, grpc.KeepaliveParams(kap))
	kep := keepalive.EnforcementPolicy{
		MinTime: ka.ServerMinInterval,
		// allow keepalive w/o rpc
		PermitWithoutStream: true,
	}
	serverOpts = append(serverOpts, grpc.KeepaliveEnforcementPolicy(kep))
	return serverOpts
}

// ClientKeepaliveOptions returns gRPC keepalive options for clients.  If
// opts is nil, the default keepalive options are returned
func ClientKeepaliveOptions(ka *KeepaliveOptions) []grpc.DialOption {
	// use default keepalive options if nil
	if ka == nil {
		ka = DefaultKeepaliveOptions
	}

	var dialOpts []grpc.DialOption
	kap := keepalive.ClientParameters{
		Time:                ka.ClientInterval,
		Timeout:             ka.ClientTimeout,
		PermitWithoutStream: true,
	}
	dialOpts = append(dialOpts, grpc.WithKeepaliveParams(kap))
	return dialOpts
}

func GetPath(key string) string {
	k := viper.GetString(key)
	if k == "" {
		return ""
	}

	return TransPath(filepath.Dir(viper.ConfigFileUsed()), k)
}

func TransPath(base, p string) string {
	if filepath.IsAbs(p) {
		return p
	}

	return filepath.Join(base, p)
}

type GRPC_Type int32

const (
	TLS_ENABLED              GRPC_Type = 0
	TLS_SERVER_KEY           GRPC_Type = 1
	TLS_SERVER_CERT          GRPC_Type = 2
	TLS_CLIENT_AUTH_REQUIRED GRPC_Type = 3
	TLS_ROOT_CERT            GRPC_Type = 4
	TLS_KEEPALIVE_INTERVAL   GRPC_Type = 5
	TLS_CONNECT_TIMEOUT      GRPC_Type = 6
)

func GetServerConfig(viperKey map[GRPC_Type]string) (ServerConfig, error) {
	secureOptions := &SecureOptions{
		UseTLS: viper.GetBool(viperKey[TLS_ENABLED]),
	}
	serverConfig := ServerConfig{SecOpts: secureOptions}
	if secureOptions.UseTLS {
		// get the certs from the file system
		serverKey, err := ioutil.ReadFile(
			GetPath(viperKey[TLS_SERVER_KEY]))
		if err != nil {

			return serverConfig, fmt.Errorf("error loading gRpcTLS serverKey (%s)", err.Error())
		}
		serverCert, err := ioutil.ReadFile(
			GetPath(viperKey[TLS_SERVER_CERT]))
		if err != nil {
			return serverConfig, fmt.Errorf("error loading gRpcTLS serverCert (%s)", err)
		}
		secureOptions.Certificate = serverCert
		secureOptions.Key = serverKey
		secureOptions.RequireClientCert = viper.GetBool(viperKey[TLS_CLIENT_AUTH_REQUIRED])
		if secureOptions.RequireClientCert {
			var clientRoots [][]byte
			for _, file := range viper.GetStringSlice(viperKey[TLS_ROOT_CERT]) {
				clientRoot, err := ioutil.ReadFile(
					TransPath(filepath.Dir(viper.ConfigFileUsed()), file))
				if err != nil {
					return serverConfig, fmt.Errorf("error loading client root CAs (%s)", err)
				}
				clientRoots = append(clientRoots, clientRoot)
			}
			secureOptions.ClientRootCAs = clientRoots
		}
		// check for root cert
		if viper.GetString(viperKey[TLS_ROOT_CERT]) != "" {
			rootCert, err := ioutil.ReadFile(
				GetPath(viperKey[TLS_ROOT_CERT]))
			if err != nil {
				return serverConfig, fmt.Errorf("error loading gRpcTLS rootCA certificate (%s)", err)
			}
			secureOptions.ServerRootCAs = [][]byte{rootCert}
		}
	}
	// get the default keepalive options
	serverConfig.KaOpts = DefaultKeepaliveOptions
	// check to see if minInterval is set for the env
	if viper.IsSet(viperKey[TLS_KEEPALIVE_INTERVAL]) {
		serverConfig.KaOpts.ServerMinInterval = viper.GetDuration(viperKey[TLS_KEEPALIVE_INTERVAL]) * time.Second
	}
	if viper.IsSet(viperKey[TLS_CONNECT_TIMEOUT]) {
		serverConfig.ConnectionTimeout = viper.GetDuration(viperKey[TLS_CONNECT_TIMEOUT]) * time.Second
	}
	return serverConfig, nil
}
