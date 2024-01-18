package npm

import (
	"github.com/charmbracelet/log"
	"os/exec"
)

func RunTailwind(input string, output string) {
	args := []string{"tailwindcss", "-i", "./views/input.css", "-o", "./views/assets/css/style.css"}
	//cmd := exec.Command("npx tailwindcss -i ./views/input.css -o ./views/assets/css/style.css")
	cmd := exec.Command("npx", args...)
	err := cmd.Run()

	if err != nil {
		log.Error("Tailwind Error:")
		log.Error(err)
	}
}
