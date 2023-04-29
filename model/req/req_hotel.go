package req

type RequestCreateHotel struct {
	HotelierID       string `json:"hotelier_id"`
	Name             string `json:"name"`
	Overview         string `json:"overview"`
	Activate         bool   `json:"activate"`
	Photos           string `json:"photos"`
	RawAddress       string `json:"raw_address"`
	BankAccount      string `json:"bankAccount"`
	BankName         string `json:"bankName"`
	BankBeneficiary  string `json:"bankBeneficiary"`
	Hotel            bool   `json:"hotel"`
	Apartment        bool   `json:"apartment"`
	Resort           bool   `json:"resort"`
	Villa            bool   `json:"villa"`
	Camping          bool   `json:"camping"`
	Motel            bool   `json:"motel"`
	Homestay         bool   `json:"homestay"`
	Beach            bool   `json:"beach"`
	Pool             bool   `json:"pool"`
	Bar              bool   `json:"bar"`
	NoSmokingRoom    bool   `json:"no_smoking_room"`
	Fitness          bool   `json:"fitness"`
	Bath             bool   `json:"bath"`
	Wifi             bool   `json:"wifi"`
	Breakfast        bool   `json:"breakfast"`
	Casio            bool   `json:"casio"`
	Parking          bool   `json:"parking"`
	District         int    `json:"district"`
	Province         int    `json:"province"`
	Ward             int    `json:"ward"`
	BussinessLicense string `json:"bussiness_license"`
}

type RequestUpdateHotel struct {
	Name            string `json:"name"`
	Overview        string `json:"overview"`
	Activate        bool   `json:"activate"`
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
	Beach           bool   `json:"beach"`
	Pool            bool   `json:"pool"`
	Bar             bool   `json:"bar"`
	NoSmokingRoom   bool   `json:"no_smoking_room"`
	Fitness         bool   `json:"fitness"`
	Bath            bool   `json:"bath"`
	Wifi            bool   `json:"wifi"`
	Breakfast       bool   `json:"breakfast"`
	Casio           bool   `json:"casio"`
	Parking         bool   `json:"parking"`
	District        int    `json:"district"`
	Province        int    `json:"province"`
	Ward            int    `json:"ward"`
}

type RequestCreatePayout struct {
	Payments string `json:"payments"`
}
