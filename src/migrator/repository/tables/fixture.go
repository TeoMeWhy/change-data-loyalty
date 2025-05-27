package tables

import (
	"migrator/models/schemas"
	"time"

	"github.com/jmoiron/sqlx"
)

func mockDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	// Create a mock table
	db.MustExec(`CREATE TABLE customer (
		uuid TEXT,
		desc_customer_name TEXT,
		cod_cpf TEXT,
		desc_email TEXT,
		id_twitch TEXT,
		id_you_tube TEXT,
		id_blue_sky TEXT,
		id_instagram TEXT,
		nr_points INTEGER,
		created_at DATETIME,
		updated_at DATETIME)
	 `)

	db.MustExec(`INSERT INTO customer VALUES ('1', 'John Doe', '12345678901', 'john@example.com', '8923742', '4723967', '3190423', '23094923', 100, '2023-01-01 00:00:00', '2023-01-01 00:00:00')`)
	db.MustExec(`INSERT INTO customer VALUES ('2', 'Jane Doe', '10987654321', 'jane@example.com', '980347312', '47239645237', '319042343124', '230949453223', 200, '2023-01-01 00:00:00', '2023-01-01 00:00:00')`)

	return db, nil

}

func makeTables() []schemas.CustomerSchema {

	time1, _ := time.Parse("2006-01-02 15:04:05", "2023-01-01 00:00:00")
	time2, _ := time.Parse("2006-01-02 15:04:05", "2023-01-01 00:00:00")

	return []schemas.CustomerSchema{
		{
			IdCustomer: "1",
			// DescCustomerName: "John Doe",
			// CodCpf:           "12345678901",
			// DescEmail:        "john@example.com",
			// IdTwitch:         "8923742",
			// IdYouTube:        "4723967",
			// IdBlueSky:        "3190423",
			// IdInstagram:      "23094923",
			NrPoints:  100,
			CreatedAt: time1,
			UpdatedAt: time2,
		},
		{
			IdCustomer: "2",
			// DescCustomerName: "Jane Doe",
			// CodCpf:           "10987654321",
			// DescEmail:        "jane@example.com",
			// IdTwitch:         "980347312",
			// IdYouTube:        "47239645237",
			// IdBlueSky:        "319042343124",
			// IdInstagram:      "230949453223",
			NrPoints:  200,
			CreatedAt: time1,
			UpdatedAt: time2,
		},
	}
}

func makeString() string {
	txt := `UUID;DescCustomerName;CodCpf;DescEmail;IdTwitch;IdYouTube;IdBlueSky;IdInstagram;NrPoints;CreatedAt;UpdatedAt
1;John Doe;12345678901;john@example.com;8923742;4723967;3190423;23094923;100;2023-01-01 00:00:00 +0000 UTC;2023-01-01 00:00:00 +0000 UTC
2;Jane Doe;10987654321;jane@example.com;980347312;47239645237;319042343124;230949453223;200;2023-01-01 00:00:00 +0000 UTC;2023-01-01 00:00:00 +0000 UTC`
	return txt
}
