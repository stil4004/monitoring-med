package s_constant

import "time"

// db --------------------------------------------------------------------------------------------------------------

const (
	TokenTable              string = " market_db.public.token "
	CommissionTable         string = " market_db.public.commission "
	RequisiteTable          string = " market_db.public.requisite "
	OfferTable              string = " market_db.public.offer "
	PermissonsTable         string = " market_db.public.permissions "
	OrderTable              string = " market_db.public.order "
	MaskTable               string = " market_db.public.user_mask"
	OrderTransactionTable   string = " market_db.public.order_transaction "
	InnerConnectionTable    string = " market_api.public.inner_connection_m "
	PaymentTable            string = " market_db.public.payment "
	MerchantCommissionTable string = " market_db.public.merchant_commission "
	ReferralProgram         string = " market_db.public.referral "
	TransactionRegister     string = " market_db.public.transaction_register "
	WithdrawCondition       string = " market_db.public.withdraw_condition "
	TCryptoRegisterTable    string = " market_db.public.t_crypto_register "
	ReceiptTable            string = " market_db.public.receipt_analyse "
	TgReceiptDataTable      string = " market_db.public.tg_receipt_data "
	CancelOrderTable        string = " market_db.public.cancel_order "
	CancelReasonTable       string = " market_db.public.cancel_reason "
	UserReceiptInfoTable    string = " market_db.public.user_receipt_info "
)

// -----------------------------------------------------------------------------------------------------------------
// Commissions -----------------------------------------------------------------------------------------------------

var (
	DefaultCommissionRatio float64 = 0.002
	WithdrawUSDTCommission float64 = 2
)

// Limit amounts ---------------------------------------------------------------------------------------------------

const (
	LiquidityMaxAmount float64 = 5000000
	LimitMaxAmount     float64 = 5000000
	LimitMinBuyAmount  float64 = 1000
	LimitMinSellAmount float64 = 500
	LimitMax           float64 = 130
	LimitMin           float64 = 80
)

// -----------------------------------------------------------------------------------------------------------------
// Payments --------------------------------------------------------------------------------------------------------

const (
	OfferPaymentsMaxLenght int64 = 5
)

// -----------------------------------------------------------------------------------------------------------------

//todo

var P2PDnd_Id int64 = 3

const ( // TODO: Query --> Request

	BaseSiteUrl string = ""

	GetUserByIdQuery           string = "/user/service/by_id/%d"
	GetUserByAccessTokenQuery  string = "/user/service/by_access_token"
	GetUserByAccessTokenWithId string = "/user/by_access_token"

	GetOffersQuery              string = "offers"
	GetOfferQuery               string = "offers/%s"
	CreateOfferQuery            string = "offers"
	UnpauseRequisitesQuery      string = "requisites/%s/unpause"
	PauseRequisitesQuery        string = "requisites/%s/pause"
	GetOrderQuery               string = "orders/%s"
	UpdateRequisitesLimitsQuery string = "requisites/%s/limits"
	GetOrdersQuery              string = "orders/%s"
	GetOrdersByOfferQuery       string = "offers/%s/orders?limit=%s&offset=%s"
	ConfirmOrderQuery           string = "orders/%s/confirm"
	CancelOrderQuery            string = "orders/%s/cancel"
	CreateRequisiteQuery        string = "requisites"
	GetRequisitesQuery          string = "requisites?user_id=%d&limit=%d&offset=%d"
	CreateOrderQuery            string = "offers/%s/execute"
	CreateOrderQuery2           string = "order/create/%s"

	TrxBalance              string = "/balance?ServiceId=USDT"
	CreateInvoiceQuery      string = "/txs/invoice"
	CreateWithdrawQuery     string = "/txs/withdraw"
	GetBalanceQuery         string = "/balance?service_client=%s"
	GetBalanceByDates       string = "/client_balances_by_dates"
	ApproveTransactionQuery string = "/txs/%s/approve"
	CancelTransactionQuery  string = "/txs/%s/cancel"
	GetBalanceToTodayQuery  string = "/client_balance_to_today/%s"

	AssignDefendantQuery string = "/appeals/%s/defend"
	FetchAppealsQuery    string = "/appeals"
	CreateAppealsQuery   string = "/appeals"
	CloseAppealQuery     string = "/appeals/%s/close"
	GetAppealQuery       string = "/appeals/%s"

	TextSendMail string = "/send_email"

	ComCreateSession string = "/sessions"

	StartDebate    string = "admin/service/appeals/add"
	AddChatWithDnd string = "admin/service/add/chat_dnd"

	AdminNickname string = "admin"
	TehNickname   string = "info"

	SendBotText string = "https://api.telegram.org/bot"

	NewOrder string = "neworder"
	Apparel  string = "apparel"
)

var (
	Communication      string = "communication"
	TransactionService string = "transaction"
	DndApi             string = "dnd_api"
	Debate             string = "debate"
	Wallet             string = "wallet"
	Market             string = "market"
	Texts              string = "texts"
	Users              string = "users"
	EmailService       string = "email_service"
	TwkApi             string = "twk_api"
	AdminService       string = "admin"
	CDN                string = "cdn"
	OffersBot          string = "offers_bot"
	CdnUrl             string = "https://apiv4.exnode.ru/cdn/"
	Stock              string = "stock"
	Controller         string = "controller"
	Pc                 string = "paycraft"
	ReceiptAnalyzer    string = "receipt_analyzer"
)

