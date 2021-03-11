package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	var telegramBotToken, telegramParseMode, ciProjectUrl, ciPipelineId, ciCommitSlug, text string
	var telegramChatId int

	app := &cli.App{
		Name:    "notifyer",
		Version: "v3.0.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "telegram_bot_token",
				DefaultText: "YOUR_TELEGRAM_BOT_TOKEN",
				Aliases:     []string{"tbt"},
				Usage:       "TELEGRAM_BOT_TOKEN",
				EnvVars:     []string{"NOTIFYER_TELEGRAM_BOT_TOKEN"},
				Destination: &telegramBotToken,
			},
			&cli.StringFlag{
				Name:        "ci_project_url",
				Aliases:     []string{"cpu"},
				Usage:       "CI_PROJECT_URL - GitLab CI",
				EnvVars:     []string{"CI_PROJECT_URL"},
				Destination: &ciProjectUrl,
			},
			&cli.StringFlag{
				Name:        "ci_pipeline_id",
				Aliases:     []string{"cpi"},
				Usage:       "CI_PIPELINE_ID - GitLab CI",
				EnvVars:     []string{"CI_PIPELINE_ID"},
				Destination: &ciPipelineId,
			},
			&cli.StringFlag{
				Name:        "ci_commit_slug",
				Aliases:     []string{"gcs"},
				Usage:       "CI_COMMIT_REF_SLUG - GitLab CI",
				EnvVars:     []string{"CI_COMMIT_REF_SLUG"},
				Destination: &ciCommitSlug,
			},
			&cli.IntFlag{
				Name:        "telegram_chat_id",
				Aliases:     []string{"tci"},
				Usage:       "TELEGRAM_CHAT_ID - Telegram",
				EnvVars:     []string{"NOTIFYER_TELEGRAM_CHAT_ID"},
				Destination: &telegramChatId,
			},
			&cli.StringFlag{
				Name:        "telegram_parse_mode",
				Aliases:     []string{"tpm"},
				Usage:       "TELEGRAM_PARSE_MODE - Telegram",
				EnvVars:     []string{"NOTIFYER_TELEGRAM_PARSE_MODE"},
				DefaultText: "HTML",
				Destination: &telegramParseMode,
			},
			&cli.StringFlag{
				Name:        "text",
				Aliases:     []string{"t"},
				Usage:       "TEXT",
				EnvVars:     []string{"NOTIFYER_TEXT"},
				Destination: &text,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "release",
				Aliases: []string{"r"},
				Usage:   "Release [ready]",
				Subcommands: []*cli.Command{
					{
						Name:  "ready",
						Usage: "release ready",
						Action: func(c *cli.Context) error {
							fmt.Println("[Notifyer] Release ready")
							sendTextToTelegramChat(telegramChatId, text, telegramBotToken, telegramParseMode)
							return nil
						},
					},
				},
			},
			{
				Name:    "deploy",
				Aliases: []string{"d"},
				Usage:   "Deploy [done]",
				Subcommands: []*cli.Command{
					{
						Name:  "done",
						Usage: "deploy done",
						Action: func(c *cli.Context) error {
							fmt.Println("[Notifyer] Deploy done")
							sendTextToTelegramChat(telegramChatId, text, telegramBotToken, telegramParseMode)
							return nil
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
