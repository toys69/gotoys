package cmdx

import (
	"context"
	"io"
	"os"
	"os/exec"
	"syscall"
)

type Command struct {
	cmd *exec.Cmd
}

func CMD(name string, args ...string) *Command {
	c := &Command{
		cmd: exec.Command(name, args...),
	}
	c.cmd.Stdout = os.Stdout
	c.cmd.Stderr = os.Stderr
	return c
}

// CMDContext creates a Command with a context.
// The provided context is used to kill the process (by calling os.Process.Kill)
// if the context becomes done before the command completes on its own.
func CMDContext(ctx context.Context, name string, args ...string) *Command {
	c := &Command{
		cmd: exec.CommandContext(ctx, name, args...),
	}
	c.cmd.Stdout = os.Stdout
	c.cmd.Stderr = os.Stderr
	return c
}

func (c *Command) set(opt func(c *exec.Cmd)) *Command {
	opt(c.cmd)
	return c
}

func (c *Command) Dir(dir string) *Command {
	return c.set(func(c *exec.Cmd) { c.Dir = dir })
}

// Env sets the environment variables for the command.
// Each element is in the form "key=value".
func (c *Command) Env(env []string) *Command {
	return c.set(func(c *exec.Cmd) { c.Env = env })
}

// Stdin sets the standard input for the command.
func (c *Command) Stdin(stdin io.Reader) *Command {
	return c.set(func(c *exec.Cmd) { c.Stdin = stdin })
}

// Stdout sets the standard output for the command.
func (c *Command) Stdout(stdout io.Writer) *Command {
	return c.set(func(c *exec.Cmd) { c.Stdout = stdout })
}

// Stderr sets the standard error for the command.
func (c *Command) Stderr(stderr io.Writer) *Command {
	return c.set(func(c *exec.Cmd) { c.Stderr = stderr })
}

// ExtraFiles specifies additional open files to be inherited by the new process.
func (c *Command) ExtraFiles(files []*os.File) *Command {
	return c.set(func(c *exec.Cmd) { c.ExtraFiles = files })
}

// SysProcAttr holds optional, OS-specific attributes for the command.
func (c *Command) SysProcAttr(attr *syscall.SysProcAttr) *Command {
	return c.set(func(c *exec.Cmd) { c.SysProcAttr = attr })
}

// Run starts the specified command and waits for it to complete.
func (c *Command) Run() error {
	return c.cmd.Run()
}

// Start starts the specified command but does not wait for it to complete.
func (c *Command) Start() error {
	return c.cmd.Start()
}

// Wait waits for the command to exit and waits for any copying to stdin or copying from stdout or stderr to complete.
func (c *Command) Wait() error {
	return c.cmd.Wait()
}

// Output runs the command and returns its standard output.
func (c *Command) Output() ([]byte, error) {
	return c.cmd.Output()
}

// CombinedOutput runs the command and returns its combined standard output and standard error.
func (c *Command) CombinedOutput() ([]byte, error) {
	return c.cmd.CombinedOutput()
}

// StdinPipe returns a pipe that will be connected to the command's standard input when the command starts.
func (c *Command) StdinPipe() (io.WriteCloser, error) {
	return c.cmd.StdinPipe()
}

// StdoutPipe returns a pipe that will be connected to the command's standard output when the command starts.
func (c *Command) StdoutPipe() (io.ReadCloser, error) {
	return c.cmd.StdoutPipe()
}

// StderrPipe returns a pipe that will be connected to the command's standard error when the command starts.
func (c *Command) StderrPipe() (io.ReadCloser, error) {
	return c.cmd.StderrPipe()
}

// String returns a human-readable description of the command.
func (c *Command) String() string {
	return c.cmd.String()
}
