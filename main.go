package main

import (
	"fmt"
	"strings"

	"github.com/alexflint/go-arg"
	"github.com/steve-care-software/steve/applications/blockchains"
	"github.com/steve-care-software/steve/applications/resources"
	"github.com/steve-care-software/steve/applications/resources/lists"
)

func main() {
	var args struct {
		Action    string `arg:"positional"  help:"register, authenticate"`
		ChunkSize uint64 `arg:"--chunk_size,env:CHUNK_SIZE" default:"1024"`
		Username  string `arg:"--username,env:USERNAME"`
		Password  string `arg:"--password,env:PASSWORD"`
		BaseDir   string `arg:"env:BASE_DIR"`
		DbFile    string `arg:"env:DB_FILENAME"`
	}

	// args parsing:
	arg.MustParse(&args)

	// engine:
	targetIdentitifer := "target_identifier.tmp"
	resourceApp, err := resources.NewBuilder().Create().
		WithBasePath(args.BaseDir).
		WithReadChunkSize(args.ChunkSize).
		WithTargetIdentifier(targetIdentitifer).
		Now()

	if err != nil {
		panic(err)
	}

	err = resourceApp.Init(args.DbFile)
	if err != nil {
		panic(err)
	}

	listApp, err := lists.NewBuilder().Create().
		WithResource(resourceApp).
		Now()

	if err != nil {
		panic(err)
	}

	application, err := blockchains.NewBuilder(
		"identities",
		"blockchains",
		"identities:by_name:",
		"units:by_blockchain_and_pubkeyhash:",
		"blockchain:by_uuid:",
		"script:by_hash:",
		"block:by_hash:",
	).Create().
		WithResource(resourceApp).
		WithList(listApp).
		Now()

	if err != nil {
		panic(err)
	}

	switch args.Action {
	case "register":
		seedWords := []string{
			"abandon",
			"abandon",
			"abandon",
			"abandon",
			"abandon",
			"abandon",
			"abandon",
			"abandon",
			"abandon",
			"abandon",
			"abandon",
			"about",
		}

		err := application.Register(args.Username, []byte(args.Password), seedWords)
		if err != nil {
			fmt.Printf("register: %s", err.Error())
			return
		}

		fmt.Printf("registered username: %s\n", args.Username)
		fmt.Printf("##########\n SEED WORDS:\n %s \n##########\n", strings.Join(seedWords, ", "))
		return
	case "authenticate":
		err := application.Authenticate(args.Username, []byte(args.Password))
		if err != nil {
			fmt.Printf("authenticate: %s", err.Error())
			return
		}

		fmt.Printf("authenticated using username: %s", args.Username)
		return
	default:
		fmt.Println("the action is invalid")
	}
}
