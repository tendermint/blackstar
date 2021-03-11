package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client/flags"
	tmcli "github.com/tendermint/tendermint/libs/cli"
)

// Default constants
const (
	chainID        = ""
	keyringBackend = "os"
	output         = "text"
	node           = "tcp://localhost:26657"
	broadcastMode  = "sync"
	trace          = false
)

var ErrWrongNumberOfArgs = fmt.Errorf("wrong number of arguments")

type ClientConfig struct {
	ChainID        string `mapstructure:"chain-id" json:"chain-id"`
	KeyringBackend string `mapstructure:"keyring-backend" json:"keyring-backend"`
	Output         string `mapstructure:"output" json:"output"`
	Node           string `mapstructure:"node" json:"node"`
	BroadcastMode  string `mapstructure:"broadcast-mode" json:"broadcast-mode"`
	Trace          bool   `mapstructure:"trace" json:"trace"`
}

// TODO Validate values in setters
func (c *ClientConfig) SetChainID(chainID string) {
	c.ChainID = chainID
}

func (c *ClientConfig) SetKeyringBackend(keyringBackend string) {
	c.KeyringBackend = keyringBackend
}

func (c *ClientConfig) SetOutput(output string) {
	c.Output = output
}

func (c *ClientConfig) SetNode(node string) {
	c.Node = node
}

func (c *ClientConfig) SetBroadcastMode(broadcastMode string) {
	c.BroadcastMode = broadcastMode
}

func (c *ClientConfig) SetTrace(trace string) error {
	boolVal, err := strconv.ParseBool(trace)
	if err != nil {
		return err
	}
	c.Trace = boolVal
	return nil
}

func DefaultClientConfig() *ClientConfig {
	return &ClientConfig{chainID, keyringBackend, output, node, broadcastMode, trace}
}

// Cmd returns a CLI command to interactively create an application CLI
// config file.
func Cmd(defaultCLIHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config <key> [value]",
		Short: "Create or query an application CLI configuration file",
		RunE:  runConfigCmd,
		Args:  cobra.RangeArgs(0, 2),
	}

	cmd.Flags().String(flags.FlagHome, defaultCLIHome,
		"set client's home directory for configuration")
	return cmd
}

func runConfigCmd(cmd *cobra.Command, args []string) error {

	v := viper.New()

	cfgPath, err := ensureCfgPath(v.GetString(flags.FlagHome))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to make config path: %v\n", err)
		return err
	}

	cliConfig, err := getClientConfig(cfgPath, v)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to get client config: %v\n", err)
		return err
	}

	switch len(args) {
	case 0:
		// print all client config fields to sdt out
		s, _ := json.MarshalIndent(cliConfig, "", "\t")
		fmt.Print(string(s))

	case 1:
		// it's a get
		// TODO implement method for get
		// should i implement getters here?
		key := args[0]
		switch key {
		case flags.FlagChainID:
			fmt.Println(cliConfig.ChainID)
		case flags.FlagKeyringBackend:
			fmt.Println(cliConfig.KeyringBackend)
		case tmcli.OutputFlag:
			fmt.Println(cliConfig.Output)
		case flags.FlagNode:
			fmt.Println(cliConfig.Node)
		case flags.FlagBroadcastMode:
			fmt.Println(cliConfig.BroadcastMode)
		case "trace":
			fmt.Println(cliConfig.Trace)
		default:
			err := errUnknownConfigKey(key)
			fmt.Fprintf(os.Stderr, "Unable to get the value for the key: %v, error:  %v\n", key, err)
			return err
		}

	case 2:
		// it's set

		key, value := args[0], args[1]

		switch key {
		case flags.FlagChainID:
			cliConfig.SetChainID(value)
		case flags.FlagKeyringBackend:
			cliConfig.SetKeyringBackend(value)
		case tmcli.OutputFlag:
			cliConfig.SetOutput(value)
		case flags.FlagNode:
			cliConfig.SetNode(value)
		case flags.FlagBroadcastMode:
			cliConfig.SetBroadcastMode(value)
		case "trace":
			if err := cliConfig.SetTrace(value); err != nil {
				fmt.Fprintf(os.Stderr, "Unable to parse value to bool, err: %v\n", err)
				return err
			}
		default:
			return errUnknownConfigKey(key)
		}

		configTemplate, err := InitConfigTemplate()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to initiate config template, err: %v\n", err)
			return err
		}

		cfgFile := path.Join(cfgPath, "config.toml")
		if err := WriteConfigFile(cfgFile, cliConfig, configTemplate); err != nil {
			fmt.Fprintf(os.Stderr, "Unable to write client config to the file, err: %v\n", err)
			return err
		}

	default:
		// print error
		err := errors.New("unable to execute config command")
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}

	return nil
}

func errUnknownConfigKey(key string) error {
	return fmt.Errorf("unknown configuration key: %q", key)
}
