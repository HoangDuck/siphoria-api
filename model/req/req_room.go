package req

type RequestCreateRoomType struct {
	ID             string
	Name           string `json:"name"`
	Description    string `json:"description"`
	BedNums        int    `json:"bed_nums"`
	BathroomNums   int    `json:"bathroom_nums"`
	Activated      bool   `json:"activated"`
	MaxChildren    int    `json:"max_children"`
	MaxAdult       int    `json:"max_adult"`
	HotelID        string `json:"hotel_id"`
	Photos         string `json:"photos"`
	Bay            bool   `json:"bay"`
	Ocean          bool   `json:"ocean"`
	City           bool   `json:"city"`
	Garden         bool   `json:"garden"`
	Lake           bool   `json:"lake"`
	Mountain       bool   `json:"mountain"`
	River          bool   `json:"river"`
	PrivateBalcony bool   `json:"private_balcony"`
	AirConditional bool   `json:"air_conditional"`
	Tivi           bool   `json:"tivi"`
	Heater         bool   `json:"heater"`
	Kitchen        bool   `json:"kitchen"`
	PrivatePool    bool   `json:"private_pool"`
	Iron           bool   `json:"iron"`
	Sofa           bool   `json:"sofa"`
	Desk           bool   `json:"desk"`
	Soundproof     bool   `json:"soundproof"`
	Towels         bool   `json:"towels"`
	Toiletries     bool   `json:"toiletries"`
	Shower         bool   `json:"shower"`
	Slippers       bool   `json:"slippers"`
	Hairdry        bool   `json:"hairdry"`
	Fruit          bool   `json:"fruit"`
	Bbq            bool   `json:"bbq"`
	Wine           bool   `json:"wine"`
	Fryer          bool   `json:"fryer"`
	KitchenTool    bool   `json:"kitchen_tool"`
}

type RequestUpdateRoomNight struct {
	RoomTypeID   string `json:"room_type_id"`
	SelectedDate string `json:"selected_date"`
	Quantity     int    `json:"quantity"`
}

type RequestUpdateRoomType struct {
	Name           string `json:"name,omitempty"`
	Description    string `json:"description,omitempty"`
	BedNums        int    `json:"bed_nums,omitempty"`
	BathroomNums   int    `json:"bathroom_nums,omitempty"`
	Activated      bool   `json:"activated,omitempty"`
	MaxChildren    int    `json:"max_children,omitempty"`
	MaxAdult       int    `json:"max_adult,omitempty"`
	Ocean          bool   `json:"ocean,omitempty"`
	Sea            bool   `json:"sea,omitempty"`
	Bay            bool   `json:"bay,omitempty"`
	City           bool   `json:"city,omitempty"`
	Garden         bool   `json:"garden,omitempty"`
	Lake           bool   `json:"lake,omitempty"`
	Mountain       bool   `json:"mountain,omitempty"`
	River          bool   `json:"river,omitempty"`
	PrivateBalcony bool   `json:"private_balcony,omitempty"`
	AirConditional bool   `json:"air_conditional,omitempty"`
	Tivi           bool   `json:"tivi,omitempty"`
	Kitchen        bool   `json:"kitchen,omitempty"`
	PrivatePool    bool   `json:"private_pool,omitempty"`
	Heater         bool   `json:"heater,omitempty"`
	Iron           bool   `json:"iron,omitempty"`
	Sofa           bool   `json:"sofa,omitempty"`
	Desk           bool   `json:"desk,omitempty"`
	Soundproof     bool   `json:"soundproof,omitempty"`
	Towels         bool   `json:"towels,omitempty"`
	Toiletries     bool   `json:"toiletries,omitempty"`
	Shower         bool   `json:"shower,omitempty"`
	Slippers       bool   `json:"slippers,omitempty"`
	Hairdry        bool   `json:"hairdry,omitempty"`
	Fuirt          bool   `json:"fuirt,omitempty"`
	Bbq            bool   `json:"bbq,omitempty"`
	Wine           bool   `json:"wine,omitempty"`
	Fryer          bool   `json:"fryer,omitempty"`
	KitchenTool    bool   `json:"kitchen_tool,omitempty"`
	IsDelete       bool   `json:"is_delete,omitempty"`
}
