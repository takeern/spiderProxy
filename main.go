package main

import (
	"net"
	"log"
	"google.golang.org/grpc"

	pb "spiderProxy/interval/serve/grpc"
	"spiderProxy/interval/dao"
)

func main() {
	lis, err := net.Listen("tcp", ":2333")
	log.Printf("listen: 2333")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBookServer(s, &dao.Server{})
	if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
	}
}

// func main() {
// 	splitBookNumber := strings.Split("/d/179/179082/", "/")
// 	url := modal.SPIDER_URL + "down?id=" +splitBookNumber[3] + "&p=1"
// 	log.Printf(url)
// 	reader, err := dao.DownloadBook(url, 0)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Print("reader:", reader)

// 	for _, file := range reader.File {
// 		rc, err := file.Open()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		for {
// 			buf := make([]byte, 1024)
// 			_, err = rc.Read(buf)
// 			log.Print(buf)
// 			if err != nil {
// 				log.Println(err)
// 				break
// 			}
// 		}
// 	}
// }
