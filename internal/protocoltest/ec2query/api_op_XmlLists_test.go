// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2query

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/ec2query/types"
	smithydocument "github.com/aws/smithy-go/document"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyrand "github.com/aws/smithy-go/rand"
	smithytesting "github.com/aws/smithy-go/testing"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"io/ioutil"
	"math"
	"net/http"
	"testing"
	"time"
)

func TestClient_XmlLists_awsEc2queryDeserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *XmlListsOutput
	}{
		// Tests for XML list serialization
		"Ec2XmlLists": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"text/xml;charset=UTF-8"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<XmlListsResponse xmlns="https://example.com/">
			    <stringList>
			        <member>foo</member>
			        <member>bar</member>
			    </stringList>
			    <stringSet>
			        <member>foo</member>
			        <member>bar</member>
			    </stringSet>
			    <integerList>
			        <member>1</member>
			        <member>2</member>
			    </integerList>
			    <booleanList>
			        <member>true</member>
			        <member>false</member>
			    </booleanList>
			    <timestampList>
			        <member>2014-04-29T18:30:38Z</member>
			        <member>2014-04-29T18:30:38Z</member>
			    </timestampList>
			    <enumList>
			        <member>Foo</member>
			        <member>0</member>
			    </enumList>
			    <intEnumList>
			        <member>1</member>
			        <member>2</member>
			    </intEnumList>
			    <nestedStringList>
			        <member>
			            <member>foo</member>
			            <member>bar</member>
			        </member>
			        <member>
			            <member>baz</member>
			            <member>qux</member>
			        </member>
			    </nestedStringList>
			    <renamed>
			        <item>foo</item>
			        <item>bar</item>
			    </renamed>
			    <flattenedList>hi</flattenedList>
			    <flattenedList>bye</flattenedList>
			    <customName>yep</customName>
			    <customName>nope</customName>
			    <flattenedListWithMemberNamespace xmlns="https://xml-member.example.com">a</flattenedListWithMemberNamespace>
			    <flattenedListWithMemberNamespace xmlns="https://xml-member.example.com">b</flattenedListWithMemberNamespace>
			    <flattenedListWithNamespace>a</flattenedListWithNamespace>
			    <flattenedListWithNamespace>b</flattenedListWithNamespace>
			    <myStructureList>
			        <item>
			            <value>1</value>
			            <other>2</other>
			        </item>
			        <item>
			            <value>3</value>
			            <other>4</other>
			        </item>
			    </myStructureList>
			    <RequestId>requestid</RequestId>
			</XmlListsResponse>
			`),
			ExpectResult: &XmlListsOutput{
				StringList: []string{
					"foo",
					"bar",
				},
				StringSet: []string{
					"foo",
					"bar",
				},
				IntegerList: []int32{
					1,
					2,
				},
				BooleanList: []bool{
					true,
					false,
				},
				TimestampList: []time.Time{
					smithytime.ParseEpochSeconds(1398796238),
					smithytime.ParseEpochSeconds(1398796238),
				},
				EnumList: []types.FooEnum{
					types.FooEnum("Foo"),
					types.FooEnum("0"),
				},
				IntEnumList: []types.IntegerEnum{
					1,
					2,
				},
				NestedStringList: [][]string{
					{
						"foo",
						"bar",
					},
					{
						"baz",
						"qux",
					},
				},
				RenamedListMembers: []string{
					"foo",
					"bar",
				},
				FlattenedList: []string{
					"hi",
					"bye",
				},
				FlattenedList2: []string{
					"yep",
					"nope",
				},
				FlattenedListWithMemberNamespace: []string{
					"a",
					"b",
				},
				FlattenedListWithNamespace: []string{
					"a",
					"b",
				},
				StructureList: []types.StructureListMember{
					{
						A: ptr.String("1"),
						B: ptr.String("2"),
					},
					{
						A: ptr.String("3"),
						B: ptr.String("4"),
					},
				},
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			serverURL := "http://localhost:8888/"
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					headers := http.Header{}
					for k, vs := range c.Header {
						for _, v := range vs {
							headers.Add(k, v)
						}
					}
					if len(c.BodyMediaType) != 0 && len(headers.Values("Content-Type")) == 0 {
						headers.Set("Content-Type", c.BodyMediaType)
					}
					response := &http.Response{
						StatusCode: c.StatusCode,
						Header:     headers,
						Request:    r,
					}
					if len(c.Body) != 0 {
						response.ContentLength = int64(len(c.Body))
						response.Body = ioutil.NopCloser(bytes.NewReader(c.Body))
					} else {

						response.Body = http.NoBody
					}
					return response, nil
				}),
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						s.Initialize.Remove(`OperationInputValidation`)
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = serverURL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			var params XmlListsInput
			result, err := client.XmlLists(context.Background(), &params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			opts := cmp.Options{
				cmpopts.IgnoreUnexported(
					middleware.Metadata{},
				),
				cmp.FilterValues(func(x, y float64) bool {
					return math.IsNaN(x) && math.IsNaN(y)
				}, cmp.Comparer(func(_, _ interface{}) bool { return true })),
				cmp.FilterValues(func(x, y float32) bool {
					return math.IsNaN(float64(x)) && math.IsNaN(float64(y))
				}, cmp.Comparer(func(_, _ interface{}) bool { return true })),
				cmpopts.IgnoreTypes(smithydocument.NoSerde{}),
			}
			if err := smithytesting.CompareValues(c.ExpectResult, result, opts...); err != nil {
				t.Errorf("expect c.ExpectResult value match:\n%v", err)
			}
		})
	}
}
