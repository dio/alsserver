package als

import (
	"io"
	"log"

	v2 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v2"
	"github.com/gogo/protobuf/jsonpb"
)

type server struct {
	marshaler jsonpb.Marshaler
}

var _ v2.AccessLogServiceServer = &server{}

// New ...
func New() v2.AccessLogServiceServer {
	return &server{}
}

func (s *server) StreamAccessLogs(stream v2.AccessLogService_StreamAccessLogsServer) error {
	log.Println("Started stream")
	for {
		in, err := stream.Recv()
		log.Println("Received value")
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		str, _ := s.marshaler.MarshalToString(in)
		log.Println(str)
	}
}
