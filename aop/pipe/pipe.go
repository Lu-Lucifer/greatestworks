// Package pipe extends os.exec, making it easier to create pipes to subcommands.
package pipe

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
)

// Cmd is drop-in replacement for exec.Cmd, extended with the Pipe method.
type Cmd struct {
	*exec.Cmd
	closeAfterStart []io.Closer // closed after Start()
	closeAfterWait  []io.Closer // closed after Wait()
}

// CommandContext returns a new Cmd. See exec.CommandContext for details.
func CommandContext(ctx context.Context, name string, arg ...string) *Cmd {
	return &Cmd{Cmd: exec.CommandContext(ctx, name, arg...)}
}

// RPipe returns the reading side of a pipe that will be connected to the
// subprocess when the command starts.
//
// The subprocess can write to the pipe using the file descriptor returned by
// RPipe. Cmd.ExtraFiles should not be modified directly if RPipe is called.
//
// Wait will close the pipe after seeing the command exit, so most callers need
// not close the pipe themselves. It is thus incorrect to call Wait before all
// reads from the pipe have completed. For the same reason, it is incorrect to
// use Run when using Pipe. See the exec.Cmd.StdoutPipe example [1] for
// idiomatic usage.
//
// [1]: https://pkg.go.dev/os/exec#example-Cmd.StdoutPipe
func (c *Cmd) RPipe() (int, io.ReadCloser, error) {
	if c.Process != nil {
		return 0, nil, fmt.Errorf("capture: RPipe after process started")
	}
	r, w, err := os.Pipe()
	if err != nil {
		return 0, nil, err
	}
	fd := c.registerPipe(r, w)
	return int(fd), r, nil
}

// WPipe returns the writer side of a pipe that will be connected to the
// subprocess when the command starts.
//
// The subprocess can read from the pipe using the file descriptor returned by
// WPipe. Cmd.ExtraFiles should not be modified directly if WPipe is called.
//
// The pipe will be closed automatically after Wait sees the command exit. A
// caller need only call Close to force the pipe to close sooner. For example,
// if the command being run will not exit until standard input is closed, the
// caller must close the pipe.
func (c *Cmd) WPipe() (int, io.WriteCloser, error) {
	if c.Process != nil {
		return 0, nil, fmt.Errorf("capture: WPipe after process started")
	}
	// TODO: StdinPipe [1] makes sure w is only closed once. Understand why
	// they do that and do the same if needed.
	//
	// [1]: https://cs.opensource.google/go/go/+/refs/tags/go1.18.4:src/os/exec/exec.go;l=593
	r, w, err := os.Pipe()
	if err != nil {
		return 0, nil, err
	}
	fd := c.registerPipe(w, r)
	return fd, w, nil
}

func (c *Cmd) registerPipe(local, remote *os.File) int {
	// From https://pkg.go.dev/os/exec#Cmd.ExtraFiles: "entry i becomes file
	// descriptor 3+i".
	fd := 3 + len(c.Cmd.ExtraFiles)
	c.Cmd.ExtraFiles = append(c.Cmd.ExtraFiles, remote)
	c.closeAfterStart = append(c.closeAfterStart, remote)
	c.closeAfterWait = append(c.closeAfterWait, local)
	return fd
}

// Start is identical to exec.Command.Start.
func (c *Cmd) Start() error {
	if err := c.Cmd.Start(); err != nil {
		return err
	}
	closeAll(&c.closeAfterStart)
	return nil
}

// Wait is identical to exec.Command.Wait.
func (c *Cmd) Wait() error {
	if err := c.Cmd.Wait(); err != nil {
		return err
	}
	closeAll(&c.closeAfterWait)
	return nil
}

// Cleanup cleans up any unused resources.
func (c *Cmd) Cleanup() {
	closeAll(&c.closeAfterStart)
	closeAll(&c.closeAfterWait)
}

func closeAll(files *[]io.Closer) {
	for _, f := range *files {
		f.Close()
	}
	*files = nil
}
