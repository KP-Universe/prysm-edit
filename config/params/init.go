package params

func init() {
	defaults := []*BeaconChainConfig{
		MainnetConfig(),
		PraterConfig(),
		MinimalSpecConfig(),
		E2ETestConfig(),
		E2EMainnetTestConfig(),
		InteropConfig(),
		HoleskyConfig(),
		// 수정 시작 지점
		KPUniverseConfig(),
		// 수정 종료 지점
		SepoliaConfig(),
	}
	configs = newConfigset(defaults...)
	// ensure that main net is always present and active by default
	if err := SetActive(MainnetConfig()); err != nil {
		panic(err)
	}
	// make sure mainnet is present and active
	m, err := ByName(MainnetName)
	if err != nil {
		panic(err)
	}
	if configs.getActive() != m {
		panic("mainnet should always be the active config at init() time")
	}
}
