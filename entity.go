package calltouch

type CallOrder struct {
	CompletedAmount int    `json:"completedAmount"` // Реальная сумма контракта
	CompletedDate   string `json:"completedDate"`   // Дата подписания контракта
	CreatedDate     string `json:"createdDate"`     // Дата создания контракта
	OrderDate       string `json:"orderDate"`       // Дата заказа
	OrderID         int    `json:"orderId"`         // ID сделки в Calltouch
	OrderNumber     string `json:"orderNumber"`     // Номер заказа
	PlannedAmount   int    `json:"plannedAmount"`   // Планируемая сумма контракта
}

type Tag struct {
	Category string   `json:"category"`
	Type     string   `json:"type"`
	Names    []string `json:"names"`
}

type Call struct {
	CallID          string         `json:"callId"`          // Уникальный идентификатор звонка в Calltouch.
	Callphase       string         `json:"callphase"`       // Фаза звонка на момент API запроса:
	Attribution     int            `json:"attribution"`     // Модель атрибуции звонков.
	CallTags        *[]Tag         `json:"callTags"`        // Теги звонков
	Date            string         `json:"date"`            // Дата звонка.
	Duration        string         `json:"duration"`        // Длительность разговора.
	CallerNumber    string         `json:"callerNumber"`    // Номер звонившего.
	RedirectNumber  string         `json:"redirectNumber"`  // Номер переадресации.
	PhoneNumber     string         `json:"phoneNumber"`     // Номер, на который звонили.
	Manager         string         `json:"manager"`         // ФИО менеджера, который был присвоен этому звонку с помощью API-метода присвоения менеджеров к лидам.
	Successful      bool           `json:"successful"`      // Статус звонка.
	UniqueCall      string         `json:"uniqueCall"`      // Уникальный звонок.
	TargetCall      string         `json:"targetCall"`      // Целевой звонок.
	UniqTargetCall  string         `json:"uniqTargetCall"`  // Уникально-целевой звонок.
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
	Orders          []CallOrder    `json:"orders"`          // Массив всех сделок, связанных со звонком.
	YandexDirect    *YandexDirect  `json:"yandexDirect"`    // Данные для Яндекс.Директа.
	GoogleAdWords   *GoogleAdWords `json:"googleAdWords"`   // Данные для GoogleAds.
	CallbackInfo    CallBackInfo   `json:"callbackInfo"`    // Данные (ФИО, email, номер телефона) из форм обратного звонка, оставленных в соц. сетях
	CtClientID      *int           `json:"ctClientId"`      // Идентификатор посетителя Calltouch. Он представляет из себя значение нашей куки _ct.
	DCM             *DCM           `json:"dcm"`             // Данные по отправке звонка с DoubleClick Campaign Manager.
	PhonesInText    *[]string      `json:"phonesInText"`    // Массив номеров телефонов, полученных из текста разговора (номера были произнесены в ходе разговора)
	CtGlobalID      *int           `json:"ctGlobalId"`      // Глобальный идентификатор посетителя Calltouch, общий для сайтов, на которых установлен скрипт Calltouch.
	SubPoolName     *string        `json:"subPoolName"`     // Название сабпула, с которым связан рекламный номер.
	StatusDetails   string         `json:"statusDetails"`   // Детализация статуса звонка.
}

type Lead struct {
	Date              int64          `json:"date"`              // Дата и время создания заявки в формате Unix Timestamp в миллисекундах. +
	Comments          []Comment      `json:"comments"`          // Комментарии к заявкам. +
	DateStr           string         `json:"dateStr"`           // Дата и время создания заявки в формате dd/mm/yyyy hh:mm:ss. -
	Manager           string         `json:"manager"`           // ФИО менеджера, который был присвоен заявке с помощью API-метода присвоения менеджеров к лидам +
	Session           Session        `json:"session"`           // Объект будет содержать вложенные объекты с описанием посещения, за которым закрепилась заявка.
	Subject           string         `json:"subject"`           // Название формы на Вашем сайте, которое Вы отправили в запросе.
	UniqTargetRequest bool           `json:"uniqTargetRequest"` // Уникально-целевая заявка.
	UniqueRequest     bool           `json:"uniqueRequest"`     // Уникальная заявка.
	YandexDirect      *YandexDirect  `json:"yandexDirect"`      // Данные для Яндекс.Директа.
	GoogleAdWords     *GoogleAdWords `json:"googleAdWords"`     // Данные для GoogleAds.
	RequestNumber     string         `json:"requestNumber"`     // Уникальный идентификатор заявки на Вашем сайте, который Вы отправили в запросе.
	RequestID         int            `json:"requestId"`         // Уникальный идентификатор заявки в Calltouch.
	Client            ClientInfo     `json:"client"`            // Объект будет содержать вложенные объекты с описанием клиента, данные по которому Вы отправили в запросе.
	Orders            []LeadOrder    `json:"orders"`            // Массив всех сделок, связанных с заявкой.
	TargetRequest     bool           `json:"targetRequest"`     // Целевая заявка.
	MapVisits         *[]MapVisits   `json:"mapVisits"`         // История посещений.
	CtClientID        *int64         `json:"ctClientId"`        // Идентификатор посетителя Calltouch. Он представляет из себя значение нашей куки _ct.
	DCM               *[]DCM         `json:"dcm"`               // Данные по отправке заявки с DoubleClick Campaign Manager.
	CtGlobalID        *int           `json:"ctGlobalId"`        // Глобальный идентификатор посетителя Calltouch, общий для сайтов, на которых установлен скрипт Calltouch.
	WidgetInfo        interface{}    `json:"widgetInfo"`        // Данные по кастомным полям заявки из виджета.
	RequestTags       *[]RequestTag  `json:"RequestTags"`       // Теги заявок
}

