package params

// 수정 시작 지점
// UseKPUniverseNetworkConfig uses the KPUniverse beacon chain specific network config.
func UseKPUniverseNetworkConfig() {
	cfg := BeaconNetworkConfig().Copy()
	cfg.ContractDeploymentBlock = 0
	cfg.BootstrapNodes = []string{ // 갱신 필요
		"enr:-MK4QHL8Ytm5TBNUjszTBm_jcLfZY_2GM65ZAp2rMSc07Z7KcI-oODQTF8s-jk12Jnbfi8LP9PP1KDW-dRNYOCfzcvGGAY2CSg9Ih2F0dG5ldHOIAAAAAAAAAACEZXRoMpDRhUOMIAAAk___________gmlkgnY0gmlwhMCoAAiJc2VjcDI1NmsxoQOsqanxPNLzIbCi46u0S4BjzjlMH8FK6lRvQS10Fs2n7IhzeW5jbmV0cwCDdGNwgjLIg3VkcIIu4A",
	}
	OverrideBeaconNetworkConfig(cfg)
}
// 수정 종료 지점


// 수정 시작 지점
// KPUniverseConfig defines the config for the KPUniverse beacon chain testnet.
func KPUniverseConfig() *BeaconChainConfig {
	cfg := MainnetConfig().Copy()
	cfg.MinGenesisTime = 1707288397 // 갱신 필요
	cfg.GenesisDelay = 0
	cfg.ConfigName = KPUniverseName
	cfg.GenesisForkVersion = []byte{0x20, 0x00, 0x00, 0x89} // []byte{0x01, 0x01, 0x70, 0x00}
	cfg.SecondsPerETH1Block = 14
	cfg.DepositChainID = 142536
	cfg.DepositNetworkID = 142536
	cfg.AltairForkEpoch = 0
	cfg.AltairForkVersion = []byte{0x20, 0x00, 0x00, 0x90}
	cfg.BellatrixForkEpoch = 0
	cfg.BellatrixForkVersion = []byte{0x20, 0x00, 0x00, 0x91}
	cfg.CapellaForkEpoch = 0
	cfg.CapellaForkVersion = []byte{0x20, 0x00, 0x00, 0x92}
	cfg.DenebForkEpoch = 100000000000
	cfg.DenebForkVersion = []byte{0x20, 0x00, 0x00, 0x93}
	cfg.TerminalTotalDifficulty = "0"
	cfg.DepositContractAddress = "0x4242424242424242424242424242424242424242"
	cfg.EjectionBalance = 28000000000 // 검증자 풀에서 추방되기위한 잔액 28이더
	cfg.InitializeForkSchedule()
	return cfg
}
// 수정 종료 지점