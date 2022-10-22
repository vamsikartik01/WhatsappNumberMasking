package helper

type WhatsappBody struct {
	Metadata   map[string]interface{} `json:"-"`
	WabaNumber string                 `queryparam:"waba_number" json:"waba_number,omitempty"`
	WabaId     string                 `json:"waba_id,omitempty"`
	Object     string                 `json:"object,omitempty"`
	Entry      []Entry                `json:"entry,omitempty"`
	Error      Error                  `json:"error,omitempty"`
}

type Entry struct {
	Id      string   `json:"id,omitempty"`
	Time    int64    `json:"time,omitempty"`
	Changes []Change `json:"changes,omitempty"`
}

type Change struct {
	Value Value  `json:"value,omitempty"`
	Field string `json:"field,omitempty"`
}

type Value struct {
	MessagingProduct             string    `json:"messaging_product,omitempty"`
	Metadata                     Metadata  `json:"metadata,omitempty"`
	Statuses                     []Status  `json:"statuses,omitempty" minlength:"1" validateeach:"true"`
	Contacts                     []Contact `json:"contacts,omitempty" minlength:"1" validateeach:"true"`
	Messages                     []Message `json:"messages,omitempty" minlength:"1" validateeach:"true"`
	MaxDailyConversationPerPhone int32     `json:"max_daily_conversation_per_phone,omitempty"`
	MaxPhoneNumbersPerBusiness   int32     `json:"max_phone_numbers_per_business,omitempty"`
	MaxPhoneNumbersPerWaba       int32     `json:"max_phone_numbers_per_waba,omitempty"`
	DisplayPhoneNumber           string    `json:"display_phone_number,omitempty"`
	Event                        string    `json:"event,omitempty"`
	CurrentLimit                 string    `json:"current_limit,omitempty"`
	OldLimit                     string    `json:"old_limit,omitempty"`
}

type Metadata struct {
	DisplayPhoneNumber string `json:"display_phone_number,omitempty"`
	PhoneNumberId      string `json:"phone_number_id,omitempty"`
}

type Status struct {
	Conversation Conversation `json:"conversation,omitempty"`
	WhatsappId   string       `json:"id,omitempty"`
	Pricing      Pricing      `json:"pricing,omitempty"`
	RecipientId  string       `json:"recipient_id,omitempty"`
	Status       string       `json:"status,omitempty"`
	TimeStamp    string       `json:"timestamp,omitempty"`
	Type         string       `json:"type,omitempty"`
	Errors       []Error      `json:"errors,omitempty"`
}

type Conversation struct {
	Id     string `json:"id,omitempty"`
	Origin Origin `json:"origin,omitempty"`
}

type Origin struct {
	Type string `json:"type,omitempty"`
}

type Pricing struct {
	Billable     bool   `json:"billable,omitempty"`
	Category     string `json:"category,omitempty"`
	PricingModel string `json:"pricing_model,omitempty"`
}

type Error struct {
	Code         int       `json:"code,omitempty"`
	Href         string    `json:"href,omitempty"`
	Title        string    `json:"title,omitempty"`
	Details      string    `json:"details,omitempty"`
	Message      string    `json:"message,omitempty"`
	Type         string    `json:"type,omitempty"`
	ErrorData    ErrorData `json:"error_data,omitempty"`
	ErrorSubcode int       `json:"error_subcode,omitempty"`
	FBTraceId    string    `json:"fbtrace_id,omitempty"`
}

type ErrorData struct {
	MessagingProduct string `json:"messaging_product,omitempty"`
	Details          string `json:"details,omitempty"`
}

type Contact struct {
	Profile    ProfileContent `json:"profile,omitempty"`
	WhatsappId string         `json:"wa_id,omitempty"`
	Addresses  []Address      `json:"addresses,omitempty"`
	Birthday   string         `json:"birthday,omitempty"`
	Emails     []Email        `json:"emails,omitempty"`
	Name       Name           `json:"name,omitempty"`
	Org        Org            `json:"org,omitempty"`
	Phones     []Phone        `json:"phones,omitempty"`
	URLs       []URL          `json:"urls,omitempty"`
	Input      string         `json:"input,omitempty"`
}

type ProfileContent struct {
	Name string `json:"name,omitempty"`
}

