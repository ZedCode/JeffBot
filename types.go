package main

import "time"

// This file contains all the API response types

// SeasonsAPIResponse is exported for encoding/decoding
type SeasonsAPIResponse struct {
	Data []struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			IsCurrentSeason bool `json:"isCurrentSeason"`
			IsOffseason     bool `json:"isOffseason"`
		} `json:"attributes"`
	} `json:"data"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Meta struct {
	} `json:"meta"`
}

// PlayersAPIResponse is exported for encoding/decoding
type PlayersAPIResponse struct {
	Data []struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			Name         string      `json:"name"`
			Stats        interface{} `json:"stats"`
			TitleID      string      `json:"titleId"`
			ShardID      string      `json:"shardId"`
			CreatedAt    time.Time   `json:"createdAt"`
			UpdatedAt    time.Time   `json:"updatedAt"`
			PatchVersion string      `json:"patchVersion"`
		} `json:"attributes"`
		Relationships struct {
			Assets struct {
				Data []interface{} `json:"data"`
			} `json:"assets"`
			Matches struct {
				Data []struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
			} `json:"matches"`
		} `json:"relationships"`
		Links struct {
			Self   string `json:"self"`
			Schema string `json:"schema"`
		} `json:"links"`
	} `json:"data"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Meta struct {
	} `json:"meta"`
}

