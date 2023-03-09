// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cdnv2iface provides an interface to enable mocking the cdnv2 service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cdnv2iface

import (
	"github.com/KscSDK/ksc-sdk-go/service/cdnv2"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

// Cdnv2API provides an interface to enable mocking the
// cdnv2.Cdnv2 service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // cdnv2.
//    func myFunc(svc cdnv2iface.Cdnv2API) bool {
//        // Make svc.GetClientRequestDataGet request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := cdnv2.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCdnv2Client struct {
//        cdnv2iface.Cdnv2API
//    }
//    func (m *mockCdnv2Client) GetClientRequestDataGet(input *map[string]interface{}) (*map[string]interface{}, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCdnv2Client{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type Cdnv2API interface {
	GetClientRequestDataGet(*map[string]interface{}) (*map[string]interface{}, error)
	GetClientRequestDataGetWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetClientRequestDataGetRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetClientRequestDataPost(*map[string]interface{}) (*map[string]interface{}, error)
	GetClientRequestDataPostWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetClientRequestDataPostRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetServerDataGet(*map[string]interface{}) (*map[string]interface{}, error)
	GetServerDataGetWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetServerDataGetRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GetServerDataPost(*map[string]interface{}) (*map[string]interface{}, error)
	GetServerDataPostWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetServerDataPostRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})
}

var _ Cdnv2API = (*cdnv2.Cdnv2)(nil)
