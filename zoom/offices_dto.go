package zoom

type CourierOffice struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Description  *string   `json:"description"`
	Street       string    `json:"street"`
	City         string    `json:"city"`
	State        *string   `json:"state"`
	PostalCode   *string   `json:"postal_code"`
	LatStr       string    `json:"lat"`
	LngStr       string    `json:"lng"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
	Phone        string    `json:"phone"`
	Fax          *string   `json:"fax"`
	Email        *string   `json:"email"`
	Website      *string   `json:"website"`
	LogoID       string    `json:"logo_id"`
	Path         *string   `json:"path"`
	MarkerID     *string   `json:"marker_id"`
	Description2 *string   `json:"description_2"`
	OpenHours    string    `json:"open_hours"`
	OpenHoursMap OpenHours `json:"open_hours_map"`
	Ordr         string    `json:"ordr"`
	Brand        *string   `json:"brand"`
	Product      *string   `json:"product"`
	Slug         *string   `json:"slug"`
	Categories   string    `json:"categories"`
	DaysStr      string    `json:"days_str"`
	SKU          string    `json:"sku"`
	MetodoPago   string    `json:"metodo_pago"`
}
