package client

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/JustSomeHack/wyze-go/models"
)

type config struct {
	WyzeEmail    string `json:"WYZE_EMAIL"`
	WyzePassword string `json:"WYZE_PASSWORD"`
	WyzeKey      string `json:"WYZE_KEY"`
	WyzeAPI      string `json:"WYZE_API"`
}

func loadConfigFile() *config {
	file, err := os.Open("../config.json")
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	c := new(config)
	err = json.Unmarshal(data, &c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func Test_wyzeClient_login(t *testing.T) {
	c := loadConfigFile()
	type fields struct {
		AuthURL  string
		BaseURL  string
		Email    string
		Password string
		KeyID    string
		APIKey   string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Should login",
			fields: fields{
				AuthURL:  AuthURL,
				BaseURL:  BaseURL,
				Email:    c.WyzeEmail,
				Password: hashPassword(c.WyzePassword),
				KeyID:    c.WyzeKey,
				APIKey:   c.WyzeAPI,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &wyzeClient{
				AuthURL:  tt.fields.AuthURL,
				BaseURL:  tt.fields.BaseURL,
				Email:    tt.fields.Email,
				Password: tt.fields.Password,
				KeyID:    tt.fields.KeyID,
				APIKey:   tt.fields.APIKey,
			}
			s.login()
		})
	}
}

func TestNewWyzeClient(t *testing.T) {
	c := loadConfigFile()
	type args struct {
		Email    string
		Password string
		KeyID    string
		APIKey   string
	}
	tests := []struct {
		name string
		args args
		want WyzeClient
	}{
		{
			name: "Should get interface",
			args: args{
				Email:    c.WyzeEmail,
				Password: c.WyzePassword,
				KeyID:    c.WyzeKey,
				APIKey:   c.WyzeAPI,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWyzeClient(tt.args.Email, tt.args.Password, tt.args.KeyID, tt.args.APIKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWyzeClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wyzeClient_getDevices(t *testing.T) {
	c := loadConfigFile()
	type fields struct {
		AuthURL      string
		BaseURL      string
		Email        string
		Password     string
		KeyID        string
		APIKey       string
		AccessToken  string
		RefreshToken string
		UserID       string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Should get devices",
			fields: fields{
				AuthURL:  AuthURL,
				BaseURL:  BaseURL,
				Email:    c.WyzeEmail,
				Password: hashPassword(c.WyzePassword),
				KeyID:    c.WyzeKey,
				APIKey:   c.WyzeAPI,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &wyzeClient{
				AuthURL:      tt.fields.AuthURL,
				BaseURL:      tt.fields.BaseURL,
				Email:        tt.fields.Email,
				Password:     tt.fields.Password,
				KeyID:        tt.fields.KeyID,
				APIKey:       tt.fields.APIKey,
				AccessToken:  tt.fields.AccessToken,
				RefreshToken: tt.fields.RefreshToken,
				UserID:       tt.fields.UserID,
			}
			s.getDevices()
		})
	}
}

func Test_wyzeClient_GetPlugs(t *testing.T) {
	c := loadConfigFile()
	type fields struct {
		AuthURL      string
		BaseURL      string
		Email        string
		Password     string
		KeyID        string
		APIKey       string
		AccessToken  string
		RefreshToken string
		UserID       string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []models.Device
		wantErr bool
	}{
		{
			name: "Should get plugs",
			fields: fields{
				AuthURL:  AuthURL,
				BaseURL:  BaseURL,
				Email:    c.WyzeEmail,
				Password: hashPassword(c.WyzePassword),
				KeyID:    c.WyzeKey,
				APIKey:   c.WyzeAPI,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &wyzeClient{
				AuthURL:      tt.fields.AuthURL,
				BaseURL:      tt.fields.BaseURL,
				Email:        tt.fields.Email,
				Password:     tt.fields.Password,
				KeyID:        tt.fields.KeyID,
				APIKey:       tt.fields.APIKey,
				AccessToken:  tt.fields.AccessToken,
				RefreshToken: tt.fields.RefreshToken,
				UserID:       tt.fields.UserID,
			}
			got, err := s.GetPlugs()
			if (err != nil) != tt.wantErr {
				t.Errorf("wyzeClient.GetPlugs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wyzeClient.GetPlugs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wyzeClient_TurnOn(t *testing.T) {
	c := loadConfigFile()
	type fields struct {
		AuthURL      string
		BaseURL      string
		Email        string
		Password     string
		KeyID        string
		APIKey       string
		AccessToken  string
		RefreshToken string
		UserID       string
	}
	type args struct {
		mac   string
		model string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Should turn on device",
			fields: fields{
				AuthURL:  AuthURL,
				BaseURL:  BaseURL,
				Email:    c.WyzeEmail,
				Password: hashPassword(c.WyzePassword),
				KeyID:    c.WyzeKey,
				APIKey:   c.WyzeAPI,
			},
			args: args{
				mac:   "",
				model: "WLPP1CFH",
			},
			want:    "SUCCESS",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &wyzeClient{
				AuthURL:      tt.fields.AuthURL,
				BaseURL:      tt.fields.BaseURL,
				Email:        tt.fields.Email,
				Password:     tt.fields.Password,
				KeyID:        tt.fields.KeyID,
				APIKey:       tt.fields.APIKey,
				AccessToken:  tt.fields.AccessToken,
				RefreshToken: tt.fields.RefreshToken,
				UserID:       tt.fields.UserID,
			}
			got, err := s.TurnOn(tt.args.mac, tt.args.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("wyzeClient.TurnOn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wyzeClient.TurnOn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wyzeClient_TurnOff(t *testing.T) {
	c := loadConfigFile()
	type fields struct {
		AuthURL      string
		BaseURL      string
		Email        string
		Password     string
		KeyID        string
		APIKey       string
		AccessToken  string
		RefreshToken string
		UserID       string
	}
	type args struct {
		mac   string
		model string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Should turn off device",
			fields: fields{
				AuthURL:  AuthURL,
				BaseURL:  BaseURL,
				Email:    c.WyzeEmail,
				Password: hashPassword(c.WyzePassword),
				KeyID:    c.WyzeKey,
				APIKey:   c.WyzeAPI,
			},
			args: args{
				mac:   "",
				model: "WLPP1CFH",
			},
			want:    "SUCCESS",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &wyzeClient{
				AuthURL:      tt.fields.AuthURL,
				BaseURL:      tt.fields.BaseURL,
				Email:        tt.fields.Email,
				Password:     tt.fields.Password,
				KeyID:        tt.fields.KeyID,
				APIKey:       tt.fields.APIKey,
				AccessToken:  tt.fields.AccessToken,
				RefreshToken: tt.fields.RefreshToken,
				UserID:       tt.fields.UserID,
			}
			got, err := s.TurnOff(tt.args.mac, tt.args.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("wyzeClient.TurnOff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wyzeClient.TurnOff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wyzeClient_getDeviceProperties(t *testing.T) {
	c := loadConfigFile()
	type fields struct {
		AuthURL      string
		BaseURL      string
		Email        string
		Password     string
		KeyID        string
		APIKey       string
		AccessToken  string
		RefreshToken string
		UserID       string
	}
	type args struct {
		mac   string
		model string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.Property
		wantErr bool
	}{
		{
			name: "Should get device properties",
			fields: fields{
				AuthURL:  AuthURL,
				BaseURL:  BaseURL,
				Email:    c.WyzeEmail,
				Password: hashPassword(c.WyzePassword),
				KeyID:    c.WyzeKey,
				APIKey:   c.WyzeAPI,
			},
			args: args{
				mac:   "",
				model: "WLPP1CFH",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &wyzeClient{
				AuthURL:      tt.fields.AuthURL,
				BaseURL:      tt.fields.BaseURL,
				Email:        tt.fields.Email,
				Password:     tt.fields.Password,
				KeyID:        tt.fields.KeyID,
				APIKey:       tt.fields.APIKey,
				AccessToken:  tt.fields.AccessToken,
				RefreshToken: tt.fields.RefreshToken,
				UserID:       tt.fields.UserID,
			}
			got, err := s.getDeviceProperties(tt.args.mac, tt.args.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("wyzeClient.getDeviceProperties() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wyzeClient.getDeviceProperties() = %v, want %v", got, tt.want)
			}
		})
	}
}
