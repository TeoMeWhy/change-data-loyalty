package schemas

import "time"

type CustomerSchema struct {
	IdCustomer string `db:"customer_id"`
	// DescCustomerName string    `db:"desc_customer_name"`
	// CodCpf           string    `db:"cod_cpf"`
	// DescEmail        string    `db:"desc_email"`
	// IdTwitch         string    `db:"id_twitch"`
	// IdYouTube        string    `db:"id_you_tube"`
	// IdBlueSky        string    `db:"id_blue_sky"`
	// IdInstagram      string    `db:"id_instagram"`
	NrPoints  int       `db:"nr_points"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type TransactionSchema struct {
	IdTransaction string    `db:"transaction_id"`
	IdCustomer    string    `db:"id_customer"`
	CreatedAt     time.Time `db:"created_at"`
	VlPoints      int       `db:"vl_points"`
	DescSysOrigin string    `db:"desc_sys_origin"`
}

type TransactionProductSchema struct {
	IdTransactionProduct string `db:"id_transaction_product"`
	IdTransaction        string `db:"id_transaction"`
	CodProduct           string `db:"cod_product"`
	QtdeProduct          int    `db:"qtde_product"`
	VlProduct            int    `db:"vl_product"`
}

type ProductSchema struct {
	IdProduct           string `db:"product_id"`
	DescProduct         string `db:"desc_product"`
	DescProductCategory string `db:"desc_product_category"`
}
