package places

import (
	"coffe-shop/database"
	"coffe-shop/services"
	"github.com/doug-martin/goqu/v9"
)

func PerformSearchPlaces(page uint64, limit uint64, search string) ([]PlacesEntity, error) {
	builder := services.LoadBuilder()
	offset := (page - 1) * limit
	query := builder.Select("id", "name", "image_url", "description", "location", "instagram", "ig_url", "contact", "menu_img_url").Prepared(true).From("places")
	if search != "" {
		query = query.Where(goqu.C("name").RegexpILike(search))
	}
	query = query.Order(goqu.I("inserted_at").Asc()).Limit(uint(limit)).Offset(uint(offset))
	sql, params, err := query.ToSQL()
	if err != nil {
		return nil, err
	}
	resp := []PlacesEntity{}
	db := database.GetDB()
	err = db.Select(&resp, sql, params...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func PerformGetPlaceByID(placeID string) (*PlacesEntity, error) {
	db := database.GetDB()
	resp := PlacesEntity{}
	query := "select id, name, image_url, description, location, instagram, ig_url, contact, menu_img_url from places where id = $1"
	err := db.Get(&resp, query, placeID)
	return &resp, err
}

func PerformInsertNewPlace(place PlacesEntity) (*PlacesEntity, error) {
	db := database.GetDB()
	builder := services.LoadBuilder()
	query := builder.Insert("places").Prepared(true).Rows(
		goqu.Record{
			"name":         place.Name,
			"location":     place.Location,
			"image_url":    place.ImageURl,
			"description":  place.Desc,
			"instagram":    place.Instagram,
			"ig_url":       place.IGURL,
			"contact":      place.Contact,
			"menu_img_url": place.MenuImageURl,
		},
	).Returning("id")
	sql, params, err := query.ToSQL()
	if err != nil {
		return nil, err
	}
	err = db.Get(&place, sql, params...)
	if err != nil {
		return nil, err
	}
	return &place, nil
}

func PerformUpdatePlaceByID(place PlacesEntity) error {
	builder := services.LoadBuilder()
	query := builder.Update("places").Prepared(true).Set(goqu.Record{
		"name":         place.Name,
		"location":     place.Location,
		"image_url":    place.ImageURl,
		"description":  place.Desc,
		"instagram":    place.Instagram,
		"ig_url":       place.IGURL,
		"contact":      place.Contact,
		"menu_img_url": place.MenuImageURl,
	}).Where(goqu.C("id").Eq(place.Id))
	sql, params, err := query.ToSQL()
	if err != nil {
		return err
	}
	db := database.GetDB()
	_, err = db.Exec(sql, params...)
	return err
}

func PerformDeletePlaceByID(placeID string) error {
	builder := services.LoadBuilder()
	query := builder.Delete("places").Prepared(true).Where(goqu.C("id").Eq(placeID))
	sql, params, err := query.ToSQL()
	if err != nil {
		return err
	}
	db := database.GetDB()
	_, err = db.Exec(sql, params...)
	return err
}

func PerformGetListOfCommentByPlaceID(placeID string) ([]CommentEntity, error) {
	query := `select 
    			p.id, p.comment, p.place_id, p.user_id, u.first_name 
			from place_comment p 
			    left join users u on p.user_id = u.id
			where p.place_id = $1`
	resp := []CommentEntity{}
	db := database.GetDB()
	err := db.Select(&resp, query, placeID)
	return resp, err
}

func PerformAddComment(comment CommentEntity) error {
	builder := services.LoadBuilder()
	query := builder.Insert("place_comment").Prepared(true).Rows(
		goqu.Record{
			"user_id":  comment.UserID,
			"place_id": comment.PlaceID,
			"comment":  comment.Comment,
		})
	sql, params, err := query.ToSQL()
	if err != nil {
		return err
	}
	db := database.GetDB()
	_, err = db.Exec(sql, params...)
	return err
}
