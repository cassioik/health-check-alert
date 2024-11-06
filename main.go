package main

import (
    "bytes"
    "crypto/tls"
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "strconv"
    "time"
)

var (
    urlToPing        = os.Getenv("URL_TO_PING")
    discordWebhookURL = os.Getenv("DISCORD_WEBHOOK_URL")
    pingInterval     = getEnvAsInt("PING_INTERVAL", 5) // Default to 5 minutes if not set
)

var httpClient = &http.Client{
    Transport: &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    },
}

func main() {
    if urlToPing == "" || discordWebhookURL == "" {
        fmt.Println(time.Now().Format(time.RFC3339), "URL_TO_PING and DISCORD_WEBHOOK_URL environment variables must be set")
        return
    }

    for {
        resp, err := httpClient.Get(urlToPing)
        if err != nil || resp.StatusCode != http.StatusOK {
            sendToDiscord(fmt.Sprintf("Ping failed: %v", err))
        } else {
            fmt.Println(time.Now().Format(time.RFC3339), "Ping successful")
        }
        if resp != nil {
            resp.Body.Close()
        }
        time.Sleep(time.Duration(pingInterval) * time.Minute)
    }
}

func sendToDiscord(message string) {
    payload := map[string]string{"content": message}
    payloadBytes, err := json.Marshal(payload)
    if err != nil {
        fmt.Printf(time.Now().Format(time.RFC3339), "Failed to marshal payload: %v\n", err)
        return
    }

    resp, err := httpClient.Post(discordWebhookURL, "application/json", bytes.NewBuffer(payloadBytes))
    if err != nil {
        fmt.Printf(time.Now().Format(time.RFC3339), "Failed to send to Discord: %v\n", err)
        return
    }
    defer resp.Body.Close()
    fmt.Println(time.Now().Format(time.RFC3339), "Message sent to Discord")
}

func getEnvAsInt(name string, defaultValue int) int {
    valueStr := os.Getenv(name)
    if valueStr == "" {
        return defaultValue
    }
    value, err := strconv.Atoi(valueStr)
    if err != nil {
        fmt.Printf("Invalid value for %s: %v. Using default value: %d\n", name, err, defaultValue)
        return defaultValue
    }
    return value
}