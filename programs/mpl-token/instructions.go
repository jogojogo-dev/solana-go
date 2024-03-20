package mpl_token

import (
	"bytes"
	"fmt"
	ag_spew "github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	ag_text "github.com/gagliardetto/solana-go/text"
	"github.com/gagliardetto/treeout"
)

var ProgramID = solana.TokenMetadataProgramID

const ProgramName = "TokenMetadata"

func init() {
	if !ProgramID.IsZero() {
		solana.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
	}
}

const (
	//Instruction_MintNewEditionFromMasterEditionViaToken uint8 = 11

	Instruction_CreateMasterEditionV3 uint8 = 17

	Instruction_CreateMetadataAccountV3 = 33
)

func InstructionIDToName(id uint8) string {
	switch id {
	//case Instruction_MintNewEditionFromMasterEditionViaToken:
	//	return "MintNewEditionFromMasterEditionViaToken"
	case Instruction_CreateMasterEditionV3:
		return "CreateMasterEditionV3"
	case Instruction_CreateMetadataAccountV3:
		return "CreateMetadataAccountV3"
	default:
		return ""
	}
}

type Instruction struct {
	bin.BaseVariant
}

func (inst *Instruction) EncodeToTree(parent treeout.Branches) {
	if enToTree, ok := inst.Impl.(ag_text.EncodableToTree); ok {
		enToTree.EncodeToTree(parent)
	} else {
		parent.Child(ag_spew.Sdump(inst))
	}
}

var InstructionImplDef = bin.NewVariantDefinition(
	bin.Uint8TypeIDEncoding,
	[]bin.VariantType{
		{
			"CreateMasterEditionV3", (*CreateMasterEditionV3)(nil),
		},
		{
			"CreateMetadataAccountV3", (*CreateMetadataAccountV3)(nil),
		},
	},
)

func (inst *Instruction) ProgramID() solana.PublicKey {
	return ProgramID
}

func (inst *Instruction) Accounts() (out []*solana.AccountMeta) {
	return inst.Impl.(solana.AccountsGettable).GetAccounts()
}

func (inst *Instruction) Data() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bin.NewBorshEncoder(buf).Encode(inst); err != nil {
		return nil, fmt.Errorf("unable to encode instruction: %w", err)
	}
	return buf.Bytes(), nil
}

func registryDecodeInstruction(accounts []*solana.AccountMeta, data []byte) (interface{}, error) {
	inst, err := DecodeInstruction(accounts, data)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func DecodeInstruction(accounts []*solana.AccountMeta, data []byte) (*Instruction, error) {
	inst := new(Instruction)
	if err := bin.NewBorshDecoder(data).Decode(inst); err != nil {
		return nil, fmt.Errorf("unable to decode instruction: %w", err)
	}
	if v, ok := inst.Impl.(solana.AccountsSettable); ok {
		err := v.SetAccounts(accounts)
		if err != nil {
			return nil, fmt.Errorf("unable to set accounts for instruction: %w", err)
		}
	}
	return inst, nil
}
