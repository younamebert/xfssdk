package apitransfer

// type TransferLink interface {
// 	EnCodeRawTransaction(fromprikey string, tx *servetxpool.StringRawTransaction) (*servetxpool.StringRawTransaction, error)
// 	DeCodeRawTransaction(dataraw string) (*servetxpool.StringRawTransaction, error)
// }

// type ApiTransfer struct{}

// // EnCodeRawTransaction Create a signed transaction
// func (transfer *ApiTransfer) EnCodeRawTransaction(fromprikey string, tx *servetxpool.StringRawTransaction) (string, error) {

// 	if err := tx.SignWithPrivateKey(fromprikey); err != nil {
// 		return "", err
// 	}
// 	return tx.Transfer2Raw()
// }

// // DeCodeRawTransaction Decode a signed transaction into a stringrawtransaction object
// func (transfer *ApiTransfer) DeCodeRawTransaction(dataraw string) (*servetxpool.StringRawTransaction, error) {

// 	databytes, err := base64.StdEncoding.DecodeString(dataraw)
// 	if err != nil {
// 		return nil, err
// 	}

// 	rawtx := &servetxpool.StringRawTransaction{}
// 	if err := json.Unmarshal(databytes, rawtx); err != nil {
// 		return nil, fmt.Errorf("failed to parse data: %s", err)
// 	}

// 	// if err := rawtx.SignWithPrivateKey(fromprikey); err != nil {
// 	// 	return nil, err
// 	// }
// 	return rawtx, nil
// }
