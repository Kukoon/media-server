package web

import (
	"errors"
	"net/http"

	"github.com/eduncan911/podcast"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/Kukoon/media-server/models"
)

// @Summary List all Channels
// @Description Show a list of all channels
// @Produce  json
// @Success 200 {array} models.Channel
// @Failure 500 {object} HTTPError
// @Router /v1/channels [get]
func (ws *Webservice) apiChannelList(c *gin.Context) {
	list := []*models.Channel{}
	if err := ws.DB.Find(&list).Error; err != nil {
		c.JSON(http.StatusInternalServerError, HTTPError{
			Message: APIErrorInternalDatabase,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &list)
}

// @Summary Show Channel
// @Description Show channel with all informations
// @Produce  json
// @Success 200 {object} models.Channel
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Router /v1/channel/{slug} [get]
// @Param slug path string false "slug or uuid of channel"
func (ws *Webservice) apiChannelGet(c *gin.Context) {
	slug := c.Params.ByName("slug")
	db := ws.DB
	obj := models.Channel{}

	uuid, err := uuid.Parse(slug)
	if err != nil {
		db = db.Where("common_name", slug)
		obj.CommonName = slug
	} else {
		obj.ID = uuid
	}
	if err := db.First(&obj).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, HTTPError{
				Message: APIErrorNotFound,
				Error:   err.Error(),
			})
			c.JSON(http.StatusNotFound, err.Error())
			return
		}
		c.JSON(http.StatusInternalServerError, HTTPError{
			Message: APIErrorInternalDatabase,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &obj)
}

/**
 *
 * Podcast
 *
 */

func (ws *Webservice) rssChannel(c *gin.Context) {
	slug := c.Params.ByName("slug")
	db := ws.DB
	obj := models.Channel{}

	uuid, err := uuid.Parse(slug)
	if err != nil {
		db = db.Where("common_name", slug)
		obj.CommonName = slug
	} else {
		obj.ID = uuid
	}
	if err := db.Preload("Recordings", func(db *gorm.DB) *gorm.DB {
		return db.Order("Recordings.created_at DESC")
	}).Preload("Recordings.Formats").First(&obj).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, HTTPError{
				Message: APIErrorNotFound,
				Error:   err.Error(),
			})
			c.JSON(http.StatusNotFound, err.Error())
			return
		}
		c.JSON(http.StatusInternalServerError, HTTPError{
			Message: APIErrorInternalDatabase,
			Error:   err.Error(),
		})
		return
	}
	pubTime := obj.Recordings[0].CreatedAt
	p := podcast.New(obj.Title, "", "", &pubTime, &pubTime)
	p.AddImage(obj.Logo)
	p.Language = "de_DE"

	for _, i := range obj.Recordings {

		url := i.Formats[0].URL
		format := podcast.MP4

		// create an Item
		item := podcast.Item{
			Title:       i.CommonName,
			Link:        url,
			Description: "Description for Episode " + i.CommonName,
			PubDate:     &i.CreatedAt,
		}
		item.AddImage(i.Poster)
		// add a Download to the Item
		item.AddEnclosure(url, format, 0)

		// add the Item and check for validation errors
		if _, err := p.AddItem(item); err != nil {
			c.JSON(http.StatusInternalServerError, HTTPError{
				Message: "Podcast Rendering Error",
				Error:   err.Error(),
			})
			return
		}
	}

	c.Writer.Header().Set("Content-Type", "application/xml")

	if err := p.Encode(c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, HTTPError{
			Message: "Podcast Rendering Error",
			Error:   err.Error(),
		})
	}
}
