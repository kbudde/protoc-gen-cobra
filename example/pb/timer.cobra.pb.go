// Code generated by protoc-gen-cobra. DO NOT EDIT.

package pb

import (
	client "github.com/NathanBaulch/protoc-gen-cobra/client"
	flag "github.com/NathanBaulch/protoc-gen-cobra/flag"
	iocodec "github.com/NathanBaulch/protoc-gen-cobra/iocodec"
	proto "github.com/golang/protobuf/proto"
	cobra "github.com/spf13/cobra"
	grpc "google.golang.org/grpc"
	io "io"
)

func TimerClientCommand(options ...client.Option) *cobra.Command {
	cfg := client.NewConfig(options...)
	cmd := &cobra.Command{
		Use:   "timer",
		Short: "Timer service client",
		Long:  "",
	}
	cfg.BindFlags(cmd.PersistentFlags())
	cmd.AddCommand(
		_TimerTickCommand(cfg),
	)
	return cmd
}

func _TimerTickCommand(cfg *client.Config) *cobra.Command {
	req := &TickRequest{}

	cmd := &cobra.Command{
		Use:   "tick",
		Short: "Tick RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), cfg.EnvVarPrefix); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), cfg.EnvVarPrefix, "TIMER", "TICK"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewTimerClient(cc)
				v := &TickRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				stm, err := cli.Tick(cmd.Context(), v)

				if err != nil {
					return err
				}

				for {
					res, err := stm.Recv()
					if err != nil {
						if err == io.EOF {
							break
						}
						return err
					}
					if err = out(res); err != nil {
						return err
					}
				}
				return nil

			})
		},
	}

	cmd.PersistentFlags().Int32Var(&req.Interval, "interval", 0, "")

	return cmd
}
