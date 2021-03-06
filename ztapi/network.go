// Copyright © 2017 uxbh
// This file is part of github.com/uxbh/ztdns.

package ztapi

import "fmt"

// GetNetworkInfo returns a Nework containing information about a ZeroTier network
func GetNetworkInfo(API, host, networkID string) *Network {
	resp := new(Network)
	url := fmt.Sprintf("%s/network/%s", host, networkID)
	getJSON(url, API, resp)
	return resp
}

// Network contains the JSON response for a request for a network
type Network struct {
	ID    string
	Type  string
	Clock int
	UI    struct {
		FlowRulesCollapsed    bool
		MembersCollapsed      bool
		MembersHelpCollapsed  bool
		RulesHelpCollapsed    bool
		SettingsCollapsed     bool
		SettingsHelpCollapsed bool
		V4EasyMode            bool
	}
	Config struct {
		ActiveMemberCount     int
		AuthTokens            []string
		AuthorizedMemberCount int
		Capabilities          []string
		Clock                 int
		CreationTime          int
		EnableBroadcast       bool
		ID                    string
		IPAssignmentPools     []struct {
			IPRangeEnd   string
			IPRangeStart string
		}
		MulticastLimit int
		Name           string
		Nwid           string
		Objtype        string
		Private        bool
		Revision       int
		Routes         []struct {
			Target string
			Via    string
		}
		Rules []struct {
			EtherType int
			Not       bool
			Or        bool
			Type      string
		}
		Tags             []string
		TotalMemberCount int
		V4AssignMode     struct {
			Zt bool
		}
		V6AssignMode struct {
			Sixplane bool `json:"6plane"`
			Rfc4193  bool
			Zt       bool
		}
	}
	RuleSource  string
	Description string
	Permissions map[string]struct {
		A bool
		D bool
		M bool
		R bool
		T string
	}
	OnlineMemberCount int
	CapabilitesByName map[string]string
	TagsByName        map[string]string
	CircuitTestEvery  int
}
