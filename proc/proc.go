package proc

import (
	"os"
	"os/exec"
	"io"
	"strings"
	"bufio"
)

type Proc struct {
	*exec.Cmd
	Stdin		chan string
	Stdout		chan string
	StderrLines	[]string
	Wait		chan error
}

func writeProc(in_q chan string, in_writer io.WriteCloser) {
	for {
		in_data, ok := <-in_q
		if !ok {
			return
		}
		if len(in_data) > 0 {
			in_writer.Write([]byte(in_data))
		}
	}
}

func readProc(out_q chan string, out_reader_raw io.ReadCloser) {
	out_reader := bufio.NewReader(out_reader_raw)
	for {
		out_line, _, err:= out_reader.ReadLine()
		if err == io.EOF {
			return
		}

		out_q <-string(out_line)
	}
}

func readProcStderr(out_s *[]string, out_reader_raw io.ReadCloser, prefix string) {
	out_reader := bufio.NewReader(out_reader_raw)
	for {
		out_line, _, err:= out_reader.ReadLine()
		if err == io.EOF {
			return
		}

		os.Stderr.WriteString(prefix)
		os.Stderr.Write(out_line)
		os.Stderr.WriteString("\n")
		out_s = append(*out_s, string(out_line))
	}
}

func New(cmd string, args ...string) *Proc {
	proc := &Proc{
		exec.Command(cmd, args...),
		make(chan string),
		make(chan string),
		[]string{},
		make(chan error),
	}

	stdin, _ := proc.StdinPipe()
	go func() {
		writeProc(proc.Stdin, stdin)
		stdin.Close()
		//close(proc.Stdin) //FIXME
	}()

	stdout, _ := proc.StdoutPipe()
	go func() {
		readProc(proc.Stdout, stdout)
		stdout.Close()
		close(proc.Stdout)
	}()

	stderr, _ := proc.StderrPipe()
	go func() {
		readProcStderr(&proc.StderrLines, stderr, cmd+": ")
		stderr.Close()
	}()

	go func() {
		proc.Wait <-proc.Run()
	}()

	return proc
}

