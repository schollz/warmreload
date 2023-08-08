// os.execute("cd /home/we/dust/code/warmreload && go build -v")
// os.execute("cd /home/we/dust/code/warmreload/warmreload --path /home/we/dust/code &")
package main

import (
	"flag"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/bep/debounce"
	"github.com/gorilla/websocket"
	"github.com/rjeczalik/notify"
	log "github.com/schollz/logger"
)

var flagAddr = flag.String("addr", "localhost", "address")
var flagPort = flag.String("port", "5555", "port")
var flagSubProtocol = flag.String("sub", "bus.sp.nanomsg.org", "sub protocol")
var flagSend = flag.String("send", "<wscat>", "format sent piped input")

var flagRecvHost, flagRecvAddress, flagHost, flagAddress, flagPath, flagIgnore string
var flagDebounce int

func init() {
	flag.StringVar(&flagIgnore, "ignore", "data", "path to ignore")
	flag.StringVar(&flagPath, "path", ".", "path to watch")
	flag.IntVar(&flagDebounce, "debounce", 200, "debounce time in milliseconds")
}

func main() {
	flag.Parse()
	// norns.rerun()
	flag.Parse()
	// Create new watcher.
	log.SetLevel("info")
	log.Info("oscnotify started")

	c := make(chan notify.EventInfo, 1)
	pathChanged := ""
	f := func() {
		log.Debugf("sending %s to %s:%d", pathChanged, flagHost, flagPort)
		errReload := reload()
		if errReload != nil {
			log.Error(errReload)
		}
	}

	debounced := debounce.New(time.Duration(flagDebounce) * time.Millisecond)

	flagPath, _ = filepath.Abs(flagPath)
	flagIgnore, _ = filepath.Abs(flagIgnore)
	filepath.Walk(flagPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && !strings.Contains(path, ".git") {
			log.Debugf("watching %s", path)
			if filepath.HasPrefix(path, flagIgnore) {
				log.Tracef("ignoring '%s'", path)
				return nil
			}
			if err := notify.Watch(path, c, notify.Write); err != nil {
				log.Error(err)
			}
		}
		return nil
	})

	defer notify.Stop(c)

	// Block until an event is received.
	for {
		ei := <-c
		log.Debugf("Got event: %s", ei.Path())
		pathChanged = ei.Path()
		debounced(f)
	}

}

func reload() (err error) {
	u := url.URL{Scheme: "ws", Host: *flagAddr + ":" + *flagPort, Path: "/"}
	var cstDialer = websocket.Dialer{
		Subprotocols:     []string{*flagSubProtocol},
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		HandshakeTimeout: 3 * time.Second,
	}

	norns, _, err := cstDialer.Dial(u.String(), nil)
	if err != nil {
		return
	}
	defer norns.Close()
	norns.WriteMessage(websocket.TextMessage, []byte("norns.script.load(norns.state.script)\n"))
	return
}
