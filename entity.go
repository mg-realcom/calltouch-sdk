package calltouch_sdk

type Order struct {
	CompletedAmount int    `json:"completedAmount"`
	CompletedDate   string `json:"completedDate"`
	CreatedDate     string `json:"createdDate"`
	OrderDate       string `json:"orderDate"`
	OrderId         int    `json:"orderId"`
	OrderNumber     string `json:"orderNumber"`
	PlannedAmount   int    `json:"plannedAmount"`
}

type Tag struct {
	Category string   `json:"category"`
	Type     string   `json:"type"`
	Names    []string `json:"names"`
}

type Call struct {
	Date            string   `json:"date"`
	CallUrl         string   `json:"callUrl"`
	City            string   `json:"city"`
	UniqueCall      bool     `json:"uniqueCall"`
	RedirectNumber  string   `json:"redirectNumber"`
	CallReferenceId string   `json:"callReferenceId"`
	UtmContent      string   `json:"utmContent"`
	YaClientId      string   `json:"yaClientId"`
	Source          string   `json:"source"`
	Medium          string   `json:"medium"`
	CallPhase       string   `json:"callphase"`
	Duration        int      `json:"duration"`
	Ref             string   `json:"ref"`
	AdditionalTags  []string `json:"additionalTags"`
	Hostname        string   `json:"hostname"`
	WaitingConnect  int      `json:"waitingConnect"`
	CtCallerId      string   `json:"ctCallerId"`
	CallbackCall    bool     `json:"callbackCall"`
	Keyword         string   `json:"keyword"`
	Successful      bool     `json:"successful"`
	Order           Order    `json:"order"`
	CallId          int      `json:"callId"`
	CallTags        []Tag    `json:"callTags"`
	UtmSource       string   `json:"utmSource"`
	SipCallId       string   `json:"sipCallId"`
	CallerNumber    string   `json:"callerNumber"`
	Ip              string   `json:"ip"`
	UtmTerm         string   `json:"utmTerm"`
	UserAgent       string   `json:"userAgent"`
	SessionId       int      `json:"sessionId"`
	UtmCampaign     string   `json:"utmCampaign"`
	Url             string   `json:"url"`
	Attrs           string   `json:"attrs"`
	PhoneNumber     string   `json:"phoneNumber"`
	UniqTargetCall  bool     `json:"uniqTargetCall"`
	TargetCall      bool     `json:"targetCall"`
	Attribution     int      `json:"attribution"`
	UtmMedium       string   `json:"utmMedium"`
	Orders          []Order  `json:"orders"`
}
