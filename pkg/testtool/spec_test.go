package testtool_test

import (
	"net/http"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type SpecForGetMethodPathTest struct {
	testtool.TestSpecCountableList
	Response map[api.Action]struct {
		Method string
		Path   string
	}
}

func (t *SpecForGetMethodPathTest) GetPathMethod(action api.Action) (string, string) {
	res := t.Response[action]
	return res.Method, res.Path
}

var _ = Describe("test spec", func() {
	var (
		c      testtool.TestSpec
		cl     *testtool.TestClient
		err    error
		reqId  string
		s1, s2 testtool.TestSpec
		slist  testtool.TestSpecList
	)
	BeforeEach(func() {
		cl = testtool.NewTestClient("", "http://localhost", nil)
		s1 = testtool.TestSpec{
			Id:     "1",
			Name:   "name1",
			Number: 100,
		}
		s2 = testtool.TestSpec{
			Id:     "2",
			Name:   "",
			Number: 0,
		}

		slist = testtool.TestSpecList{
			Items: []testtool.TestSpec{s1, s2},
		}
	})
	Describe("TestSpec", func() {
		Context("Read", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests/1", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "743C1E2428264E368D37233049AD49B5",
					"result": {
						"id": "1",
						"name": "name1",
						"number": 100
					}
				}`)))
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests/2", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "6B69B707901B479AAF83477B31340FEE",
					"result": {
						"id": "2",
						"name": "",
						"number": 0
					}
				}`)))
			})
			When("returns id=r1", func() {
				BeforeEach(func() {
					c = testtool.TestSpec{
						Id: "1",
					}
					reqId, err = cl.Read(&c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("743C1E2428264E368D37233049AD49B5"))
					Expect(c).To(Equal(s1))
				})
			})
			When("returns id=r2", func() {
				BeforeEach(func() {
					c = testtool.TestSpec{
						Id: "2",
					}
					reqId, err = cl.Read(&c)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("6B69B707901B479AAF83477B31340FEE"))
					Expect(c).To(Equal(s2))
				})
			})
		})
		Context("Create", func() {
			var (
				id1, bs1 = testtool.CreateAsyncResponse()
			)
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPost, "http://localhost/tests", httpmock.NewBytesResponder(202, bs1))
			})
			When("create", func() {
				BeforeEach(func() {
					s := testtool.TestSpec{
						Id:     "3",
						Name:   "name 3",
						Number: 3,
					}
					reqId, err = cl.Create(&s, nil)
				})
				It("returns job_id", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id1))
				})
				It("post json", func() {
					Expect(cl.RequestBody["/tests"]).To(MatchJSON(`{
							"name": "name 3",
							"number": 3
						}`))
				})
			})
			When("create empty", func() {
				BeforeEach(func() {
					s := testtool.TestSpec{
						Id:     "3",
						Name:   "",
						Number: 0,
					}
					reqId, err = cl.Create(&s, nil)
				})
				It("returns job_id", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id1))
				})
				It("post json", func() {
					Expect(cl.RequestBody["/tests"]).To(MatchJSON(`{
							"name": "",
							"number": 0
						}`))
				})
			})
		})
		Context("Update", func() {
			var (
				id1, bs1 = testtool.CreateAsyncResponse()
				id2, bs2 = testtool.CreateAsyncResponse()
			)
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/tests/1", httpmock.NewBytesResponder(202, bs1))
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/tests/2", httpmock.NewBytesResponder(202, bs2))
			})
			When("test 1", func() {
				BeforeEach(func() {
					reqId, err = cl.Update(&s1, nil)
				})
				It("returns job_id", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id1))
				})
				It("post json", func() {
					Expect(cl.RequestBody["/tests/1"]).To(MatchJSON(`{
						"name": "name1",
						"number": 100
					}`))
				})
			})
			When("test 2", func() {
				BeforeEach(func() {
					reqId, err = cl.Update(&s2, nil)
				})
				It("returns job_id", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id2))
				})
				It("post json", func() {
					Expect(cl.RequestBody["/tests/2"]).To(MatchJSON(`{
							"name": "",
							"number": 0
					}`))
				})
			})
		})
		Context("Delete", func() {
			var (
				id1, bs1 = testtool.CreateAsyncResponse()
				id2, bs2 = testtool.CreateAsyncResponse()
			)
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodDelete, "http://localhost/tests/1", httpmock.NewBytesResponder(202, bs1))
				httpmock.RegisterResponder(http.MethodDelete, "http://localhost/tests/2", httpmock.NewBytesResponder(202, bs2))
			})
			When("remove1", func() {
				BeforeEach(func() {
					reqId, err = cl.Delete(&s1)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id1))
				})
			})
			When("remove 2", func() {
				BeforeEach(func() {
					reqId, err = cl.Delete(&s2)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id2))
				})
			})
		})
		Context("Cancel", func() {
			var (
				id1, bs1 = testtool.CreateAsyncResponse()
				id2, bs2 = testtool.CreateAsyncResponse()
			)
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodDelete, "http://localhost/tests/1/cancel", httpmock.NewBytesResponder(202, bs1))
				httpmock.RegisterResponder(http.MethodDelete, "http://localhost/tests/2/cancel", httpmock.NewBytesResponder(202, bs2))
			})
			When("cancel edit 1", func() {
				BeforeEach(func() {
					reqId, err = cl.Cancel(&s1)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id1))
				})
			})
			When("cancel edit 2", func() {
				BeforeEach(func() {
					reqId, err = cl.Cancel(&s2)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id2))
				})
			})
		})
		Context("Apply", func() {
			var (
				id1, bs1 = testtool.CreateAsyncResponse()
				id2, bs2 = testtool.CreateAsyncResponse()
			)
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/tests/1/apply", httpmock.NewBytesResponder(202, bs1))
				httpmock.RegisterResponder(http.MethodPatch, "http://localhost/tests/2/apply", httpmock.NewBytesResponder(202, bs2))
			})
			When("apply edit 1", func() {
				BeforeEach(func() {
					reqId, err = cl.Apply(&s1, nil)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id1))
				})
			})
			When("apply edit 2", func() {
				BeforeEach(func() {
					reqId, err = cl.Apply(&s2, nil)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal(id2))
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *testtool.TestSpec
			testtool.TestDeepCopyObject(&s1, nilSpec)
			testtool.TestGetName(&s1, "tests")
			testtool.TestGetGroup(&s1, "test")
			Context("", func() {
				When("action is ActionCreate", func() {
					testtool.TestGetPathMethod(&s1, api.ActionCreate, http.MethodPost, "/tests")
				})
				When("action is ActionRead", func() {
					testtool.TestGetPathMethod(&s1, api.ActionRead, http.MethodGet, "/tests/1")
				})
				When("action is ActionUpdate", func() {
					testtool.TestGetPathMethod(&s1, api.ActionUpdate, http.MethodPatch, "/tests/1")
				})
				When("action is ActionDelete", func() {
					testtool.TestGetPathMethod(&s1, api.ActionDelete, http.MethodDelete, "/tests/1")
				})
				When("action is ActionCancel", func() {
					testtool.TestGetPathMethod(&s1, api.ActionCancel, http.MethodDelete, "/tests/1/cancel")
				})
				When("action is ActionApply", func() {
					testtool.TestGetPathMethod(&s1, api.ActionApply, http.MethodPatch, "/tests/1/apply")
				})
				When("action is other", func() {
					testtool.TestGetPathMethod(&s1, api.ActionList, "", "")
				})
			})
		})
	})
	Describe("TestSpecList", func() {
		var (
			c testtool.TestSpecList
		)
		Context("List", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "0987808B8D38403C8A1FBE66D72D68F7",
					"results": [
						{
							"id": "1",
							"name": "name1",
							"number": 100	
						},
						{
							"id": "2",
							"name": "",
							"number": 0	
						}
					]
				}`)))
			})
			When("returns list ", func() {
				BeforeEach(func() {
					c = testtool.TestSpecList{}
					reqId, err = cl.List(&c, nil)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("0987808B8D38403C8A1FBE66D72D68F7"))
					Expect(c).To(Equal(slist))
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *testtool.TestSpecList
			testtool.TestDeepCopyObject(&slist, nilSpec)
			testtool.TestGetName(&slist, "tests")
			testtool.TestGetGroup(&slist, "test")
			testtool.TestGetPathMethodForList(&slist, "/tests")
		})
		Context("api.ListSpec common test", func() {
			testtool.TestGetItems(&slist, &slist.Items)
			testtool.TestLen(&slist, 2)
			Context("Index", func() {
				When("index exist", func() {
					It("returns item", func() {
						Expect(slist.Index(0)).To(Equal(s1))
					})
				})
			})
		})
	})
	Describe("TestSpecCountableList", func() {
		var (
			c     testtool.TestSpecCountableList
			slist testtool.TestSpecCountableList
		)
		BeforeEach(func() {
			slist = testtool.TestSpecCountableList{
				TestSpecList: testtool.TestSpecList{
					Items: []testtool.TestSpec{s1, s2},
				},
			}
		})
		Context("List", func() {
			BeforeEach(func() {
				httpmock.RegisterResponder(http.MethodGet, "http://localhost/tests", httpmock.NewBytesResponder(200, []byte(`{
					"request_id": "0987808B8D38403C8A1FBE66D72D68F7",
					"results": [
						{
							"id": "1",
							"name": "name1",
							"number": 100	
						},
						{
							"id": "2",
							"name": "",
							"number": 0	
						}
					]
				}`)))
			})
			When("returns list ", func() {
				BeforeEach(func() {
					c = testtool.TestSpecCountableList{}
					reqId, err = cl.List(&c, nil)
				})
				It("returns normal", func() {
					Expect(err).To(Succeed())
					Expect(reqId).To(Equal("0987808B8D38403C8A1FBE66D72D68F7"))
					Expect(c).To(Equal(slist))
				})
			})
		})
		Context("api.Spec common test", func() {
			var nilSpec *testtool.TestSpecCountableList
			testtool.TestDeepCopyObject(&slist, nilSpec)
			testtool.TestGetName(&slist, "tests")
			testtool.TestGetGroup(&slist, "test")
			testtool.TestGetPathMethodForCountableList(&slist, "/tests")
		})
		Context("api.ListSpec common test", func() {
			testtool.TestGetItems(&slist, &slist.Items)
			testtool.TestLen(&slist, 2)
			Context("Index", func() {
				When("index exist", func() {
					It("returns item", func() {
						Expect(slist.Index(0)).To(Equal(s1))
					})
				})
			})
		})
		Context("api.CountableListSpec common test", func() {
			testtool.TestGetMaxLimit(&slist, 10000)
			testtool.TestClearItems(&slist)
			testtool.TestAddItem(&slist, s1)
		})
	})
	Context("TestGetPathMethod", func() {
		var (
			s = &SpecForGetMethodPathTest{
				Response: map[api.Action]struct {
					Method string
					Path   string
				}{
					api.ActionCreate: {http.MethodPost, "/tests"},
				},
			}
		)
		testtool.TestGetPathMethod(s, api.ActionCreate, http.MethodPost, "/tests")
		testtool.TestGetPathMethod(s, api.ActionList, "", "")
	})
	Context("TestGetPathMethodForSpec", func() {
		var (
			s = &SpecForGetMethodPathTest{
				Response: map[api.Action]struct {
					Method string
					Path   string
				}{
					api.ActionCreate: {http.MethodPost, "/tests"},
					api.ActionRead:   {http.MethodGet, "/tests/1"},
					api.ActionUpdate: {http.MethodPatch, "/tests/1"},
					api.ActionDelete: {http.MethodDelete, "/tests/1"},
					api.ActionCancel: {http.MethodDelete, "/tests/cancel"},
				},
			}
		)
		testtool.TestGetPathMethodForSpec(s, "/tests", "/tests/1")
	})
	Context("TestGetPathMethodForList", func() {
		var (
			s = &SpecForGetMethodPathTest{
				Response: map[api.Action]struct {
					Method string
					Path   string
				}{
					api.ActionList: {http.MethodGet, "/tests"},
				},
			}
		)
		testtool.TestGetPathMethodForList(s, "/tests")
	})
	Context("TestGetPathMethodForCountableList", func() {
		var (
			s = &SpecForGetMethodPathTest{
				Response: map[api.Action]struct {
					Method string
					Path   string
				}{
					api.ActionList:  {http.MethodGet, "/tests"},
					api.ActionCount: {http.MethodGet, "/tests/count"},
				},
			}
		)
		testtool.TestGetPathMethodForCountableList(s, "/tests")
	})
})
