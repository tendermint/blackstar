package types

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

//-----------------------------------------------------------------------------
// Coin

// Coin hold some amount of one currency.
//
// CONTRACT: A coin will never hold a negative amount of any denomination.
type Coin struct {
	Denom  string `json:"denom"`
	Amount Uint   `json:"amount"`
}

func NewCoin(denom string, amount Uint) Coin {
	return Coin{
		Denom:  denom,
		Amount: amount,
	}
}

func NewUInt64Coin(denom string, amount uint64) Coin {
	return NewCoin(denom, NewUint(amount))
}

// String provides a human-readable representation of a coin
func (coin Coin) String() string {
	return fmt.Sprintf("%v%v", coin.Amount, coin.Denom)
}

// SameDenomAs returns true if the two coins are the same denom
func (coin Coin) SameDenomAs(other Coin) bool {
	return (coin.Denom == other.Denom)
}

// IsZero returns if this represents no money
func (coin Coin) IsZero() bool {
	return coin.Amount.IsZero()
}

// IsGTE returns true if they are the same type and the receiver is
// an equal or greater value
func (coin Coin) IsGTE(other Coin) bool {
	return coin.SameDenomAs(other) && (!coin.Amount.LT(other.Amount))
}

// IsLT returns true if they are the same type and the receiver is
// a smaller value
func (coin Coin) IsLT(other Coin) bool {
	return coin.SameDenomAs(other) && coin.Amount.LT(other.Amount)
}

// IsEqual returns true if the two sets of Coins have the same value
func (coin Coin) IsEqual(other Coin) bool {
	return coin.SameDenomAs(other) && (coin.Amount.Equal(other.Amount))
}

// Adds amounts of two coins with same denom
func (coin Coin) Plus(coinB Coin) Coin {
	if !coin.SameDenomAs(coinB) {
		return coin
	}
	return Coin{coin.Denom, coin.Amount.Add(coinB.Amount)}
}

// Subtracts amounts of two coins with same denom.
func (coin Coin) Minus(coinB Coin) Coin {
	if !coin.SameDenomAs(coinB) {
		return coin
	}
	return Coin{coin.Denom, coin.Amount.Sub(coinB.Amount)}
}

// SafeMinus subtracts coinB from coin. In addition, a boolean is returned
// indicating if integer overflow occurred.
func (coin Coin) SafeMinus(coinB Coin) (Coin, bool) {
	if !coin.SameDenomAs(coinB) {
		return coin, false
	}

	res, overflow := coin.Amount.SafeSub(coinB.Amount)
	return Coin{coin.Denom, res}, overflow
}

//-----------------------------------------------------------------------------
// Coins

type (
	coinSumOp int

	// Coins is a set of Coin, one per currency
	Coins []Coin
)

const (
	coinSumOpAdd coinSumOp = iota
	coinSumOpSub
)

func (coins Coins) String() string {
	if len(coins) == 0 {
		return ""
	}

	out := ""
	for _, coin := range coins {
		out += fmt.Sprintf("%v,", coin.String())
	}
	return out[:len(out)-1]
}

// IsValid asserts the Coins are sorted, and don't have 0 amounts.
func (coins Coins) IsValid() bool {
	switch len(coins) {
	case 0:
		return true
	case 1:
		return !coins[0].IsZero()
	default:
		lowDenom := coins[0].Denom
		for _, coin := range coins[1:] {
			if coin.Denom <= lowDenom {
				return false
			}
			if coin.IsZero() {
				return false
			}
			// we compare each coin against the last denom
			lowDenom = coin.Denom
		}
		return true
	}
}

// Plus adds two sets of coins.
//
// e.g.
// {2A} + {A, 2B} = {3A, 2B}
// {2A} + {0B} = {2A}
//
// CONTRACT: Plus will never return Coins where one Coin has a 0 amount.
func (coins Coins) Plus(coinsB Coins) Coins {
	res, _ := sumCoins(coins, coinsB, coinSumOpAdd) // should not panic
	return res
}

// Minus subtracts a set of coins from another.
//
// e.g.
// {2A} - {A, 2B} = {A, 2B}
// {2A} - {0B} = {2A}
// {A, B} - {A} = {B}
//
// CONTRACT
// - Minus will never return Coins where one Coin has a 0 amount.
// - Minus will panic on unsigned integer overflow
func (coins Coins) Minus(coinsB Coins) Coins {
	return sumCoins(coins, coinsB, coinSumOpSub)
}

// IsAllGT returns True iff for every denom in coins, the denom is present at a
// greater amount in coinsB.
func (coins Coins) IsAllGT(coinsB Coins) bool {
	diff := coins.Minus(coinsB)
	if len(diff) == 0 {
		return false
	}
	return diff.IsPositive()
}

// IsAllGTE returns True iff for every denom in coins, the denom is present at an
// equal or greater amount in coinsB.
func (coins Coins) IsAllGTE(coinsB Coins) bool {
	diff := coins.Minus(coinsB)
	if len(diff) == 0 {
		return true
	}
	return diff.IsNotNegative()
}

// IsAllLT returns True iff for every denom in coins, the denom is present at
// a smaller amount in coinsB.
func (coins Coins) IsAllLT(coinsB Coins) bool {
	diff := coinsB.Minus(coins)
	if len(diff) == 0 {
		return false
	}
	return diff.IsPositive()
}

// IsAllLTE returns True iff for every denom in coins, the denom is present at
// a smaller or equal amount in coinsB.
func (coins Coins) IsAllLTE(coinsB Coins) bool {
	diff := coinsB.Minus(coins)
	if len(diff) == 0 {
		return true
	}
	return diff.IsNotNegative()
}

