package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*TokenI)(nil), nil)

	cdc.RegisterConcrete(&Token{}, "checkers/token/Token", nil)

	cdc.RegisterConcrete(&MsgIssueToken{}, "checkers/token/MsgIssueToken", nil)
	cdc.RegisterConcrete(&MsgEditToken{}, "checkers/token/MsgEditToken", nil)
	cdc.RegisterConcrete(&MsgMintToken{}, "checkers/token/MsgMintToken", nil)
	cdc.RegisterConcrete(&MsgBurnToken{}, "checkers/token/MsgBurnToken", nil)
	cdc.RegisterConcrete(&MsgTransferTokenOwner{}, "checkers/token/MsgTransferTokenOwner", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgIssueToken{},
		&MsgEditToken{},
		&MsgMintToken{},
		&MsgBurnToken{},
		&MsgTransferTokenOwner{},
	)
	registry.RegisterInterface(
		"checkers.token.TokenI",
		(*TokenI)(nil),
		&Token{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
