package teamspeak

import (
	"dynts-bann3r/src/config"
	"dynts-bann3r/src/utils"
	"log"
	"strconv"
	"strings"

	"github.com/multiplay/go-ts3"
)

type Teamspeak struct {
	Client        *ts3.Client
	AdminGroupIds []int
}

var state = make(map[string]string)

func New(connection config.Connection, adminGroupIds []int) Teamspeak {
	client, err := ts3.NewClient(connection.Host + ":" + strconv.Itoa(connection.Port))

	if err != nil {
		log.Fatal(err)
	}

	if err := client.Login(connection.User, connection.Password); err != nil {
		log.Fatal(err)
	}

	if err := client.Use(connection.ServerId); err != nil {
		log.Fatal(err)
	}

	t := Teamspeak{Client: client, AdminGroupIds: adminGroupIds}
	return t
}

func (t Teamspeak) Refresh() {
	if err := t.refreshServerInfo(); err != nil {
		log.Printf("[ERROR] Server info could not be refreshed: %v \n", err)
	}

	if err := t.refreshAdminsOnline(); err != nil {
		log.Printf("[ERROR] Admins online count could not be refreshed: %v \n", err)
	}
}

func (t Teamspeak) State() map[string]string {
	return state
}

func (t Teamspeak) refreshServerInfo() error {
	if serverInfo, err := t.Client.Server.Info(); err != nil {
		return err
	} else {
		state["MAX_CLIENTS"] = strconv.Itoa(serverInfo.MaxClients)
		state["REAL_CLIENTS_ONLINE"] = strconv.Itoa(serverInfo.ClientsOnline - serverInfo.QueryClientsOnline)
		state["CLIENTS_ONLINE"] = strconv.Itoa(serverInfo.ClientsOnline)
		state["QUERY_CLIENTS_ONLINE"] = strconv.Itoa(serverInfo.QueryClientsOnline)

		state["SERVER_NAME"] = serverInfo.Name
		state["SERVER_PORT"] = strconv.Itoa(serverInfo.Port)

		state["TIME_HH"] = utils.GetHours()
		state["TIME_MM"] = utils.GetMinutes()
		state["TIME_SS"] = utils.GetSeconds()

		return nil
	}
}

func (t Teamspeak) refreshAdminsOnline() error {
	onlineClients, err := t.getOnlineClients()

	if err != nil {
		return err
	}

	onlineClientsInGroups := 0

	for _, gid := range t.AdminGroupIds {
		clientsInGroup, err := t.getClientIdsInGroup(gid)
		if err != nil {
			return err
		}

		onlineClientsInGroup, err := countOnlineClientsInGroup(clientsInGroup, onlineClients)
		if err != nil {
			return err
		}

		onlineClientsInGroups += onlineClientsInGroup
	}

	state["ADMIN_CLIENTS_ONLINE"] = strconv.Itoa(onlineClientsInGroups)

	return nil
}

func (t Teamspeak) getOnlineClients() ([]*ts3.OnlineClient, error) {
	if clients, err := t.Client.Server.ClientList(); err != nil {
		return nil, err
	} else {
		return clients, nil
	}
}

func countOnlineClientsInGroup(clientsInGroup []int, onlineClients []*ts3.OnlineClient) (int, error) {
	onlineClientsInGroup := 0

	for _, clientIdInGroup := range clientsInGroup {
		for _, onlineClient := range onlineClients {
			if onlineClient.DatabaseID == clientIdInGroup {
				onlineClientsInGroup++
			}
		}
	}

	return onlineClientsInGroup, nil
}

func (t Teamspeak) getClientIdsInGroup(gid int) ([]int, error) {
	cmd := ts3.NewCmd("servergroupclientlist")
	arg := ts3.NewArg("sgid", gid)
	cmd.WithArgs(arg)

	var clientIds []int

	out, err := t.Client.ExecCmd(cmd)

	if err == nil && len(out) > 0 {
		strIds := strings.ReplaceAll(out[0], "cldbid=", "")
		ids := strings.Split(strIds, "|")

		for _, strId := range ids {
			id, err := strconv.Atoi(strId)

			if err != nil {
				return nil, err
			}

			clientIds = append(clientIds, id)
		}

		return clientIds, nil
	}

	return nil, err
}
