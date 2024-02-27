package command

import (
	"write-api/internal/server/domain/vo"
)

type CreateWalletCommand struct {
	BaseCommand
	Wallet vo.Wallet
}
