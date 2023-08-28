package daemon

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"path"
	"strings"
	"sync"
	"syscall"

	"github.com/gosolid/solid/pkg/client"
	"github.com/gosolid/solid/pkg/common"
	"github.com/gosolid/solid/pkg/solid"
	"github.com/grexie/isolates"
	"github.com/hashicorp/yamux"
)

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

var daemon bool
var socket string

func newFlagSet(errorHandling flag.ErrorHandling) *flag.FlagSet {
	flagSet := flag.NewFlagSet("daemon", errorHandling)
	flagSet.BoolVar(&daemon, "daemon", true, "")
	flagSet.StringVar(&socket, "socket", "/var/run/com.grexie.Solid.socket", "daemon communication socket")

	return flagSet
}

func NewFlagSet() *flag.FlagSet {
	flagSet := newFlagSet(flag.ContinueOnError)
	buf := bytes.NewBuffer([]byte{})
	flagSet.SetOutput(buf)
	return flagSet
}

func Run(flagSet *flag.FlagSet) error {
	if flagSet.Arg(0) == "install" {
		return installDaemon()
	} else if flagSet.Arg(0) == "uninstall" {
		return uninstallDaemon()
	} else if flagSet.Arg(0) == "start" {
		return startDaemon()
	} else if flagSet.Arg(0) == "stop" {
		return stopDaemon()
	}

	isolates.Initialize()
	signal.Ignore(syscall.SIGSEGV)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	running := true

	if err := os.MkdirAll(path.Dir(socket), 0755); err != nil {
		return err
	} else if os.Chdir(path.Dir(socket)); err != nil {
		return err
	} else if err := os.RemoveAll(socket); err != nil {
		return err
	} else if addr, err := net.ResolveUnixAddr("unix", socket); err != nil {
		return err
	} else if listener, err := net.ListenUnix("unix", addr); err != nil {
		return err
	} else if err := os.Chmod(socket, 0666); err != nil {
		return err
	} else {
		fmt.Println("🚀 solid daemon running at", listener.Addr())

		go func() {
			processes := []int{}

			defer func() {
				listener.Close()
				for _, pid := range processes {
					syscall.Kill(pid, syscall.SIGINT)
				}
				os.Remove(socket)
				c <- nil
			}()

			for running {
				if conn, err := listener.AcceptUnix(); err != nil && !errors.Is(err, net.ErrClosed) {
					log.Println(err)
				} else if errors.Is(err, net.ErrClosed) {
					running = false
					return
				} else {
					go func() {
						var options client.TaskOptions
						if !path.IsAbs(options.Script) {
							options.Script = path.Join(options.WorkingDirectory, options.Script)
							options.Script = path.Clean(options.Script)
						}

						if session, err := yamux.Server(conn, nil); err != nil {
							log.Println(err)
							return
						} else if stream, err := session.Accept(); err != nil {
							log.Println(err)
							return
						} else if data, err := io.ReadAll(stream); err != nil {
							log.Println(err)
							return
						} else if err := json.Unmarshal(data, &options); err != nil {
							log.Println(err)
							return
						} else {
							var stdin, stdout, stderr net.Conn

							var wg sync.WaitGroup
							wg.Add(3)

							for i := 0; i < 3; i++ {
								go func() {
									fd := make([]byte, 1)
									if conn, err := session.Accept(); err != nil {
										log.Println(err)
										return
									} else if _, err := conn.Read(fd); err != nil {
										log.Println(err)
									} else {
										switch fd[0] {
										case 0:
											stdin = conn
											break
										case 1:
											stdout = conn
											break
										case 2:
											stderr = conn
											break
										default:
											panic("invalid stream identifier")
										}
									}
									wg.Done()
								}()
							}

							wg.Wait()

							if fs, script, err := common.CreateFS(path.Join(options.WorkingDirectory, options.Script)); err != nil {
								log.Println(err)
								return
							} else if task, err := solid.NewTask(solid.TaskOptions{
								FS:       fs,
								Requires: options.Requires,
								Env:      options.Env,
								Script:   script,
								Stdin:    stdin,
								Stdout:   stdout,
								Stderr:   stderr,
							}); err != nil {
								isolates.For(task.GetIsolate().GetExecutionContext()).Error(err)
								return
							} else {
								task.OnTerminate(func() error {
									return conn.Close()
								})

								go func() {
									if err := task.Start(); err != nil {
										log.Println(err)
										isolates.For(task.GetIsolate().GetExecutionContext()).Error(err)
										return
									}
								}()
							}
							// command, _ := os.Executable()
							// cmd := exec.Command(command, "run")

							// cmd.SysProcAttr = &syscall.SysProcAttr{
							// 	Foreground: false,
							// }

							// cmd.Dir = "/Users/tim/src/grexie/solid"
							// cmd.Stdin = conn
							// cmd.Stdout = conn
							// cmd.Stderr = conn

							// // raw, _ := conn.SyscallConn()
							// // var cred unix.Xucred
							// // raw.Control(func(fd uintptr) {
							// // 	c, _ := unix.GetsockoptXucred(int(fd), unix.SOL_LOCAL, unix.LOCAL_PEERCRED)
							// // 	cred = *c
							// // })

							// // cmd.SysProcAttr.Credential = &syscall.Credential{
							// // 	Uid:    cred.Uid,
							// // 	Gid:    cred.Groups[0],
							// // 	Groups: cred.Groups[:],
							// // }

							// if err := cmd.Start(); err != nil {
							// 	log.Println(err)
							// } else {
							// 	processes = append(processes, cmd.Process.Pid)

							// 	if err := cmd.Wait(); err != nil {
							// 		log.Println(err)
							// 	} else if err := conn.Close(); err != nil {
							// 		log.Println(err)
							// 	}

							// 	for i, pid := range processes {
							// 		if pid == cmd.Process.Pid {
							// 			processes = append(processes[:i], processes[i+1:]...)
							// 		}
							// 	}
							// }
						}
					}()
				}
			}
		}()

		for a := <-c; a != nil; a = <-c {
			running = false
			listener.Close()
		}
	}

	return nil
}
