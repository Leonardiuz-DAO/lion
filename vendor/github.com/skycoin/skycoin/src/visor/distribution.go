package visor

import "github.com/skycoin/skycoin/src/coin"

const (
	// Maximum supply of skycoins
	MaxCoinSupply uint64 = 888888888 // 888,888,888 million

	FirstCoinSupply uint64 = 88888888 // 88,888,888 million

	// Number of distribution addresses
	DistributionAddressesTotal uint64 = 101

	DistributionAddressInitialBalance uint64 = (MaxCoinSupply - FirstCoinSupply) / (DistributionAddressesTotal - 1)

	// Initial number of unlocked addresses
	InitialUnlockedCount uint64 = DistributionAddressesTotal

	// Number of addresses to unlock per unlock time interval
	UnlockAddressRate uint64 = 5

	// Unlock time interval, measured in seconds
	// Once the InitialUnlockedCount is exhausted,
	// UnlockAddressRate addresses will be unlocked per UnlockTimeInterval
	UnlockTimeInterval uint64 = 60 * 60 * 24 * 365 // 1 year
)

// func init() {
// 	if MaxCoinSupply%DistributionAddressesTotal != 0 {
// 		panic("MaxCoinSupply should be perfectly divisible by DistributionAddressesTotal")
// 	}
// }

// Returns a copy of the hardcoded distribution addresses array.
// Each address has 1,000,000 coins. There are 100 addresses.
func GetDistributionAddresses() []string {
	addrs := make([]string, len(distributionAddresses))
	for i := range distributionAddresses {
		addrs[i] = distributionAddresses[i]
	}
	return addrs
}

// Returns distribution addresses that are unlocked, i.e. they have spendable outputs
func GetUnlockedDistributionAddresses() []string {
	// The first InitialUnlockedCount (30) addresses are unlocked by default.
	// Subsequent addresses will be unlocked at a rate of UnlockAddressRate (5) per year,
	// after the InitialUnlockedCount (30) addresses have no remaining balance.
	// The unlock timer will be enabled manually once the
	// InitialUnlockedCount (30) addresses are distributed.

	// NOTE: To have automatic unlocking, transaction verification would have
	// to be handled in visor rather than in coin.Transactions.Visor(), because
	// the coin package is agnostic to the state of the blockchain and cannot reference it.
	// Instead of automatic unlocking, we can hardcode the timestamp at which the first 30%
	// is distributed, then compute the unlocked addresses easily here.

	addrs := make([]string, InitialUnlockedCount)
	for i := range distributionAddresses[:InitialUnlockedCount] {
		addrs[i] = distributionAddresses[i]
	}
	return addrs
}

// Returns distribution addresses that are locked, i.e. they have unspendable outputs
func GetLockedDistributionAddresses() []string {
	// TODO -- once we reach 30% distribution, we can hardcode the
	// initial timestamp for releasing more coins
	addrs := make([]string, DistributionAddressesTotal-InitialUnlockedCount)
	for i := range distributionAddresses[InitialUnlockedCount:] {
		addrs[i] = distributionAddresses[InitialUnlockedCount+uint64(i)]
	}
	return addrs
}

// Returns true if the transaction spends locked outputs
func TransactionIsLocked(inUxs coin.UxArray) bool {
	lockedAddrs := GetLockedDistributionAddresses()
	lockedAddrsMap := make(map[string]struct{})
	for _, a := range lockedAddrs {
		lockedAddrsMap[a] = struct{}{}
	}

	for _, o := range inUxs {
		uxAddr := o.Body.Address.String()
		if _, ok := lockedAddrsMap[uxAddr]; ok {
			return true
		}
	}

	return false
}

