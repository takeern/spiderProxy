package modal

const SPIDER_URL = "https://www.aixdzs.com/"

const DESC_URL = SPIDER_URL + "bsearch?q="

const BOOK_READ_URL = "https://read.aixdzs.com/"

const HTTP_TRY_REQUEST_TIMES = 3

const (
	GETBOOKDESC	= iota
	GETBOOKLIST
	GETBOOKDATA
	GETBOOKALLDATA
)

type Server struct {}
