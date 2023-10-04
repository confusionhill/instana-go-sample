package places

type PlacesEntity struct {
	Id           string `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	ImageURl     string `json:"image_url" db:"image_url"`
	Desc         string `json:"desc" db:"description"`
	Location     string `json:"location" db:"location"`
	Instagram    string `json:"instagram" db:"instagram"`
	IGURL        string `json:"ig_url" db:"ig_url"`
	Contact      string `json:"contact" db:"contact"`
	MenuImageURl string `db:"menu_img_url"`
}

type PlaceResponseEntity struct {
	Id           string `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	ImageURl     string `json:"image_url" db:"image_url"`
	Desc         string `json:"desc" db:"description"`
	Location     string `json:"location" db:"location"`
	Instagram    string `json:"instagram" db:"instagram"`
	IGURL        string `json:"ig_url" db:"ig_url"`
	Contact      string `json:"contact" db:"contact"`
	MenuImageURl string `json:"menu_img_url" db:"menu_img_url"`
}

type CommentEntity struct {
	Id        string `json:"id" db:"id"`
	UserID    string `json:"user_id" db:"user_id"`
	FirstName string `json:"first_name" db:"first_name"`
	PlaceID   string `json:"place_id" db:"place_id"`
	Comment   string `json:"comment" db:"comment"`
}

type DeletePlaceEntity struct {
	Id string `json:"place_id" db:"id"`
}
