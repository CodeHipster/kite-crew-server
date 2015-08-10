package data

import (
	"fmt"

	"github.com/thijsoostdam/kitecrewserver/model"
)

var currentInviteId int
var invites model.Invites

// Give us some seed data
func init() {
	StoreInvite(model.Invite{InviterUsername: "Kees", Code: "12345a"})
	StoreInvite(model.Invite{InviterUsername: "Gertrude", Code: "12345b"})
}

func GetInvites() model.Invites {
	return invites
}

func FindInvite(code string) *model.Invite {
	for _, i := range invites {
		if i.Code == code {
			return &i
		}
	}
	// return nil if not found
	return nil
}

func StoreInvite(i model.Invite) error {
	if FindInvite(i.Code) != nil {
		fmt.Errorf("Invite already exists with code :%s", i.Code)
	}
	currentInviteId += 1
	invites = append(invites, i)
	return nil
}

func RemoveInvite(username string) error {
	foundOne := false
	for i, inv := range invites {
		if inv.InviterUsername == username {
			invites = append(invites[:i], invites[i+1:]...)
			foundOne = true
		}
	}
	if foundOne {
		return nil
	} else {
		return fmt.Errorf("Could not find any Invite for %s to delete", username)
	}
}
