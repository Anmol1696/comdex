package types

func NewGenesisState(borrowAsset []BorrowAsset, borrowInterestTracker []BorrowInterestTracker, lendAsset []LendAsset, pool []Pool, assetToPairMapping []AssetToPairMapping, poolAssetLBMapping []PoolAssetLBMapping, lendRewardsTracker []LendRewardsTracker, userAssetLendBorrowMapping []UserAssetLendBorrowMapping, reserveBuybackAssetData []ReserveBuybackAssetData, extended_Pair []Extended_Pair, auctionParams []AuctionParams, assetRatesParams []AssetRatesParams, stableBorrowMapping StableBorrowMapping, params Params) *GenesisState {
	return &GenesisState{
		BorrowAsset:                borrowAsset,
		BorrowInterestTracker:      borrowInterestTracker,
		LendAsset:                  lendAsset,
		Pool:                       pool,
		AssetToPairMapping:         assetToPairMapping,
		PoolAssetLBMapping:         poolAssetLBMapping,
		LendRewardsTracker:         lendRewardsTracker,
		UserAssetLendBorrowMapping: userAssetLendBorrowMapping,
		ReserveBuybackAssetData:    reserveBuybackAssetData,
		Extended_Pair:              extended_Pair,
		AuctionParams:              auctionParams,
		AssetRatesParams:           assetRatesParams,
		StableBorrowMapping:        stableBorrowMapping,
		Params:                     params,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(
		[]BorrowAsset{},
		[]BorrowInterestTracker{},
		[]LendAsset{},
		[]Pool{},
		[]AssetToPairMapping{},
		[]PoolAssetLBMapping{},
		[]LendRewardsTracker{},
		[]UserAssetLendBorrowMapping{},
		[]ReserveBuybackAssetData{},
		[]Extended_Pair{},
		[]AuctionParams{},
		[]AssetRatesParams{},
		StableBorrowMapping{},
		DefaultParams(),
	)
}

func (m *GenesisState) Validate() error {
	return nil
}
