package app

import (
	"fmt"
	mflags "github.com/TryRpc/component/pkg/cli/flag"
	"github.com/TryRpc/component/pkg/cli/term"
	"github.com/TryRpc/component/pkg/errors"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"runtime"
	"strings"
)

type App struct {
	basename    string
	name        string
	description string
	Options     CliOption
	runFunc     RunFunc
	cmd         *cobra.Command
}

// Option 定义可选参数，初始化应用
type Option func(app *App)

func WithOptions(opt CliOption) Option {
	return func(app *App) {
		app.Options = opt
	}
}

// RunFunc 定义回调函数
type RunFunc func(basename string) error

func WithRunFunc(runFunc RunFunc) Option {
	return func(app *App) {
		app.runFunc = runFunc
	}
}

func WithDescription(desc string) Option {
	return func(app *App) {
		app.description = desc
	}
}

func NewApp(basename, name string, opts ...Option) *App {
	app := &App{
		basename: basename,
		name:     name,
	}
	for _, o := range opts {
		o(app)
	}
	app.buildCommand()
	return app
}

func (app *App) buildCommand() {
	cmd := cobra.Command{
		Use:   FormatBaseName(app.basename),
		Short: app.name,
		Long:  app.description,
	}
	cmd.SetOutput(os.Stdout)
	cmd.Flags().SortFlags = true
	mflags.InitFlags(cmd.Flags())
	if app.runFunc != nil {
		cmd.RunE = app.runCommand
	}

	var namedFlagSets mflags.NamedFlagSets
	if app.Options != nil {
		namedFlagSets = app.Options.Flags()
		fs := cmd.Flags()
		for _, f := range namedFlagSets.FlagSets {
			fs.AddFlagSet(f)
		}
		usageFmt := "Usage:\n %s\n"
		//获取窗口高度
		cols, _, _ := term.TerminalSize(cmd.OutOrStdout())
		//根据宽高打印
		cmd.SetHelpFunc(func(command *cobra.Command, i []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "%s\n\n"+usageFmt, cmd.Long, cmd.UseLine())
			mflags.PrintSections(cmd.OutOrStdout(), namedFlagSets, cols)
		})
		cmd.SetUsageFunc(func(command *cobra.Command) error {
			fmt.Fprintf(cmd.OutOrStdout(), usageFmt, cmd.UseLine())
			mflags.PrintSections(cmd.OutOrStdout(), namedFlagSets, cols)
			return nil
		})
	}
	readConfig(app.basename, namedFlagSets.FlagSet("global"))
	namedFlagSets.FlagSet("global").BoolP("help", "h", false, fmt.Sprintf("help for %s", cmd.Name()))
	app.cmd = &cmd
}
func (app *App) runCommand(cmd *cobra.Command, args []string) error {
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return err
	}
	if err := viper.Unmarshal(app.Options); err != nil {
		return err
	}
	if app.Options != nil {
		if errs := app.Options.Validate(); len(errs) != 0 {
			return errors.NewAggregate(errs)
		}
	}
	if app.runFunc != nil {
		app.runFunc(app.basename)
	}
	return nil
}

func FormatBaseName(basename string) string {
	if runtime.GOOS == "windows" {
		basename = strings.ToLower(basename)
		basename = strings.TrimSuffix(basename, ".exe")
	}
	return basename
}
func (a *App) Run() {
	if err := a.cmd.Execute(); err != nil {
		fmt.Printf("%v %v\n", color.RedString("Error:"), err)
		os.Exit(1)
	}
}
