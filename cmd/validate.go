package cmd

import (
	"fmt"

	"github.com/ghdwlsgur/gossl/internal"
	"github.com/spf13/cobra"
)

func setDomain(args []string) (string, error) {
	if len(args) < 1 || args[0] == "" {
		return "", fmt.Errorf("please enter your domain. ex) gossl validate naver.com")
	}
	return args[0], nil
}

var (
	validateCommand = &cobra.Command{
		Use:   "validate",
		Short: "Proxy the A record ip address of the cache server to review the application of the certificate.",
		Long:  "Proxy the A record ip address of the cache server to review the application of the certificate.",
		Run: func(_ *cobra.Command, args []string) {
			var (
				err error
			)

			domain, err := setDomain(args)
			if err != nil {
				panicRed(err)
			}

			ips, err := internal.GetRecordIPv4(domain)
			if err != nil {
				panicRed(err)
			}

			for _, ip := range ips {
				err = internal.GetCertificateInfo(ip, domain)
				if err != nil {
					panicRed(err)
				}
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(validateCommand)
}