// PlayerStatsFromAPI is exported for encoding/decoding
type PlayerStatsFromAPI struct {
	Data struct {
		Type       string `json:"type"`
		Attributes struct {
			GameModeStats struct {
				Duo struct {
					Assists             int `json:"assists"`
					Boosts              int `json:"boosts"`
					DBNOs               int `json:"dBNOs"`
					DailyKills          int `json:"dailyKills"`
					DamageDealt         int `json:"damageDealt"`
					Days                int `json:"days"`
					HeadshotKills       int `json:"headshotKills"`
					Heals               int `json:"heals"`
					KillPoints          int `json:"killPoints"`
					Kills               int `json:"kills"`
					LongestKill         int `json:"longestKill"`
					LongestTimeSurvived int `json:"longestTimeSurvived"`
					Losses              int `json:"losses"`
					MaxKillStreaks      int `json:"maxKillStreaks"`
					MostSurvivalTime    int `json:"mostSurvivalTime"`
					Revives             int `json:"revives"`
					RideDistance        int `json:"rideDistance"`
					RoadKills           int `json:"roadKills"`
					RoundMostKills      int `json:"roundMostKills"`
					RoundsPlayed        int `json:"roundsPlayed"`
					Suicides            int `json:"suicides"`
					TeamKills           int `json:"teamKills"`
					TimeSurvived        int `json:"timeSurvived"`
					Top10S              int `json:"top10s"`
					VehicleDestroys     int `json:"vehicleDestroys"`
					WalkDistance        int `json:"walkDistance"`
					WeaponsAcquired     int `json:"weaponsAcquired"`
					WeeklyKills         int `json:"weeklyKills"`
					WinPoints           int `json:"winPoints"`
					Wins                int `json:"wins"`
				} `json:"duo"`
				DuoFpp struct {
					Assists             int `json:"assists"`
					Boosts              int `json:"boosts"`
					DBNOs               int `json:"dBNOs"`
					DailyKills          int `json:"dailyKills"`
					DamageDealt         int `json:"damageDealt"`
					Days                int `json:"days"`
					HeadshotKills       int `json:"headshotKills"`
					Heals               int `json:"heals"`
					KillPoints          int `json:"killPoints"`
					Kills               int `json:"kills"`
					LongestKill         int `json:"longestKill"`
					LongestTimeSurvived int `json:"longestTimeSurvived"`
					Losses              int `json:"losses"`
					MaxKillStreaks      int `json:"maxKillStreaks"`
					MostSurvivalTime    int `json:"mostSurvivalTime"`
					Revives             int `json:"revives"`
					RideDistance        int `json:"rideDistance"`
					RoadKills           int `json:"roadKills"`
					RoundMostKills      int `json:"roundMostKills"`
					RoundsPlayed        int `json:"roundsPlayed"`
					Suicides            int `json:"suicides"`
					TeamKills           int `json:"teamKills"`
					TimeSurvived        int `json:"timeSurvived"`
					Top10S              int `json:"top10s"`
					VehicleDestroys     int `json:"vehicleDestroys"`
					WalkDistance        int `json:"walkDistance"`
					WeaponsAcquired     int `json:"weaponsAcquired"`
					WeeklyKills         int `json:"weeklyKills"`
					WinPoints           int `json:"winPoints"`
					Wins                int `json:"wins"`
				} `json:"duo-fpp"`
				Solo struct {
					Assists             int `json:"assists"`
					Boosts              int `json:"boosts"`
					DBNOs               int `json:"dBNOs"`
					DailyKills          int `json:"dailyKills"`
					DamageDealt         int `json:"damageDealt"`
					Days                int `json:"days"`
					HeadshotKills       int `json:"headshotKills"`
					Heals               int `json:"heals"`
					KillPoints          int `json:"killPoints"`
					Kills               int `json:"kills"`
					LongestKill         int `json:"longestKill"`
					LongestTimeSurvived int `json:"longestTimeSurvived"`
					Losses              int `json:"losses"`
					MaxKillStreaks      int `json:"maxKillStreaks"`
					MostSurvivalTime    int `json:"mostSurvivalTime"`
					Revives             int `json:"revives"`
					RideDistance        int `json:"rideDistance"`
					RoadKills           int `json:"roadKills"`
					RoundMostKills      int `json:"roundMostKills"`
					RoundsPlayed        int `json:"roundsPlayed"`
					Suicides            int `json:"suicides"`
					TeamKills           int `json:"teamKills"`
					TimeSurvived        int `json:"timeSurvived"`
					Top10S              int `json:"top10s"`
					VehicleDestroys     int `json:"vehicleDestroys"`
					WalkDistance        int `json:"walkDistance"`
					WeaponsAcquired     int `json:"weaponsAcquired"`
					WeeklyKills         int `json:"weeklyKills"`
					WinPoints           int `json:"winPoints"`
					Wins                int `json:"wins"`
				} `json:"solo"`
				SoloFpp struct {
					Assists             int     `json:"assists"`
					Boosts              int     `json:"boosts"`
					DBNOs               int     `json:"dBNOs"`
					DailyKills          int     `json:"dailyKills"`
					DamageDealt         float64 `json:"damageDealt"`
					Days                int     `json:"days"`
					HeadshotKills       int     `json:"headshotKills"`
					Heals               int     `json:"heals"`
					KillPoints          float64 `json:"killPoints"`
					Kills               int     `json:"kills"`
					LongestKill         float64 `json:"longestKill"`
					LongestTimeSurvived float64 `json:"longestTimeSurvived"`
					Losses              int     `json:"losses"`
					MaxKillStreaks      int     `json:"maxKillStreaks"`
					MostSurvivalTime    float64 `json:"mostSurvivalTime"`
					Revives             int     `json:"revives"`
					RideDistance        int     `json:"rideDistance"`
					RoadKills           int     `json:"roadKills"`
					RoundMostKills      int     `json:"roundMostKills"`
					RoundsPlayed        int     `json:"roundsPlayed"`
					Suicides            int     `json:"suicides"`
					TeamKills           int     `json:"teamKills"`
					TimeSurvived        float64 `json:"timeSurvived"`
					Top10S              int     `json:"top10s"`
					VehicleDestroys     int     `json:"vehicleDestroys"`
					WalkDistance        float64 `json:"walkDistance"`
					WeaponsAcquired     int     `json:"weaponsAcquired"`
					WeeklyKills         int     `json:"weeklyKills"`
					WinPoints           float64 `json:"winPoints"`
					Wins                int     `json:"wins"`
				} `json:"solo-fpp"`
				Squad struct {
					Assists             int `json:"assists"`
					Boosts              int `json:"boosts"`
					DBNOs               int `json:"dBNOs"`
					DailyKills          int `json:"dailyKills"`
					DamageDealt         int `json:"damageDealt"`
					Days                int `json:"days"`
					HeadshotKills       int `json:"headshotKills"`
					Heals               int `json:"heals"`
					KillPoints          int `json:"killPoints"`
					Kills               int `json:"kills"`
					LongestKill         int `json:"longestKill"`
					LongestTimeSurvived int `json:"longestTimeSurvived"`
					Losses              int `json:"losses"`
					MaxKillStreaks      int `json:"maxKillStreaks"`
					MostSurvivalTime    int `json:"mostSurvivalTime"`
					Revives             int `json:"revives"`
					RideDistance        int `json:"rideDistance"`
					RoadKills           int `json:"roadKills"`
					RoundMostKills      int `json:"roundMostKills"`
					RoundsPlayed        int `json:"roundsPlayed"`
					Suicides            int `json:"suicides"`
					TeamKills           int `json:"teamKills"`
					TimeSurvived        int `json:"timeSurvived"`
					Top10S              int `json:"top10s"`
					VehicleDestroys     int `json:"vehicleDestroys"`
					WalkDistance        int `json:"walkDistance"`
					WeaponsAcquired     int `json:"weaponsAcquired"`
					WeeklyKills         int `json:"weeklyKills"`
					WinPoints           int `json:"winPoints"`
					Wins                int `json:"wins"`
				} `json:"squad"`
				SquadFpp struct {
					Assists             int `json:"assists"`
					Boosts              int `json:"boosts"`
					DBNOs               int `json:"dBNOs"`
					DailyKills          int `json:"dailyKills"`
					DamageDealt         int `json:"damageDealt"`
					Days                int `json:"days"`
					HeadshotKills       int `json:"headshotKills"`
					Heals               int `json:"heals"`
					KillPoints          int `json:"killPoints"`
					Kills               int `json:"kills"`
					LongestKill         int `json:"longestKill"`
					LongestTimeSurvived int `json:"longestTimeSurvived"`
					Losses              int `json:"losses"`
					MaxKillStreaks      int `json:"maxKillStreaks"`
					MostSurvivalTime    int `json:"mostSurvivalTime"`
					Revives             int `json:"revives"`
					RideDistance        int `json:"rideDistance"`
					RoadKills           int `json:"roadKills"`
					RoundMostKills      int `json:"roundMostKills"`
					RoundsPlayed        int `json:"roundsPlayed"`
					Suicides            int `json:"suicides"`
					TeamKills           int `json:"teamKills"`
					TimeSurvived        int `json:"timeSurvived"`
					Top10S              int `json:"top10s"`
					VehicleDestroys     int `json:"vehicleDestroys"`
					WalkDistance        int `json:"walkDistance"`
					WeaponsAcquired     int `json:"weaponsAcquired"`
					WeeklyKills         int `json:"weeklyKills"`
					WinPoints           int `json:"winPoints"`
					Wins                int `json:"wins"`
				} `json:"squad-fpp"`
			} `json:"gameModeStats"`
		} `json:"attributes"`
		Relationships struct {
			MatchesSoloFPP struct {
				Data []struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
			} `json:"matchesSoloFPP"`
			MatchesDuo struct {
				Data []interface{} `json:"data"`
			} `json:"matchesDuo"`
			MatchesDuoFPP struct {
				Data []interface{} `json:"data"`
			} `json:"matchesDuoFPP"`
			MatchesSquad struct {
				Data []interface{} `json:"data"`
			} `json:"matchesSquad"`
			MatchesSquadFPP struct {
				Data []interface{} `json:"data"`
			} `json:"matchesSquadFPP"`
			Season struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
			} `json:"season"`
			Player struct {
				Data struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
			} `json:"player"`
			MatchesSolo struct {
				Data []interface{} `json:"data"`
			} `json:"matchesSolo"`
		} `json:"relationships"`
	} `json:"data"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Meta struct {
	} `json:"meta"`
}
