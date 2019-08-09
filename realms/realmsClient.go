package realms

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/Tnze/go-mc/bot"
)

var (
	serverList    realmsServers
	sInfo         serverInfo
	sessionCookie string

	client http.Client = http.Client{}
)

func SetCookie(uuid, name, accessToken string) {
	sessionCookie = "sid=token:" + accessToken + ":" + uuid + ",user=" + name + ",version=" + bot.MCVersion
}

type realmsServers struct {
	Servers []struct {
		ID                   int         `json:"id"`
		RemoteSubscriptionID string      `json:"remoteSubscriptionId"`
		Owner                string      `json:"owner"`
		OwnerUUID            string      `json:"ownerUUID"`
		Name                 string      `json:"name"`
		Motd                 string      `json:"motd"`
		DefaultPermission    string      `json:"defaultPermission"`
		State                string      `json:"state"`
		DaysLeft             int         `json:"daysLeft"`
		Expired              bool        `json:"expired"`
		ExpiredTrial         bool        `json:"expiredTrial"`
		GracePeriod          bool        `json:"gracePeriod"`
		WorldType            string      `json:"worldType"`
		Players              interface{} `json:"players"`
		MaxPlayers           int         `json:"maxPlayers"`
		MinigameName         interface{} `json:"minigameName"`
		MinigameID           interface{} `json:"minigameId"`
		MinigameImage        interface{} `json:"minigameImage"`
		ActiveSlot           int         `json:"activeSlot"`
		Slots                interface{} `json:"slots"`
		Member               bool        `json:"member"`
		ClubID               interface{} `json:"clubId"`
	} `json:"servers"`
}

type serverInfo struct {
	Address          string `json:"address"`
	ResourcePackURL  string `json:"resourcePackUrl"`
	ResourcePackHash string `json:"resourcePackHash"`
}

func ListWorlds(realmsName string) (RealmsID int, err error) {

	url := "https://pc.realms.minecraft.net/worlds"

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-agent", "go-mc")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", sessionCookie)

	if err != nil {
		err = fmt.Errorf("make request error: %v", err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("make request error: %v", err)
		return
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(body, &serverList)
	if err != nil {
		err = fmt.Errorf("unmarshal json data fail: %v", err)
		return
	}

	if strings.EqualFold("", realmsName) {
		for _, v := range serverList.Servers {
			fmt.Print("Name:", v.Name)
			fmt.Println("\tID:", v.ID)
		}
		return
	}

	for _, v := range serverList.Servers {

		if strings.EqualFold(v.Name, realmsName) {
			fmt.Print("查找到Realms名称:", v.Name)
			fmt.Println("\tID:", v.ID)
			RealmsID = v.ID
			return
		}

	}
	panic("找不到Realms Name.请检查是否正确.")
}

func Join(serverID int) (address string, err error) {

	url := "https://pc.realms.minecraft.net/worlds/v1/" + strconv.Itoa(serverID) + "/join/pc"
	// fmt.Println("Join URL:", url)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-agent", "go-mc")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", sessionCookie)
	if err != nil {
		err = fmt.Errorf("make request error: %v", err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("make request error: %v", err)
		return
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != 200 {
		err = fmt.Errorf("Failed to get realms server address: %v", string(body))
		return
	}

	// fmt.Println(string(body))
	// fmt.Println(res)

	err = json.Unmarshal(body, &sInfo)
	if err != nil {
		err = fmt.Errorf("Unmarshal json data fail: %v", err)
		return
	}

	address = sInfo.Address
	return
}