var distributionAddresses = [DistributionAddressesTotal]string{
	"QgkemaxJ9iNezZqFJkVAv12RtJrsp73QrJ",
	"2m36ufkA9rrGzHXK8wBox9m2gVTcZQucosC",
	"2bcxjCJNvLNw6g1eH31GpASJaAq9jjncDnu",
	"ErDhqpTQHZKdhzghEG4N3A8MpQAeaWMFNB",
	"2CteNR6E52nCu5g2e6RKg4bdbgaJKhpUnZh",
	"2R8h5ZfUzbJMdpJrjqFQt2jofgQ5CW1287G",
	"u3iKh8gnwsY7LT4RDPsqjj7R64ud71MPSm",
	"iGow6rSvt1kC2K5EBijiZbdFCZcJEogN6W",
	"2mEgRV5x55TbGEfRGuPF2Khky6NJL5GHUSj",
	"36M7CaDvSZemt9HEtPMtbRtqwFYr8UMsKu",
	"LwGbyxTXRTs7cTkoCkuAvNg4zuRVc6tfD3",
	"bnCFoGmLBu566FZwwt2TE74yeB7EY6nqJQ",
	"owpRoW2Qayo3jKKiASGrmdrX4CEk781d4E",
	"EtxNLUH7UtJgKRLrZvxxsrjqMoFMZaPrWX",
	"kLwcYH9hPPCiErZAs5YhXELsbBdhqi1zpQ",
	"NbqHCMmqQDg9o41C329RHNPHkBJYysoeA3",
	"RL8Tt7pko7XpxuQmfaUWda2xyJHkBKZe4g",
	"sC1ucZaeATFRh1sfbuTSdD4kkqTAgmhPeQ",
	"ebSAmM3RUYTtR6RizRRUupBsADK9pmaqj7",
	"BFaVKiCiAtsXxAEq1T3KWPPJQdGhza4Uyd",
	"5jotwrsdidLhWg5FboN3QsDyKF9b87sYzH",
	"bzbFxjnGRv7dk3EtDp6TVD5pFx5dV8Rw6v",
	"2QpVfyan9VAA8YsgMiqCsDs6kAF21ZiSPgR",
	"xewDKsDBVk1zgenEjrXsSAYKLbKKcCUnJT",
	"Euw4nW8gFpEaZFb1khC4AjBCfeoqN27NHi",
	"uGixmTFxiw8SodeC7yNsTctybbAY3iQDiz",
	"PpnnUinHBArJoxEtcEN2JhynWyY3YRcjmQ",
	"Zy6AzM68ppVrWzEjCEzwskZz1YkDYmGu7f",
	"RVcmTz9LH5pgQVfU5bY8hAVUgdbagC2W2P",
	"ro2tzSWmnXVbgBxuc1BLmoicsTUtFfM22m",
	"Xc17H9tBroySssRgfTrT7PxE54hKQKYoyX",
	"12dHzY1AmirA3bVbSMPp34Ugr8BxNeY1zg",
	"AVhx1qXCT5hB9UdM247NtRicnZZmRFAv1W",
	"ojJv8YT8795T6shmpadPRbXeGCmPQiA8yu",
	"2CirtJ9VXLd9pEbKzickjC3ie4VH9rxFFy7",
	"2i9EVBqWswv6Zp1kkWaPMbPivnX5H1qLFyb",
	"23YFQmFD56ZJgvFXXpZ2EFNTPbpG8DmPmga",
	"2L8o2H8zdVtWBdDRCCLMWjGh9jefcdsk3Zu",
	"VeSffYoM2b8C35g6gXeWkGN5Ea2M2hCcds",
	"3bG9zzFzZGQd2FaGKkkfxzF2BbdUSCuAiq",
	"2EVPnogbCeTX3PKMT6ihW5Mwd857ALRw61J",
	"F3vWRyYoE45eE8FF6cpGVT6vUsNGTYNC7v",
	"2FcirHs5bnuimZdnWnrdfp5Bh15AsNvXtWn",
	"2ckHkZazJKEEhdQbTPaMwee9fkqBQ2tZzZc",
	"CknJoyEVxX68k1R1WMtkZ84rytVYqyYggU",
	"AgujmvHCUCPdW3eoVXg633j5xRVpDKqZLb",
	"9pB3SMrfvnYhJBHoceDVEiCrw8ujcXf1wz",
	"rqD6yaPJDcHqeUwt6wE5sspqt6kurscqMr",
	"WXZWVBFxbKt3AFSAqXziAaTUrik2LnSitC",
	"NnT5n4ndXrzc6Euy3YfxurRr4zZ7oEmdRz",
	"2d5zUyV2QQoYR2aCQsLn2bx5mSCQHc2RdNY",
	"jkGCDYU3GHqpZ6DM65N6HTjbuyUWPzJLHn",
	"qgfchpdMdQqpD4kM4PhLgcKMEifiVU8R62",
	"2jQsedfPAbMwTuoC3GGEdAvhim2u9tAbRys",
	"YSDAVsZjiWX83wgVVws5TqPPbU41idV51x",
	"JmbASXMkutUW6Q5kjSLb5fnDfm697PtzVf",
	"29sdU4CoM5XEEXqii9WCr3gGFsP3zsCTJN",
	"zYceWuLfc8digeDNKGhVMfDhhBHGr6tSJv",
	"25bP6c5rFSkm92oteADuhZKZqwAQZdzhzxZ",
	"3wPBvGmD35BQSFLbRqfaX1J2Soc5avPC3V",
	"9tQ9GP9NydUDz2ZRs7r47m3HNv54zGtpgM",
	"2NTVUSyUVNZC8wwi5CH6Mijz97RFE98HYSV",
	"MqLj8VxoMunb2hGhSEPNSewZFyeZNqytC8",
	"A7G42yh18mycfYiLo6KMzEyJEHtSbmQAW6",
	"5nn78BQRwBVncfScQTJtMqSmKjfXtN7SSH",
	"2MRiKbq7n2iSkb7RDPWS1p1FTFHPNZaYhWj",
	"KvrbdZgUuJcqCrucm69j1w746yB3vePrfQ",
	"2S36oDCViNJAXAyrLqCMnfxLtqEq9ka9HGS",
	"gHBreWug6zwnaRmRB6RP1fHErFLHzAP9AQ",
	"X9mRTCYXaFoLTZ75BzeQhqA5uRAPYr1SVS",
	"LMskJjETwR86XBQCAD5JLGujjid4pQwkaL",
	"2G2uLyuq4diWMHayBBR5fZzZKeUQeYt4o31",
	"2AijqkcYqQt6fyTcXP7jmXb7VXZToH1jRhh",
	"cVQiqKuKRv2h9ZH9UgL5u7uXY5ggsD3GFX",
	"7xiDr8anF3RCEqpNLaZ64pLu1W12T3cVzX",
	"sCweZdyW6iZqRPX5RArvaCWn1vYb6zTiSR",
	"2PwDXwhanUwMrWUPHixreRnFzHKbgoduTUS",
	"2mE7CurC9NiRuG6sC7p4DaoTMk5rFkEAPqs",
	"YeniPAs6AuBqcAuUJ3DYfq5Q9jmLC5Aovz",
	"2m46YTeiWGgExRBPNwPPzXtwui23Xn4Q7Zr",
	"244EbtTS46CEqvy5rkgu12EaMLPPMeYCyfw",
	"UvjbmtfjpRKJJ7UXFweKDHxeLfrp3uHgbi",
	"N6joYyTesijnnG3ETP258TA6PzsMWvcKGZ",
	"27G5xkKcAme5dCP435xFcs6iLo5GvGHvWPN",
	"CpwvUyz3aJCBMV3szdcbzbMF53N1R8dYvj",
	"2KteQNUDt9xRuX8pGGMzJF9zxDZDDGgud8c",
	"2Mw4DCzwuufZ86yAVKu819Ac4Y7aFXYipip",
	"ZVfRhYJqMWUWEi3UTJ69z7c5X8NX38ffRQ",
	"4nBMDAifxuYiubusmcq3T2WsvTgiidxfT8",
	"F2SFZgyd8ecethzfFmcWgA6ManBAQf6srv",
	"2PXeygFzcf5yPpgJwvTkpnGpHAgth1MEDqQ",
	"arcgPjakgvBkojqntzQM1SfjZyTKDxMNh8",
	"2KKdKKY5P1HX3uYiTonTXBmctBjFDaMwbL9",
	"2PoF6HtJFksGQCvpPncNjWT8SBMQn8tHMQq",
	"iSnYxoff1hT3e4LGz9i6wCiu9A9MQWPZph",
	"27ke7fkDsxYzXdGLuDZyLF6tmy5JhP8afTu",
	"D7szUDpY4SSuRZ9sfGU1teCSy6tF3ox6zL",
	"NxjqiCkwEVT581JXLzeAtmmWWGsmT8dEZV",
	"WYSyRxvfC7JyiouesMC9ya22GGu4pbdGky",
	"7E5JoyN26Qsvnpt3Cn463S2WHa6Jvj32eh",
	"H1Ry6zpBNpLYPJvxwwycGNf26hPx3zkc8d",
}
