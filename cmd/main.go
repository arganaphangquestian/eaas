package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/arganaphangquestian/eaas/data"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"github.com/muesli/termenv"
)

var (
	DATABASE_URL string
	VAULT_TOKEN  string
	queries      *data.Queries
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Can't Load .env Config")
	}
	DATABASE_URL = fmt.Sprintf("postgres://%s:%s@localhost:5432/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)
	VAULT_TOKEN = os.Getenv("VAULT_DEV_ROOT_TOKEN_ID")
}

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, DATABASE_URL)
	if err != nil {
		log.Fatal("Can't connect to database")
	}
	queries = data.New(conn)
	p := tea.NewProgram(newModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

type Menu int

const (
	IDLE Menu = iota
	LIST
	SEED
	ADD
)

func (m Menu) String() string {
	return [...]string{"Idle", "List", "Add", "Seed"}[m]
}

func (m Menu) View() string {
	return [...]string{`What do you want?

0. Menu
1. List
2. Seed
3. Add`, `List Customer Data`, `Add Customer Data`, `Seed Customer Data`}[m]
}

type model struct {
	menu Menu
}

func newModel() model {
	return model{
		menu: IDLE,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "ctrl+0":
			m.menu = IDLE
		case "ctrl+1":
			m.menu = LIST
		case "ctrl+2":
			m.menu = SEED
		case "ctrl+3":
			m.menu = ADD
		}
	}
	return m, nil
}

func (m model) View() string {
	termenv.ClearScreen()
	s := m.menu.View()
	s += "\n\ntype `ctrl+c` or `q` to quit"
	return s
}