var (
	False bool = false
	True  bool = true
)

var (
	Russia         string = "RU"
	DefaultCountry string = Russia
)

const ( // TODO: add all queries
	GetInnerConnectionQuery string = "SELECT * FROM market_db.public.inner_connection WHERE name = $1"
	GetServiceByPublicQuery string = "SELECT * FROM market_db.public.inner_connection WHERE public = $1"

	ShowAllActiveCurrencyMethodsQuery string = "SELECT * FROM market_db.public.currency_method WHERE currency_id = $1 AND is_active = true"

	GetOfferByInternalIdQuery  string = "SELECT * FROM market_db.public.offer WHERE internal_id = $1"
	GetOfferByDExternalIdQuery string = "SELECT * FROM market_db.public.offer WHERE d_external_id = $1"
	GetOfferByIdQuery          string = "SELECT * FROM market_db.public.offer WHERE id = $1"
)

var (
	OfferCreated         string = "CREATED"
	OfferTransactionSend string = "ACCEPTED"
	OfferSuccess         string = "SUCCESS"
)

var (
	OrderCreated         string = "CREATED"
	OrderFinishStart     string = "FINISHING"
	OrderTransactionSend string = "FIXED"
	OrderDndSend         string = "ACCEPTED"
	OrderChatCreated     string = "CHAT"
	OrderSuccess         string = "SUCCESS"

	OrderPending    string = "PENDING"
	OrderValidation string = "VALIDATING"

	OrderFirstApprove string = "1 APPROVE"

	OfferTackerApproved string = "TAKER APPROVED"
	OrderTakerApproved  string = "TAKER APPROVED"

	OfferDeleted string = "OFFER DELETED"
	OrderDeleted string = "ORDER DELETED"

	OrderTakerApprovedDndSend         string = "TAKER ACCEPTED"
	OrderTakerApprovedTransactionSend string = "TAKER FIXED"
	OrderMakerApprovedTransactionSend string = "MAKER FIXED"

	OrderCanceledStart                string = "CANCELLED START"
	OrderMakerCanceledTransactionSend string = "MAKER CANCELLED FIXED"
	OrderTakerCanceledTransactionSend string = "TAKER CANCELLED FIXED"
	OrderCanceledDnd                  string = "DND CANCELLED"
	OrderCanceled                     string = "CANCELLED"
	OrderClosed                       string = "CLOSED"
	OrderCancelViaNoLiquidity         string = "NO LIQUIDITY"

	OrderDone string = "DONE"
)

const (
	GetClientOrdersById string = "market/statistic/get_client_orders"
	VerifyTotp          string = "user/service/verify/totp/with_id"
	AddComment          string = "user/statistic/add_comment"

	SignUp string = "user/sign_up"
	SignIn string = "user/sign_in"
	Logout string = "user/logout"

	VerifyCode   string = "user/confirm/email/code/withdraw"
	VerifyAccess string = "user/token/valid_access"
	AddNotice    string = "user/service/add_notice"

	GetClientNickname string = "/user/service/get_client_nickname/%d"
)

var (
	ActiveStatus        string = "SUCCESS, 1 APPROVE"
	ActiveCreatedStatus string = "CREATED, SUCCESS, 1 APPROVE, PENDING"
	ConfirmWithdraw     string = "confirm_withdraw"
)

const (
	TimeMaker                         = int64(time.Minute) * 15
	TimeTaker                         = int64(time.Minute) * 30
	TimeAutoClose                     = time.Second * 150
	TimeAutoCloseCreatedOrders        = time.Minute * 30
	Timezone                   string = "MOS"
	TimeToLoadServiceData             = time.Minute * 1
)

const (
	Buy  string = "buy"
	Sell string = "sell"
)

// --------------------------------------------------------------------------------------------------------------------------------

const (
	RollbackCancelTypeInnerDB uint16 = uint16(iota + 1)
)

const (
	TypeBuyId  uint16 = 3
	TypeSellId uint16 = 4
)

// -------------------------------------------------------------------------------------------------------------------------------

const (
	OriginMarket int64 = 1
	DndOriginId  int64 = 2
)

const (
	TokenHandlerError string = "incorrect token"
)

const (
	ServiceTable string = " market_db.public.service "
)

const SetDebug bool = false

const (
	LogMilliSecondsAwaitTime = time.Millisecond * 50
)

const (
	OTCOrigin string = "otc"
	AOOrigin  string = "ao"
)

const (
	StatisticCache         string = "statCacheBy%s"
	UserCache              string = "userCacheBy%s"
	UserCacheShort         string = "userShortCacheBy%s"
	UserCacheOTCShort      string = "userShortCacheOTCBy%s"
	StatisticCacheLiveTime        = time.Minute * 1
	UserCacheLiveTime             = time.Minute * 1
)

const (
	MutexGroup1 uint64 = 1 // ["InitOrder"]
)

const (
	MakerRole string = "maker"
	TakerRole string = "taker"
)

const (
	InternalError string = "internal_error"
	OK            string = "ok"
	EmptyData            = "no data"
)
