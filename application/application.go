package application

import (
	"context"
	"fmt"
	"interview/application/api"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	ctx        context.Context
	HTTPServer *fiber.App
}

func New(ctx context.Context, version string) (r App, err error) {
	r.ctx = ctx

	r.HTTPServer = api.BuildAPI(version)
	return
}

func (a App) Run(port int) error {
	errs := make(chan error)
	defer func() {
		close(errs)
	}()

	// Start server
	go func() {
		if err := a.HTTPServer.Listen(":" + strconv.Itoa(port)); err != nil {
			errs <- err
		}
	}()
	defer func() {
		if err := a.HTTPServer.Shutdown(); err != nil {
			fmt.Println(err)
		}
	}()

	signalC := make(chan os.Signal, 1)
	defer close(signalC)
	signal.Notify(signalC, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case err := <-errs:
		return err
	case s := <-signalC:
		return fmt.Errorf("signal %s", s)
	case <-a.ctx.Done():
	}

	return nil
}
