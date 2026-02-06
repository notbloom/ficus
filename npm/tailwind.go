package npm

import (
	"os/exec"

	"github.com/charmbracelet/log"
	"github.com/notbloom/ficus/config"
)

func RunTailwind(cfg *config.Config) error {
	args := []string{"tailwindcss", "-i", cfg.InputCSSPath(), "-o", cfg.OutputCSSPath()}
	cmd := exec.Command("npx", args...)
	err := cmd.Run()

	if err != nil {
		log.Error("Tailwind Error:", "error", err)
		return err
	}

	log.Debug("Tailwind compiled", "output", cfg.OutputCSSPath())
	return nil
}
