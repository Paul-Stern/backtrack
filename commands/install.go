package commands

import (
	"fmt"
	"os/exec"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var installCompletionsCmd = &cobra.Command{
	Use:   "install",
	Short: "install completions",
	Run: func(cmd *cobra.Command, args []string) {
		InstallCompletions()
	},
}

func InstallCompletions() {
	if err := rootCmd.GenBashCompletionFile("backtrack-completion.bash"); err != nil {
		log.Fatal().Err(err).Msg("Failed to generate bash completions")
	}

	// mkdir bash-completion dir
	cmd := exec.Command("mkdir", "-p", "~/.local/share/bash-completion/completions")
	if err := cmd.Run(); err != nil {
		log.Fatal().Err(err).Msg("Failed to create bash-completion dir")
	}
	// move bash-completion file
	cmd = exec.Command("mv", "backtrack-completion.bash", "~/.local/share/bash-completion/completions/backtrack-completion.bash")
	if err := cmd.Run(); err != nil {
		log.Fatal().Err(err).Msg("Failed to move bash-completion file")
	}

	fmt.Println("bash completions generated")
	fmt.Println("use 'source ~/.bashrc' or restart shell to load completions")

}
