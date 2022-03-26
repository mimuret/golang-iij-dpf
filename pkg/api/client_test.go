package api_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	. "github.com/mimuret/golang-iij-dpf/pkg/testtool"
)

var _ api.SearchParams = &TestParams{}

type TestParams struct {
	Values url.Values
	Err    error
}

func (t *TestParams) GetValues() (url.Values, error) {
	return t.Values, t.Err
}

func (t *TestParams) GetOffset() int32 {
	return 0
}

func (t *TestParams) SetOffset(int32) {
}

func (t *TestParams) GetLimit() int32 {
	return 100
}

func (t *TestParams) SetLimit(int32) {
}

type ErrSpec struct {
	TestSpec
	ErrMarshalJSON   error
	ErrUnMarshalJSON error
}

func (t *ErrSpec) MarshalJSON() ([]byte, error) {
	if t.ErrMarshalJSON != nil {
		return nil, t.ErrMarshalJSON
	}
	return json.Marshal(t.TestSpec)
}

func (t *ErrSpec) UnmarshalJSON(bs []byte) error {
	if t.ErrUnMarshalJSON != nil {
		return t.ErrUnMarshalJSON
	}
	return json.Unmarshal(bs, &t.TestSpec)
}

var _ = Describe("Client", func() {
	var (
		testSpec      *TestSpec
		errSpec       *ErrSpec
		listTestSpec  *TestSpecList
		countableList *TestSpecCountableList
		c             *api.Client
		reqId         string
		err           error
	)
	BeforeEach(func() {
		testSpec = &TestSpec{}
		errSpec = &ErrSpec{}
		listTestSpec = &TestSpecList{}
		countableList = &TestSpecCountableList{}
		c = api.NewClient("token", "http://localhost", nil)
		httpmock.RegisterNoResponder(httpmock.NewBytesResponder(404, []byte(`{
				"request_id": "F96C25C6B13E49E59F0093BCD8D731AB",
				"error_type": "NotFound",
				"error_message": "Specified resource not found."
			}`)))
	})
	Context("NewClient", func() {
		BeforeEach(func() {
			c = api.NewClient("", "", nil)
		})
		It("Endpoint = DefaultEndpoint", func() {
			Expect(c.Endpoint).To(Equal(api.DefaultEndpoint))
		})
	})
	Context("Do", func() {
		BeforeEach(func() {
			httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests/F2246BD8617444329A40470AEC7B00B9", httpmock.NewBytesResponder(200, []byte(`{
				"request_id": "0CAC3AEFB6334ECCA7B70AF76D73508B",
				"result": {
					"id": "F2246BD8617444329A40470AEC7B00B9",
					"name": "test1",
					"number": 1
				}
			}`)))
		})
		AfterEach(func() {
			httpmock.Reset()
		})
		When("the API succesfully", func() {
			BeforeEach(func() {
				testSpec.ID = "F2246BD8617444329A40470AEC7B00B9"
				reqId, err = c.Do(context.Background(), testSpec, api.ActionRead, nil, nil)
			})
			It("reqId is not empty", func() {
				Expect(reqId).Should(Equal("0CAC3AEFB6334ECCA7B70AF76D73508B"))
			})
			It("err is empty", func() {
				Expect(err).To(Succeed())
			})
		})
		When("action is ActionCount", func() {
			When("spec is not CountableListSpec", func() {
				BeforeEach(func() {
					reqId, err = c.Do(context.Background(), listTestSpec, api.ActionCount, nil, nil)
				})
				It("reqId is empty", func() {
					Expect(reqId).Should(Equal(""))
				})
				It("err is not empty", func() {
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(MatchRegexp("spec is not CountableListSpec"))
				})
			})
		})
		When("spec.GetPathMethod is empty", func() {
			BeforeEach(func() {
				reqId, err = c.Do(context.Background(), listTestSpec, api.ActionApply, nil, nil)
			})
			It("reqId is empty", func() {
				Expect(reqId).Should(Equal(""))
			})
			It("err is not empty", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("not support action"))
			})
		})
		When("SearchParams.GetValues() is error", func() {
			BeforeEach(func() {
				param := &TestParams{Err: fmt.Errorf("err")}
				reqId, err = c.Do(context.Background(), testSpec, api.ActionApply, nil, param)
			})
			It("reqId is empty", func() {
				Expect(reqId).Should(Equal(""))
			})
			It("err is not empty", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("failed to get search params:"))
			})
		})
		When("body is not nil, but action is not any of ActionCreate, ActionUpdate and ActionApply.", func() {
			BeforeEach(func() {
				reqId, err = c.Do(context.Background(), testSpec, api.ActionRead, testSpec, nil)
			})
			It("reqId is empty", func() {
				Expect(reqId).Should(Equal(""))
			})
			It("err is not empty", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("not support action"))
			})
		})
		When("action is ActionCreate. body is not nil,but failed to encode to json", func() {
			BeforeEach(func() {
				errSpec.ErrMarshalJSON = fmt.Errorf("fail")
				reqId, err = c.Do(context.Background(), errSpec, api.ActionCreate, errSpec, nil)
			})
			It("reqId is empty", func() {
				Expect(reqId).Should(Equal(""))
			})
			It("err is not empty", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("failed to encode body to json"))
			})
		})
		When("action is ActionUpdate. body is not nil. but but failed to encode to json", func() {
			BeforeEach(func() {
				errSpec.ErrMarshalJSON = fmt.Errorf("fail")
				reqId, err = c.Do(context.Background(), errSpec, api.ActionUpdate, errSpec, nil)
			})
			It("reqId is empty", func() {
				Expect(reqId).Should(Equal(""))
			})
			It("err is not empty", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("failed to encode body to json"))
			})
		})
		When("action ActionApply. body is not nil. but failed to encode to json", func() {
			BeforeEach(func() {
				errSpec.ErrMarshalJSON = fmt.Errorf("fail")
				reqId, err = c.Do(context.Background(), errSpec, api.ActionApply, errSpec, nil)
			})
			It("reqId is empty", func() {
				Expect(reqId).Should(Equal(""))
			})
			It("err is not empty", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("failed to encode body to json"))
			})
		})
		When("endpoint url is invalid", func() {
			BeforeEach(func() {
				cl := api.NewClient("", "", nil)
				cl.Endpoint = "\n"
				reqId, err = cl.Do(context.Background(), errSpec, api.ActionRead, nil, nil)
			})
			It("reqId is empty", func() {
				Expect(reqId).Should(Equal(""))
			})
			It("err is not empty", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("failed to create http request"))
			})
		})
		When("failed to get http response", func() {
			BeforeEach(func() {
				httpmock.Reset()
				cl := api.NewClient("", "https://localhost/hoge", nil)
				reqId, err = cl.Do(context.Background(), errSpec, api.ActionRead, nil, nil)
			})
			It("reqId is empty", func() {
				Expect(reqId).Should(Equal(""))
			})
			It("err is not empty", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("failed to get http response"))
			})
		})
		Context("status code >= 400", func() {
			When("recv body is not json format", func() {
				BeforeEach(func() {
					httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests/Unavailable", httpmock.NewBytesResponder(503, []byte(`Service Unavailable`)))
					testSpec.ID = "Unavailable"
					reqId, err = c.Do(context.Background(), testSpec, api.ActionRead, nil, nil)
				})
				It("reqId is empty", func() {
					Expect(reqId).Should(Equal(""))
				})
				It("err is not empty", func() {
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(MatchRegexp("failed to request: status code"))
				})
			})
			When("recv body is json format", func() {
				BeforeEach(func() {
					httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests/ERROR-BAD_REQUEST", httpmock.NewBytesResponder(400, []byte(`{
						"request_id": "F96C25C6B13E49E59F0093BCD8D731AB",
						"error_type": "fail",
						"error_message": "error message"
					}`)))
					testSpec.ID = "ERROR-BAD_REQUEST"
					reqId, err = c.Do(context.Background(), testSpec, api.ActionRead, nil, nil)
				})
				It("reqId is not empty", func() {
					Expect(reqId).Should(Equal("F96C25C6B13E49E59F0093BCD8D731AB"))
				})
				It("err is not empty", func() {
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(MatchRegexp("ErrorType: fail Message: error message"))
				})
			})
		})
	})
	Context("Read", func() {
		BeforeEach(func() {
			httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests/F2246BD8617444329A40470AEC7B00B9", httpmock.NewBytesResponder(200, []byte(`{
				"request_id": "0CAC3AEFB6334ECCA7B70AF76D73508B",
				"result": {
					"id": "F2246BD8617444329A40470AEC7B00B9",
					"name": "test1",
					"number": 1
				}
			}`)))
			httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests/86E5EB52E2294614B35F01943AB7239A", httpmock.NewBytesResponder(200, []byte(`{
				"request_id": "86E5EB52E2294614B35F01943AB7239A",
				"result": {
					"id": "86E5EB52E2294614B35F01943AB7239A",
					"name": "failed",
					"number": "1"
				}
			}`)))
			httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests/NOTFOUND", httpmock.NewBytesResponder(404, []byte(`{
				"request_id": "F96C25C6B13E49E59F0093BCD8D731AB",
				"error_type": "NotFound",
				"error_message": "Specified resource not found."
			}`)))
			httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests/41874381FD914D88B24A759155B5F0F8", httpmock.NewBytesResponder(200, []byte(`{
				"request_id": "41874381FD914D88B24A759155B5F0F8",
				"results": [
					{
						"id": "86E5EB52E2294614B35F01943AB7239A",
						"name": "failed",
						"number": "1"	
					}
				]
			}`)))
			httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests/83870F11B3F24AE89888C86F497EF697", httpmock.NewBytesResponder(200, []byte(`{
				"request_id": "83870F11B3F24AE89888C86F497EF697",
				"id": "83870F11B3F24AE89888C86F497EF697",
				"name": "failed",
				"number": 1
			}`)))
			httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests/72AA019055524D09A4C367DD96DCE0F9", httpmock.NewBytesResponder(200, []byte(`{
				"request_id": "72AA019055524D09A4C367DD96DCE0F9",
				"id": "72AA019055524D09A4C367DD96DCE0F9",
				"name": "failed",
				"number": "1"
			}`)))
			httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests/NOT_JSON", httpmock.NewBytesResponder(200, []byte(`{1}`)))
			httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests/TIMEOUT", func(r *http.Request) (*http.Response, error) {
				time.Sleep(2 * time.Second)
				return httpmock.NewStringResponse(500, "ERROR"), nil
			})
		})
		AfterEach(func() {
			httpmock.Reset()
		})
		When("response timeout", func() {
			BeforeEach(func() {
				testSpec.ID = "TIMEOUT"
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				_, err = c.Read(ctx, testSpec)
			})
			It("is timeout", func() {
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(context.DeadlineExceeded))
			})
		})
		When("cancel", func() {
			BeforeEach(func() {
				testSpec.ID = "TIMEOUT"
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				cancel()
				_, err = c.Read(ctx, testSpec)
			})
			It("is timeout", func() {
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(context.Canceled))
			})
		})
		When("resource is exist", func() {
			BeforeEach(func() {
				testSpec.ID = "F2246BD8617444329A40470AEC7B00B9"
				reqId, err = c.Read(context.Background(), testSpec)
			})
			It("reqId is not empty", func() {
				Expect(reqId).Should(Equal("0CAC3AEFB6334ECCA7B70AF76D73508B"))
			})
			It("err is empty", func() {
				Expect(err).To(Succeed())
			})
			It("can get values", func() {
				Expect(testSpec.ID).To(Equal("F2246BD8617444329A40470AEC7B00B9"))
				Expect(testSpec.Name).To(Equal("test1"))
				Expect(testSpec.Number).To(Equal(int64(1)))
			})
		})
		When("resource is exist, but schema error", func() {
			BeforeEach(func() {
				testSpec.ID = "86E5EB52E2294614B35F01943AB7239A"
				reqId, err = c.Read(context.Background(), testSpec)
			})
			It("reqId is not empty", func() {
				Expect(reqId).Should(Equal("86E5EB52E2294614B35F01943AB7239A"))
			})
			It("err is not empty", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("failed to parse response result"))
			})
		})
		When("resource is exist, but response is list format", func() {
			BeforeEach(func() {
				testSpec.ID = "41874381FD914D88B24A759155B5F0F8"
				reqId, err = c.Read(context.Background(), testSpec)
			})
			It("reqId is not empty", func() {
				Expect(reqId).Should(Equal("41874381FD914D88B24A759155B5F0F8"))
			})
			It("err is not empty", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("not support ListSpec"))
			})
		})
		When("response is no frame format", func() {
			BeforeEach(func() {
				testSpec.ID = "83870F11B3F24AE89888C86F497EF697"
				reqId, err = c.Read(context.Background(), testSpec)
			})
			It("reqId is not empty", func() {
				Expect(reqId).Should(Equal("83870F11B3F24AE89888C86F497EF697"))
			})
			It("err is not empty", func() {
				Expect(testSpec.ID).To(Equal("83870F11B3F24AE89888C86F497EF697"))
				Expect(testSpec.Name).To(Equal("failed"))
				Expect(testSpec.Number).To(Equal(int64(1)))
			})
		})
		When("response is no frame format, but invalid schema", func() {
			BeforeEach(func() {
				testSpec.ID = "72AA019055524D09A4C367DD96DCE0F9"
				reqId, err = c.Read(context.Background(), testSpec)
			})
			It("reqId is not empty", func() {
				Expect(reqId).Should(Equal("72AA019055524D09A4C367DD96DCE0F9"))
			})
			It("err is not empty", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("failed to parse response result"))
			})
		})
		When("resource is not exist", func() {
			BeforeEach(func() {
				testSpec.ID = "NOTFOUND"
				reqId, err = c.Read(context.Background(), testSpec)
			})
			It("reqId is not empty", func() {
				Expect(reqId).Should(Equal("F96C25C6B13E49E59F0093BCD8D731AB"))
			})
			It("err is not empty", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("ErrorType: NotFound"))
			})
		})
		When("response is not json", func() {
			BeforeEach(func() {
				testSpec.ID = "NOT_JSON"
				reqId, err = c.Read(context.Background(), testSpec)
			})
			It("err is not empty", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("failed to parse get response"))
			})
		})
	})
	Context("List", func() {
		When("resource exist", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "0CAC3AEFB6334ECCA7B70AF76D73508B",
					"results": [
						{
							"id": "F2246BD8617444329A40470AEC7B00B9",
							"name": "test1",
							"number": 1
						},
						{
							"id": "C79CE3B1C87B47FA9BC618E6C40C3BD1",
							"name": "test2",
							"number": 2
						}
					]
				}`)))
				reqId, err = c.List(context.Background(), listTestSpec, nil)
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("reqId is not empty", func() {
				Expect(reqId).Should(Equal("0CAC3AEFB6334ECCA7B70AF76D73508B"))
			})
			It("err is empty", func() {
				Expect(err).To(Succeed())
			})
			It("can get values", func() {
				Expect(*listTestSpec).To(Equal(TestSpecList{
					Items: []TestSpec{
						{
							ID:     "F2246BD8617444329A40470AEC7B00B9",
							Name:   "test1",
							Number: 1,
						},
						{
							ID:     "C79CE3B1C87B47FA9BC618E6C40C3BD1",
							Name:   "test2",
							Number: 2,
						},
					},
				}))
			})
		})
		When("resource exist, but schema invalid", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "0CAC3AEFB6334ECCA7B70AF76D73508B",
					"results": [
						{
							"id": "F2246BD8617444329A40470AEC7B00B9",
							"name": "test1",
							"number": "1"
						},
						{
							"id": "C79CE3B1C87B47FA9BC618E6C40C3BD1",
							"name": "test2",
							"number": "2"
						}
					]
				}`)))
				reqId, err = c.List(context.Background(), listTestSpec, nil)
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("reqId is not empty", func() {
				Expect(reqId).Should(Equal("0CAC3AEFB6334ECCA7B70AF76D73508B"))
			})
			It("err is not empty", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("failed to parse list response results"))
			})
		})
		When("resource not exist", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests", httpmock.NewBytesResponder(404, []byte(`{
					"request_id": "F96C25C6B13E49E59F0093BCD8D731AB",
					"error_type": "NotFound",
					"error_message": "Specified resource not found."
				}`)))
				reqId, err = c.List(context.Background(), listTestSpec, nil)
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("reqId is not empty", func() {
				Expect(reqId).Should(Equal("F96C25C6B13E49E59F0093BCD8D731AB"))
			})
			It("err is not empty", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("ErrorType: NotFound"))
			})
		})
	})
	Context("ListAll", func() {
		BeforeEach(func() {
			httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests?limit=1", httpmock.NewBytesResponder(200, []byte(`{
				"request_id": "B7B871E6357E453287E2D7A9590D195D",
				"results": [
					{
						"id": "F2246BD8617444329A40470AEC7B00B9",
						"name": "test1",
						"number": 1
					}
				]
			}`)))
			httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests?limit=1&offset=1", httpmock.NewBytesResponder(200, []byte(`{
				"request_id": "5F234DDF1D8E47D1BB26745D13FAE459",
				"results": [
					{
						"id": "C79CE3B1C87B47FA9BC618E6C40C3BD1",
						"name": "test2",
						"number": 2
					}
				]
			}`)))
			httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests/count", httpmock.NewBytesResponder(200, []byte(`{
				"request_id": "E5173B43AE7F4F8AB71168CDBEC75605",
				"result": {
					"count": 2
				}
			}`)))
		})
		AfterEach(func() {
			httpmock.Reset()
		})
		When("resource exist", func() {
			BeforeEach(func() {
				reqId, err = c.ListAll(context.Background(), countableList, &api.CommonSearchParams{Limit: 1})
			})
			It("reqId is not empty", func() {
				Expect(reqId).Should(Equal("5F234DDF1D8E47D1BB26745D13FAE459"))
			})
			It("err is empty", func() {
				Expect(err).To(Succeed())
			})
			It("can get values", func() {
				Expect(*countableList).To(Equal(TestSpecCountableList{
					Count: api.Count{Count: 2},
					TestSpecList: TestSpecList{
						Items: []TestSpec{
							{
								ID:     "F2246BD8617444329A40470AEC7B00B9",
								Name:   "test1",
								Number: 1,
							},
							{
								ID:     "C79CE3B1C87B47FA9BC618E6C40C3BD1",
								Name:   "test2",
								Number: 2,
							},
						},
					},
				}))
			})
		})
		When("resource not exist", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests", httpmock.NewBytesResponder(404, []byte(`{
					"request_id": "F96C25C6B13E49E59F0093BCD8D731AB",
					"error_type": "NotFound",
					"error_message": "Specified resource not found."
				}`)))
				reqId, err = c.ListAll(context.Background(), countableList, nil)
			})
			It("reqId is not empty", func() {
				Expect(reqId).Should(Equal("F96C25C6B13E49E59F0093BCD8D731AB"))
			})
			It("err is not empty", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("ErrorType: NotFound"))
			})
		})
		When("count value is invalid", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests/count", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "E5173B43AE7F4F8AB71168CDBEC75605",
					"result": {
						"count": "2"
					}
				}`)))
				reqId, err = c.ListAll(context.Background(), countableList, nil)
			})
			It("reqId is not empty", func() {
				Expect(reqId).Should(Equal("E5173B43AE7F4F8AB71168CDBEC75605"))
			})
			It("err is not empty", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})
	Context("Count", func() {
		When("resource exist", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests/count?limit=1", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "0CAC3AEFB6334ECCA7B70AF76D73508B",
					"result": {
						"count": 4
					}
				}`)))
				reqId, err = c.Count(context.Background(), countableList, &api.CommonSearchParams{Limit: 1})
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("reqId is not empty", func() {
				Expect(reqId).Should(Equal("0CAC3AEFB6334ECCA7B70AF76D73508B"))
			})
			It("err is empty", func() {
				Expect(err).To(Succeed())
			})
			It("can get values", func() {
				Expect(countableList.Count.Count).To(Equal(int32(4)))
			})
		})
		When("resource exist, but count format invalid", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests/count?limit=1", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "0CAC3AEFB6334ECCA7B70AF76D73508B",
					"result": {
						"count": "4"
					}
				}`)))
				reqId, err = c.Count(context.Background(), countableList, &api.CommonSearchParams{Limit: 1})
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("reqId is not empty", func() {
				Expect(reqId).Should(Equal("0CAC3AEFB6334ECCA7B70AF76D73508B"))
			})
			It("err is not empty", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("failed to parse count response result"))
			})
		})
		When("resource not exist", func() {
			BeforeEach(func() {
				reqId, err = c.Count(context.Background(), countableList, nil)
			})
			It("reqId is not empty", func() {
				Expect(reqId).Should(Equal("F96C25C6B13E49E59F0093BCD8D731AB"))
			})
			It("err is not empty", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("ErrorType: NotFound"))
			})
		})
	})
	Context("Create", func() {
		AfterEach(func() {
			httpmock.Reset()
		})
		When("request is successful", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPost, "http://localhost/tests", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "8C246D4DF69F4473AD9F90ADE0D4DAF2",
					"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/8C246D4DF69F4473AD9F90ADE0D4DAF2"
				}`)))
				reqId, err = c.Create(context.Background(), testSpec, nil)
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("reqId is not empty", func() {
				Expect(reqId).Should(Equal("8C246D4DF69F4473AD9F90ADE0D4DAF2"))
			})
			It("err is empty", func() {
				Expect(err).To(Succeed())
			})
		})
		When("request is failed", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPost, "http://localhost/tests", httpmock.NewBytesResponder(400, []byte(`{
					"request_id": "61AC452CA8164E768D69BD878D2E0479",
					"error_type": "ParamaterError",
					"error_message": "There are invalid parameters.",
					"error_details": [
						{
							"code": "invalid",
							"attribute": "schema"
						}
					]				
				}`)))
				reqId, err = c.Create(context.Background(), testSpec, nil)
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("reqId is not empty", func() {
				Expect(reqId).Should(Equal("61AC452CA8164E768D69BD878D2E0479"))
			})
			It("err is not empty", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("ErrorType: ParamaterError"))
			})
		})
	})
	Context("Update", func() {
		AfterEach(func() {
			httpmock.Reset()
		})
		BeforeEach(func() {
			testSpec.ID = "F2BF2DCC-D94C-4436-B199-7C7377184D06"
		})
		When("request is successful", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/tests/F2BF2DCC-D94C-4436-B199-7C7377184D06", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "70F6B7C158BE4D31AB5201CEB9175185",
					"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/70F6B7C158BE4D31AB5201CEB9175185"
				}`)))
				reqId, err = c.Update(context.Background(), testSpec, nil)
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("reqId is not empty", func() {
				Expect(reqId).Should(Equal("70F6B7C158BE4D31AB5201CEB9175185"))
			})
			It("err is empty", func() {
				Expect(err).To(Succeed())
			})
		})
		When("request is failed", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/tests/F2BF2DCC-D94C-4436-B199-7C7377184D06", httpmock.NewBytesResponder(400, []byte(`{
					"request_id": "9732ACCB6C09406AB7961AAC7241DA8F",
					"error_type": "ParamaterError",
					"error_message": "There are invalid parameters.",
					"error_details": [
						{
							"code": "invalid",
							"attribute": "schema"
						}
					]				
				}`)))
				reqId, err = c.Update(context.Background(), testSpec, nil)
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("reqId is not empty", func() {
				Expect(reqId).Should(Equal("9732ACCB6C09406AB7961AAC7241DA8F"))
			})
			It("err is not empty", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("ErrorType: ParamaterError"))
			})
		})
	})
	Context("Delete", func() {
		AfterEach(func() {
			httpmock.Reset()
		})
		BeforeEach(func() {
			testSpec.ID = "665DB8DD-0280-44B6-B4C6-FEAFCDD90E8B"
		})
		When("request is successful", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodDelete, "http://localhost/tests/665DB8DD-0280-44B6-B4C6-FEAFCDD90E8B", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "64D5485D09F2440C880609406F2257DE",
					"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/64D5485D09F2440C880609406F2257DE"
				}`)))
				reqId, err = c.Delete(context.Background(), testSpec)
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("reqId is not empty", func() {
				Expect(reqId).Should(Equal("64D5485D09F2440C880609406F2257DE"))
			})
			It("err is empty", func() {
				Expect(err).To(Succeed())
			})
		})
		When("request is failed", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodDelete, "http://localhost/tests/665DB8DD-0280-44B6-B4C6-FEAFCDD90E8B", httpmock.NewBytesResponder(400, []byte(`{
					"request_id": "C33C12FCF7B3466EBA729981630A09E5",
					"error_type": "ParamaterError",
					"error_message": "There are invalid parameters.",
					"error_details": [
						{
							"code": "invalid",
							"attribute": "schema"
						}
					]				
				}`)))
				reqId, err = c.Delete(context.Background(), testSpec)
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("reqId is not empty", func() {
				Expect(reqId).Should(Equal("C33C12FCF7B3466EBA729981630A09E5"))
			})
			It("err is not empty", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("ErrorType: ParamaterError"))
			})
		})
	})
	Context("Cancel", func() {
		AfterEach(func() {
			httpmock.Reset()
		})
		BeforeEach(func() {
			testSpec.ID = "BD1BC291-5763-408B-9D2E-A70434E4A810"
		})
		When("request is successful", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodDelete, "http://localhost/tests/BD1BC291-5763-408B-9D2E-A70434E4A810/cancel", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "EE891520C5894EC2B75B6002F01AA76F",
					"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/EE891520C5894EC2B75B6002F01AA76F"
				}`)))
				reqId, err = c.Cancel(context.Background(), testSpec)
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("reqId is not empty", func() {
				Expect(reqId).Should(Equal("EE891520C5894EC2B75B6002F01AA76F"))
			})
			It("err is empty", func() {
				Expect(err).To(Succeed())
			})
		})
		When("request is failed", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodDelete, "http://localhost/tests/BD1BC291-5763-408B-9D2E-A70434E4A810/cancel", httpmock.NewBytesResponder(400, []byte(`{
					"request_id": "86D4FEC9B34A4C1F9DE3E48285C860DB",
					"error_type": "ParamaterError",
					"error_message": "There are invalid parameters.",
					"error_details": [
						{
							"code": "invalid",
							"attribute": "schema"
						}
					]				
				}`)))
				reqId, err = c.Cancel(context.Background(), testSpec)
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("reqId is not empty", func() {
				Expect(reqId).Should(Equal("86D4FEC9B34A4C1F9DE3E48285C860DB"))
			})
			It("err is not empty", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("ErrorType: ParamaterError"))
			})
		})
	})
	Context("Apply", func() {
		AfterEach(func() {
			httpmock.Reset()
		})
		BeforeEach(func() {
			testSpec.ID = "BD1BC291-5763-408B-9D2E-A70434E4A810"
		})
		When("request is successful", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/tests/BD1BC291-5763-408B-9D2E-A70434E4A810/apply", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "EE891520C5894EC2B75B6002F01AA76F",
					"jobs_url": "https://dpi.dns-platform.jp/v1/jobs/EE891520C5894EC2B75B6002F01AA76F"
				}`)))
				reqId, err = c.Apply(context.Background(), testSpec, nil)
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("reqId is not empty", func() {
				Expect(reqId).Should(Equal("EE891520C5894EC2B75B6002F01AA76F"))
			})
			It("err is empty", func() {
				Expect(err).To(Succeed())
			})
		})
		When("request is failed", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/tests/BD1BC291-5763-408B-9D2E-A70434E4A810/apply", httpmock.NewBytesResponder(400, []byte(`{
					"request_id": "86D4FEC9B34A4C1F9DE3E48285C860DB",
					"error_type": "ParamaterError",
					"error_message": "There are invalid parameters.",
					"error_details": [
						{
							"code": "invalid",
							"attribute": "schema"
						}
					]				
				}`)))
				reqId, err = c.Apply(context.Background(), testSpec, nil)
			})
			AfterEach(func() {
				httpmock.Reset()
			})
			It("reqId is not empty", func() {
				Expect(reqId).Should(Equal("86D4FEC9B34A4C1F9DE3E48285C860DB"))
			})
			It("err is not empty", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("ErrorType: ParamaterError"))
			})
		})
	})
	Context("WatchRead", func() {
		BeforeEach(func() {
			httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests/F2246BD8617444329A40470AEC7B00B9", httpmock.NewBytesResponder(200, []byte(`{
				"request_id": "0CAC3AEFB6334ECCA7B70AF76D73508B",
				"result": {
					"id": "F2246BD8617444329A40470AEC7B00B9",
					"name": "test1",
					"number": 1
				}
			}`)))
			testSpec.ID = "F2246BD8617444329A40470AEC7B00B9"
		})
		AfterEach(func() {
			httpmock.Reset()
		})
		When("change value", func() {
			BeforeEach(func() {
				err = c.WatchRead(context.Background(), time.Second, testSpec)
			})
			It("get change value", func() {
				Expect(err).To(Succeed())
				Expect(*testSpec).To(Equal(TestSpec{
					ID:     "F2246BD8617444329A40470AEC7B00B9",
					Name:   "test1",
					Number: 1,
				}))
			})
		})
		When("timeout", func() {
			BeforeEach(func() {
				_, err = c.Read(context.Background(), testSpec)
				Expect(err).To(Succeed())
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
				defer cancel()
				err = c.WatchRead(ctx, time.Second, testSpec)
			})
			It("is timeout", func() {
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(context.DeadlineExceeded))
			})
		})
		When("interval less than 1s", func() {
			BeforeEach(func() {
				err = c.WatchRead(context.Background(), time.Microsecond, testSpec)
			})
			It("returns error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("interval must greater than equals to 1s"))
			})
		})
		When("read error", func() {
			BeforeEach(func() {
				testSpec.ID = "NOT_FOUND"
				err = c.WatchRead(context.Background(), time.Second, testSpec)
			})
			It("return err", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("ErrorType: NotFound"))
			})
		})
	})
	Context("WatchList", func() {
		BeforeEach(func() {
			httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests", httpmock.NewBytesResponder(200, []byte(`{
				"request_id": "0CAC3AEFB6334ECCA7B70AF76D73508B",
				"results": [
					{
						"id": "F2246BD8617444329A40470AEC7B00B9",
						"name": "test1",
						"number": 1
					},
					{
						"id": "C79CE3B1C87B47FA9BC618E6C40C3BD1",
						"name": "test2",
						"number": 2
					}
				]
			}`)))
		})
		AfterEach(func() {
			httpmock.Reset()
		})
		When("change value", func() {
			BeforeEach(func() {
				err = c.WatchList(context.Background(), time.Second, listTestSpec, nil)
			})
			It("get change value", func() {
				Expect(err).To(Succeed())
				Expect(*listTestSpec).To(Equal(TestSpecList{
					Items: []TestSpec{
						{
							ID:     "F2246BD8617444329A40470AEC7B00B9",
							Name:   "test1",
							Number: 1,
						},
						{
							ID:     "C79CE3B1C87B47FA9BC618E6C40C3BD1",
							Name:   "test2",
							Number: 2,
						},
					},
				}))
			})
		})
		When("timeout", func() {
			BeforeEach(func() {
				_, err = c.List(context.Background(), listTestSpec, nil)
				Expect(err).To(Succeed())
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
				defer cancel()
				err = c.WatchList(ctx, time.Second, listTestSpec, nil)
			})
			It("is timeout", func() {
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(context.DeadlineExceeded))
			})
		})
		When("read error", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests", httpmock.NewBytesResponder(404, []byte(`{
					"request_id": "F96C25C6B13E49E59F0093BCD8D731AB",
					"error_type": "NotFound",
					"error_message": "Specified resource not found."
				}`)))
				err = c.WatchList(context.Background(), time.Second, listTestSpec, nil)
			})
			It("return err", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("ErrorType: NotFound"))
			})
		})
	})
	Context("WatchListAll", func() {
		BeforeEach(func() {
			httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests?limit=1", httpmock.NewBytesResponder(200, []byte(`{
				"request_id": "B7B871E6357E453287E2D7A9590D195D",
				"results": [
					{
						"id": "F2246BD8617444329A40470AEC7B00B9",
						"name": "test1",
						"number": 1
					}
				]
			}`)))
			httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests?limit=1&offset=1", httpmock.NewBytesResponder(200, []byte(`{
				"request_id": "5F234DDF1D8E47D1BB26745D13FAE459",
				"results": [
					{
						"id": "C79CE3B1C87B47FA9BC618E6C40C3BD1",
						"name": "test2",
						"number": 2
					}
				]
			}`)))
			httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests/count", httpmock.NewBytesResponder(200, []byte(`{
				"request_id": "E5173B43AE7F4F8AB71168CDBEC75605",
				"result": {
					"count": 2
				}
			}`)))
		})
		AfterEach(func() {
			httpmock.Reset()
		})
		When("change value", func() {
			BeforeEach(func() {
				err = c.WatchListAll(context.Background(), time.Second, countableList, &api.CommonSearchParams{Limit: 1})
			})
			It("get change value", func() {
				Expect(err).To(Succeed())
				Expect(*countableList).To(Equal(TestSpecCountableList{
					Count: api.Count{Count: 2},
					TestSpecList: TestSpecList{
						Items: []TestSpec{
							{
								ID:     "F2246BD8617444329A40470AEC7B00B9",
								Name:   "test1",
								Number: 1,
							},
							{
								ID:     "C79CE3B1C87B47FA9BC618E6C40C3BD1",
								Name:   "test2",
								Number: 2,
							},
						},
					},
				}))
			})
		})
		When("timeout", func() {
			BeforeEach(func() {
				_, err = c.ListAll(context.Background(), countableList, &api.CommonSearchParams{Limit: 1})
				Expect(err).To(Succeed())
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
				defer cancel()
				err = c.WatchListAll(ctx, time.Second, countableList, &api.CommonSearchParams{Limit: 1})
			})
			It("is timeout", func() {
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(context.DeadlineExceeded))
			})
		})
		When("read error", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests?limit=1&offset=1", httpmock.NewBytesResponder(404, []byte(`{
					"request_id": "F96C25C6B13E49E59F0093BCD8D731AB",
					"error_type": "NotFound",
					"error_message": "Specified resource not found."
				}`)))
				err = c.WatchListAll(context.Background(), time.Second, countableList, &api.CommonSearchParams{Limit: 1})
			})
			It("return err", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("ErrorType: NotFound"))
			})
		})
	})
})
