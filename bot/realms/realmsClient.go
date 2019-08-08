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

func New(uuid, name, accessToken string) {
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

func ListWorlds(realmsName string) int {

	url := "https://pc.realms.minecraft.net/worlds"

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Cookie", sessionCookie)
	req.Header.Add("cache-control", "no-cache")

	if err != nil {
		err = fmt.Errorf("make request error: %v", err)

	}
	res, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("make request error: %v", err)
		return -1
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(body, &serverList)
	if err != nil {
		err = fmt.Errorf("unmarshal json data fail: %v", err)
		return -1
	}
	if strings.EqualFold("", realmsName) {
		for _, v := range serverList.Servers {
			fmt.Print("Name:", v.Name)
			fmt.Println("\tID:", v.ID)
		}
		return 1
	}

	for _, v := range serverList.Servers {

		if strings.EqualFold(v.Name, realmsName) {
			fmt.Print("查找到Realms名称:", v.Name)
			fmt.Println("\tID:", v.ID)
			return v.ID
		}

	}
	panic("找不到Realms Name对应Realms ID.请检查Realms Name是否正确.")
}

func Join(serverID int) string {

	url := "https://pc.realms.minecraft.net/worlds/v1/" + strconv.Itoa(serverID) + "/join/pc"
	//fmt.Println("Join URL:", url)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Cookie", sessionCookie)
	if err != nil {
		err = fmt.Errorf("make request error: %v", err)
		return ""
	}
	res, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("make request error: %v", err)
		return ""
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	if strings.EqualFold(string(body), "Retry again later") {
		panic("Failed to get Realms Server ID: Retry again later")
	}

	fmt.Println(string(body))

	err = json.Unmarshal(body, &sInfo)
	if err != nil {
		err = fmt.Errorf("Unmarshal json data fail: %v", err)
		return ""
	}

	return sInfo.Address
}
