type Block struct {
	Index        int
	Timestamp    int64
	Transactions []Transaction
	PrevHash     []byte
	Hash         []byte
	Validator    string // Add this field
}
