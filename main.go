package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/atotto/clipboard"
	"github.com/keybase/go-keychain"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

var accountName string
var otpUri string
var add bool
var delete bool

func init() {
	flag.StringVar(&accountName, "account", "", "Name to be used for account")
	flag.StringVar(&otpUri, "url", "", "Value to add")
	flag.BoolVar(&add, "add", false, "Add mode")
	flag.BoolVar(&delete, "delete", false, "delete mode")
	flag.Parse()
	if accountName == "" {
		fmt.Println("account can't be empty")
		os.Exit(1)
	}
	if add {
		if accountName == "" || otpUri == "" {
			fmt.Println("account and value can't be empty")
			os.Exit(1)
		}
	}
}
func main() {
	if err := execute(); err != nil {
		os.Exit(1)
	}
}
func execute() error {
	serviceName := "TOTPGenerator"
	label := accountName
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService(serviceName)
	item.SetAccount(accountName)
	item.SetAccessGroup("com.zero.zero")
	if add {
		item.SetLabel(label)
		item.SetData([]byte(otpUri))
		item.SetSynchronizable(keychain.SynchronizableNo)
		item.SetAccessible(keychain.AccessibleAfterFirstUnlockThisDeviceOnly)
		err := keychain.AddItem(item)
		if err != nil {
			fmt.Println("Failed to add:", err)
			return err
		}
		fmt.Println("Added:", serviceName, accountName)
	} else if delete {
		err := keychain.DeleteGenericPasswordItem(serviceName, accountName)
		if err != nil {
			fmt.Println("Failed to delete:", err)
		}
		fmt.Println("Deleted:", serviceName, accountName)
		return nil
	}

	item.SetMatchLimit(keychain.MatchLimitOne)
	item.SetReturnData(true)
	results, err := keychain.QueryItem(item)
	if err != nil {
		fmt.Println("Failed to fetch item:", err)
		return err
	}
	for _, result := range results {
		key, err := otp.NewKeyFromURL(string(result.Data))
		if err != nil {
			fmt.Println("Failed to create otp:", err)
			return err
		}
		passCode, err := totp.GenerateCodeCustom(key.Secret(), time.Now(), totp.ValidateOpts{
			Period:    uint(key.Period()),
			Digits:    key.Digits(),
			Algorithm: key.Algorithm(),
		})
		if err != nil {
			fmt.Println("Failed to generate code:", err)
			return err
		}
		clipboard.WriteAll(passCode)
		fmt.Println("PassCode copied to clipboard")
	}
	return nil
}
