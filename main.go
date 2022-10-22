package main

import (
	"Projectgo1/helper"
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type MessageReq struct {
	MessagingProductReq string  `json:"messaging_product,omitempty"`
	RecipientType       string  `json:"recipient_type,omitempty"`
	To                  string  `json:"to,omitempty"`
	Type                string  `json:"type,omitmepty"`
	TextReq             TextReq `json:"text,omitempty"`
}

type TempMessageReq struct {
	MessagingProductReq string   `json:"messaging_product,omitempty"`
	RecipientType       string   `json:"recipient_type,omitempty"`
	To                  string   `json:"to,omitempty"`
	Type                string   `json:"type,omitmepty"`
	Template            Template `json:"template,omitempty"`
}

type Template struct {
	Name     string   `json:"name,omitempty"`
	Language Language `json:"language,omitempty"`
}

type Language struct {
	Code string `json:"code,omitempty"`
}

type TextReq struct {
	PreviewURL bool   `json:"preview_url,omitempty"`
	Body       string `json:"body,omitempty"`
}

type StatusReq struct {
	MessagingProductReq string `json:"messaging_product,omitempty"`
	MessageId           string `json:"message_id,omitempty"`
	StatusReq           string `json:"status,omitempty"`
}

type Response struct {
	MessagingProductReq string           `json:"messaging_product,omitempty"`
	Contacts            []helper.Contact `json:"contacts,omitempty"`
	Error               helper.Error     `json:"error,omitempty"`
	Messages            []helper.Message `json:"messages,omitempty"`
	StatusCode          string           `json:"status_code,omitempty"`
}

var database = make(map[string]string)

func main() {
	fmt.Println("Hello World!")
	e := echo.New()
	e.GET("/", getReq)
	e.GET("/webhook", webhookReq)
	e.GET("/hitpost", hitPost)
	e.POST("/webhook", icmPayload)
	e.Logger.Fatal(e.Start(":2552"))
}

func getReq(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func hitPost(c echo.Context) error {
	obj := new(TempMessageReq)
	obj.MessagingProductReq = "whatsapp"
	obj.RecipientType = "individual"
	obj.To = "917095101959"
	obj.Type = "template"
	obj.Template.Name = "promoverify "
	obj.Template.Language.Code = "en"
	jsonStr, err := json.Marshal(obj)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println("string object", string(jsonStr))
	}
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, "https://graph.facebook.com/v15.0/105396722334906/messages", bytes.NewBuffer(jsonStr))
	r.Header.Add("Authorization", "Bearer EAAIglDnk6OoBADazOW3spDSR5zfm2RTC3THYUdcBSLPb4Gd7ijpePEPd1PqSNFmE4Las5kSTx70bW3bAZAZBs402owNkS4huV0M9OBA2yZCdVDiXyqrpp46W534s2bwm749PPYPEQ4N2ZAGPH48B07ItzSElu721RpWeLBaAfUAbLQrDuFpjUkxSk4nkZBQsVmrV3mHhyosYdQlghxZCSI")
	r.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(r)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(resp.Status, string(body))
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["json"])
	fmt.Println("request success")
	return c.String(http.StatusOK, "Request sent!")
}

func webhookReq(r echo.Context) error {
	mode := r.QueryParam("hub.mode")
	token := r.QueryParam("hub.verify_token")
	challenge := r.QueryParam("hub.challenge")

	if mode == "subscribe" && token == "testwebhook" {
		fmt.Sprintln("webhook registered")
		return r.String(http.StatusOK, challenge)
	}
	return r.String(http.StatusOK, "Hello, World!")
}

