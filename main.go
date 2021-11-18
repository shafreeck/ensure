package main

import (
	"context"
	"fmt"
	"math"
	"os"
	"os/exec"
	"time"

	"github.com/shafreeck/cortana"
	"github.com/shafreeck/retry"
)

func run(cmd *exec.Cmd) error {
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return retry.Retriable(err)
	}
	return nil
}

func main() {
	args := struct {
		Retries int           `cortana:"--retries, -, -1, number of retries if executed failed, -1 means unlimited"`
		Timeout time.Duration `cortana:"--wait, -w, -1s, time duration before return, -1s means unlimited"`
		Command string        `cortana:"command, -, -"`
		// only used for help info
		_ []string `cortana:"args"`
	}{}

	sep := len(os.Args)
	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--help", "-h":
			continue
		case "--retries", "-r":
			fallthrough
		case "--wait", "-w":
			i++
			continue
		}
		sep = i + 1
		break
	}

	cortana.Parse(&args, cortana.WithArgs(os.Args[1:sep]))

	var cancel context.CancelFunc
	ctx := context.Background()
	if args.Timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, args.Timeout)
		defer cancel()
	}

	if args.Retries < 0 {
		args.Retries = math.MaxInt64
	}

	err := retry.EnsureN(ctx, args.Retries, func() error {
		cmd := exec.Command(args.Command, os.Args[sep:]...)
		return run(cmd)
	})
	if err != nil {
		fmt.Println(err)
	}
}
