package mpl_token

import (
	"github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/text/format"
	"github.com/gagliardetto/treeout"
)

type CreateMasterEditionV3 struct {
	MaxSupply *uint64 `bin:"optional"`

	///   0. `[writable]` edition
	///   1. `[writable]` mint
	///   2. `[signer]` update_authority
	///   3. `[signer]` mint_authority
	///   4. `[writable, signer]` payer
	///   5. `[writable]` metadata
	///   6. `[optional]` token_program (default to `TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA`)
	///   7. `[optional]` system_program (default to `11111111111111111111111111111111`)
	solana.AccountMetaSlice `bin:"-"`
}

// NewCreateMasterEditionV3InstructionBuilder creates a new `CreateMasterEditionV3` instruction builder.
func NewCreateMasterEditionV3InstructionBuilder() *CreateMasterEditionV3 {
	inst := &CreateMasterEditionV3{
		AccountMetaSlice: make(solana.AccountMetaSlice, 8),
	}
	inst.AccountMetaSlice[6] = solana.Meta(solana.TokenProgramID)
	inst.AccountMetaSlice[7] = solana.Meta(solana.SystemProgramID)
	return inst
}

func (inst *CreateMasterEditionV3) SetMaxSupply(maxSupply uint64) *CreateMasterEditionV3 {
	inst.MaxSupply = &maxSupply
	return inst
}

func (inst *CreateMasterEditionV3) SetEdition(edition solana.PublicKey) *CreateMasterEditionV3 {
	inst.AccountMetaSlice[0] = solana.Meta(edition).WRITE()
	return inst
}

func (inst *CreateMasterEditionV3) SetMint(mint solana.PublicKey) *CreateMasterEditionV3 {
	inst.AccountMetaSlice[1] = solana.Meta(mint).WRITE()
	return inst
}

func (inst *CreateMasterEditionV3) SetUpdateAuthority(updateAuthority solana.PublicKey) *CreateMasterEditionV3 {
	inst.AccountMetaSlice[2] = solana.Meta(updateAuthority).SIGNER()
	return inst
}

func (inst *CreateMasterEditionV3) SetMintAuthority(mintAuthority solana.PublicKey) *CreateMasterEditionV3 {
	inst.AccountMetaSlice[3] = solana.Meta(mintAuthority).SIGNER()
	return inst
}

func (inst *CreateMasterEditionV3) SetPayer(payer solana.PublicKey) *CreateMasterEditionV3 {
	inst.AccountMetaSlice[4] = solana.Meta(payer).WRITE().SIGNER()
	return inst
}

func (inst *CreateMasterEditionV3) SetMetadata(metadata solana.PublicKey) *CreateMasterEditionV3 {
	inst.AccountMetaSlice[5] = solana.Meta(metadata).WRITE()
	return inst
}

func (inst CreateMasterEditionV3) Build() *Instruction {
	return &Instruction{
		BaseVariant: bin.BaseVariant{
			TypeID: bin.TypeIDFromUint8(Instruction_CreateMasterEditionV3),
			Impl:   inst,
		},
	}
}

func (inst *CreateMasterEditionV3) EncodeToTree(parent treeout.Branches) {
	parent.Child(format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch treeout.Branches) {
			programBranch.Child(format.Instruction("CreateMasterEditionV3")).
				//
				ParentFunc(func(instructionBranch treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params").ParentFunc(func(paramsBranch treeout.Branches) {
						paramsBranch.Child(format.Param("MaxSupply (OPT)", inst.MaxSupply))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=8]").ParentFunc(func(accountsBranch treeout.Branches) {
						accountsBranch.Child(format.Meta("     edition", inst.AccountMetaSlice[0]))
						accountsBranch.Child(format.Meta("        mint", inst.AccountMetaSlice[1]))
						accountsBranch.Child(format.Meta("update_authority", inst.AccountMetaSlice[2]))
						accountsBranch.Child(format.Meta("mint_authority", inst.AccountMetaSlice[3]))
						accountsBranch.Child(format.Meta("       payer", inst.AccountMetaSlice[4]))
						accountsBranch.Child(format.Meta("    metadata", inst.AccountMetaSlice[5]))
						accountsBranch.Child(format.Meta("token_program", inst.AccountMetaSlice[6]))
						accountsBranch.Child(format.Meta("system_program", inst.AccountMetaSlice[7]))
					})
				})
		})
}
