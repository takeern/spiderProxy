package dao

import(
	"io/ioutil"
	"crypto/tls"
	"net/url"
	"regexp"
	"errors"
	"net/http"
	"fmt"
	"context"

	"spiderProxy/interval/modal"
	pb "spiderProxy/interval/grpc"
)

type Server struct{}

type BookDesc struct {
	BookName             string   `protobuf:"bytes,1,opt,name=BookName,proto3" json:"BookName,omitempty"`
	BookState            string   `protobuf:"bytes,2,opt,name=BookState,proto3" json:"BookState,omitempty"`
	BookIntro            string   `protobuf:"bytes,3,opt,name=BookIntro,proto3" json:"BookIntro,omitempty"`
	BookNumber           string   `protobuf:"bytes,4,opt,name=BookNumber,proto3" json:"BookNumber,omitempty"`
}

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

// func getUrl (type string, url string)
// 解析html
func getHtml(url string, times int) (string, error) {
	if (times > modal.HTTP_TRY_REQUEST_TIMES) {
		return "", errors.New("try many times")
	}

	if (times != 0) {
		fmt.Println("http req try %d", times)
	}

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("http err", err.Error())
		times ++
		return getHtml(url, times)
	}

	defer res.Body.Close()
	Body, err := ioutil.ReadAll(res.Body)

	return string(Body), err
} 

func (s *Server) GetBookDesc(ctx context.Context, req *pb.GetBookDescReq) (*pb.GetBookDescResp, error) {
	url := "https://www.aixdzs.com/bsearch?q=" + url.QueryEscape(req.BookName)
	html, err := getHtml(url, 0)
	var resp *pb.GetBookDescResp

	if (len(html) == 0) {
		return resp, errors.New("string length = 0")
	}

	if err != nil {
		return resp, err
	}

	re := regexp.MustCompile(`b_name"><a href="(.{1,}?)"[\s\S]{1,}?"_blank">(.{1,}?)<[\s\S]{1,}?p">(.{1,}?)<[\s\S]{1,}?b_intro">([\s\S]{1,}?)<`)
	params := re.FindAllSubmatch([]byte(html), -1)
	for _, param := range params {
		resp.BooksDesc = append(resp.BooksDesc, &pb.BookDesc{
			BookName: string(param[2]),
			BookNumber: string(param[1]),
			BookIntro: string(param[4]),
			BookState: string(param[3]),
		})
	}
	return resp, nil
}