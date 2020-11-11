package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/aidtechnology/affinityctl/client"
	"github.com/spf13/viper"
	"go.bryk.io/x/cli"
)

type entry struct {
	ID    string                 `json:"id"`
	Name  string                 `json:"name"`
	Email string                 `json:"email"`
	Doc   map[string]interface{} `json:"document"`
}

func dirExist(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func getPIN(sdk *client.SDK) (string, error) {
	pin, err := cli.ReadSecure("Enter PIN:\n")
	if err != nil {
		return "", err
	}
	confirmation, err := cli.ReadSecure("Confirm PIN:\n")
	if err != nil {
		return "", err
	}
	if !bytes.Equal(pin, confirmation) {
		return "", errors.New("different values provided")
	}
	if len(pin) == 0 {
		p, err := sdk.DID.GetMaterial()
		if err != nil {
			return "", err
		}
		fmt.Printf("Random PIN generated: %s\n", p)
		pin = []byte(p)
	}
	return string(pin), nil
}

func store(did string, name string, email string, doc []byte) error {
	content := make(map[string]interface{})
	if err := json.Unmarshal(doc, &content); err != nil {
		return err
	}
	rec := &entry{
		ID:    did,
		Name:  name,
		Email: email,
		Doc:   content,
	}
	js, _ := json.MarshalIndent(rec, "", "  ")
	home := viper.GetString("client.home")
	fn := fmt.Sprintf("%s.json", name)
	return ioutil.WriteFile(path.Join(home, fn), js, 0400)
}

func details(name string) (*entry, error) {
	home := viper.GetString("client.home")
	contents, err := ioutil.ReadFile(filepath.Clean(filepath.Join(home, name+".json")))
	if err != nil {
		return nil, fmt.Errorf("no DID with name: %s", name)
	}
	rec := new(entry)
	if err = json.Unmarshal(contents, rec); err != nil {
		return nil, errors.New("failed to decode available information")
	}
	return rec, nil
}

func list() []*entry {
	var list []*entry
	_ = filepath.Walk(viper.GetString("client.home"), func(f string, info os.FileInfo, err error) error {
		if err != nil || info == nil || info.IsDir() {
			return nil
		}
		if filepath.Ext(f) != ".json" {
			return nil
		}
		contents, err := ioutil.ReadFile(filepath.Clean(f))
		if err != nil {
			return nil
		}
		rec := new(entry)
		if err = json.Unmarshal(contents, rec); err != nil {
			return nil
		}
		list = append(list, rec)
		return nil
	})
	return list
}

func sdkClient() (*client.SDK, error) {
	opts := client.DefaultOptions()
	opts.Key = viper.GetString("client.key")
	return client.New(opts)
}
