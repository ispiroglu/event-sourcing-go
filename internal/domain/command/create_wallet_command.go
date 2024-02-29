package command

import (
	"write-api/internal/common/eventsourcing"
	"write-api/internal/domain/vo"
)

type CreateWalletCommand struct {
	*eventsourcing.BaseCommand
	Wallet vo.Wallet
}
