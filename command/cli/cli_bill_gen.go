// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-commands'; DO NOT EDIT

package cli

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/imdario/mergo"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/completion"
	"github.com/sacloud/usacloud/command/funcs"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/schema"
	"gopkg.in/urfave/cli.v2"
)

func init() {
	csvParam := params.NewCsvBillParam()
	listParam := params.NewListBillParam()

	cliCommand := &cli.Command{
		Name:  "bill",
		Usage: "A manage commands of Bill",
		Action: func(c *cli.Context) error {
			comm := c.App.Command("list")
			if comm != nil {
				return comm.Action(c)
			}
			return cli.ShowSubcommandHelp(c)
		},
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "year",
				Usage: "set year",
			},
			&cli.IntFlag{
				Name:  "month",
				Usage: "set month",
			},
			&cli.StringFlag{
				Name:  "param-template",
				Usage: "Set input parameter from string(JSON)",
			},
			&cli.StringFlag{
				Name:  "param-template-file",
				Usage: "Set input parameter from file",
			},
			&cli.BoolFlag{
				Name:  "generate-skeleton",
				Usage: "Output skelton of parameter JSON",
			},
			&cli.StringFlag{
				Name:    "output-type",
				Aliases: []string{"out"},
				Usage:   "Output type [table/json/csv/tsv]",
			},
			&cli.StringSliceFlag{
				Name:    "column",
				Aliases: []string{"col"},
				Usage:   "Output columns(using when '--output-type' is in [csv/tsv] only)",
			},
			&cli.BoolFlag{
				Name:    "quiet",
				Aliases: []string{"q"},
				Usage:   "Only display IDs",
			},
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"fmt"},
				Usage:   "Output format(see text/template package document for detail)",
			},
			&cli.StringFlag{
				Name:  "format-file",
				Usage: "Output format from file(see text/template package document for detail)",
			},
			&cli.StringFlag{
				Name:  "query",
				Usage: "JMESPath query(using when '--output-type' is json only)",
			},
		},
		Subcommands: []*cli.Command{
			{
				Name:  "csv",
				Usage: "Csv Bill",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "param-template",
						Usage: "Set input parameter from string(JSON)",
					},
					&cli.StringFlag{
						Name:  "param-template-file",
						Usage: "Set input parameter from file",
					},
					&cli.BoolFlag{
						Name:  "generate-skeleton",
						Usage: "Output skelton of parameter JSON",
					},
					&cli.BoolFlag{
						Name:  "no-header",
						Usage: "set output header flag",
					},
					&cli.StringFlag{
						Name:    "bill-output",
						Aliases: []string{"file"},
						Usage:   "set bill-detail output path",
					},
					&cli.Int64Flag{
						Name:  "bill-id",
						Usage: "set bill ID",
					},
				},
				ShellComplete: func(c *cli.Context) {

					if c.NArg() < 3 { // invalid args
						return
					}

					if err := checkConfigVersion(); err != nil {
						return
					}
					if err := applyConfigFromFile(c); err != nil {
						return
					}

					// c.Args() == arg1 arg2 arg3 -- [cur] [prev] [commandName]
					args := c.Args().Slice()
					commandName := args[c.NArg()-1]
					prev := args[c.NArg()-2]
					cur := args[c.NArg()-3]

					// set real args
					realArgs := args[0 : c.NArg()-3]

					// Validate global params
					command.GlobalOption.Validate(false)

					// set default output-type
					// when params have output-type option and have empty value
					var outputTypeHolder interface{} = csvParam
					if v, ok := outputTypeHolder.(command.OutputTypeHolder); ok {
						if v.GetOutputType() == "" {
							v.SetOutputType(command.GlobalOption.DefaultOutputType)
						}
					}

					// build command context
					ctx := command.NewContext(c, realArgs, csvParam)

					// Set option values
					if c.IsSet("param-template") {
						csvParam.ParamTemplate = c.String("param-template")
					}
					if c.IsSet("param-template-file") {
						csvParam.ParamTemplateFile = c.String("param-template-file")
					}
					if c.IsSet("generate-skeleton") {
						csvParam.GenerateSkeleton = c.Bool("generate-skeleton")
					}
					if c.IsSet("no-header") {
						csvParam.NoHeader = c.Bool("no-header")
					}
					if c.IsSet("bill-output") {
						csvParam.BillOutput = c.String("bill-output")
					}
					if c.IsSet("bill-id") {
						csvParam.BillId = c.Int64("bill-id")
					}

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completion.FlagNames(c, commandName)
								return
							} else {
								completion.BillCsvCompleteArgs(ctx, csvParam, cur, prev, commandName)
								return
							}
						}

						// cleanup flag name
						name := prev
						for {
							if !strings.HasPrefix(name, "-") {
								break
							}
							name = strings.Replace(name, "-", "", 1)
						}

						// flag is exists? , is BoolFlag?
						exists := false
						for _, flag := range c.App.Command(commandName).Flags {

							for _, n := range flag.Names() {
								if n == name {
									exists = true
									break
								}
							}

							if exists {
								if _, ok := flag.(*cli.BoolFlag); ok {
									if strings.HasPrefix(cur, "-") {
										completion.FlagNames(c, commandName)
										return
									} else {
										completion.BillCsvCompleteArgs(ctx, csvParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									completion.BillCsvCompleteFlags(ctx, csvParam, name, cur)
									return
								}
							}
						}
						// here, prev is wrong, so noop.
					} else {
						if strings.HasPrefix(cur, "-") {
							completion.FlagNames(c, commandName)
							return
						} else {
							completion.BillCsvCompleteArgs(ctx, csvParam, cur, prev, commandName)
							return
						}
					}
				},
				Action: func(c *cli.Context) error {

					if err := checkConfigVersion(); err != nil {
						return err
					}
					if err := applyConfigFromFile(c); err != nil {
						return err
					}

					csvParam.ParamTemplate = c.String("param-template")
					csvParam.ParamTemplateFile = c.String("param-template-file")
					strInput, err := command.GetParamTemplateValue(csvParam)
					if err != nil {
						return err
					}
					if strInput != "" {
						p := params.NewCsvBillParam()
						err := json.Unmarshal([]byte(strInput), p)
						if err != nil {
							return fmt.Errorf("Failed to parse JSON: %s", err)
						}
						mergo.Merge(csvParam, p, mergo.WithOverride)
					}

					// Set option values
					if c.IsSet("param-template") {
						csvParam.ParamTemplate = c.String("param-template")
					}
					if c.IsSet("param-template-file") {
						csvParam.ParamTemplateFile = c.String("param-template-file")
					}
					if c.IsSet("generate-skeleton") {
						csvParam.GenerateSkeleton = c.Bool("generate-skeleton")
					}
					if c.IsSet("no-header") {
						csvParam.NoHeader = c.Bool("no-header")
					}
					if c.IsSet("bill-output") {
						csvParam.BillOutput = c.String("bill-output")
					}
					if c.IsSet("bill-id") {
						csvParam.BillId = c.Int64("bill-id")
					}

					// Validate global params
					if errors := command.GlobalOption.Validate(false); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					var outputTypeHolder interface{} = csvParam
					if v, ok := outputTypeHolder.(command.OutputTypeHolder); ok {
						if v.GetOutputType() == "" {
							v.SetOutputType(command.GlobalOption.DefaultOutputType)
						}
					}

					// Generate skeleton
					if csvParam.GenerateSkeleton {
						csvParam.GenerateSkeleton = false
						csvParam.FillValueToSkeleton()
						d, err := json.MarshalIndent(csvParam, "", "\t")
						if err != nil {
							return fmt.Errorf("Failed to Marshal JSON: %s", err)
						}
						fmt.Fprintln(command.GlobalOption.Out, string(d))
						return nil
					}

					// Validate specific for each command params
					if errors := csvParam.Validate(); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := command.NewContext(c, c.Args().Slice(), csvParam)

					// Run command with params
					return funcs.BillCsv(ctx, csvParam)

				},
			},
			{
				Name:    "list",
				Aliases: []string{"ls", "find"},
				Usage:   "List Bill (default)",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:  "year",
						Usage: "set year",
					},
					&cli.IntFlag{
						Name:  "month",
						Usage: "set month",
					},
					&cli.StringFlag{
						Name:  "param-template",
						Usage: "Set input parameter from string(JSON)",
					},
					&cli.StringFlag{
						Name:  "param-template-file",
						Usage: "Set input parameter from file",
					},
					&cli.BoolFlag{
						Name:  "generate-skeleton",
						Usage: "Output skelton of parameter JSON",
					},
					&cli.StringFlag{
						Name:    "output-type",
						Aliases: []string{"out"},
						Usage:   "Output type [table/json/csv/tsv]",
					},
					&cli.StringSliceFlag{
						Name:    "column",
						Aliases: []string{"col"},
						Usage:   "Output columns(using when '--output-type' is in [csv/tsv] only)",
					},
					&cli.BoolFlag{
						Name:    "quiet",
						Aliases: []string{"q"},
						Usage:   "Only display IDs",
					},
					&cli.StringFlag{
						Name:    "format",
						Aliases: []string{"fmt"},
						Usage:   "Output format(see text/template package document for detail)",
					},
					&cli.StringFlag{
						Name:  "format-file",
						Usage: "Output format from file(see text/template package document for detail)",
					},
					&cli.StringFlag{
						Name:  "query",
						Usage: "JMESPath query(using when '--output-type' is json only)",
					},
				},
				ShellComplete: func(c *cli.Context) {

					if c.NArg() < 3 { // invalid args
						return
					}

					if err := checkConfigVersion(); err != nil {
						return
					}
					if err := applyConfigFromFile(c); err != nil {
						return
					}

					// c.Args() == arg1 arg2 arg3 -- [cur] [prev] [commandName]
					args := c.Args().Slice()
					commandName := args[c.NArg()-1]
					prev := args[c.NArg()-2]
					cur := args[c.NArg()-3]

					// set real args
					realArgs := args[0 : c.NArg()-3]

					// Validate global params
					command.GlobalOption.Validate(false)

					// set default output-type
					// when params have output-type option and have empty value
					var outputTypeHolder interface{} = listParam
					if v, ok := outputTypeHolder.(command.OutputTypeHolder); ok {
						if v.GetOutputType() == "" {
							v.SetOutputType(command.GlobalOption.DefaultOutputType)
						}
					}

					// build command context
					ctx := command.NewContext(c, realArgs, listParam)

					// Set option values
					if c.IsSet("year") {
						listParam.Year = c.Int("year")
					}
					if c.IsSet("month") {
						listParam.Month = c.Int("month")
					}
					if c.IsSet("param-template") {
						listParam.ParamTemplate = c.String("param-template")
					}
					if c.IsSet("param-template-file") {
						listParam.ParamTemplateFile = c.String("param-template-file")
					}
					if c.IsSet("generate-skeleton") {
						listParam.GenerateSkeleton = c.Bool("generate-skeleton")
					}
					if c.IsSet("output-type") {
						listParam.OutputType = c.String("output-type")
					}
					if c.IsSet("column") {
						listParam.Column = c.StringSlice("column")
					}
					if c.IsSet("quiet") {
						listParam.Quiet = c.Bool("quiet")
					}
					if c.IsSet("format") {
						listParam.Format = c.String("format")
					}
					if c.IsSet("format-file") {
						listParam.FormatFile = c.String("format-file")
					}
					if c.IsSet("query") {
						listParam.Query = c.String("query")
					}

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completion.FlagNames(c, commandName)
								return
							} else {
								completion.BillListCompleteArgs(ctx, listParam, cur, prev, commandName)
								return
							}
						}

						// cleanup flag name
						name := prev
						for {
							if !strings.HasPrefix(name, "-") {
								break
							}
							name = strings.Replace(name, "-", "", 1)
						}

						// flag is exists? , is BoolFlag?
						exists := false
						for _, flag := range c.App.Command(commandName).Flags {

							for _, n := range flag.Names() {
								if n == name {
									exists = true
									break
								}
							}

							if exists {
								if _, ok := flag.(*cli.BoolFlag); ok {
									if strings.HasPrefix(cur, "-") {
										completion.FlagNames(c, commandName)
										return
									} else {
										completion.BillListCompleteArgs(ctx, listParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									completion.BillListCompleteFlags(ctx, listParam, name, cur)
									return
								}
							}
						}
						// here, prev is wrong, so noop.
					} else {
						if strings.HasPrefix(cur, "-") {
							completion.FlagNames(c, commandName)
							return
						} else {
							completion.BillListCompleteArgs(ctx, listParam, cur, prev, commandName)
							return
						}
					}
				},
				Action: func(c *cli.Context) error {

					if err := checkConfigVersion(); err != nil {
						return err
					}
					if err := applyConfigFromFile(c); err != nil {
						return err
					}

					listParam.ParamTemplate = c.String("param-template")
					listParam.ParamTemplateFile = c.String("param-template-file")
					strInput, err := command.GetParamTemplateValue(listParam)
					if err != nil {
						return err
					}
					if strInput != "" {
						p := params.NewListBillParam()
						err := json.Unmarshal([]byte(strInput), p)
						if err != nil {
							return fmt.Errorf("Failed to parse JSON: %s", err)
						}
						mergo.Merge(listParam, p, mergo.WithOverride)
					}

					// Set option values
					if c.IsSet("year") {
						listParam.Year = c.Int("year")
					}
					if c.IsSet("month") {
						listParam.Month = c.Int("month")
					}
					if c.IsSet("param-template") {
						listParam.ParamTemplate = c.String("param-template")
					}
					if c.IsSet("param-template-file") {
						listParam.ParamTemplateFile = c.String("param-template-file")
					}
					if c.IsSet("generate-skeleton") {
						listParam.GenerateSkeleton = c.Bool("generate-skeleton")
					}
					if c.IsSet("output-type") {
						listParam.OutputType = c.String("output-type")
					}
					if c.IsSet("column") {
						listParam.Column = c.StringSlice("column")
					}
					if c.IsSet("quiet") {
						listParam.Quiet = c.Bool("quiet")
					}
					if c.IsSet("format") {
						listParam.Format = c.String("format")
					}
					if c.IsSet("format-file") {
						listParam.FormatFile = c.String("format-file")
					}
					if c.IsSet("query") {
						listParam.Query = c.String("query")
					}

					// Validate global params
					if errors := command.GlobalOption.Validate(false); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					var outputTypeHolder interface{} = listParam
					if v, ok := outputTypeHolder.(command.OutputTypeHolder); ok {
						if v.GetOutputType() == "" {
							v.SetOutputType(command.GlobalOption.DefaultOutputType)
						}
					}

					// Generate skeleton
					if listParam.GenerateSkeleton {
						listParam.GenerateSkeleton = false
						listParam.FillValueToSkeleton()
						d, err := json.MarshalIndent(listParam, "", "\t")
						if err != nil {
							return fmt.Errorf("Failed to Marshal JSON: %s", err)
						}
						fmt.Fprintln(command.GlobalOption.Out, string(d))
						return nil
					}

					// Validate specific for each command params
					if errors := listParam.Validate(); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := command.NewContext(c, c.Args().Slice(), listParam)

					// Run command with params
					return funcs.BillList(ctx, listParam)

				},
			},
		},
	}

	// build Category-Resource mapping
	AppendResourceCategoryMap("bill", &schema.Category{
		Key:         "billing",
		DisplayName: "Billing",
		Order:       70,
	})

	// build Category-Command mapping

	AppendCommandCategoryMap("bill", "csv", &schema.Category{
		Key:         "default",
		DisplayName: "",
		Order:       2147483647,
	})
	AppendCommandCategoryMap("bill", "list", &schema.Category{
		Key:         "default",
		DisplayName: "",
		Order:       2147483647,
	})

	// build Category-Param mapping

	AppendFlagCategoryMap("bill", "csv", "bill-id", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	AppendFlagCategoryMap("bill", "csv", "bill-output", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("bill", "csv", "generate-skeleton", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("bill", "csv", "no-header", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("bill", "csv", "param-template", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("bill", "csv", "param-template-file", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("bill", "list", "column", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("bill", "list", "format", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("bill", "list", "format-file", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("bill", "list", "generate-skeleton", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("bill", "list", "month", &schema.Category{
		Key:         "filter",
		DisplayName: "Filter options",
		Order:       2147483587,
	})
	AppendFlagCategoryMap("bill", "list", "output-type", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("bill", "list", "param-template", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("bill", "list", "param-template-file", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("bill", "list", "query", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("bill", "list", "quiet", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("bill", "list", "year", &schema.Category{
		Key:         "filter",
		DisplayName: "Filter options",
		Order:       2147483587,
	})

	// append command to GlobalContext
	Commands = append(Commands, cliCommand)
}
