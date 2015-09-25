package main

import (
	"io"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/chzyer/readline"
)

func usage(w io.Writer) {
	io.WriteString(w, `
sayhello: start to display oneline log per second
bye: quit
`[1:])
}

var completer = readline.NewPrefixCompleter(
	readline.PcItem("say",
		readline.PcItem("hello"),
		readline.PcItem("bye"),
	),
	readline.PcItem("bye"),
	readline.PcItem("help"),
	readline.PcItem("go",
		readline.PcItem("build"),
		readline.PcItem("install"),
		readline.PcItem("test"),
	),
)

func main() {
	l, err := readline.NewEx(&readline.Config{
		Prompt:       "\033[31m»\033[0m ",
		HistoryFile:  "/tmp/readline.tmp",
		AutoComplete: completer,
	})
	if err != nil {
		panic(err)
	}
	defer l.Close()

	log.SetOutput(l.Stderr())
	for {
		line, err := l.Readline()
		if err != nil {
			break
		}

		switch {
		case line == "help":
			usage(l.Stderr())
		case strings.HasPrefix(line, "say"):
			line := strings.TrimSpace(line[3:])
			if len(line) == 0 {
				log.Println("say what?")
				break
			}
			go func() {
				for _ = range time.Tick(time.Second) {
					log.Println(line)
				}
			}()
		case line == "bye":
			goto exit
		case line == "":
		default:
			log.Println("you said:", strconv.Quote(line))
		}
	}
exit:
}
