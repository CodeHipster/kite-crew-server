package model

type Invite struct {
	InviterUsername string
	Code            string
}

type Invites []Invite
