package model

type Response = struct {
	HTMLAttributions HTMLAttributions `json:"html_attributions"`
	Result           Result           `json:"result"`
}

type HTMLAttributions = []string

type Result = struct {
	AddressComponents    AddressComponents `json:"address_components"`
	AdrAddress           string            `json:"adr_address"`
	BusinessStatus       string            `json:"business_status"`
	FormattedAddress     string            `json:"formatted_address"`
	FormattedPhoneNumber string            `json:"formatted_phone_number"`
	Geometry             Geometry          `json:"geometry"`
	Rating               int               `json:"rating"`
	UserRatingsTotal     int               `json:"user_ratings_total"`
	Reviews              Reviews           `json:"reviews"`
	OpeningHours         OpeningHours      `json:"opening_hours"`
	Photos               []PlacePhoto      `json:"photos"`
}

type AddressComponents []struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Geometry struct {
	Location Coordinates `json:"location"`
	Viewport struct {
		NorthEast Coordinates `json:"northeast"`
		SouthWest Coordinates `json:"southwest"`
	} `json:"viewport"`
}

type Reviews []struct {
	AuthorName              string `json:"author_name"`
	AuthorURL               string `json:"author_url"`
	Language                string `json:"language"`
	ProfilePhotoURL         string `json:"profile_photo_url"`
	Rating                  int    `json:"rating"`
	RelativeTimeDescription string `json:"relative_time_description"`
	Text                    string `json:"text"`
	Time                    int    `json:"time"`
}

type PlaceOpeningHoursPeriodDetail struct {
	Day  int    `json:"day"`
	Time string `json:"time"`
}

type PlaceOpeningHoursPeriod struct {
	Open  PlaceOpeningHoursPeriodDetail `json:"open"`
	Close PlaceOpeningHoursPeriodDetail `json:"close"`
}

type OpeningHours struct {
	OpenNow     bool                      `json:"open_now"`
	Periods     []PlaceOpeningHoursPeriod `json:"periods"`
	WeekDayText []string                  `json:"weekday_text"`
}

type PlacePhoto struct {
	Height           int      `json:"height"`
	HtmlAttributions []string `json:"html_attributions"`
	PhotoReference   string   `json:"photo_reference"`
	Width            int      `json:"width"`
}
