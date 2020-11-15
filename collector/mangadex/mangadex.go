package mangadex

import (
	"GoManga/collector"
	"GoManga/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type MangaDexConfig struct {
	URL string
}

type MangaDex struct {
	config MangaDexConfig
}

func CreateMangaDex(config MangaDexConfig) MangaDex {
	return MangaDex{config: config}
}

func (md MangaDex) GetName() string {
	return "MangaDex"
}

func (md MangaDex) GetTitle(id int) (db.Manga, error) {
	manga := db.Manga{}

	//Execute Get request
	response, err := http.Get(fmt.Sprint(md.config.URL, "manga/", id))
	if err != nil {
		return manga, err
	}

	//Map body
	data := Manga{}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return manga, err
	}
	err = json.Unmarshal(responseData, &data)
	if err != nil {
		return manga, err
	}

	return db.Manga{
		MangaId:     data.Data.Id,
		Name:        data.Data.Title,
		Description: data.Data.Description,
		Source:      md.GetName(),
		Language:    data.Data.Publication.Language,
	}, err
}

func (md MangaDex) GetChapters(manga db.Manga) ([]collector.Chapter, error) {
	chapters := []collector.Chapter{}

	//Get Chapters
	response, err := http.Get(fmt.Sprint(md.config.URL, "manga/", manga.MangaId, "/chapters"))
	if err != nil {
		return chapters, err
	}

	//Map body
	data := Chapters{}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return chapters, err
	}
	err = json.Unmarshal(responseData, &data)
	if err != nil {
		return chapters, err
	}

	chapters = make([]collector.Chapter, len(data.Data.Chapter))
	for i, chapterId := range data.Data.Chapter {
		chapter, err := getChapter(chapterId.Id, md)
		if err != nil {
			log.Print("Couldn't get chapter", strconv.Itoa(chapterId.Id), "error:", err)
		} else {
			chapters[i] = collector.Chapter{
				TLGroup:  chapter.Data.Group[0].Name,
				Language: chapter.Data.Chapter,
				Volume:   chapter.Data.Volume,
				Chapter:  chapter.Data.Chapter,
			}
			pages := make([]string, len(chapter.Data.Pages))
			for j, page := range chapter.Data.Pages {
				pages[j] = fmt.Sprint(chapter.Data.Server, chapter.Data.Hash, "/", page)
			}
			chapters[i].Pages = pages
		}
	}
	log.Print(chapters[0].Pages[1])
	return chapters, nil
}

func getChapter(chapterId int, md MangaDex) (Chapter, error) {
	response, err := http.Get(fmt.Sprint(md.config.URL, "chapter/", strconv.Itoa(chapterId)))
	if err != nil {
		return Chapter{}, err
	} else {
		chapter := Chapter{}
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return chapter, err
		}
		err = json.Unmarshal(responseData, &chapter)
		return chapter, err
	}
}

type Manga struct {
	Data struct {
		Id          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Publication struct {
			Language string `json:"data.publication.language"`
		} `json:"publication"`
		Cover string `json:"cover"`
	} `json"data"`
}

type Chapters struct {
	Data struct {
		Chapter []struct {
			Id int `json:"id"`
		} `json:"chapters"`
	} `json:"data"`
}

type Chapter struct {
	Data struct {
		Hash    string `json:"hash"`
		Volume  string `json:"volume"`
		Chapter string `json:"chapter"`
		Group   []struct {
			Name string `json:"name"`
		} `json:"groups"`
		Pages  []string `json:"pages"`
		Server string   `json:"server"`
	} `json:"data"`
}
