package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"

	"github.com/nahuelmarianolosada/test_project/conversions"
	dateUtils "github.com/nahuelmarianolosada/test_project/dateutils"

	"io"
	"log"
	"net/http"
	"os"
	"time"

	stringUtils "github.com/nahuelmarianolosada/test_project/stringutils"
)

type ResponseUser struct {
	ID               int    `json:"id"`
	Nickname         string `json:"nickname"`
	RegistrationDate string `json:"registration_date"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Gender           string `json:"gender"`
	CountryID        string `json:"country_id"`
	Email            string `json:"email"`
	Identification   struct {
		Number string `json:"number"`
		Type   string `json:"type"`
	} `json:"identification"`
	InternalTags []interface{} `json:"internal_tags"`
	Address      struct {
		Address string `json:"address"`
		City    string `json:"city"`
		State   string `json:"state"`
		ZipCode string `json:"zip_code"`
	} `json:"address"`
	Phone struct {
		AreaCode  string `json:"area_code"`
		Extension string `json:"extension"`
		Number    string `json:"number"`
		Verified  bool   `json:"verified"`
	} `json:"phone"`
	AlternativePhone struct {
		AreaCode  string `json:"area_code"`
		Extension string `json:"extension"`
		Number    string `json:"number"`
	} `json:"alternative_phone"`
	UserType         string      `json:"user_type"`
	Tags             []string    `json:"tags"`
	Logo             interface{} `json:"logo"`
	Points           int         `json:"points"`
	SiteID           string      `json:"site_id"`
	Permalink        string      `json:"permalink"`
	ShippingModes    []string    `json:"shipping_modes"`
	SellerExperience string      `json:"seller_experience"`
	BillData         struct {
		AcceptCreditNote interface{} `json:"accept_credit_note"`
	} `json:"bill_data"`
	SellerReputation struct {
		LevelID           interface{} `json:"level_id"`
		PowerSellerStatus interface{} `json:"power_seller_status"`
		Transactions      struct {
			Canceled  int    `json:"canceled"`
			Completed int    `json:"completed"`
			Period    string `json:"period"`
			Ratings   struct {
				Negative int `json:"negative"`
				Neutral  int `json:"neutral"`
				Positive int `json:"positive"`
			} `json:"ratings"`
			Total int `json:"total"`
		} `json:"transactions"`
		Metrics struct {
			Sales struct {
				Period    string `json:"period"`
				Completed int    `json:"completed"`
			} `json:"sales"`
			Claims struct {
				Period string `json:"period"`
				Rate   int    `json:"rate"`
			} `json:"claims"`
			DelayedHandlingTime struct {
				Period string `json:"period"`
				Rate   int    `json:"rate"`
			} `json:"delayed_handling_time"`
		} `json:"metrics"`
	} `json:"seller_reputation"`
	BuyerReputation struct {
		CanceledTransactions int           `json:"canceled_transactions"`
		Tags                 []interface{} `json:"tags"`
		Transactions         struct {
			Canceled struct {
				Paid  interface{} `json:"paid"`
				Total interface{} `json:"total"`
			} `json:"canceled"`
			Completed   interface{} `json:"completed"`
			NotYetRated struct {
				Paid  interface{} `json:"paid"`
				Total interface{} `json:"total"`
				Units interface{} `json:"units"`
			} `json:"not_yet_rated"`
			Period  string      `json:"period"`
			Total   interface{} `json:"total"`
			Unrated struct {
				Paid  interface{} `json:"paid"`
				Total interface{} `json:"total"`
			} `json:"unrated"`
		} `json:"transactions"`
	} `json:"buyer_reputation"`
	Status struct {
		Billing struct {
			Allow bool          `json:"allow"`
			Codes []interface{} `json:"codes"`
		} `json:"billing"`
		Buy struct {
			Allow            bool          `json:"allow"`
			Codes            []interface{} `json:"codes"`
			ImmediatePayment struct {
				Reasons  []interface{} `json:"reasons"`
				Required bool          `json:"required"`
			} `json:"immediate_payment"`
		} `json:"buy"`
		ConfirmedEmail bool `json:"confirmed_email"`
		ShoppingCart   struct {
			Buy  string `json:"buy"`
			Sell string `json:"sell"`
		} `json:"shopping_cart"`
		ImmediatePayment bool `json:"immediate_payment"`
		List             struct {
			Allow            bool          `json:"allow"`
			Codes            []interface{} `json:"codes"`
			ImmediatePayment struct {
				Reasons  []interface{} `json:"reasons"`
				Required bool          `json:"required"`
			} `json:"immediate_payment"`
		} `json:"list"`
		Mercadoenvios          string `json:"mercadoenvios"`
		MercadopagoAccountType string `json:"mercadopago_account_type"`
		MercadopagoTcAccepted  bool   `json:"mercadopago_tc_accepted"`
		RequiredAction         string `json:"required_action"`
		Sell                   struct {
			Allow            bool          `json:"allow"`
			Codes            []interface{} `json:"codes"`
			ImmediatePayment struct {
				Reasons  []interface{} `json:"reasons"`
				Required bool          `json:"required"`
			} `json:"immediate_payment"`
		} `json:"sell"`
		SiteStatus         string `json:"site_status"`
		UserType           string `json:"user_type"`
		VerificationStatus string `json:"verification_status"`
	} `json:"status"`
	SecureEmail string `json:"secure_email"`
	Credit      struct {
		Consumed      int    `json:"consumed"`
		CreditLevelID string `json:"credit_level_id"`
		Rank          string `json:"rank"`
	} `json:"credit"`
	Context struct {
	} `json:"context"`
}

func mainOld() {

	fmt.Println(stringUtils.Reverse("hola"))

	fmt.Println("es Nuevo? " + conversions.BoolToString(true))

	current := time.Now()
	fmt.Println(dateUtils.ToString(current))
	fmt.Println("Hoy: ", dateUtils.FormatShort(current))
	fmt.Println("Dentro de un mes: ", dateUtils.FormatShort(dateUtils.AddMonths(current, 1)))
	fmt.Println("Mes pasado: ", dateUtils.FormatShort(dateUtils.AddMonths(current, -1)))
	fmt.Println("MM-DD-YYYY hh:mm:ss :", dateUtils.FormatLong(current))

	dateToValidated := current.Add(-24 * time.Hour)
	fmt.Println(dateToValidated)

	fmt.Println(current.Local())
	fmt.Println(current.Location() == time.UTC)

	dateToValidate2 := current.AddDate(0, 0, -1)
	fmt.Println(dateToValidate2)

	/*
		fmt.Println(stringUtils.IsEmpty(""))
		fmt.Println(stringUtils.GetOSEnvironment())

		helloWorld := " Hello, world "
		fmt.Printf("Largo: %d \t '%s'\n", len(helloWorld), helloWorld)
		fmt.Printf("Largo: %d \t '%s'\n",len(stringUtils.Trimmer(helloWorld)), stringUtils.Trimmer(helloWorld))*/

	/*list := readDataFromCSV("user_ids.csv")

	for _,value := range list {
		if value != "" {
			fmt.Println(`"` + value + `",` + getEmailByUserId(value))
		}
	}*/

}

func getEmailByUserId(userId string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://api.internal.ml.com/users/"+userId, nil)
	if err != nil {
		fmt.Errorf("[ERROR]", err)
	}
	req.Header.Add("X-Caller-Scopes", `admin`)
	req.Header.Add("X-Admin-Id", `embeddedSearch`)
	resp, err := client.Do(req)

	defer resp.Body.Close()

	userData := ResponseUser{}
	json.NewDecoder(resp.Body).Decode(&userData)
	return userData.Email
}

func readDataFromCSV(fileName string) []string {
	var listResponse []string
	csvFile, _ := os.Open(fileName)
	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		listResponse = append(listResponse, line[0])
	}

	return listResponse
}