type Address struct {
	Street      string `json:"street,omitempty"`
	City        string `json:"city,omitempty"`
	State       string `json:"state,omitempty"`
	ZIP         string `json:"zip,omitempty"`
	Country     string `json:"country,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	Type        string `json:"type,omitempty"`
}

type Email struct {
	Email string `json:"email,omitempty"`
	Type  string `json:"type,omitempty"`
}

type Name struct {
	FormattedName string `json:"formatted_name,omitempty"`
	FirstName     string `json:"first_name,omitempty"`
	LastName      string `json:"last_name,omitempty"`
	MiddleName    string `json:"middle_name,omitempty"`
	Suffix        string `json:"suffix,omitempty"`
	Prefix        string `json:"prefix,omitempty"`
}

type Org struct {
	Company    string `json:"company,omitempty"`
	Department string `json:"department,omitempty"`
	Title      string `json:"title,omitempty"`
}

type Phone struct {
	Phone string `json:"phone,omitempty"`
	Type  string `json:"type,omitempty"`
	WAId  string `json:"wa_id,omitempty"`
}

type URL struct {
	URL  string `json:"url,omitempty"`
	Type string `json:"type,omitempty"`
}

type Message struct {
	Context     Context     `json:"context,omitempty"`
	From        string      `json:"from,omitempty"`
	Id          string      `json:"id,omitempty"`
	GroupId     string      `json:"group_id,omitempty"`
	Text        TextContent `json:"text,omitempty"`
	TimeStamp   string      `json:"timestamp,omitempty"`
	Type        string      `json:"type,omitempty"`
	GC          GC          `json:"gc,omitempty"`
	Image       Image       `json:"image,omitempty"`
	Audio       Audio       `json:"audio,omitempty"`
	Voice       Audio       `json:"voice,omitempty"`
	Video       Video       `json:"video,omitempty"`
	Sticker     Sticker     `json:"sticker,omitempty"`
	Document    Document    `json:"document,omitempty"`
	Location    Location    `json:"location,omitempty"`
	Contacts    []Contact   `json:"contacts,omitempty"`
	Interactive Interactive `json:"interactive,omitempty"`
	Errors      []Error     `json:"errors,omitempty"`
}

type TextContent struct {
	Body string `json:"body,omitempty"`
}

type Context struct {
	From                string   `json:"from,omitempty"`
	MessageId           string   `json:"id,omitempty"`
	Forwarded           bool     `json:"forwarded,omitempty"`
	FrequentlyForwarded bool     `json:"frequently_forwarded,omitempty"`
	GroupId             string   `json:"group_id,omitempty"`
	Mentions            []string `json:"mentions,omitempty"`
}

type GC struct {
	MessageDeleted int    `json:"messages_deleted,omitempty"`
	Node           string `json:"node,omitempty"`
	Status         string `json:"status,omitempty"`
}

type Image struct {
	MimeType string `json:"mime_type,omitempty"`
	SHA256   string `json:"sha256,omitempty"`
	Id       string `json:"id,omitempty"`
	Caption  string `json:"caption,omitempty"`
	File     string `json:"file,omitempty"`
}

type Audio struct {
	MimeType string `json:"mime_type,omitempty"`
	SHA256   string `json:"sha256,omitempty"`
	Id       string `json:"id,omitempty"`
	Voice    bool   `json:"voice,omitempty"`
	File     string `json:"file,omitempty"`
}

type Video struct {
	MimeType string `json:"mime_type,omitempty"`
	SHA256   string `json:"sha256,omitempty"`
	Id       string `json:"id,omitempty"`
	Status   string `json:"status,omitempty"`
}

type Sticker struct {
	MimeType string `json:"mime_type,omitempty"`
	SHA256   string `json:"sha256,omitempty"`
	Id       string `json:"id,omitempty"`
	Animated bool   `json:"animated,omitempty"`
	Status   string `json:"status,omitempty"`
}

type Document struct {
	Caption  string `json:"caption,omitempty"`
	FileName string `json:"filename,omitempty"`
	MimeType string `json:"mime_type,omitempty"`
	SHA256   string `json:"sha256,omitempty"`
	Id       string `json:"id,omitempty"`
	File     string `json:"file,omitempty"`
}

type Location struct {
	Address   string  `json:"address,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Name      string  `json:"name,omitempty"`
	URL       string  `json:"url,omitempty"`
}

type Interactive struct {
	Type        string      `json:"type,omitempty"`
	ListReply   ListReply   `json:"list_reply,omitempty"`
	ButtonReply ButtonReply `json:"button_reply,omitempty"`
}

type ListReply struct {
	Id          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type ButtonReply struct {
	Id    string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}
