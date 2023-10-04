package places

import (
	"coffe-shop/services"
	"coffe-shop/services/guard"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
)

func GetPlaceByID(c echo.Context) error {
	placeID := c.Param("id")
	resp, err := PerformGetPlaceByID(placeID)
	if err != nil {
		log.Info(err.Error())
		return c.String(http.StatusNotFound, "place not found")
	}
	return c.JSON(http.StatusOK, services.ResponseDTO{
		ResponseCode: http.StatusOK,
		Message:      "ok",
		Data:         resp,
	})
}

func GetPlaces(c echo.Context) error {
	strPage := c.QueryParam("page")
	var page uint64 = 1
	var limit uint64 = 10
	nPage, err := strconv.ParseUint(strPage, 10, 64)
	if err == nil {
		page = nPage
	}
	strLimit := c.QueryParam("limit")
	nLimit, err := strconv.ParseUint(strLimit, 10, 64)
	if err == nil {
		limit = nLimit
	}
	strSearch := c.QueryParam("search")
	resp, err := PerformSearchPlaces(page, limit, strSearch)
	if err != nil {
		log.Info(err)
	}
	return c.JSON(http.StatusOK, services.ResponseDTO{
		ResponseCode: http.StatusOK,
		Message:      "ok",
		Data:         resp,
	})
}

func InsertPlace(a guard.AuthGuardRequest) error {
	req := PlacesEntity{}
	if err := a.EchoCtx.Bind(&req); err != nil {
		return a.EchoCtx.String(http.StatusBadRequest, "missing request field(s)")
	}
	resp, err := PerformInsertNewPlace(req)
	if err != nil {
		log.Info(err.Error())
		return a.EchoCtx.String(http.StatusBadRequest, "cannot add new place")
	}
	return a.EchoCtx.JSON(http.StatusOK, services.ResponseDTO{
		ResponseCode: http.StatusOK,
		Message:      "added",
		Data:         resp,
	})
}

func UpdatePlace(a guard.AuthGuardRequest) error {
	req := PlacesEntity{}
	if err := a.EchoCtx.Bind(&req); err != nil {
		return a.EchoCtx.String(http.StatusBadRequest, "missing request field(s)")
	}
	err := PerformUpdatePlaceByID(req)
	if err != nil {
		log.Info(err)
		return a.EchoCtx.String(http.StatusInternalServerError, "cannot update place")
	}
	return a.EchoCtx.JSON(http.StatusOK, services.ResponseDTO{
		ResponseCode: http.StatusOK,
		Message:      "ok",
		Data:         req,
	})
}

func DeletePlace(a guard.AuthGuardRequest) error {
	req := DeletePlaceEntity{}
	if err := a.EchoCtx.Bind(&req); err != nil {
		return a.EchoCtx.String(http.StatusBadRequest, "missing request field(s)")
	}
	err := PerformDeletePlaceByID(req.Id)
	if err != nil {
		log.Info(err)
		return a.EchoCtx.String(http.StatusInternalServerError, "cannot delete place")
	}
	return a.EchoCtx.JSON(http.StatusAccepted, services.ResponseDTO{
		ResponseCode: http.StatusAccepted,
		Message:      "deleted",
		Data: map[string]interface{}{
			"place_id": req.Id,
		},
	})
}

func GetCommentByPlaceID(c echo.Context) error {
	placeID := c.Param("id")
	comments, err := PerformGetListOfCommentByPlaceID(placeID)
	if err != nil {
		return c.String(http.StatusNotFound, "cannot find comment(s)")
	}
	return c.JSON(http.StatusOK, services.ResponseDTO{
		ResponseCode: http.StatusOK,
		Message:      "ok",
		Data:         comments,
	})
}

func AddComment(g guard.AuthGuardRequest) error {
	req := CommentEntity{}
	if err := g.EchoCtx.Bind(&req); err != nil {
		return g.EchoCtx.String(http.StatusBadRequest, "missing request field(s)")
	}
	err := PerformAddComment(req)
	if err != nil {
		log.Info(err)
		return g.EchoCtx.String(http.StatusInternalServerError, "cannot add comment")
	}
	return g.EchoCtx.JSON(http.StatusAccepted, services.ResponseDTO{
		ResponseCode: http.StatusAccepted,
		Message:      "added",
	})
}
