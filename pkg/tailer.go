package tailer

import "github.com/segmentio/go-loggly"
import "github.com/segmentio/go-log"
import "bufio"
import "io"
import "os"

// Tailer
type Tailer struct {
	r    io.Reader
	w    io.Writer
	l    *loggly.Client
	exit chan bool
}

// NewTailer creates a new tailer reading from `r`
// and writing to the loggly client `l`.
func NewTailer(r io.Reader, l *loggly.Client) *Tailer {
	return &Tailer{r, os.Stdout, l, make(chan bool)}
}

// Start tailing.
func (t *Tailer) Start() {
	go t.Tail()
}

// Stop tailing and flush loggly.
func (t *Tailer) Stop() {
	log.Info("stopping")
	close(t.exit)

	log.Info("flushing")
	t.l.Flush()

	log.Info("stopped")
}

// Tail the reader.
func (t *Tailer) Tail() {
	buf := bufio.NewReader(t.r)

	for {
		select {
		case <-t.exit:
			log.Info("exiting")
			return
		default:
			line, err := buf.ReadBytes('\n')

			_, err = t.w.Write(line)

			if err != nil {
				log.Error("failed to write: %s", err)
				break
			}

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Error("failed to read: %s", err)
				break
			}

			_, err = t.l.Write(line)

			if err != nil {
				log.Error("failed to write to loggly: %s", err)
				break
			}
		}
	}
}
