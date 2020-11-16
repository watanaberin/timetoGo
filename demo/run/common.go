package common

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

var ErrTimeOut = errors.New("Over Time") //超时
var ErrInterrupt = errors.New("Interrupted")

type Runner struct {
	tasks     []func(int) //切片
	complete  chan error
	timeout   <-chan time.Time //时间内完成才ok 否则中断
	interrupt chan os.Signal
}

func New(tm time.Duration) *Runner {
	return &Runner{
		complete:  make(chan error),        //
		timeout:   time.After(tm),          //同time.Time告知我们时间结束
		interrupt: make(chan os.Signal, 1), //至少收到一个中断所以有缓冲（防止阻塞）存疑
	}

}
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) isInterupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt) //停止信道
		return true
	default:
		return false

	}
}
func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.isInterupt() {
			return ErrInterrupt
		}
		task(id) //run the task
	}
	return nil
}

func (r *Runner) Start() error {
	//希望接受的信号 os？
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err //完成
	case <-r.interrupt:
		return ErrTimeOut
	}
}