// IsZero returns true if there are no coins or all coins are zero.
func (coins Coins) IsZero() bool {
	for _, coin := range coins {
		if !coin.IsZero() {
			return false
		}
	}
	return true
}

// IsEqual returns true if the two sets of Coins have the same value
func (coins Coins) IsEqual(coinsB Coins) bool {
	if len(coins) != len(coinsB) {
		return false
	}

	sort.Sort(coins)
	sort.Sort(coinsB)

	for i := 0; i < len(coins); i++ {
		if coins[i].Denom != coinsB[i].Denom || !coins[i].Amount.Equal(coinsB[i].Amount) {
			return false
		}
	}

	return true
}

// Empty returns true if there are no coins and false otherwise.
func (coins Coins) Empty() bool {
	return len(coins) == 0
}

// Returns the amount of a denom from coins
func (coins Coins) AmountOf(denom string) Uint {
	switch len(coins) {
	case 0:
		return ZeroUint()

	case 1:
		coin := coins[0]
		if coin.Denom == denom {
			return coin.Amount
		}
		return ZeroUint()

	default:
		midIdx := len(coins) / 2 // 2:1, 3:1, 4:2
		coin := coins[midIdx]

		if denom < coin.Denom {
			return coins[:midIdx].AmountOf(denom)
		} else if denom == coin.Denom {
			return coin.Amount
		} else {
			return coins[midIdx+1:].AmountOf(denom)
		}
	}
}

func sumCoins(coinsA, coinsB Coins, op coinSumOp) Coins {
	sum := ([]Coin)(nil)
	indexA, indexB := 0, 0
	lenA, lenB := len(coinsA), len(coinsB)

	for {
		if indexA == lenA {
			if indexB == lenB {
				// return nil coins if both sets are empty
				return sum
			}

			// return set B (excluding zero coins) if set A is empty
			return append(sum, filterZeroCoins(coinsB[indexB:])...)
		} else if indexB == lenB {
			// return set A (excluding zero coins) if set B is empty
			return append(sum, filterZeroCoins(coinsA[indexA:])...)
		}

		coinA, coinB := coinsA[indexA], coinsB[indexB]

		switch strings.Compare(coinA.Denom, coinB.Denom) {
		case -1: // coin A denom < coin B denom
			if coinA.IsZero() {
				// ignore 0 sum coin
			} else {
				sum = append(sum, coinA)
			}

			indexA++

		case 0: // coin A denom == coin B denom
			var res Coin

			if op == coinSumOpAdd {
				res = coinA.Plus(coinB)
			} else if op == coinSumOpSub {
				// will panic on overflow
				res = coinA.Minus(coinB)
			}

			if res.IsZero() {
				// ignore 0 sum coin
			} else {
				sum = append(sum, res)
			}

			indexA++
			indexB++

		case 1: // coin A denom > coin B denom
			if coinB.IsZero() {
				// ignore 0 sum coin
			} else {
				sum = append(sum, coinB)
			}

			indexB++
		}
	}
}

// filterZeroCoins removes all zero coins from the given coin set.
func filterZeroCoins(coins Coins) Coins {
	var res Coins

	for _, coin := range coins {
		if !coin.IsZero() {
			res = append(res, coin)
		}
	}

	return res
}

//-----------------------------------------------------------------------------
// Sort interface

//nolint
func (coins Coins) Len() int           { return len(coins) }
func (coins Coins) Less(i, j int) bool { return coins[i].Denom < coins[j].Denom }
func (coins Coins) Swap(i, j int)      { coins[i], coins[j] = coins[j], coins[i] }

var _ sort.Interface = Coins{}

// Sort is a helper function to sort the set of coins inplace
func (coins Coins) Sort() Coins {
	sort.Sort(coins)
	return coins
}

//-----------------------------------------------------------------------------
// Parsing

var (
	// Denominations can be 3 ~ 16 characters long.
	reDnm  = `[[:alpha:]][[:alnum:]]{2,15}`
	reAmt  = `[[:digit:]]+`
	reSpc  = `[[:space:]]*`
	reCoin = regexp.MustCompile(fmt.Sprintf(`^(%s)%s(%s)$`, reAmt, reSpc, reDnm))
)

// ParseCoin parses a cli input for one coin type, returning errors if invalid.
// This returns an error on an empty string as well.
func ParseCoin(coinStr string) (coin Coin, err error) {
	coinStr = strings.TrimSpace(coinStr)

	matches := reCoin.FindStringSubmatch(coinStr)
	if matches == nil {
		err = fmt.Errorf("invalid coin expression: %s", coinStr)
		return
	}
	denomStr, amountStr := matches[2], matches[1]

	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return
	}

	return Coin{denomStr, NewUint(uint64(amount))}, nil
}

// ParseCoins will parse out a list of coins separated by commas.
// If nothing is provided, it returns nil Coins.
// Returned coins are sorted.
func ParseCoins(coinsStr string) (coins Coins, err error) {
	coinsStr = strings.TrimSpace(coinsStr)
	if len(coinsStr) == 0 {
		return nil, nil
	}

	coinStrs := strings.Split(coinsStr, ",")
	for _, coinStr := range coinStrs {
		coin, err := ParseCoin(coinStr)
		if err != nil {
			return nil, err
		}
		coins = append(coins, coin)
	}

	// Sort coins for determinism.
	coins.Sort()

	// Validate coins before returning.
	if !coins.IsValid() {
		return nil, fmt.Errorf("parseCoins invalid: %#v", coins)
	}

	return coins, nil
}