func icmPayload(r echo.Context) error {
	jsonBody := new(helper.WhatsappBody)
	if err := r.Bind(&jsonBody); err != nil {
		log.Error("empty json body")
		return err
	}
	log.Error("success")
	fmt.Printf("%v", jsonBody)
	Parser(jsonBody)
	jsonStr, err := json.Marshal(jsonBody)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println("string object", string(jsonStr))
	}
	client := &http.Client{}
	if jsonBody != nil && jsonBody.Entry != nil && jsonBody.Entry[0].Changes != nil && jsonBody.Entry[0].Changes[0].Value.Statuses != nil && jsonBody.Entry[0].Changes[0].Value.Statuses[0].Status == "read" {
		reqObj := new(StatusReq)
		reqObj.MessagingProductReq = "whatsapp"
		reqObj.StatusReq = "read"
		_, found := database[jsonBody.Entry[0].Changes[0].Value.Statuses[0].WhatsappId]
		if found {
			reqObj.MessageId = database[jsonBody.Entry[0].Changes[0].Value.Statuses[0].WhatsappId]
			delete(database, jsonBody.Entry[0].Changes[0].Value.Statuses[0].WhatsappId)
		} else {
			reqObj.MessageId = ""
		}
		jsonStr1, err := json.Marshal(reqObj)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		} else {
			fmt.Println("string object", string(jsonStr1))
		}
		rq, _ := http.NewRequest(http.MethodPost, "https://graph.facebook.com/v15.0/105396722334906/messages", bytes.NewBuffer(jsonStr1))
		rq.Header.Add("Authorization", "Bearer EAAIglDnk6OoBADazOW3spDSR5zfm2RTC3THYUdcBSLPb4Gd7ijpePEPd1PqSNFmE4Las5kSTx70bW3bAZAZBs402owNkS4huV0M9OBA2yZCdVDiXyqrpp46W534s2bwm749PPYPEQ4N2ZAGPH48B07ItzSElu721RpWeLBaAfUAbLQrDuFpjUkxSk4nkZBQsVmrV3mHhyosYdQlghxZCSI")
		rq.Header.Add("Content-Type", "application/json")
		resp, err := client.Do(rq)
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(resp.Status, string(body))
		resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}

		var res map[string]interface{}

		json.NewDecoder(resp.Body).Decode(&res)

		fmt.Println(res["json"])
		fmt.Println("request success")
	}

	if jsonBody != nil && jsonBody.Entry != nil && jsonBody.Entry[0].Changes != nil && jsonBody.Entry[0].Changes[0].Value.Messages != nil {
		reqObj := new(MessageReq)
		if jsonBody.Entry[0].Changes[0].Value.Messages[0].From == "919390271925" {
			reqObj.To = "917095101959"
		} else if jsonBody.Entry[0].Changes[0].Value.Messages[0].From == "917095101959" {
			reqObj.To = "919390271925"
		}
		reqObj.MessagingProductReq = "whatsapp"
		reqObj.RecipientType = "individual"
		if jsonBody.Entry[0].Changes[0].Value.Messages[0].Type == "text" {
			reqObj.Type = jsonBody.Entry[0].Changes[0].Value.Messages[0].Type
			reqObj.TextReq.Body = jsonBody.Entry[0].Changes[0].Value.Messages[0].Text.Body
		} else {
			reqObj.Type = "text"
			reqObj.TextReq.Body = "start!"
		}
		reqObj.TextReq.PreviewURL = false

		jsonStr1, err := json.Marshal(reqObj)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		} else {
			fmt.Println("string object", string(jsonStr1))
		}
		rq, _ := http.NewRequest(http.MethodPost, "https://graph.facebook.com/v15.0/105396722334906/messages", bytes.NewBuffer(jsonStr1))
		rq.Header.Add("Authorization", "Bearer EAAIglDnk6OoBADazOW3spDSR5zfm2RTC3THYUdcBSLPb4Gd7ijpePEPd1PqSNFmE4Las5kSTx70bW3bAZAZBs402owNkS4huV0M9OBA2yZCdVDiXyqrpp46W534s2bwm749PPYPEQ4N2ZAGPH48B07ItzSElu721RpWeLBaAfUAbLQrDuFpjUkxSk4nkZBQsVmrV3mHhyosYdQlghxZCSI")
		rq.Header.Add("Content-Type", "application/json")
		resp, err := client.Do(rq)
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(resp.Status, string(body))
		if err != nil {
			log.Fatal(err)
		}

		res := new(Response)

		err = json.Unmarshal([]byte(string(body)), &res)
		if err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()
		fmt.Println(res.Messages[0].Id)
		fmt.Println("request success")
		database[res.Messages[0].Id] = jsonBody.Entry[0].Changes[0].Value.Messages[0].Id
		fmt.Println(res.Messages[0].Id, database[res.Messages[0].Id])
	}

	return r.String(http.StatusOK, fmt.Sprintf("Recieved!"))
}

func Parser(Body *helper.WhatsappBody) {
	fmt.Println("username  is ", Body.Object, Body.Entry[0].Changes[0].Value.MessagingProduct, Body.Entry[0].Id)
}
