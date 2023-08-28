package client

import "C"

import (
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"log"
	"net"
	"os"
	"path"
	"strings"
	"sync"

	"os/signal"
	"syscall"

	"github.com/gosolid/solid/pkg/common"
	"github.com/gosolid/solid/pkg/solid"
	"github.com/grexie/isolates"
	"github.com/hashicorp/yamux"
)

type TaskOptions struct {
	WorkingDirectory string            `v8:"cwd"`
	Requires         []string          `v8:"requires"`
	Script           string            `v8:"script"`
	Env              map[string]string `v8:"env"`
}

var modulePath *string

type stringFlags []string

func (s *stringFlags) String() string {
	return strings.Join(*s, ",")
}

func (s *stringFlags) Set(value string) error {
	*s = append(*s, value)
	return nil
}

type QuietFlagSet struct {
	*flag.FlagSet
}

func (f *QuietFlagSet) failf(format string, a ...interface{}) error {
	return nil
}

var standalone bool
var socket string
var requires stringFlags

func newFlagSet(errorHandling flag.ErrorHandling) *flag.FlagSet {
	flagSet := flag.NewFlagSet("run", errorHandling)
	flagSet.StringVar(&socket, "socket", "/var/run/com.grexie.Solid.socket", "daemon communication socket")
	flagSet.BoolVar(&standalone, "standalone", false, "run solid in standalone mode")
	flagSet.Var(&requires, "r", "Require modules before running program")

	return flagSet
}

func NewFlagSet() *flag.FlagSet {
	flagSet := newFlagSet(flag.ContinueOnError)
	buf := bytes.NewBuffer([]byte{})
	flagSet.SetOutput(buf)
	return flagSet
}

func Run(flagSet *flag.FlagSet) error {
	isolates.Initialize()
	signal.Ignore(syscall.SIGSEGV)

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	options := TaskOptions{
		WorkingDirectory: wd,
		Requires:         requires,
		Script:           flagSet.Arg(0),
	}

	getenvironment := func(data []string, getkeyval func(item string) (key, val string)) map[string]string {
		items := make(map[string]string)
		for _, item := range data {
			key, val := getkeyval(item)
			items[key] = val
		}
		return items
	}
	options.Env = getenvironment(os.Environ(), func(item string) (key, val string) {
		splits := strings.SplitN(item, "=", 2)
		key = splits[0]
		val = splits[1]
		return
	})

	if standalone {
		if fs, script, err := common.CreateFS(path.Join(options.WorkingDirectory, options.Script)); err != nil {
			log.Println(err)
			return err
		} else if task, err := solid.NewTask(solid.TaskOptions{
			FS:       fs,
			Requires: options.Requires,
			Env:      options.Env,
			Script:   script,
			Stdin:    os.Stdin,
			Stdout:   os.Stdout,
			Stderr:   os.Stderr,
		}); err != nil {
			isolates.For(task.GetExecutionContext()).Error(err)
			return nil
		} else if err := task.Start(); err != nil {
			isolates.For(task.GetExecutionContext()).Error(err)
			return nil
		} else {
			if exitCode, err := task.Wait(task.GetExecutionContext()); err != nil {
				log.Println(err)
				os.Exit(exitCode)
			} else {
				os.Exit(exitCode)
			}
			return nil
		}
	} else if addr, err := net.ResolveUnixAddr("unix", socket); err != nil {
		return err
	} else if conn, err := net.DialUnix("unix", nil, addr); err != nil {
		return err
	} else if session, err := yamux.Client(conn, nil); err != nil {
		return err
	} else if stream, err := session.Open(); err != nil {
		return err
	} else if options, err := json.Marshal(options); err != nil {
		return err
	} else if _, err := stream.Write(options); err != nil {
		return err
	} else {
		stream.Close()

		var wg sync.WaitGroup
		wg.Add(3)

		go func() {
			if stdin, err := session.Open(); err != nil {
				return
			} else {
				stdin.Write([]byte{0x0})
				io.Copy(stdin, os.Stdin)
			}
			wg.Done()
		}()

		go func() {
			if stdout, err := session.Open(); err != nil {
				return
			} else {
				stdout.Write([]byte{0x1})
				io.Copy(os.Stdout, stdout)
			}
			wg.Done()
		}()

		go func() {
			if stderr, err := session.Open(); err != nil {
				return
			} else {
				stderr.Write([]byte{0x2})
				io.Copy(os.Stderr, stderr)
			}
			wg.Done()
		}()

		wg.Wait()
	}

	return nil
}
