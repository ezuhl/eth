package model

import (
	"strconv"
)

type ChainHeadResponse struct {
	HeadSlot                   string `json:"headSlot"`
	HeadEpoch                  string `json:"headEpoch"`
	HeadBlockRoot              string `json:"headBlockRoot"`
	FinalizedSlot              string `json:"finalizedSlot"`
	FinalizedEpoch             string `json:"finalizedEpoch"`
	FinalizedBlockRoot         string `json:"finalizedBlockRoot"`
	JustifiedSlot              string `json:"justifiedSlot"`
	JustifiedEpoch             string `json:"justifiedEpoch"`
	JustifiedBlockRoot         string `json:"justifiedBlockRoot"`
	PreviousJustifiedSlot      string `json:"previousJustifiedSlot"`
	PreviousJustifiedEpoch     string `json:"previousJustifiedEpoch"`
	PreviousJustifiedBlockRoot string `json:"previousJustifiedBlockRoot"`
}

func (ch *ChainHeadResponse) GetHeadSlot() (int64, error) {
	headSlot, err := strconv.ParseInt(ch.HeadSlot, 10, 64)
	if err != nil {
		return 0, err
	}

	return headSlot, nil
}

func (ch *ChainHeadResponse) GetFinalizedSlot() (int64, error) {
	headSlot, err := strconv.ParseInt(ch.FinalizedSlot, 10, 64)
	if err != nil {
		return 0, err
	}

	return headSlot, nil
}
