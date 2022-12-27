package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/96368a/NewApi/model"
	"github.com/96368a/NewApi/services"
	"io/ioutil"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

func get(url string) (map[string]any, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var data map[string]any
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("反序列化失败", err)
		return nil, err
	}
	return data, nil
}

func getArtist(id int64) (*model.Artist, error) {
	url := fmt.Sprintf("https://y.233c.cn/artist/detail?id=%d", id)
	data, err := get(url)
	if err != nil {
		return nil, err
	}
	if data["code"].(float64) != 200 {
		return nil, errors.New("获取歌手信息失败")
	}
	artist := model.Artist{}
	data1, err := json.Marshal(data["data"].(map[string]any)["artist"])
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data1, &artist)
	return &artist, nil
}

func getAlbum(id int64) (*model.Album, error) {
	url := fmt.Sprintf("https://y.233c.cn/album?id=%d", id)
	data, err := get(url)
	if err != nil {
		return nil, err
	}
	if data["code"].(float64) != 200 {
		return nil, errors.New("获取失败")
	}
	album := model.Album{}
	data1, err := json.Marshal(data["album"])
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data1, &album)
	return &album, nil
}

func getSong(id int64) (*model.Song, error) {
	url := fmt.Sprintf("https://y.233c.cn/song/detail?ids=%d", id)
	data, err := get(url)
	if err != nil {
		return nil, err
	}
	if data["code"].(float64) != 200 {
		return nil, errors.New("获取失败")
	}
	song := model.Song{}
	data1, err := json.Marshal(data["songs"].([]any)[0])
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data1, &song)
	return &song, nil
}

func getSongs(ids []int64) ([]model.Song, error) {
	idstr := ""
	for _, id := range ids {
		idstr += fmt.Sprintf("%d,", id)
	}
	idstr = idstr[:len(idstr)-1]

	url := fmt.Sprintf("https://y.233c.cn/song/detail?ids=%s", idstr)
	data, err := get(url)
	if err != nil {
		return nil, err
	}
	if data["code"].(float64) != 200 {
		return nil, errors.New("获取失败")
	}
	songs := []model.Song{}
	data1, err := json.Marshal(data["songs"])
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data1, &songs)
	return songs, nil
}

func getPlaylist(id int64) (*model.Playlist, error) {
	url := fmt.Sprintf("https://y.233c.cn/playlist/detail?id=%d", id)
	data, err := get(url)
	if err != nil {
		return nil, err
	}
	if data["code"].(float64) != 200 {
		return nil, errors.New("获取失败")
	}
	songList := model.Playlist{}
	data1, err := json.Marshal(data["playlist"])
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data1, &songList)
	return &songList, nil
}

func addSong(id int64) error {
	song, err := getSong(id)
	if err != nil {
		return err
	}
	//添加歌手
	for _, ar := range song.Ar {
		artist := services.GetOneArtist(ar.ID)
		if artist == nil {
			arr, err := getArtist(ar.ID)
			if err != nil {
				return err
			}
			services.AddArtist(*arr)
		}
	}
	//添加专辑
	album := services.GetOneAlbum(song.Al.ID)
	if album == nil {
		alb, err := getAlbum(song.Al.ID)
		if err != nil {
			return err
		}
		services.AddAlbum(*alb)
	}
	oneSong := services.GetOneSong(id)
	if oneSong == nil {
		services.AddSong(*song)
	}
	fmt.Println(song)
	return nil
}

func addPlaylist(id int64) error {
	playlist, err := getPlaylist(id)
	if err != nil {
		return err
	}
	var num int = len(playlist.TrackIDS)
	rand.Seed(time.Now().Unix())
	if len(playlist.TrackIDS) > 30 {
		num = rand.Intn(len(playlist.TrackIDS)-30) + 30
	}
	for i, song := range playlist.TrackIDS {
		if i < num {
			addSong(song.ID)
		}
	}
	onePlaylist := services.GetOnePlaylist(id)
	if onePlaylist == nil {
		services.AddPlaylist(*playlist)
	}
	return nil
}

func TestAddData(t *testing.T) {
	//err := addPlaylist(3778678)
	//if err != nil {
	//	t.Error(err)
	//}http://y.233c.cn/personalized?limit=102
	url := "https://y.233c.cn/top/playlist/highquality?limit=101"
	data, err := get(url)
	if err != nil {
		t.Error(err)
	}
	if data["code"].(float64) != 200 {
		t.Error("获取失败")
	}
	playlists := make([]model.Playlist, 0)
	data1, err := json.Marshal(data["playlists"])
	if err != nil {
		t.Error(err)
	}
	err = json.Unmarshal(data1, &playlists)
	fmt.Printf("%v", playlists)
	for _, playlist := range playlists {
		addPlaylist(playlist.ID)
	}

}

func TestAddSong(t *testing.T) {
	id := int64(347230)
	song, err := getSong(id)
	if err != nil {
		t.Error(err)
	}
	//添加歌手
	for _, ar := range song.Ar {
		artist := services.GetOneArtist(ar.ID)
		if artist == nil {
			arr, err := getArtist(ar.ID)
			if err != nil {
				t.Error(err)
			}
			services.AddArtist(*arr)
		}
	}
	//添加专辑
	album := services.GetOneAlbum(song.Al.ID)
	if album == nil {
		alb, err := getAlbum(song.Al.ID)
		if err != nil {
			t.Error(err)
		}
		services.AddAlbum(*alb)
	}
	oneSong := services.GetOneSong(id)
	if oneSong == nil {
		services.AddSong(*song)
	}
	fmt.Println(song)
}

func TestSongDetail(t *testing.T) {
	//artist, err := getArtist(11972054)
	//if err != nil {
	//	return
	//}
	//fmt.Println(artist)
	//album, err := getAlbum(32311)
	//if err != nil {
	//	return
	//}
	//fmt.Println(album)
	//song, err := getSong(347230)
	//if err != nil {
	//	return
	//}
	//fmt.Println(song)
	playlist, err := getPlaylist(3779629)
	if err != nil {
		return
	}
	fmt.Println(playlist)
	ids := []int64{}
	for _, song := range playlist.TrackIDS {
		ids = append(ids, song.ID)
	}
	songs, err := getSongs(ids)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(songs)
}
