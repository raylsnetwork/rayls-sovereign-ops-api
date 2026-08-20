package app

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/raylsnetwork/rayls-sovereign-ops-api/flags"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/logger"
	"github.com/raylsnetwork/rayls-sovereign-ops-api/withstack"
)

var (
	rootCMD = &cobra.Command{
		Use: "",
	}

	runCMD = &cobra.Command{
		Use:   "run",
		Short: "Run API",
		Long:  "Run API",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := RunAPI(viper.GetString(flags.ConfigFlagName)); err != nil {
				return withstack.Wrap(err)
			}
			return nil
		},
	}

	workerCMD = &cobra.Command{
		Use:   "worker",
		Short: "Run background worker",
		Long:  "Run blockchain indexer, event listeners, and NATS consumers",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := RunWorker(viper.GetString(flags.ConfigFlagName)); err != nil {
				return withstack.Wrap(err)
			}
			return nil
		},
	}

	identityCMD = &cobra.Command{
		Use:   "identity",
		Short: "Run the shared identity service",
		Long: "Run the identity service: login (SIWE, OAuth, email), the user record and " +
			"session JWTs, shared by every chain. Owns no chain state — roles and signing " +
			"wallets are resolved per chain by each ops-api.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := RunIdentity(viper.GetString(flags.ConfigFlagName)); err != nil {
				return withstack.Wrap(err)
			}
			return nil
		},
	}
)

func init() {
	flags.BindFlags(runCMD)
	flags.BindFlags(workerCMD)
	flags.BindFlags(identityCMD)
}

func Execute() {
	rootCMD.AddCommand(runCMD, workerCMD, identityCMD)
	if err := rootCMD.Execute(); err != nil {
		logger.Error("failed to execute root cmd", "error", err)
	}
}
