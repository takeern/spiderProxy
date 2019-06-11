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
	"strings"
	"archive/zip"
	"log"
	"bytes"

	"spiderProxy/interval/modal"
	pb "spiderProxy/interval/serve/grpc"
)

type Server struct{}

func init() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

func (s *Server) GetBookDesc(ctx context.Context, req *pb.GetBookDescReq) (*pb.GetBookDescResp, error) {
	params, err := GetBookDesc(req.BookName)
	var resp *pb.GetBookDescResp
	resp = new(pb.GetBookDescResp)

	for _, param := range params {
		resp.BooksDesc = append(resp.BooksDesc, &pb.BookDesc{
			BookName: string(param[2]),
			BookNumber: string(param[1]),
			BookIntro: string(param[4]),
			BookState: string(param[3]),
		})
	}
	return resp, err
}

func (s *Server) GetBookData(ctx context.Context, req *pb.GetBookDataReq) (*pb.GetBookDataResp, error) {
	var resp *pb.GetBookDataResp
	resp = new(pb.GetBookDataResp)
	bookData, err := GetBookData(req.BookHref, req.BookNumber)

	resp.BookData = bookData
	resp.BookNumber = req.BookNumber

	return resp, err
}

func (s *Server) GetBookList(ctx context.Context, req *pb.GetBookListReq) (*pb.GetBookListResp, error) {
	var resp *pb.GetBookListResp
	resp = new(pb.GetBookListResp)

	params, err := GetBookList(req.BookNumber)
	if err != nil {
		return resp, err
	}

	for _, param := range params {
		resp.BookList = append(resp.BookList, &pb.BookChapter{
			Href: string(param[1]),
			Length: string(param[2]),
			Title: string(param[3]),
		})
	}

	resp.BookNumber = req.BookNumber

	return resp, nil
}

func (s *Server) DownloadBook(req *pb.DownloadBookReq, stream pb.Book_DownloadBookServer) (error) {
	splitBookNumber := strings.Split(req.BookNumber, "/")
	url := modal.DESC_URL + "down?id=" +splitBookNumber[3] + "&p=1"
	
	reader, err := DownloadBook(url, 0)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range reader.File {
		rc, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		buf := make([]byte, 1024)
		_, err = rc.Read(buf)
		if err != nil {
			log.Println(err)
			break
		}

		stream.Send(&pb.DownloadBookResp{
			BookData: buf,
		})
	}
	return nil
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

func GetBookDesc(BookName string) ([][][]byte, error) {
	url := modal.DESC_URL + url.QueryEscape(BookName)
	html, err := getHtml(url, 0)
	var params [][][]byte

	if (len(html) == 0) {
		return params, errors.New("string length = 0")
	}

	if err != nil {
		return params, err
	}

	re := regexp.MustCompile(`b_name"><a href="(.{1,}?)"[\s\S]{1,}?"_blank">(.{1,}?)<[\s\S]{1,}?p">(.{1,}?)<[\s\S]{1,}?b_intro">([\s\S]{1,}?)<`)
	params = re.FindAllSubmatch([]byte(html), -1)
	return params, nil
}

func GetBookData(BookHref string, BookNumber string) (string, error) {
	var bookData string
	if (len(BookNumber) == 0) {
		return bookData, errors.New("string BookNumber undfined")
	}

	if (len(BookHref) == 0) {
		return bookData, errors.New("string BookHref undfined")
	}

	splitBookNumber := strings.Split(BookNumber, "/")
	url := modal.BOOK_READ_URL + splitBookNumber[2] + "/"+ splitBookNumber[3] + "/" + BookHref

	html, err := getHtml(url, 0)
	if err != nil {
		return bookData, err
	}

	re := regexp.MustCompile(`<div class="content">([\w\W]*?)<\/div>`)
	params := re.FindAllSubmatch([]byte(html), -1)

	return string(params[0][1]), nil
}

func GetBookList(BookNumber string) ([][][]byte, error) {
	splitBookNumber := strings.Split(BookNumber, "/")
	url := modal.BOOK_READ_URL + splitBookNumber[2] + "/"+ splitBookNumber[3]
	html, err := getHtml(url, 0)

	var params [][][]byte

	if err != nil {
		return params, err
	}

	re := regexp.MustCompile(`<li class="chapter"><a href="([\s\S]{1,}?)"[\s]?title="字数:([\d]{1,}?)">(.{1,}?)(<\/a)`)
	params = re.FindAllSubmatch([]byte(html), -1)
	return params, nil
}

func DownloadBook(url string, times int) (*zip.Reader, error) {
	var reader *zip.Reader
	if (times > 3) {
		return reader, errors.New("try many times")
	}
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		fmt.Println("http err", err.Error())
		times ++
		DownloadBook(url, times)
	}
	content_zipped, _ := ioutil.ReadAll(res.Body)

	reader, _ = zip.NewReader(bytes.NewReader(content_zipped), int64(len(content_zipped)))

	return reader, err
}