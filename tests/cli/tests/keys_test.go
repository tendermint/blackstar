
func TestCLIKeysAddRecover(t *testing.T) {
	t.Parallel()
	f := helpers.InitFixtures(t)

	exitSuccess, _, _ := f.KeysAddRecover("empty-mnemonic", "")
	require.False(t, exitSuccess)

	exitSuccess, _, _ = f.KeysAddRecover("test-recover", "dentist task convince chimney quality leave banana trade firm crawl eternal easily")
	require.True(t, exitSuccess)
	require.Equal(t, "cosmos1qcfdf69js922qrdr4yaww3ax7gjml6pdds46f4", f.KeyAddress("test-recover").String())

	// Cleanup testing directories
	f.Cleanup()
}

func TestCLIKeysAddRecoverHDPath(t *testing.T) {
	t.Parallel()
	f := helpers.InitFixtures(t)

	f.KeysAddRecoverHDPath("test-recoverHD1", "dentist task convince chimney quality leave banana trade firm crawl eternal easily", 0, 0)
	require.Equal(t, "cosmos1qcfdf69js922qrdr4yaww3ax7gjml6pdds46f4", f.KeyAddress("test-recoverHD1").String())

	f.KeysAddRecoverHDPath("test-recoverH2", "dentist task convince chimney quality leave banana trade firm crawl eternal easily", 1, 5)
	require.Equal(t, "cosmos1pdfav2cjhry9k79nu6r8kgknnjtq6a7rykmafy", f.KeyAddress("test-recoverH2").String())

	f.KeysAddRecoverHDPath("test-recoverH3", "dentist task convince chimney quality leave banana trade firm crawl eternal easily", 1, 17)
	require.Equal(t, "cosmos1909k354n6wl8ujzu6kmh49w4d02ax7qvlkv4sn", f.KeyAddress("test-recoverH3").String())

	f.KeysAddRecoverHDPath("test-recoverH4", "dentist task convince chimney quality leave banana trade firm crawl eternal easily", 2, 17)
	require.Equal(t, "cosmos1v9plmhvyhgxk3th9ydacm7j4z357s3nhtwsjat", f.KeyAddress("test-recoverH4").String())

	// Cleanup testing directories
	f.Cleanup()
}
