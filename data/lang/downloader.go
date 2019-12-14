package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

//go:generate go run $GOFILE
//go:generate go fmt ./...
func main() {
	// from {https://launchermeta.mojang.com/mc/game/version_manifest.json}.assetIndex.url
	versionURL := "https://launchermeta.mojang.com/v1/packages/e8016c24200e6dd1b9001ec5204d4332bae24c38/1.15.json"
	log.Print("start generating lang packages")

	resp, err := http.Get(versionURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var list struct {
		Objects map[string]struct {
			Hash string
			Size int64
		}
	}

	err = json.NewDecoder(resp.Body).Decode(&list)
	if err != nil {
		log.Fatal(err)
	}

	for i, v := range list.Objects {
		if strings.HasPrefix(i, "minecraft/lang/") {
			name := i[len("minecraft/lang/") : len(i)-len(".json")]
			lang(name, v.Hash, v.Size)
		}
	}
}

func lang(name, hash string, size int64) {
	log.Print("generating ", name, " package")

	//download language
	LangURL := "http://resources.download.minecraft.net/" + hash[:2] + "/" + hash
	resp, err := http.Get(LangURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// read language
	var LangMap map[string]string
	err = json.NewDecoder(resp.Body).Decode(&LangMap)
	if err != nil {
		log.Fatal("unmarshal json fail: ", err)
	}
	trans(LangMap)

	pName := strings.ReplaceAll(name, "_", "-")

	// mkdir
	err = os.Mkdir(pName, 0777)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	f, err := os.OpenFile(filepath.Join(pName, name+".go"), os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// write code
	_, err = fmt.Fprintf(f,
		"package %s\n\n"+
			"import \"github.com/Tnze/go-mc/chat\"\n\n"+
			"func init() { chat.SetLanguage(Map) }\n\n"+
			"var Map = %#v\n",
		name, LangMap)

	if err != nil {
		log.Fatal(err)
	}
}

var javaN = regexp.MustCompile(`%[0-9]\$s`)

// Java use %2$s to refer to the second arg, but Golang use %2s, so we need this
func trans(m map[string]string) {
	//replace "%[0-9]\$s" with "%[0-9]s"
	for i := range m {
		c := m[i]
		if javaN.MatchString(c) {
			m[i] = javaN.ReplaceAllStringFunc(c, func(s string) string {
				var index int
				_, err := fmt.Sscanf(s, "%%%d$s", &index)
				if err != nil {
					log.Fatal(err)
				}

				return fmt.Sprintf("%%%ds", index)
			})
		}
	}
}
