package main

import (
	"crypto/tls"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"sort"

	pb "github.com/sssergei/BaseService/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var target = flag.String("l", ":7100", "Specify the port that the server is listening on")
var username = flag.String("u", "scaramoucheX2", "Specify the username to auth the call")
var password = flag.String("p", "Can-you-do-the-fandango?", "Specify the password to auth the call")

type byVersion []*pb.ReleaseInfo

func (b byVersion) Len() int           { return len(b) }
func (b byVersion) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b byVersion) Less(i, j int) bool { return b[i].GetVersion() < b[j].GetVersion() }

func main() {

	flag.Parse()

	// create gRPC TLS credentials
	creds := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true, // using self signed certificate for demo, for more secure connections see https://bbengfort.github.io/programmer/2017/03/03/secure-grpc.html
	})

	grpcAuth := &basicAuthCreds{
		username: *username,
		password: *password,
	}

	conn, err := grpc.Dial(*target,
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(grpcAuth),
	)
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}

	client := pb.NewGoReleaseServiceClient(conn)

	ctx := context.Background()
	rsp, err := client.ListReleases(ctx, &pb.ListReleasesRequest{})

	if err != nil {
		log.Fatalf("ListReleases err: %v", err)
	}

	releases := rsp.GetReleases()
	if len(releases) > 0 {
		sort.Sort(byVersion(releases))

		fmt.Printf("Version\tRelease Date\tRelease Notes\n")
	} else {
		fmt.Println("No releases found")
	}
	for _, ri := range releases {
		//fmt.Printf("%s\t%s\t%s\n", ri.GetVersion(), ri.GetReleaseDate(), ri.GetReleaseNotesURL())
		fmt.Printf("%s\t%s\n", ri.GetVersion(), ri.GetReleaseDate())
	}

	rsp1, err := client.GetReleaseInfo(ctx, &pb.GetReleaseInfoRequest{Version: "1.5"})
	if err != nil {
		log.Fatalf("ReleaseInfo err: %v", err.Error())
	}
	fmt.Println("%s\t%s\n", rsp1.GetVersion(), rsp1.GetReleaseDate())

	rsp2, err := client.SayHello(ctx, &pb.SayHelloRequest{Name: "Jonh"})
	if err != nil {
		log.Fatalf("SayHello err: %v", err.Error())
	}

	fmt.Println("%s\n", rsp2.GetMessage())

	rsp3, err := client.InsertUser(ctx, &pb.InsertUserRequest{Id: "9", Name: "Jonh7", Surname: "surn name7", Othername: "Other name7"})
	if err != nil {
		log.Fatalf("InsertUser err: %v", err)
	}
	fmt.Println("%s\n", rsp3.GetId())

	rsp4, err := client.DeleteUser(ctx, &pb.DeleteUserRequest{Id: "7"})
	if err != nil {
		log.Fatalf("ListReleases err: %v", err)
	}
	fmt.Println("%s\n", rsp4.GetId())
}

// basicAuthCreds is an implementation of credentials.PerRPCCredentials
// that transforms the username and password into a base64 encoded value similar
// to HTTP Basic xxx
type basicAuthCreds struct {
	username, password string
}

// GetRequestMetadata sets the value for "authorization" key
func (b *basicAuthCreds) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "Basic " + basicAuth(b.username, b.password),
	}, nil
}

// RequireTransportSecurity should be true as even though the credentials are base64, we want to have it encrypted over the wire.
func (b *basicAuthCreds) RequireTransportSecurity() bool {
	return true
}

//helper function
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
