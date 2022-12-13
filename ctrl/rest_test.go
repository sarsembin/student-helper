package ctrl

import (
	"fmt"
	"math/rand"
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestUniPostAndGet(t *testing.T) {
	client := resty.New()

	postResp, err := client.R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetBody([]byte(`{
			"title": "title5",
			"address": "address",
			"country": "country",
			"region": "region",
			"scholarships": "scholarships",
			"maleToFemale": "maleToFemale",
			"numberOfStudents": 12,
			"tuitFee": "tuitFee",
			"percentageOfInternationalStudents": 10.0,
			"description": "description",
			"studentsPerStaff": 1.1,
			"acceptanceRate": 12.1,
			"averageACTComposite": 123,
			"averageACTEnglish": 123,
			"rank": "25",
			"averageACTMath": 123,
			"averageSATEvidenceBasedReadingWriting": 123,
			"averageSATMath": 123
		}`)).
		Post("http://localhost:9000" + "/api/studentHelper" + "/university")
	if err != nil {
		t.Error(err)
	}

	if postResp.StatusCode() != http.StatusOK {
		t.Errorf("not 200: %v", postResp)
	}

	getResp, err := client.R().
		EnableTrace().
		Get("http://localhost:9000" + "/api/studentHelper" + "/university")
	if err != nil {
		t.Error(err)
	}

	if getResp.StatusCode() != http.StatusOK {
		t.Errorf("not 200: %v", getResp)
	}

	if getResp.Body() == nil {
		t.Errorf("nil body: %v", getResp)
	}
}

func TestComment(t *testing.T) {
	client := resty.New()

	_, err := client.R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetBody([]byte(`{	"content": "this is bad",
		"user_id": 2,
		"university_id": 3}`)).
		Post("http://localhost:9000" + "/user" + "/login")
	if err != nil {
		t.Error(err)
	}
}

func TestRegisterAndLogin(t *testing.T) {
	client := resty.New()

	rand := rand.Int()

	respRegister, err := client.R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetBody([]byte(fmt.Sprintf(`{"password": "amogusamogus",
		"first_name": "kuka",
		"last_name": "abdu",
		"username": "jhk_%v",
		"email": "kuka%v@email.com"}`, rand, rand))).
		Post("http://localhost:9000" + "/user" + "/register")
	if err != nil {
		t.Error(err)
	}

	if respRegister.StatusCode() != http.StatusOK {
		t.Errorf("not 200: %v", respRegister)
	}

	respLogin, err := client.R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetBody([]byte(fmt.Sprintf(`{"email":"kuka%v@email.com", "password":"amogusamogus"}`, rand))).
		Post("http://localhost:9000" + "/user/login")
	if err != nil {
		t.Error(err)
	}

	if respLogin.StatusCode() != http.StatusOK {
		t.Errorf("not 200: %v", respLogin)
	}
}
