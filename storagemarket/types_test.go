package storagemarket_test

import (
	"testing"

	"github.com/xiaomdong/go-fil-markets/storagemarket"
)

func TestDealStagesNil(t *testing.T) {
	var ds *storagemarket.DealStages
	ds.GetStage("none")                                  // no panic.
	ds.AddStageLog("MyStage", "desc", "duration", "msg") // no panic.
}
