package web

import (
	"errors"
	"net/http"

	"github.com/eduncan911/podcast"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/Kukoon/media-server/models"
)

/**
 *
 * Podcast
 *
 */

type PodcastFormat string

const (
	VideoBestPodcastFormat PodcastFormat = "video_best"
	AudioBestPodcastFormat PodcastFormat = "audio_best"
)

func (p PodcastFormat) IsValid() bool {
	switch p {
	case VideoBestPodcastFormat, AudioBestPodcastFormat:
		return true
	}
	return false
}

func (p PodcastFormat) IsVideo() bool {
	switch p {
	case VideoBestPodcastFormat:
		return true
	}
	return false
}

func (ws *Webservice) rssChannel(c *gin.Context) {
	slug := c.Params.ByName("slug")
	language := c.Params.ByName("lang")
	formatStr := PodcastFormat(c.Params.ByName("format"))

	db := ws.DB

	obj := models.Channel{}

	if !formatStr.IsValid() {
		c.String(http.StatusBadRequest, "no valid file format for podcasts")
		return
	}

	isVideo := formatStr.IsVideo()
	format := podcast.MP4
	if !formatStr.IsVideo() {
		format = podcast.MP3
	}

	// by name or id
	uuid, err := uuid.Parse(slug)
	if err != nil {
		db = db.Where("common_name", slug)
		obj.CommonName = slug
	} else {
		obj.ID = uuid
	}

	if err := db.Preload("Recordings.RecordingLang", func(db *gorm.DB) *gorm.DB {
		return db.Where("lang", language)
	}).Preload("Recordings", func(db *gorm.DB) *gorm.DB {
		return db.Order("Recordings.created_at DESC")
	}).Preload("Recordings.Formats", func(db *gorm.DB) *gorm.DB {
		return db.Where("is_video", isVideo)
	}).First(&obj).Error; err != nil {
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
	p.Language = language

	for _, recording := range obj.Recordings {

		if recording.RecordingLang == nil {
			continue
		}

		recordingFormat := recording.Formats[0]
		description := markdown.ToHTML([]byte(recording.RecordingLang.Description), nil, nil)

		// create an Item
		item := podcast.Item{
			Title:       recording.RecordingLang.Title,
			Link:        recordingFormat.URL,
			Description: string(description),
			PubDate:     &recording.CreatedAt,
		}
		item.AddImage(recording.Poster)
		// add a Download to the Item
		item.AddEnclosure(recordingFormat.URL, format, recordingFormat.Bytes)

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
