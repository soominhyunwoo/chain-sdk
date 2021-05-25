package store

import (
	dbm "github.com/soominhyunwoo/tm-db"

	"github.com/soominhyunwoo/chain-sdk/store/cache"
	"github.com/soominhyunwoo/chain-sdk/store/rootmulti"
	"github.com/soominhyunwoo/chain-sdk/store/types"
)

func NewCommitMultiStore(db dbm.DB) types.CommitMultiStore {
	return rootmulti.NewStore(db)
}

func NewCommitKVStoreCacheManager() types.MultiStorePersistentCache {
	return cache.NewCommitKVStoreCacheManager(cache.DefaultCommitKVStoreCacheSize)
}
