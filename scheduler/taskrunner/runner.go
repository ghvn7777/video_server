package taskrunner

import "log"

type Runner struct {
	Controlller controlChan
	Error controlChan
	Data dataChan
	dataSize int
	longLived bool
	Dispatcher fn
	Executor fn
}

func NewRunner(size int, longlived bool, d fn, e fn) *Runner {
	return &Runner{
		Controlller: make(chan string, 1),
		Error: make(chan string, 1),
		Data: make(chan interface{}, size),
		longLived: longlived,
		dataSize: size,
		Dispatcher: d,
		Executor: e,
	}
}

func (r *Runner) startDispatch() {
	defer func() {
		if !r.longLived {
			close(r.Controlller)
			close(r.Data)
			close(r.Error)
		}
	}()

	for {
		select {
		case c := <- r.Controlller:
			if c == READY_TO_DISPATCH {
				err := r.Dispatcher(r.Data)
				if err != nil {
					log.Printf("Dispatcher error: %v", err)
					r.Error <- CLOSE
				} else {
					r.Controlller <- READY_TO_EXECUTE
				}
			} else if c == READY_TO_EXECUTE {
				err := r.Executor(r.Data)
				if err != nil {
					log.Printf("Executor error: %v", err)
					r.Error <- CLOSE
				} else {
					r.Controlller <- READY_TO_DISPATCH
				}
			} else {
				log.Printf("unknow controller")
				r.Error <- CLOSE
			}
		case e := <- r.Error:
			if e == CLOSE {
				return
			}
		default:

		}
	}
}

func (r *Runner) StartAll() {
	r.Controlller <- READY_TO_DISPATCH
	r.startDispatch()
}