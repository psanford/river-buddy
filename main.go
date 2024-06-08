package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/raitonoberu/riverpipe/client"
	"github.com/raitonoberu/riverpipe/client/event"
)

var doGotoNextWorkspace = flag.Bool("ws-next", false, "goto next workspace")
var doGotoPrevWorkspace = flag.Bool("ws-prev", false, "goto next workspace")

func main() {
	flag.Parse()

	if !*doGotoNextWorkspace && !*doGotoPrevWorkspace {
		flag.Usage()
		os.Exit(1)
	}

	client, err := client.New()
	if err != nil {
		panic(err)
	}
	defer client.Release()

	ch := make(chan event.Event, 16)
	go client.Run(ch)

	for evt := range ch {
		switch e := evt.(type) {
		case event.FocusedTags:
			tag := e.Tags & 511

			nextTag := tag << 1
			if nextTag > 256 {
				nextTag = 1
			}
			prevTag := tag >> 1
			if prevTag == 0 {
				prevTag = 256
			}

			if *doGotoNextWorkspace {
				tag = nextTag
			}
			if *doGotoPrevWorkspace {
				tag = prevTag
			}

			fmt.Println("riverctl", "set-focused-tags", strconv.Itoa(int(tag)))
			exec.Command("riverctl", "set-focused-tags", strconv.Itoa(int(tag))).Run()
			return
		}
	}
}

// {"event":"view_tags","args":{"tags":[2,1,2147483664,2147483652,2147483652,1,1,1,1,1]}}
// {"event":"focused_tags","args":{"tags":2147483664}}
// {"event":"urgent_tags","args":{"tags":0}}
// {"event":"layout_name","args":{"name":"rivertile - left"}}
// {"event":"mode","args":{"name":"normal"}}
// {"event":"focused_output","args":{"output":6}}
// {"event":"focused_view","args":{"title":"foot"}}
