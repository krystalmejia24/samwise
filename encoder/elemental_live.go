package encoder

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/krystalmejia24/samwise/db"
)

type ElementalLive struct {
	Encoder db.Encoder
	Client
}

const (
	cuePointPath = "/live_events/$eventId/cue_point"
	cuePointData = "<cue_point><get_current_time>1</get_current_time></cue_point>"
)

func (e *ElementalLive) InsertEvent(id3, scte []byte) {
	go e.InsertId3(id3)
	go e.InsertScte(scte)
}

func (e *ElementalLive) InsertScte(scte []byte) error {
	return nil
}

func (e *ElementalLive) InsertId3(id3 []byte) error {
	return nil
}

func (e *ElementalLive) getCuePointCurrentTime(ctx context.Context, result interface{}) error {
	return e.request(ctx, http.MethodPost, cuePointPath, cuePointData, result)
}

func (e *ElementalLive) getHeaders(path string) map[string]string {
	expires := strconv.Itoa(int(time.Now().Add(time.Second * 120).Unix()))

	user, apiKey := e.Encoder.Config.User, e.Encoder.Config.APIKey

	prehash := md5.Sum([]byte(fmt.Sprintf("%s%s%s%s", path, user, apiKey, expires)))
	authHash := md5.Sum(append([]byte(apiKey), prehash[:]...))

	return map[string]string{
		"X-Auth-User":    user,
		"X-Auth-Expires": expires,
		"X-Auth-Key":     hex.EncodeToString(authHash[:]),
		"Accept":         "application/xml",
		"Content_type":   "application/xml",
	}
}

func (e *ElementalLive) request(ctx context.Context, method, path, data string, result interface{}) (err error) {
	u := fmt.Sprintf("%s/%s", e.Encoder.IP, path)

	var req *http.Request
	switch method {
	case http.MethodGet:
		req, err = http.NewRequest(method, u, nil)
	case http.MethodPost:
		body := new(bytes.Buffer)
		err = json.NewEncoder(body).Encode(data)
		if err != nil {
			return err
		}

		req, err = http.NewRequest(method, u, body)
	}

	if err != nil {
		return err
	}

	for k, v := range e.getHeaders(path) {
		req.Header.Add(k, v)
	}

	ctx, cancel := context.WithTimeout(ctx, e.Client.Timeout)
	defer cancel()

	resp, err := e.Client.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode/100 > 3 {
		return StatusError{resp.StatusCode, resp.Status, string(body)}
	}

	if err = json.Unmarshal(body, &result); err != nil {
		return err
	}

	return nil
}
