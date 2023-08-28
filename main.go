package main

import (
	"bytes"
	"flag"
	"fmt"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"runtime"

	"github.com/gosolid/solid/pkg/client"
	"github.com/gosolid/solid/pkg/daemon"
)

var _daemon bool
var _pprof bool

func newFlagSet(errorHandling flag.ErrorHandling) *flag.FlagSet {
	flagSet := flag.NewFlagSet("main", errorHandling)
	flagSet.BoolVar(&_pprof, "pprof", false, "enables the pprof profiling server")
	flagSet.BoolVar(&_daemon, "daemon", false, "runs the solid execution daemon")

	return flagSet
}

func NewFlagSet() *flag.FlagSet {
	flagSet := newFlagSet(flag.ContinueOnError)
	buf := bytes.NewBuffer([]byte{})
	flagSet.SetOutput(buf)
	return flagSet
}

func main() {
	if _pprof {
		go func() {
			runtime.SetBlockProfileRate(100)

			if addr, err := net.ResolveTCPAddr("tcp", "localhost:0"); err != nil {
				panic(err)
			} else if listener, err := net.ListenTCP("tcp", addr); err != nil {
				panic(err)
			} else {
				mux := http.NewServeMux()

				mux.HandleFunc("/debug/pprof/", pprof.Index)
				mux.HandleFunc("/debug/pprof/{action}", pprof.Index)
				mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

				fmt.Printf("🐞 go tool pprof -http : http://%s/debug/pprof/block\n", listener.Addr())

				if err := http.Serve(listener, mux); err != nil {
					panic(err)
				}
			}
		}()
	}

	flagSet := NewFlagSet()
	flagSet.Parse(os.Args[1:])
	if _daemon {
		flagSet := daemon.NewFlagSet()
		flagSet.Parse(os.Args[1:])
		if err := daemon.Run(flagSet); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		flagSet := client.NewFlagSet()
		flagSet.Parse(os.Args[1:])
		if err := client.Run(flagSet); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
