package db

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func defaultEncoder() *Encoder {
	return &Encoder{
		IP: "10.10.10.10",
		Config: &Config{
			User: "User",
			APIKey:   "api-key",
		},
		Stream: &[]Stream{
			Stream{ID: 1},
			Stream{ID: 2},
		},
	}
}

func defaultRepo() *Redis {
	//add mock https://github.com/elliotchance/redismock
	return NewRedis("localhost:6379")
}

func TestRedis_CreateEncoder(t *testing.T) {
	tests := []struct {
		name    string
		e       *Encoder
		wantErr bool
	}{
		{
			name:    "when encoder is created, it is properly stored in redis and no error is thrown",
			e:       defaultEncoder(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateEncoder(tt.e, defaultRepo()); (err != nil) != tt.wantErr {
				t.Errorf("Redis.CreateEncoder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedis_GetEncoder(t *testing.T) {
	expectedEncoder := defaultEncoder()
	wantEnc, err := json.Marshal(expectedEncoder)
	if err != nil {
		fmt.Println(err)
	}

	tests := []struct {
		name    string
		ip      string
		want    interface{}
		wantErr bool
	}{
		{
			name:    "get encoder is properly returned in redis and no error is thrown",
			ip:      expectedEncoder.IP,
			want:    string(wantEnc),
			wantErr: false,
		},
		{
			name:    "get encoder properly returns error if encoder not found",
			ip:      "some ip",
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetEncoder(tt.ip, defaultRepo())
			if (err != nil) != tt.wantErr {
				t.Errorf("Redis.GetEncoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Redis.GetEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepo_DeleteEncoder(t *testing.T) {
	tests := []struct {
		name    string
		ip      string
		want    int64
		wantErr bool
	}{
		{
			name:    "delete encoder is properly returned in redis and no error is thrown",
			ip:      defaultEncoder().IP,
			want:    1,
			wantErr: false,
		},
		{
			name:    "get encoder properly returns error if encoder not found",
			ip:      "some ip",
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteEncoder(tt.ip, defaultRepo())
			if (err != nil) != tt.wantErr {
				t.Errorf("Repo.DeleteEncoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Repo.DeleteEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}