type Comment struct {
	CommentID int    `json:"commentId"` // ID комментария
	RequestID int    `json:"requestId"` // ID заявки, по которой оставлен комментарий
	Comment   string `json:"comment"`   // Комментарий
	PartyID   int    `json:"partyId"`   // ID пользователя, оставившего комментарий
	PartyName string `json:"partyName"` // Логин пользователя, оставившего комментарий
}

type Session struct {
	SessionID   int         `json:"sessionId"`   // Идентификатор сессии Calltouch, который Вы отправили в запросе ранее +
	Keywords    string      `json:"keywords"`    // Ключевой запрос +
	City        string      `json:"city"`        // Город посетителя (определяется по его IP-адресу)
	IP          string      `json:"ip"`          // IP-адрес
	Source      string      `json:"source"`      // Источник перехода
	Medium      string      `json:"medium"`      // Канал перехода
	Ref         string      `json:"ref"`         // Адрес страницы, с которой был совершен реферальный переход на Ваш отслеживаемый сайт
	URL         string      `json:"url"`         // Адрес входа на сайт (может отличаться от страницы, с которой в итоге был совершен звонок)
	UtmSource   string      `json:"utmSource"`   // Значение utm-метки utm_source
	UtmMedium   string      `json:"utmMedium"`   // Значение utm-метки utm_medium
	UtmTerm     string      `json:"utmTerm"`     // Значение utm-метки utm_term
	UtmContent  string      `json:"utmContent"`  // Значение utm-метки utm_content
	UtmCampaign string      `json:"utmCampaign"` // Значение utm-метки utm_campaign
	GuaClientID string      `json:"guaClientId"` // Идентификатор Google Client ID (присутствует, если настроена интеграция с Google Analytics)
	Attrs       interface{} `json:"attrs"`       // Сторонние параметры, переданные заранее в статистику Calltouch.
	Attribution int         `json:"attribution"` // Текущая модель атрибуции, согласно которой определился источник заявки
	YaClientID  string      `json:"yaClientId"`  // Идентификатор Yandex Client ID (присутствует, если настроена интеграция с Яндекс.Метрика)
	CtGlobalID  *int        `json:"ctGlobalId"`  // Глобальный идентификатор посетителя Calltouch, общий для сайтов, на которых установлен скрипт Calltouch
	Browser     string      `json:"browser"`     // Браузер
}

type ClientInfo struct {
	ClientID int    `json:"clientId"` // Идентификатор клиента внутри Calltouch
	Fio      string `json:"fio"`      // Имя клиента
	Phones   []struct {
		PhoneNumber string `json:"phoneNumber"` // Номер телефона клиента
		PhoneType   string `json:"phoneType"`
	} `json:"phones"`
	Contacts []struct {
		ContactType  string `json:"contactType"`
		ContactValue string `json:"contactValue"` // Почта клиента
	} `json:"contacts"`
}

type LeadOrder struct {
	OrderID     int    `json:"orderId"`     // Идентификатор сделки внутри Calltouch
	DateCreated int64  `json:"dateCreated"` // Дата и время создания сделки
	Status      string `json:"status"`      // Статус сделки
	Sum         string `json:"sum"`         // Бюджет сделки
	OrderNumber string `json:"orderNumber"` // Уникальный идентификатор сделки в CRM
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
	ProfileIDDCM              int     `json:"profileIdDCM"`              // Идентификатор профиля DCM
	FloodlightConfigurationID string  `json:"floodlightConfigurationId"` // Конфигурация из DCM
	FloodlightActivityID      string  `json:"floodlightActivityId"`      // Идентификатор флудлайта
	RequestStatus             string  `json:"requestStatus"`             // Статус отправки заявки в DCM.
	RequestErrors             *string `json:"requestErrors"`             // Описание ошибки отправки заявки в DCM
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
	Channel int    `json:"channel"` // Канал (1 - оператор, 0 - клиент)
	Time    string `json:"time"`    // Временная метка начала фразы в формате ММ:СС
	Message string `json:"message"` // Текст фразы
}

type CallBackInfo struct {
	Fields []ValueField `json:"fields"`
}

type ValueField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
