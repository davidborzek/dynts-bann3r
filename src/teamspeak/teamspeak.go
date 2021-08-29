package teamspeak

import (
	"dynts-bann3r/src/config"
	"log"
	"strconv"
	"strings"

	"github.com/multiplay/go-ts3"
)

func Login(connection config.Connection) *ts3.Client {
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

	return client
}

func CountOnlineClients(client *ts3.Client) (string, error) {
	if serverInfo, err := client.Server.Info(); err != nil {
		return "", err
	} else {
		return strconv.Itoa(serverInfo.ClientsOnline - serverInfo.QueryClientsOnline), nil
	}
}

func GetMaxClients(client *ts3.Client) (string, error) {
	if serverInfo, err := client.Server.Info(); err != nil {
		return "", err
	} else {
		return strconv.Itoa(serverInfo.MaxClients), nil
	}
}

func GetServerName(client *ts3.Client) (string, error) {
	if serverInfo, err := client.Server.Info(); err != nil {
		return "", err
	} else {
		return serverInfo.Name, nil
	}
}

func GetServerPort(client *ts3.Client) (string, error) {
	if serverInfo, err := client.Server.Info(); err != nil {
		return "", err
	} else {
		return strconv.Itoa(serverInfo.Port), nil
	}
}

func CountOnlineClientsInGroups(client *ts3.Client, gIds []string) (string, error) {
	onlineClients, err := getOnlineClients(client)

	if err != nil {
		return "", err
	}

	onlineClientsInGroups := 0

	for _, gid := range gIds {
		onlineClientsInGroup, err := countOnlineClientsInGroup(client, gid, onlineClients)

		if err != nil {
			return "", err
		}

		onlineClientsInGroups += onlineClientsInGroup
	}

	return strconv.Itoa(onlineClientsInGroups), nil
}

func getOnlineClients(client *ts3.Client) ([]*ts3.OnlineClient, error) {
	if clients, err := client.Server.ClientList(); err != nil {
		return nil, err
	} else {
		return clients, nil
	}
}

func countOnlineClientsInGroup(client *ts3.Client, gid string, onlineClients []*ts3.OnlineClient) (int, error) {
	intGid, err := strconv.Atoi(gid)

	if err != nil {
		return 0, err
	}

	clientsInGroup, err := getClientIdsInGroup(client, intGid)

	if err != nil {
		return 0, err
	}

	onlineClientsInGroup := 0

	for _, clientIdInGroup := range clientsInGroup {
		for _, onlineClient := range onlineClients {
			if err != nil {
				return 0, err
			}

			if onlineClient.DatabaseID == clientIdInGroup {
				onlineClientsInGroup++
			}
		}
	}

	return onlineClientsInGroup, nil

}

func getClientIdsInGroup(client *ts3.Client, gid int) ([]int, error) {
	cmd := ts3.NewCmd("servergroupclientlist")
	arg := ts3.NewArg("sgid", gid)
	cmd.WithArgs(arg)

	var clientIds []int

	out, err := client.ExecCmd(cmd)

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
