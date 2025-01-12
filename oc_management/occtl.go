package oc_management

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

// Occtl occtl command
type Occtl struct{}

// OcctlInterface occtl command methods
type OcctlInterface interface {
	Reload(c context.Context) error
	OnlineUsers(c context.Context) (*[]OcctlUser, error)
	Disconnect(c context.Context, username string) error
	ShowIPBans(c context.Context) (*[]IPBan, error)
	ShowIPBansPoints(c context.Context) (*[]IPBanPoints, error)
	UnBanIP(c context.Context, ip string) error
	ShowStatus(c context.Context) (string, error)
	ShowIRoutes(c context.Context) (*[]IRoute, error)
	ShowUser(c context.Context, username string) (*[]OcctlUser, error)
}

// NewOcctl Create New Occtl command obj
func NewOcctl() *Occtl {
	return &Occtl{}
}

// Reload server configuration reload
func (o *Occtl) Reload(c context.Context) error {
	_, err := OcctlExec(c, "reload")
	if err != nil {
		return err
	}
	return nil
}

// OnlineUsers list of online users
func (o *Occtl) OnlineUsers(c context.Context) (*[]OcctlUser, error) {
	var users []OcctlUser
	result, err := OcctlExec(c, "-j show users")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(result, &users)
	if err != nil {
		return nil, err
	}
	return &users, nil
}

// Disconnect expire user session. On disconnected users raise error
func (o *Occtl) Disconnect(c context.Context, username string) error {
	_, err := OcctlExec(c, fmt.Sprintf("disconnect user %s", username))
	if err != nil {
		return errors.New("failed to disconnect user " + username)
	}
	return nil
}

// ShowIPBans List of banned IPs
func (o *Occtl) ShowIPBans(c context.Context) (*[]IPBan, error) {
	var ipBans []IPBan
	result, err := OcctlExec(c, "-j show ip bans")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(result, &ipBans)
	if err != nil {
		return nil, err
	}
	return &ipBans, nil
}

// ShowIPBansPoints List of baned IPs with points
func (o *Occtl) ShowIPBansPoints(c context.Context) (*[]IPBanPoints, error) {
	var ipBansPoint []IPBanPoints
	result, err := OcctlExec(c, "-j show ip bans points")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(result, &ipBansPoint)
	if err != nil {
		return nil, err
	}
	return &ipBansPoint, nil
}

// UnBanIP unban banned IP
func (o *Occtl) UnBanIP(c context.Context, ip string) error {
	_, err := OcctlExec(c, fmt.Sprintf("unban ip %s", ip))
	if err != nil {
		return err
	}
	return nil
}

// ShowStatus server status
func (o *Occtl) ShowStatus(c context.Context) (string, error) {
	result, err := OcctlExec(c, "show status")
	if err != nil {
		return "", err
	}
	return string(result), nil
}

// ShowIRoutes list user IP routes
func (o *Occtl) ShowIRoutes(c context.Context) (*[]IRoute, error) {
	result, err := OcctlExec(c, "-j show iroutes")
	if err != nil {
		return nil, err
	}
	var routes []IRoute
	err = json.Unmarshal(result, &routes)
	if err != nil {
		return nil, err
	}
	return &routes, nil
}

// ShowUser show user info with extra data
func (o *Occtl) ShowUser(c context.Context, username string) (*[]OcctlUser, error) {
	var user *[]OcctlUser
	result, err := OcctlExec(c, fmt.Sprintf("-j show user %s", username))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(result, &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
