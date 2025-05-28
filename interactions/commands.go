package interactions

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
)

type Application struct {
	Id   string
	Auth string
}

func (app *Application) GetCommands() (string, error) {
	r, err := http.NewRequest("GET", "https://discord.com/api/v10/applications/"+app.Id+"/commands", bytes.NewBuffer([]byte(``)))
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	r.Header.Add("Authorization", app.Auth)
	res, err := client.Do(r)
	if err != nil {
		return "", err
	} else if res.StatusCode != 200 {
		body, _ := io.ReadAll(res.Body)
		slog.Info("Getting global commands has encountered error", "Status", res.Status, "Return_body", body)
	}

	body, _ := io.ReadAll(res.Body)

	defer res.Body.Close()
	return string(body[0 : len(body)-1]), nil
}
