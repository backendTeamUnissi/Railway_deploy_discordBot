// config.go
package config

import (
    "fmt"
    // "log"
    "os"
	"strings"

    // "github.com/joho/godotenv"
)

type Config struct {
    DiscordToken          string
    DiscordTextChannelID  string
    DiscordVoiceChannelID string
    FirestoreCredentials  string
}

func LoadConfig() (*Config, error) {
    // .env ファイルの読み込み
    // err := godotenv.Load()
    // if err != nil {
    //     log.Println("No .env file found, reading environment variables")
    // }

    cfg := &Config{
        DiscordToken:          os.Getenv("DISCORDTOKEN"),
        DiscordTextChannelID:  os.Getenv("DISCORDTEXTCHANNELID"),
        DiscordVoiceChannelID: os.Getenv("DISCORDVOICECHANNELID"),
        FirestoreCredentials:  os.Getenv("FIRESTORE_CREDENTIALS_FILE"),
    }

    // 環境変数のバリデーション
	var missingVars []string

	if cfg.DiscordToken == "" {
		missingVars = append(missingVars, "DISCORD_TOKEN")
	}
	if cfg.DiscordTextChannelID == "" {
		missingVars = append(missingVars, "DISCORD_TEXT_CHANNEL_ID")
	}
	if cfg.DiscordVoiceChannelID == "" {
		missingVars = append(missingVars, "DISCORD_VOICE_CHANNEL_ID")
	}
	if cfg.FirestoreCredentials == "" {
		missingVars = append(missingVars, "FIRESTORE_CREDENTIALS")
	}
	
	if len(missingVars) > 0 {
		return nil, fmt.Errorf("以下の必須環境変数が設定されていません: %s", strings.Join(missingVars, ", "))
	}

    return cfg, nil
}
