package req

type RequestCreateHotel struct {
	ID              string
	HotelierID      string `json:"hotelier_id"`
	Name            string `json:"name"`
	Overview        string `json:"overview"`
	Activate        bool   `json:"activate"`
	Photos          string `json:"photos"`
	RawAddress      string `json:"raw_address"`
	BankAccount     string `json:"bankAccount"`
	BankName        string `json:"bankName"`
	BankBeneficiary string `json:"bankBeneficiary"`
	Hotel           bool   `json:"hotel"`
	Apartment       bool   `json:"apartment"`
	Resort          bool   `json:"resort"`
	Villa           bool   `json:"villa"`
	Camping         bool   `json:"camping"`
	Motel           bool   `json:"motel"`
	Homestay        bool   `json:"homestay"`
	Spa             bool   `json:"spa"`
	Beach           bool   `json:"beach"`
	Pool            bool   `json:"pool"`
	Bar             bool   `json:"bar"`
	NoSmokingRoom   bool   `json:"no_smoking_room"`
	Fitness         bool   `json:"fitness"`
	Bath            bool   `json:"bath"`
	Wifi            bool   `json:"wifi"`
	Breakfast       bool   `json:"breakfast"`
	Casino          bool   `json:"casio"`
	Parking         bool   `json:"parking"`
	District        int    `json:"district"`
	Province        int    `json:"province"`
	Ward            int    `json:"ward"`
	BusinessLicense string `json:"bussiness_license"`
}

type RequestUpdateHotel struct {
	Name            string `json:"name,omitempty"`
	Overview        string `json:"overview,omitempty"`
	Activate        bool   `json:"activate,omitempty"`
	RawAddress      string `json:"raw_address,omitempty"`
	BankAccount     string `json:"bankAccount,omitempty"`
	BankName        string `json:"bankName,omitempty"`
	BankBeneficiary string `json:"bankBeneficiary,omitempty"`
	Hotel           bool   `json:"hotel,omitempty"`
	Apartment       bool   `json:"apartment,omitempty"`
	Resort          bool   `json:"resort,omitempty"`
	Villa           bool   `json:"villa,omitempty"`
	Camping         bool   `json:"camping,omitempty"`
	Motel           bool   `json:"motel,omitempty"`
	Homestay        bool   `json:"homestay,omitempty"`
	Beach           bool   `json:"beach,omitempty"`
	Pool            bool   `json:"pool,omitempty"`
	Bar             bool   `json:"bar,omitempty"`
	NoSmokingRoom   bool   `json:"no_smoking_room,omitempty"`
	Fitness         bool   `json:"fitness,omitempty"`
	Bath            bool   `json:"bath,omitempty"`
	Wifi            bool   `json:"wifi,omitempty"`
	Spa             bool   `json:"spa,omitempty"`
	Breakfast       bool   `json:"breakfast,omitempty"`
	Casio           bool   `json:"casio,omitempty"`
	Parking         bool   `json:"parking,omitempty"`
	District        int    `json:"district,omitempty"`
	Province        int    `json:"province,omitempty"`
	Ward            int    `json:"ward,omitempty"`
}

type RequestCreatePayout struct {
	Payments string `json:"payments"`
}
