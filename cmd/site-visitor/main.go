package main

import (
	"fmt"
	"os"
	"context"

	"github.com/spf13/cobra"
	"github.com/ystkfujii/site-visitor/pkg/visitor"

)


func cmdNewSiteVisitorCommand() *cobra.Command {
	//var err error
	cmd := &cobra.Command{
		Use: "site-visitor",
		Short: "Launch a site-visitor",
		Long: "hogehoge",
		Run: func(_ *cobra.Command, _ []string) {
			siteVisitor := visitor.NewSiteVisitor()

			_ = siteVisitor.Run(context.Background())
		},
	}
	return cmd
}


func main() {
	command := cmdNewSiteVisitorCommand()

	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %+v\n", err)
		os.Exit(1)
	}
}
