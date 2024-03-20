package mpl_token

import (
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/text/format"
	"github.com/gagliardetto/treeout"
)

const (
	UseMethod_Burn bin.BorshEnum = iota
	UseMethod_Multiple
	UseMethod_Single
)

type Creator struct {
	Address  solana.PublicKey
	Verified bool
	Share    uint8
}

type Collection struct {
	Verified bool
	Key      solana.PublicKey
}

type Uses struct {
	UseMethod bin.BorshEnum
	Remaining uint64
	Total     uint64
}

type DataV2 struct {
	Name                 string
	Symbol               string
	Uri                  string
	SellerFeeBasisPoints uint16
	Creators             *[]Creator  `bin:"optional"`
	Collection           *Collection `bin:"optional"`
	Uses                 *Uses       `bin:"optional"`
}

type CollectionDetails struct {
	Kind bin.BorshEnum `bin:"enum"`
	V1   *struct {
		Size uint64
	}
	V2 *struct {
		Padding [8]byte
	}
}

type CreateMetadataAccountV3 struct {
	Data              DataV2
	IsMutable         bool
	CollectionDetails *CollectionDetails `bin:"optional"`

	///   0. `[writable]` metadata
	///   1. `[]` mint
	///   2. `[signer]` mint_authority
	///   3. `[writable, signer]` payer
	///   4. `[signer]` update_authority
	///   5. `[optional]` system_program (default to `11111111111111111111111111111111`)
	solana.AccountMetaSlice `bin:"-"`
}

func NewCreateMetadataAccountV3InstructionBuilder() *CreateMetadataAccountV3 {
	nd := &CreateMetadataAccountV3{
		AccountMetaSlice: make(solana.AccountMetaSlice, 6),
	}
	nd.AccountMetaSlice[5] = solana.Meta(solana.SystemProgramID)
	return nd
}

func (inst *CreateMetadataAccountV3) SetData(data DataV2) *CreateMetadataAccountV3 {
	inst.Data = data
	return inst
}

func (inst *CreateMetadataAccountV3) SetIsMutable(isMutable bool) *CreateMetadataAccountV3 {
	inst.IsMutable = isMutable
	return inst
}

func (inst *CreateMetadataAccountV3) SetCollectionDetails(collectionDetails CollectionDetails) *CreateMetadataAccountV3 {
	inst.CollectionDetails = &collectionDetails
	return inst
}

func (inst *CreateMetadataAccountV3) SetMetadata(metadata solana.PublicKey) *CreateMetadataAccountV3 {
	inst.AccountMetaSlice[0] = solana.Meta(metadata).WRITE()
	return inst
}

func (inst *CreateMetadataAccountV3) SetMint(mint solana.PublicKey) *CreateMetadataAccountV3 {
	inst.AccountMetaSlice[1] = solana.Meta(mint)
	return inst
}

func (inst *CreateMetadataAccountV3) SetMintAuthority(mintAuthority solana.PublicKey) *CreateMetadataAccountV3 {
	inst.AccountMetaSlice[2] = solana.Meta(mintAuthority).SIGNER()
	return inst
}

func (inst *CreateMetadataAccountV3) SetPayer(payer solana.PublicKey) *CreateMetadataAccountV3 {
	inst.AccountMetaSlice[3] = solana.Meta(payer).WRITE().SIGNER()
	return inst
}

func (inst *CreateMetadataAccountV3) SetUpdateAuthority(updateAuthority solana.PublicKey, asSigner bool) *CreateMetadataAccountV3 {
	if asSigner {
		inst.AccountMetaSlice[4] = solana.Meta(updateAuthority).SIGNER()
	} else {
		inst.AccountMetaSlice[4] = solana.Meta(updateAuthority)
	}
	return inst
}

func (inst *CreateMetadataAccountV3) Build() *Instruction {
	return &Instruction{
		BaseVariant: bin.BaseVariant{
			Impl:   inst,
			TypeID: bin.TypeIDFromUint8(Instruction_CreateMetadataAccountV3),
		},
	}
}

func (inst CreateMetadataAccountV3) EncodeToTree(parent treeout.Branches) {
	parent.Child(format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch treeout.Branches) {
			programBranch.Child(format.Instruction("CreateMetadataAccountV3")).
				//
				ParentFunc(func(instructionBranch treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child(format.Param("IsMutable", inst.IsMutable))
					if inst.CollectionDetails != nil {
						instructionBranch.Child(format.Param("CollectionDetails", *inst.CollectionDetails))
					}

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch treeout.Branches) {
						accountsBranch.Child(format.Meta("metadata", inst.AccountMetaSlice[0]))
						accountsBranch.Child(format.Meta("mint", inst.AccountMetaSlice[1]))
						accountsBranch.Child(format.Meta("mintAuthority", inst.AccountMetaSlice[2]))
						accountsBranch.Child(format.Meta("payer", inst.AccountMetaSlice[3]))
						accountsBranch.Child(format.Meta("updateAuthority", inst.AccountMetaSlice[4]))
						accountsBranch.Child(format.Meta("systemProgram", inst.AccountMetaSlice[5]))
					})
				})
		})
}
