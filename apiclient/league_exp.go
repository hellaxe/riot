package apiclient

import (
	"context"
	"fmt"

	"github.com/yuhanfang/riot/constants/queue"
	"github.com/yuhanfang/riot/constants/region"
	"github.com/yuhanfang/riot/constants/tier"
)

type LeagueEntry struct {
	QueueType    string     `json:"queueType",datastore:",noindex"`
	SummonerName string     `json:"summonerName",datastore:",noindex"`
	HotStreak    bool       `json:"hotStreak",datastore:",noindex"`
	MiniSeries   MiniSeries `json:"miniSeries",datastore:",noindex"`
	Wins         int        `json:"wins",datastore:",noindex"`
	Veteran      bool       `json:"veteran",datastore:",noindex"`
	Losses       int        `json:"losses",datastore:",noindex"`
	Rank         string     `json:"rank",datastore:",noindex"`
	LeagueID     string     `json:"leagueId",datastore:",noindex"`
	Inactive     bool       `json:"inactive",datastore:",noindex"`
	FreshBlood   bool       `json:"freshBlood",datastore:",noindex"`
	LeagueName   string     `json:"leagueName",datastore:",noindex"`
	SummonerId   string     `json:"summonerId",datastore:",noindex"`
	Tier         tier.Tier  `json:"tier",datastore:",noindex"`
	LeaguePoints int        `json:"leaguePoints",datastore:",noindex"`
}

func (c *client) GetLeagueExpEntries(ctx context.Context, r region.Region, q queue.Queue, tier string, division string, page int) ([]LeagueEntry, error) {
	var res []LeagueEntry
	_, err := c.dispatchAndUnmarshal(ctx, r, "/lol/league-exp/v4/entries", fmt.Sprintf("/%s/%s/%s?page=%d", q.String(), tier, division, page), nil, &res)

	return res, err
}
