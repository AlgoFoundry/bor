package cli

import (
	"strings"

	"github.com/ethereum/go-ethereum/internal/cli/flagset"
	"github.com/ethereum/go-ethereum/internal/cli/server"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/core/rawdb"
)

// DBInspectCommand is for inspecting database storage
type DBInspectCommand struct {
	*Meta

	datadir 	string
	cache 		uint64
}

// MarkDown implements cli.MarkDown interface
func (c *DBInspectCommand) MarkDown() string {
	items := []string{
		"# db",
		"The ```bor db inspect``` command will inspect the storage size for each type of data in the database",
		c.Flags().MarkDown(),
	}

	return strings.Join(items, "\n\n")
}

// Help implements the cli.Command interface
func (c *DBInspectCommand) Help() string {
	return `Usage: bor db inspect <datadir> <cache>

  This command iterates the entire database within given datadir, with optional cache (default 1024 -MB-)`
}

// Synopsis implements the cli.Command interface
func (c *DBInspectCommand) Synopsis() string {
	return "Inspect storage size of databases"
}

func (c *DBInspectCommand) Flags() *flagset.Flagset {
	flags := c.NewFlagSet("dbinspect")

	flags.Uint64Flag(&flagset.Uint64Flag{
		Name:    "cache",
		Usage:   "Megabytes of memory allocated to internal caching",
		Value:   &c.cache,
		Default: 1024.0,
		Group:   "Cache",
	})

	return flags
}

func (c *DBInspectCommand) Run(args []string) int {
	var (
		prefix []byte
		start  []byte
	)

	flags := c.Flags()

	// parse datadir
	if err := flags.Parse(args); err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	datadir := c.dataDir
	if datadir == "" {
		c.UI.Error("datadir is required")
		return 1
	}

	// Create the node
	node, err := node.New(&node.Config{
		DataDir: datadir,
	})

	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	dbHandles, err := server.MakeDatabaseHandles()
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	chaindb, err := node.OpenDatabase(chaindataPath, int(c.cache), dbHandles, "", true)

	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	rawdb.InspectDatabase(chaindb, prefix, start, false)
	return 0
}