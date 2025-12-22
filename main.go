package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ethanolivertroy/cmvp-tui/internal/tui"
)

// Version is set at build time via ldflags
var version = "dev"

func main() {
	showVersion := flag.Bool("version", false, "Print version and exit")
	flag.BoolVar(showVersion, "v", false, "Print version and exit (shorthand)")
	flag.Parse()

	if *showVersion {
		fmt.Printf("cmvp %s\n", version)
		return
	}

	p := tea.NewProgram(
		tui.NewModel(),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
}
