package cli

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/internal/cli/flagset"
	"github.com/ethereum/go-ethereum/internal/cli/server"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/ethdb"
)

// DBCompactCommand is for compacting database storage
type DBCompactCommand struct {
	*Meta

	datadir 	string
	cache 		uint64
}

// MarkDown implements cli.MarkDown interface
func (c *DBCompactCommand) MarkDown() string {
	items := []string{
		"# db",
		"The ```bor db compact``` command will compact database size",
		c.Flags().MarkDown(),
	}

	return strings.Join(items, "\n\n")
}

// Help implements the cli.Command interface
func (c *DBCompactCommand) Help() string {
	return `Usage: bor db compact <datadir> <cache>

  This command compacting the entire database size within given datadir, with optional cache (default 1024 -MB-)`
}

// Synopsis implements the cli.Command interface
func (c *DBCompactCommand) Synopsis() string {
	return "Compacting storage size of databases"
}

func (c *DBCompactCommand) Flags() *flagset.Flagset {
	flags := c.NewFlagSet("dbcompact")

	flags.Uint64Flag(&flagset.Uint64Flag{
		Name:    "cache",
		Usage:   "Megabytes of memory allocated to internal caching",
		Value:   &c.cache,
		Default: 1024.0,
		Group:   "Cache",
	})

	return flags
}

func (c *DBCompactCommand) Run(args []string) int {
	flags := c.Flags()

	// parse datadir
	if err := flags.Parse(args); err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	datadir := c.dataDir
	if datadir == "" {
		c.UI.Error("--datadir is required")
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

	log.Info("Opening database")
	chaindb, err := node.OpenDatabase(chaindataPath, int(c.cache), dbHandles, "", false)

	if err != nil {
		log.Info("err --- ", "err", err)
		c.UI.Error(err.Error())
		return 1
	}

	log.Info("Stats before compaction")
	showLeveldbStats(chaindb)
	
	log.Info("Triggering compaction")
	if err := chaindb.Compact(nil, nil); err != nil {
		log.Info("Compact err", "error", err)
		return 1
	}
	log.Info("Stats after compaction")
	showLeveldbStats(chaindb)
	
	return 0
}

func showLeveldbStats(db ethdb.Stater) {
	if stats, err := db.Stat("leveldb.stats"); err != nil {
		log.Warn("Failed to read database stats", "error", err)
	} else {
		fmt.Println(stats)
	}
	if ioStats, err := db.Stat("leveldb.iostats"); err != nil {
		log.Warn("Failed to read database iostats", "error", err)
	} else {
		fmt.Println(ioStats)
	}
}