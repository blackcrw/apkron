package cli

import (
	"sync"

	"github.com/blackcrw/apkron/internal"
	"github.com/blackcrw/apkron/internal/apktool"
	"github.com/blackcrw/apkron/internal/printer"
	temp_path "github.com/blackcrw/apkron/internal/temp"
	"github.com/blackcrw/apkron/tools/interesting"
	"github.com/spf13/cobra"
)

var (
	sync_wait_group sync.WaitGroup

	root = &cobra.Command {
		Use:     internal.PROJECT_NAME,
		Long:    internal.PROJECT_DESCRIPTION,
		Version: internal.PROJECT_VERSION,
		Run:     root_run,
	}
)

func OnInitialize() {
	flags()
	printer.CError(root.Execute())
}

func flags() {
	cobra.OnInitialize(internal.InitApkron)

	root.PersistentFlags().StringP("app", "a", " ", "Application to perform the operation" + printer.MSG_REQUIRED)
	root.PersistentFlags().BoolP("verbose", "v", false, "Verbosity mode.")

	root.MarkPersistentFlagRequired("app")
}

func root_run(cmd *cobra.Command, args []string) {
	var flag_app, _ = cmd.Flags().GetString("app")

	var temporary_directory, err = temp_path.CreateTempDirectory()
	defer temp_path.DeleteTempDirectory(temporary_directory)

	{	printer.CError(err)
		printer.Info("Temporary Directory Created", temporary_directory)	}

	sync_wait_group.Add(1)

	go func() {
		{	var _, err = apktool.Decompile(flag_app, temporary_directory)
			printer.CError(err)	}
		
		defer sync_wait_group.Done()
	}()

	printer.Done("Decompiling application!\n")
	
	sync_wait_group.Wait()
	
	printer.Info("Application Permissions:\n")
	
	if perms, err := interesting.AndroidManifest(temporary_directory); err != nil {
		printer.CError(err)
	} else if len(perms.Permissions) > 0 {
		for _, x := range perms.Permissions {
			printer.Done("Name:", x.Name)
			printer.PrinterPrefixColor("   -", printer.MAGENTA, "Description:", x.Description)
			printer.PrinterPrefixColor("   -", printer.MAGENTA, "Match:", x.Match, "\n")
		}
	} else {
		printer.Danger("No permissions found\n")
	}
}