package tools

import (
	"encoding/json"
	"io"
	"net/http"
)

func FetchData(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, target)
}

func FetchArtistData(baseURL string) ([]Card, error) {
	var api APIindex
	if err := FetchData(baseURL, &api); err != nil {
		return nil, err
	}
	var artists []Artist
	if err := FetchData(api.Artists, &artists); err != nil {
		return nil, err
	}
	var location APIlocations
	if err := FetchData(api.Locations, &location); err != nil {
		return nil, err
	}
	var dates APIdates
	if err := FetchData(api.Dates, &dates); err != nil {
		return nil, err
	}
	var relation APIrelations
	if err := FetchData(api.Relations, &relation); err != nil {
		return nil, err
	}

	var cards []Card
	for i, artist := range artists {
		cards = append(cards, Card{
			Id:           artist.Id,
			Image:        artist.Image,
			Name:         artist.Name,
			Members:      artist.Members,
			CreationDate: artist.CreationDate,
			FirstAlbum:   artist.FirstAlbum,
			Locations:    location.Index[i].Location,
			ConcertDates: dates.Index[i].Dates,
			Relation:     relation.Index[i].DatesLocations,
		})
	}
	return cards, nil
}
