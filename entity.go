package calltouch

type Order struct {
	CompletedAmount int    `json:"completedAmount"`
	CompletedDate   string `json:"completedDate"`
	CreatedDate     string `json:"createdDate"`
	OrderDate       string `json:"orderDate"`
	OrderID         int    `json:"orderId"`
	OrderNumber     string `json:"orderNumber"`
	PlannedAmount   int    `json:"plannedAmount"`
}

type Tag struct {
	Category string   `json:"category"`
	Type     string   `json:"type"`
	Names    []string `json:"names"`
}

type Call struct {
	CallID          int            `json:"callId"`          // Уникальный идентификатор звонка в Calltouch.
	Callphase       string         `json:"callphase"`       // Фаза звонка на момент API запроса:
	Attribution     int            `json:"attribution"`     // Модель атрибуции звонков.
	CallTags        *[]Tag         `json:"callTags"`        // Теги звонков
	Date            string         `json:"date"`            // Дата звонка.
	Duration        int            `json:"duration"`        // Длительность разговора.
	CallerNumber    string         `json:"callerNumber"`    // Номер звонившего.
	RedirectNumber  string         `json:"redirectNumber"`  // Номер переадресации.
	PhoneNumber     string         `json:"phoneNumber"`     // Номер, на который звонили.
	Manager         string         `json:"manager"`         // ФИО менеджера, который был присвоен этому звонку с помощью API-метода присвоения менеджеров к лидам.
	Successful      bool           `json:"successful"`      // Статус звонка.
	UniqueCall      bool           `json:"uniqueCall"`      // Уникальный звонок.
	TargetCall      bool           `json:"targetCall"`      // Целевой звонок.
	UniqTargetCall  bool           `json:"uniqTargetCall"`  // Уникально-целевой звонок.
	CallbackCall    bool           `json:"callbackCall"`    // Обратный звонок.
	City            string         `json:"city"`            // Город, в котором находится клиент, совершивший звонок.
	Source          string         `json:"source"`          // Источник
	Medium          string         `json:"medium"`          // Тип трафика.
	Keyword         string         `json:"keyword"`         // Ключевое слово.
	URL             string         `json:"url"`             // Адрес, по которому попали на сайт.
	CallURL         string         `json:"callUrl"`         // Адрес страницы, находясь на которой, посетитель совершил звонок.
	Ref             string         `json:"ref"`             // Адрес, с которого перешли на сайт.
	Hostname        string         `json:"hostname"`        // Отслеживаемый домен или поддомен ресурса, на который был осуществлен переход (например: yoursite.ru).
	UtmSource       string         `json:"utmSource"`       // Значение utm-метки utm_source.
	UtmMedium       string         `json:"utmMedium"`       // Значение utm-метки utm_medium.
	UtmCampaign     string         `json:"utmCampaign"`     // Значение utm-метки utm_campaign.
	UtmContent      string         `json:"utmContent"`      // Значение utm-метки utm_content.
	UtmTerm         string         `json:"utmTerm"`         // Значение utm-метки utm_term.
	SessionID       int            `json:"sessionId"`       // Уникальный идентификатор сессии Calltouch.
	CtCallerID      string         `json:"ctCallerId"`      // Уникальный идентификатор клиента в Calltouch.
	ClientID        *string        `json:"clientId"`        //Уникальный идентификатор Universal Analytics.
	YaClientID      *string        `json:"yaClientId"`      //Уникальный идентификатор Яндекс.Метрика.
	SIPCallID       string         `json:"sipCallId"`       // Уникальный идентификатор сеанса связи с АТС Calltouch.
	UserAgent       string         `json:"userAgent"`       // Информация об устройстве, с которого зашли на сайт.
	IP              string         `json:"ip"`              // IP-адрес посетителя.
	WaitingConnect  int            `json:"waitingConnect"`  // Время ожидания ответа.
	CallReferenceID string         `json:"callReferenceId"` // Уникальный ID звонка с Вашей АТС, переданный в параметре callid API-запроса
	MapVisits       *[]MapVisits   `json:"mapVisits"`       // История посещений.
	Attrs           interface{}    `json:"attrs"`           // Сторонние параметры, переданные заранее в статистику Calltouch.
	Comments        *[]Comment     `json:"comments"`        // Комментарии к звонкам, оставленные в плеере журнала звонков.
	Phrases         *[]Phrase      `json:"phrases"`         // массив фраз из звонка (соблюдая последовательность)
	AdditionalTags  []ValueField   `json:"additionalTags"`  // Дополнительные параметры отслеживания платного трафика.
	Orders          []Order        `json:"orders"`          // Массив всех сделок, связанных со звонком.
	YandexDirect    *YandexDirect  `json:"yandexDirect"`    // Данные для Яндекс.Директа.
	GoogleAdWords   *GoogleAdWords `json:"googleAdWords"`   // Данные для GoogleAds.
	CallbackInfo    CallBackInfo   `json:"callbackInfo"`    // Данные (ФИО, email, номер телефона) из форм обратного звонка, оставленных в соц. сетях
	CtClientID      *int           `json:"ctClientId"`      // Идентификатор посетителя Calltouch. Он представляет из себя значение нашей куки _ct.
	DCM             *DCM           `json:"dcm"`             // Данные по отправке звонка с DoubleClick Campaign Manager.
	PhonesInText    *[]string      `json:"phonesInText"`    //Массив номеров телефонов, полученных из текста разговора (номера были произнесены в ходе разговора)
	CtGlobalID      *int           `json:"ctGlobalId"`      // Глобальный идентификатор посетителя Calltouch, общий для сайтов, на которых установлен скрипт Calltouch.
	SubPoolName     *string        `json:"subPoolName"`     // Название сабпула, с которым связан рекламный номер.
	StatusDetails   string         `json:"statusDetails"`   // Детализация статуса звонка.
}

