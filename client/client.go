package client

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/JustSomeHack/wyze-go/models"
	"github.com/google/uuid"
)

type WyzeClient interface {
	GetCameras() ([]models.Device, error)
	GetPlugs() ([]models.Device, error)
	TurnOff(mac string, model string) ([]byte, error)
	TurnOn(mac string, model string) ([]byte, error)
}

type wyzeClient struct {
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

func NewWyzeClient(Email string, Password string, KeyID string, APIKey string) WyzeClient {
	Password = hashPassword(Password)
	return &wyzeClient{
		AuthURL:  AuthURL,
		BaseURL:  BaseURL,
		Email:    Email,
		Password: Password,
		KeyID:    KeyID,
		APIKey:   APIKey,
	}
}

func (s *wyzeClient) GetCameras() ([]models.Device, error) {
	devices, err := s.getDevices()
	if err != nil {
		return nil, err
	}
	plugs := make([]models.Device, 0)
	for _, d := range devices {
		if d.ProductType == "Camera" {
			plugs = append(plugs, d)
		}
	}
	return plugs, nil
}

func (s *wyzeClient) GetPlugs() ([]models.Device, error) {
	devices, err := s.getDevices()
	if err != nil {
		return nil, err
	}
	plugs := make([]models.Device, 0)
	for _, d := range devices {
		if d.ProductType == "Plug" {
			plugs = append(plugs, d)
		}
	}
	return plugs, nil
}

func (s *wyzeClient) TurnOff(mac string, model string) ([]byte, error) {
	return s.updateDevice(mac, model, "P3", "0")
}

func (s *wyzeClient) TurnOn(mac string, model string) ([]byte, error) {
	return s.updateDevice(mac, model, "P3", "1")
}

func (s *wyzeClient) getDevices() ([]models.Device, error) {
	s.login()
	data := map[string]string{
		"sv":                "c417b62d72ee44bf933054bdca183e77",
		"access_token":      s.AccessToken,
		"app_name":          "com.hualai",
		"app_ver":           "com.hualai___2.19.14",
		"app_version":       "2.19.14",
		"phone_id":          uuid.New().String(),
		"phone_system_type": "2",
		"sc":                "a626948714654991afd3c0dbd7cdb901",
		"ts":                fmt.Sprintf("%d", (time.Now().UnixMilli() + 10000)),
	}
	body, err := s.sendRequest("app/v2/home_page/get_object_list", data)
	if err != nil {
		return nil, err
	}

	objectResponse := new(models.ObjectListResponse)
	err = json.Unmarshal(body, &objectResponse)
	if err != nil {
		return nil, err
	}

	return objectResponse.Data.DeviceList, nil
}

func (s *wyzeClient) login() {
	url := fmt.Sprintf("%s/%s", s.AuthURL, "api/user/login")
	data := map[string]string{
		"email":    s.Email,
		"password": s.Password,
		"nonce":    fmt.Sprintf("%d", (time.Now().UnixMilli() + 10000)),
	}
	requestData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	request, err := http.NewRequest("POST", url, bytes.NewReader(requestData))
	if err != nil {
		log.Fatal(err)
	}

	hash := hmac.New(md5.New, []byte{})

	request.Header.Add("user-agent", "wyze-sdk-2.2.0")
	request.Header.Add("accept", "*/*")
	request.Header.Add("appid", "9319141212m2ik")
	request.Header.Add("appinfo", "wyze_android_2.19.14")
	request.Header.Add("phoneid", uuid.New().String())
	request.Header.Add("requestid", fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%d", (time.Now().UnixMilli()+10000))))))
	request.Header.Add("signature2", fmt.Sprintf("%x", hash.Sum(requestData)))
	request.Header.Add("x-api-key", "RckMFKbsds5p6QY3COEXc2ABwNTYY0q18ziEiSEm")
	request.Header.Add("content-type", "application/json")
	request.Header.Add("apikey", s.APIKey)
	request.Header.Add("keyid", s.KeyID)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != 200 {
		log.Fatalln("failed to authenticate to API")
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	result := map[string]string{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
	}

	s.AccessToken = result["access_token"]
	s.RefreshToken = result["refresh_token"]
	s.UserID = result["user_id"]
}

func (s *wyzeClient) sendRequest(path string, data map[string]string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", s.BaseURL, path)

	requestData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", url, bytes.NewReader(requestData))
	if err != nil {
		return nil, err
	}

	request.Header.Add("user-agent", "wyze-sdk-2.2.0")
	request.Header.Add("accept", "*/*")
	request.Header.Add("content-type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("failed to authenticate to API")
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, err
}

func (s *wyzeClient) updateDevice(mac string, model string, pid string, pvalue string) ([]byte, error) {
	s.login()
	data := map[string]string{
		"sv":                "44b6d5640c4d4978baba65c8ab9a6d6e",
		"access_token":      s.AccessToken,
		"app_name":          "com.hualai",
		"app_ver":           "com.hualai___2.19.14",
		"app_version":       "2.19.14",
		"phone_id":          uuid.New().String(),
		"phone_system_type": "2",
		"sc":                "a626948714654991afd3c0dbd7cdb901",
		"ts":                fmt.Sprintf("%d", (time.Now().UnixMilli() + 10000)),
		"device_mac":        mac,
		"device_model":      model,
		"pid":               pid,
		"pvalue":            pvalue,
	}
	body, err := s.sendRequest("app/v2/device/set_property", data)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func hashPassword(Password string) string {
	pass := md5.Sum([]byte(Password))
	Password = fmt.Sprintf("%x", pass)
	pass = md5.Sum([]byte(Password))
	Password = fmt.Sprintf("%x", pass)
	pass = md5.Sum([]byte(Password))
	return fmt.Sprintf("%x", pass)
}
