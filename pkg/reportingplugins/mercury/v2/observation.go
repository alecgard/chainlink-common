package mercury_v2

import (
	"math/big"

	"github.com/smartcontractkit/libocr/commontypes"
)

var _ ParsedAttributedObservation = parsedAttributedObservation{}

type parsedAttributedObservation struct {
	Timestamp uint32
	Observer  commontypes.OracleID

	BenchmarkPrice *big.Int
	PricesValid    bool

	MaxFinalizedTimestamp      uint32
	MaxFinalizedTimestampValid bool

	LinkFee      *big.Int
	LinkFeeValid bool

	NativeFee      *big.Int
	NativeFeeValid bool
}

func NewParsedAttributedObservation(ts uint32, observer commontypes.OracleID,
	bp *big.Int, pricesValid bool, mfts uint32,
	mftsValid bool, linkFee *big.Int, linkFeeValid bool, nativeFee *big.Int, nativeFeeValid bool) ParsedAttributedObservation {
	return parsedAttributedObservation{
		Timestamp: ts,
		Observer:  observer,

		BenchmarkPrice: bp,
		PricesValid:    pricesValid,

		MaxFinalizedTimestamp:      mfts,
		MaxFinalizedTimestampValid: mftsValid,

		LinkFee:      linkFee,
		LinkFeeValid: linkFeeValid,

		NativeFee:      nativeFee,
		NativeFeeValid: nativeFeeValid,
	}
}

func (pao parsedAttributedObservation) GetTimestamp() uint32 {
	return pao.Timestamp
}

func (pao parsedAttributedObservation) GetObserver() commontypes.OracleID {
	return pao.Observer
}

func (pao parsedAttributedObservation) GetBenchmarkPrice() (*big.Int, bool) {
	return pao.BenchmarkPrice, pao.PricesValid
}

func (pao parsedAttributedObservation) GetBid() (*big.Int, bool) {
	panic("current observation doesn't contain the field")
}

func (pao parsedAttributedObservation) GetAsk() (*big.Int, bool) {
	panic("current observation doesn't contain the field")
}

func (pao parsedAttributedObservation) GetMaxFinalizedTimestamp() (uint32, bool) {
	return pao.MaxFinalizedTimestamp, pao.MaxFinalizedTimestampValid
}

func (pao parsedAttributedObservation) GetLinkFee() (*big.Int, bool) {
	return pao.LinkFee, pao.LinkFeeValid
}

func (pao parsedAttributedObservation) GetNativeFee() (*big.Int, bool) {
	return pao.NativeFee, pao.NativeFeeValid
}