type Lead struct {
	Date              int64          `json:"date"`
	Comments          []Comment      `json:"comments"`
	RequestType       string         `json:"requestType"`
	DateStr           string         `json:"dateStr"`
	Manager           string         `json:"manager"`
	Session           Session        `json:"session"`
	Subject           string         `json:"subject"`
	UniqTargetRequest bool           `json:"uniqTargetRequest"`
	UniqueRequest     bool           `json:"uniqueRequest"`
	YandexDirect      *YandexDirect  `json:"yandexDirect"`
	GoogleAdWords     *GoogleAdWords `json:"googleAdWords"`
	RequestNumber     string         `json:"requestNumber"`
	RequestID         int            `json:"requestId"`
	Client            ClientInfo     `json:"client"`
	SiteID            int            `json:"siteId"`
	Orders            []LeadOrder    `json:"orders"`
	TargetRequest     bool           `json:"targetRequest"`
	Status            string         `json:"status"`
	Order             interface{}    `json:"order"`
	MapVisits         []MapVisits    `json:"mapVisits"`
	RequestURL        *string        `json:"requestUrl"`
	CtClientID        *int64         `json:"ctClientId"`
	DCM               *[]DCM         `json:"dcm"`
	CtGlobalID        *int           `json:"ctGlobalId"`
	WidgetInfo        interface{}    `json:"widgetInfo"`
	RequestTags       *[]RequestTag  `json:"RequestTags"`
}

type Comment struct {
	CommentID int    `json:"commentId"`
	RequestID int    `json:"requestId"`
	Comment   string `json:"comment"`
	PartyID   int    `json:"partyId"`
	PartyName string `json:"partyName"`
}

type Session struct {
	SessionID      int           `json:"sessionId"`
	Keywords       string        `json:"keywords"`
	City           string        `json:"city"`
	IP             string        `json:"ip"`
	Source         string        `json:"source"`
	Medium         string        `json:"medium"`
	Ref            string        `json:"ref"`
	URL            string        `json:"url"`
	UtmSource      string        `json:"utmSource"`
	UtmMedium      string        `json:"utmMedium"`
	UtmTerm        string        `json:"utmTerm"`
	UtmContent     string        `json:"utmContent"`
	UtmCampaign    string        `json:"utmCampaign"`
	GuaClientID    string        `json:"guaClientId"`
	SessionDate    string        `json:"sessionDate"`
	Attrs          interface{}   `json:"attrs"`
	Attribution    int           `json:"attribution"`
	YaClientID     string        `json:"yaClientId"`
	AdditionalTags []interface{} `json:"additionalTags"`
	CtGlobalID     *int          `json:"ctGlobalId"`
	Browser        string        `json:"browser"`
}

