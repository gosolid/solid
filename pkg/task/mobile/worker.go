package main

// func startWorker() {
// 	startup := make(chan bool)

// 	go func() {
// 		isolates.Initialize()
// 		os.Setenv("PORT", "3086")

// 		signal.Ignore(syscall.SIGSEGV)

// 		wd, err := os.Getwd()
// 		if err != nil {
// 			return err
// 		}

// 		options := TaskOptions{
// 			WorkingDirectory: wd,
// 			Requires:         requires,
// 			Script:           flagSet.Arg(0),
// 		}

// 		firstRun := true

// 		var restart func()
// 		restart = func() {
// 			defer func() {
// 				if r := recover(); r != nil {
// 					fmt.Println("recovered in callback handler", r)
// 					restart()
// 				}
// 			}()

// 			for {
// 				if fs, script, err := common.CreateFS(path.Join(options.WorkingDirectory, options.Script)); err != nil {
// 					log.Println(err)
// 					return err
// 				} else if task, err := solid.NewTask(solid.TaskOptions{
// 					FS:       fs,
// 					Requires: options.Requires,
// 					Env:      options.Env,
// 					Script:   script,
// 					Stdin:    os.Stdin,
// 					Stdout:   os.Stdout,
// 					Stderr:   os.Stderr,
// 				}); err != nil {
// 					isolates.For(task.GetExecutionContext()).Error(err)
// 					return nil
// 				} else if err := task.Start(); err != nil {
// 					isolates.For(task.GetExecutionContext()).Error(err)
// 					return nil
// 				} else {
// 					if exitCode, err := task.Wait(task.GetExecutionContext()); err != nil {
// 						log.Println(err)
// 						os.Exit(exitCode)
// 					} else {
// 						os.Exit(exitCode)
// 					}
// 					return nil
// 				}
// 			}
// 		}

// 		restart()
// 	}()

// 	<-startup
// }
