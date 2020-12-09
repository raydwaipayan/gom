package process

import (
	"log"
	"os"
	"path/filepath"
)

// Proc wraps parameters unique to the
// particular process and methods for manipulating
// the process
type Proc struct {
	Name    string
	Cmd     string
	Argv    []string
	Path    string
	process *os.Process
}

// Start will execute the given Cmd and stores the process`
func (proc *Proc) Start() error {
	out := filepath.Join(proc.Path, "gom.out")
	outfile, err := os.Create(out)
	if err != nil {
		log.Print(err)
		return err
	}

	procAttr := &os.ProcAttr{
		Dir: proc.Path,
		Env: os.Environ(),
		Files: []*os.File{
			os.Stdin,
			outfile,
			outfile,
		},
	}

	process, err := os.StartProcess(proc.Cmd, proc.Argv, procAttr)
	if err != nil {
		log.Print(err)
		return err
	}
	proc.process = process
	return nil
}

// Stop stops the given process. Sends a SIGKILL
func (proc *Proc) Stop() error {
	err := proc.process.Signal(os.Kill)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

// GracefullyStop stops the given process. Sends a SIGINT
// Note that SIGINT doesn't work on windows
func (proc *Proc) GracefullyStop() error {
	err := proc.process.Signal(os.Interrupt)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

// Restart restarts the given process.
// Exits the process gracefully with SIGINT and starts it again
func (proc *Proc) Restart() error {
	err := proc.GracefullyStop()
	if err != nil {
		log.Print(err)
		return err
	}

	err = proc.Start()
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}