type ClientInfo struct {
	ClientID int    `json:"clientId"`
	Fio      string `json:"fio"`
	Phones   []struct {
		PhoneNumber string `json:"phoneNumber"`
		PhoneType   string `json:"phoneType"`
	} `json:"phones"`
	Contacts []struct {
		ContactType  string `json:"contactType"`
		ContactValue string `json:"contactValue"`
	} `json:"contacts"`
}

type LeadOrder struct {
	OrderID       int         `json:"orderId"`
	CallID        *int        `json:"callId"`
	DateCreated   int64       `json:"dateCreated"`
	Status        string      `json:"status"`
	RealSum       string      `json:"realSum"`
	Offered       interface{} `json:"offered"`
	Sent          string      `json:"sent"`
	Sum           string      `json:"sum"`
	IsMarked      interface{} `json:"isMarked"`
	CommentsCount int         `json:"commentsCount"`
	CurrentAmount int         `json:"currentAmount"`
	OrderNumber   string      `json:"orderNumber"`
	OrderSum      string      `json:"orderSum"`
	OrderStatus   string      `json:"orderStatus"`
	OrderComments string      `json:"orderComments"`
	Session       struct {
		Keywords       string        `json:"keywords"`
		City           string        `json:"city"`
		IP             string        `json:"ip"`
		Browser        string        `json:"browser"`
		Source         string        `json:"source"`
		Medium         string        `json:"medium"`
		Ref            string        `json:"ref"`
		URL            string        `json:"url"`
		UtmSource      string        `json:"utmSource"`
		UtmMedium      string        `json:"utmMedium"`
		UtmTerm        string        `json:"utmTerm"`
		UtmContent     string        `json:"utmContent"`
		UtmCampaign    string        `json:"utmCampaign"`
		GuaClientID    interface{}   `json:"guaClientId"`
		YaClientID     interface{}   `json:"yaClientId"`
		SessionID      int           `json:"sessionId"`
		AdditionalTags []interface{} `json:"additionalTags"`
		Attribution    int           `json:"attribution"`
		Attrs          interface{}   `json:"attrs"`
	} `json:"session"`
}

type YandexDirect struct {
	CampaignID int   `json:"campaignId"`
	AdGroupID  int64 `json:"adGroupId"`
	AdID       int64 `json:"adId"`
	CriteriaID int   `json:"criteriaId"`
}

type GoogleAdWords struct {
	CampaignID int   `json:"campaignId"`
	AdGroupID  int   `json:"adGroupId"`
	CreativeID int   `json:"creativeId"`
	CriteriaID int64 `json:"criteriaId"`
}

type DCM struct {
	ProfileIDDCM              int    `json:"profileIdDCM"`
	FloodlightConfigurationID string `json:"floodlightConfigurationId"`
	FloodlightActivityID      string `json:"floodlightActivityId"`
	RequestStatus             string `json:"requestStatus"`
	RequestErrors             string `json:"requestErrors"`
}

type MapVisits struct {
	UtmSource      string        `json:"utmSource"`
	SessionDate    string        `json:"sessionDate"`
	City           string        `json:"city"`
	IP             string        `json:"ip"`
	UtmTerm        string        `json:"utmTerm"`
	UtmContent     string        `json:"utmContent"`
	UserAgent      string        `json:"userAgent"`
	SessionID      int           `json:"sessionId"`
	Source         string        `json:"source"`
	Medium         string        `json:"medium"`
	UtmCampaign    string        `json:"utmCampaign"`
	URL            string        `json:"url"`
	Ref            string        `json:"ref"`
	AdditionalTags []interface{} `json:"additionalTags"`
	UtmMedium      string        `json:"utmMedium"`
	GuaClientID    string        `json:"guaClientId"`
	Keyword        string        `json:"keyword"`
	CtClientID     int64         `json:"ctClientId"`
	CtGlobalID     *int          `json:"ctGlobalId"`
}

type RequestTag []struct {
	Category string   `json:"category"`
	Type     string   `json:"type"`
	Names    []string `json:"names"`
}

type Phrase struct {
	Channel int    `json:"channel"`
	Time    string `json:"time"`
	Message string `json:"message"`
}

type CallBackInfo struct {
	Fields []ValueField `json:"fields"`
}

type ValueField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
