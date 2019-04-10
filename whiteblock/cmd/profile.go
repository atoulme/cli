package cmd

import (
	util "../util"
	"encoding/json"
	"fmt"
	"strconv"
)

var (
	profile Profile
)
var (
	org_key OrganizationApiKey
)

type Organization struct {
	Id        int                      `json:"id"`
	Name      string                   `json:"name"`
	CreatedAt string                   `json:"created_at"`
	UpdatedAt string                   `json:"updated_at"`
	Biomes    []map[string]interface{} `json:"biomes"`
}

type OrganizationApiKey struct {
	Id            int                    `json:"id"`
	Organization  Organization           `json:"organization"`
}

type Profile struct {
	Id            int                      `json:"id"`
	Email         string                   `json:"email"`
	EmailVerified interface{}              `json:"email_verified"`
	Name          string                   `json:"name"`
	Picture       string                   `json:"picture"`
	CreatedAt     string                   `json:"created_at"`
	UpdatedAt     string                   `json:"updated_at"`
	SshKeys       []map[string]interface{} `json:"ssh_keys"`
	Organizations []Organization           `json:"organizations"`
}

func LoadOrganizationApiKey() error {
	rawKey, err := util.ReadStore("org_key")
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawKey, &org_key)
	if err != nil {
		return err
	}

	return nil
}

func LoadProfile() error {
	rawProfile, err := util.ReadStore("profile")
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawProfile, &profile)
	if err != nil {
		return err
	}

	return nil
}

func LoadBiomeAddress() error {
	var org Organization
	if len(profile.Organizations) == 0 {
		return fmt.Errorf("No availible organizations")
	}
	//Grab organization
	if util.StoreExists("organization") {
		rawOrgName, err := util.ReadStore("organization")
		if err != nil {
			return err
		}
		orgName := string(rawOrgName)
		i := 0
		//Allow for automatic detect of organization id
		orgId, err := strconv.Atoi(orgName)
		isOrgId := (err == nil)

		for i = 0; i < len(profile.Organizations); i++ {
			if (isOrgId && profile.Organizations[i].Id == orgId) ||
				(!isOrgId && profile.Organizations[i].Name == orgName) {
				org = profile.Organizations[i]
				break
			}
		}
		if i == len(profile.Organizations) {
			return fmt.Errorf("Could not find organization")
		}
	} else {
		org = profile.Organizations[0]
	}

	var biome map[string]interface{}
	if len(org.Biomes) == 0 {
		return fmt.Errorf("No availible biomes")
	}
	//Dont bother searching for biome if organization is not defined
	if !util.StoreExists("organization") || !util.StoreExists("biome") {
		biome = org.Biomes[0]
	} else {
		rawBiomeName, err := util.ReadStore("biome")
		if err != nil {
			return err
		}
		biomeName := string(rawBiomeName)
		i := 0
		//Allow for automatic detect of organization id
		biomeId, err := strconv.Atoi(biomeName)
		isBiomeId := (err == nil)

		for i = 0; i < len(org.Biomes); i++ {
			if (isBiomeId && int(org.Biomes[i]["id"].(float64)) == biomeId) ||
				(!isBiomeId && org.Biomes[i]["alias"].(string) == biomeName) {
				biome = org.Biomes[i]
				break
			}
		}
		if i == len(org.Biomes) {
			return fmt.Errorf("Could not find biome")
		}
	}
	serverAddr = biome["host"].(string) + ":5001"
	//fmt.Println(serverAddr)
	return nil
}
