package barcodes

type Barcode struct {
	Barcode string `json:"barcode,omitempty" firestore:"barcode"`
	ID      string `json:"id,omitempty" firestore:"id"`
}
