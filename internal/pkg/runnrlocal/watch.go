package runnrlocal

import (
	"fmt"
	"github.com/mdev5000/appwatchertools"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type watchCommand struct {
}

func (w *watchCommand) Run(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	r := &runner{
		appWatcher: appwatchertools.NewAppWatcher(),
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	r.workingDir = wd
	r.appWatcher.Dir = wd

	buildPath := filepath.Join(r.workingDir, ".tmpbins")
	if err := os.MkdirAll(buildPath, 0775); err != nil {
		return err
	}
	r.serverExePath = filepath.Join(buildPath, "app")

	// Only watch application files.
	r.appWatcher.FileFilter = func(path string) (bool, error) {
		return r.isApplicationPath(path), nil
	}

	r.appWatcher.OnChangeFn = r.onChange

	// OnChange will rebuild the application and remake the exe ./app, after that
	// we can run the application.
	r.appWatcher.CommandFn = func() *exec.Cmd {
		// Usually you would run the compiled application here.
		// return exec.Command("./app")
		//return exec.Command("echo", "running app")
		return exec.Command(r.serverExePath)
	}

	// Run the watcher.
	return r.appWatcher.Run(ctx)
}

func WatchCommand() *cobra.Command {
	c := &watchCommand{}
	cmd := &cobra.Command{
		Use:  "watch",
		RunE: c.Run,
	}
	return cmd
}

type runner struct {
	appWatcher    *appwatchertools.AppWatcher
	workingDir    string
	serverExePath string
}

// When a file changes run build and main applications and is there's no errors
// start the application.
func (r *runner) onChange(files []string, isInit bool) bool {
	// Usually you could build the app here:
	//
	//if err := r.appWatcher.RunCommand("go", "build", "-o", "app", "main/main.go"); err != nil {
	//	r.appWatcher.ExeLogger.Error(err)
	//	return false
	//}

	if !isInit && allFilesAre(r.isGeneratedFile, files) {
		return false
	}

	if isInit || anyFileIs(r.isFrontendFile, files) {
		if isInit || anyFileIs(r.isVectyTemplate, files) {
			fmt.Println("rebuild frontend templates...")
			if err := r.appWatcher.RunCommand("make", "gen.frontend"); err != nil {
				r.appWatcher.ExeLogger.Error(err)
				return false
			}
		}

		fmt.Println("rebuilding frontend ...")
		if err := r.appWatcher.RunCommand("make", "build.frontend"); err != nil {
			r.appWatcher.ExeLogger.Error(err)
			return false
		}
		fmt.Println("rebuilding frontend (done)")
	}

	//if anyFileIs(r.isBackendFile, files) {
	//
	//}
	fmt.Println("rebuilding backend ...")
	if err := r.appWatcher.RunCommand("go", "build", "-o",
		r.serverExePath, "./cmd/server/main.go"); err != nil {
		r.appWatcher.ExeLogger.Error(err)
		return false
	}
	fmt.Println("rebuilding backend (done)")
	return true
}

func allFilesAre(fn func(path string) bool, files []string) bool {
	for _, f := range files {
		if !fn(f) {
			return false
		}
	}
	return true
}

func anyFileIs(fn func(path string) bool, files []string) bool {
	for _, f := range files {
		if fn(f) {
			return true
		}
	}
	return false
}

// Limit what files and directories are watched
func (r *runner) isApplicationPath(path string) bool {
	return r.isFrontendFile(path) ||
		r.isBackendFile(path)
}

func (r *runner) isBackendFile(path string) bool {
	return strings.HasPrefix(path, filepath.Join(r.workingDir, "webb"))
}

func (r *runner) isFrontendFile(path string) bool {
	return strings.HasPrefix(path, filepath.Join(r.workingDir, "webf"))
}

func (r *runner) isVectyTemplate(path string) bool {
	return strings.HasSuffix(path, ".vtpl")
}

func (r *runner) isGeneratedFile(path string) bool {
	return strings.HasSuffix(path, ".vtpl.go")
}
