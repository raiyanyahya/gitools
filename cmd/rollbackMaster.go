package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var commitid string
var rollbackMasterCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Updates you master branch to a previous commit",
	Long:  "This updates your master branch with an older commit. The git history will not be overwritten.",
	Run:   rollbackMaster,
}

func executeCommand() error {
	cmd := exec.Command("git", "reset", commitid)
	_, err := cmd.Output()
	if err != nil {
		return err
	}
	fmt.Println("reset to the commit id...")
	cmd = exec.Command("git", "reset", "--soft", "HEAD@{1}")
	_, err = cmd.Output()
	if err != nil {
		return err
	}
	fmt.Println("reset the HEAD...")
	commitmsg := fmt.Sprintf("Revert to %v", commitid)
	cmd = exec.Command("git", "commit", "-m", commitmsg)
	_, err = cmd.Output()
	if err != nil {
		return err
	}
	fmt.Println("commit changes...")
	cmd = exec.Command("git", "reset", "--hard")
	_, err = cmd.Output()
	if err != nil {
		return err
	}
	fmt.Println("DONE.")
	return nil
}
func rollbackMaster(cmd *cobra.Command, args []string) {
	fmt.Println("removing changes and moving to a fresh copy of the master branch...")
	err := executeCommand()
	if err != nil {
		print(err)
	}
	fmt.Println("Please verify the change and run a git push with the -f flag")
}

func init() {
	rollbackMasterCmd.Flags().StringVarP(&commitid, "commit", "c", "", "the commit id you want to rollback the master to.")
	rollbackMasterCmd.MarkFlagRequired("commit")
	rootCmd.AddCommand(rollbackMasterCmd)
}
