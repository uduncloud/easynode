package service

type DbMonitorAddressInterface interface {
	NewToken(token *NodeToken) error
	UpdateToken(token string, nodeToken *NodeToken) error
	GetNodeTokenByEmail(email string) (error, *NodeToken)
	AddMonitorAddress(blockchain int64, address *MonitorAddress) error
	GetAddressByToken(blockchain int64, token string) ([]*MonitorAddress, error)
	NewTx(blockchain int64, tx []*Tx) error
	NewBlock(blockchain int64, block []*Block) error
	NewReceipt(blockchain int64, receipt []*Receipt) error
	NewSubTx(blockchain int64, tx []*SubTx) error
}
