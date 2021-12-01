package testtool_test

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/mimuret/golang-iij-dpf/pkg/api"
	"github.com/mimuret/golang-iij-dpf/pkg/testtool"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("testclient.go", func() {
	Context("TestClient", func() {
		var (
			err   error
			reqId string
			cl    *testtool.TestClient
			ok    bool
			s     testtool.TestSpec
			slist testtool.TestSpecCountableList
			count int
		)
		BeforeEach(func() {
			ok = false
			s = testtool.TestSpec{
				Id: "1",
			}
			slist = testtool.TestSpecCountableList{}
			cl = testtool.NewTestClient("token", "http://localhost", nil)
			// for count
			httpmock.RegisterRegexpResponder(http.MethodGet, regexp.MustCompile(".*"), httpmock.NewErrorResponder(fmt.Errorf("error")))
			httpmock.RegisterRegexpResponder(http.MethodPatch, regexp.MustCompile(".*"), httpmock.NewErrorResponder(fmt.Errorf("error")))
			httpmock.RegisterRegexpResponder(http.MethodPost, regexp.MustCompile(".*"), httpmock.NewErrorResponder(fmt.Errorf("error")))
			httpmock.RegisterRegexpResponder(http.MethodDelete, regexp.MustCompile(".*"), httpmock.NewErrorResponder(fmt.Errorf("error")))
		})
		Context("Read", func() {
			When("exist ReadFunc", func() {
				BeforeEach(func() {
					cl.ReadFunc = func(s api.Spec) (requestId string, err error) {
						ok = true
						v := s.(*testtool.TestSpec)
						v.Name = "changed"
						return "", nil
					}
					Expect(cl.Read(&s)).To(Succeed())
				})
				It("run ReadFunc", func() {
					Expect(ok).To(BeTrue())
				})
				It("can change value", func() {
					Expect(s.Name).To(Equal("changed"))
				})
			})
			When("not exist ReadFunc", func() {
				BeforeEach(func() {
					Expect(cl.Read(&s)).To(Succeed())
					count = httpmock.GetTotalCallCount()
				})
				It("run TestClient.Client.Read", func() {
					Expect(ok).To(BeFalse())
					Expect(count).To(Equal(1))
				})
			})
			When("not exist ReadFunc or Client", func() {
				BeforeEach(func() {
					cl.Client = nil
					reqId, err = cl.Read(&s)
				})
				It("returns empty request id", func() {
					Expect(reqId).To(BeEmpty())
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
			})
		})
		Context("List", func() {
			When("exist ListFunc", func() {
				BeforeEach(func() {
					cl.ListFunc = func(list api.ListSpec, keywords api.SearchParams) (requestId string, err error) {
						ok = true
						v := list.(*testtool.TestSpecCountableList)
						v.AddItem(s)
						return "", nil
					}
					Expect(cl.List(&slist, nil)).To(Succeed())
				})
				It("run ListFunc", func() {
					Expect(ok).To(BeTrue())
				})
				It("can edit", func() {
					Expect(slist.Len()).To(Equal(1))
					Expect(slist.Index(0)).To(Equal(s))
				})
			})
			When("not exist ListFunc", func() {
				BeforeEach(func() {
					Expect(cl.List(&slist, nil)).To(Succeed())
					count = httpmock.GetTotalCallCount()
				})
				It("run TestClient.Client.List", func() {
					Expect(ok).To(BeFalse())
					Expect(count).To(Equal(1))
				})
			})
			When("not exist ListFunc or Client", func() {
				BeforeEach(func() {
					cl.Client = nil
					reqId, err = cl.List(&slist, nil)
				})
				It("returns empty request id", func() {
					Expect(reqId).To(BeEmpty())
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
			})
		})
		Context("ListAll", func() {
			When("exist ListAllFunc", func() {
				BeforeEach(func() {
					cl.ListAllFunc = func(list api.CountableListSpec, keywords api.SearchParams) (requestId string, err error) {
						ok = true
						v := list.(*testtool.TestSpecCountableList)
						v.AddItem(s)
						return "", nil
					}
					Expect(cl.ListAll(&slist, nil)).To(Succeed())
				})
				It("run ListAllFunc", func() {
					Expect(ok).To(BeTrue())
				})
				It("can edit", func() {
					Expect(slist.Len()).To(Equal(1))
					Expect(slist.Index(0)).To(Equal(s))
				})
			})
			When("not exist ListAllFunc", func() {
				BeforeEach(func() {
					Expect(cl.ListAll(&slist, nil)).To(Succeed())
					count = httpmock.GetTotalCallCount()
				})
				It("run TestClient.Client.ListAll", func() {
					Expect(ok).To(BeFalse())
					Expect(count).To(Equal(1))
				})
			})
			When("not exist ListAllFunc or Client", func() {
				BeforeEach(func() {
					cl.Client = nil
					reqId, err = cl.ListAll(&slist, nil)
				})
				It("returns empty request id", func() {
					Expect(reqId).To(BeEmpty())
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
			})
		})
		Context("Count", func() {
			When("exist CountFunc", func() {
				BeforeEach(func() {
					cl.CountFunc = func(list api.CountableListSpec, keywords api.SearchParams) (requestId string, err error) {
						ok = true
						v := list.(*testtool.TestSpecCountableList)
						v.SetCount(100)
						return "", nil
					}
					Expect(cl.Count(&slist, nil)).To(Succeed())
				})
				It("run CountFunc", func() {
					Expect(ok).To(BeTrue())
				})
				It("can edit", func() {
					Expect(slist.GetCount()).To(Equal(int32(100)))
				})
			})
			When("not exist CountFunc", func() {
				BeforeEach(func() {
					Expect(cl.Count(&slist, nil)).To(Succeed())
					count = httpmock.GetTotalCallCount()
				})
				It("run TestClient.Client.Count", func() {
					Expect(ok).To(BeFalse())
					Expect(count).To(Equal(1))
				})
			})
			When("not exist CountFunc or Client", func() {
				BeforeEach(func() {
					cl.Client = nil
					reqId, err = cl.Count(&slist, nil)
				})
				It("returns empty request id", func() {
					Expect(reqId).To(BeEmpty())
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
			})
		})
		Context("Update", func() {
			When("exist UpdateFunc", func() {
				BeforeEach(func() {
					cl.UpdateFunc = func(s api.Spec, body interface{}) (requestId string, err error) {
						ok = true
						return "", nil
					}
					Expect(cl.Update(&s, nil)).To(Succeed())
				})
				It("run UpdateFunc", func() {
					Expect(ok).To(BeTrue())
				})
			})
			When("not exist UpdateFunc", func() {
				BeforeEach(func() {
					Expect(cl.Update(&s, nil)).To(Succeed())
					count = httpmock.GetTotalCallCount()
				})
				It("run TestClient.Client.Update", func() {
					Expect(ok).To(BeFalse())
					Expect(count).To(Equal(1))
				})
			})
			When("not exist UpdateFunc or Client", func() {
				BeforeEach(func() {
					cl.Client = nil
					reqId, err = cl.Update(&s, nil)
				})
				It("returns empty request id", func() {
					Expect(reqId).To(BeEmpty())
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
			})
		})
		Context("Create", func() {
			When("exist CreateFunc", func() {
				BeforeEach(func() {
					cl.CreateFunc = func(s api.Spec, body interface{}) (requestId string, err error) {
						ok = true
						return "", nil
					}
					Expect(cl.Create(&s, nil)).To(Succeed())
				})
				It("run CreateFunc", func() {
					Expect(ok).To(BeTrue())
				})
			})
			When("not exist CreateFunc", func() {
				BeforeEach(func() {
					Expect(cl.Create(&s, nil)).To(Succeed())
					count = httpmock.GetTotalCallCount()
				})
				It("run TestClient.Client.Create", func() {
					Expect(ok).To(BeFalse())
					Expect(count).To(Equal(1))
				})
			})
			When("not exist CreateFunc or Client", func() {
				BeforeEach(func() {
					cl.Client = nil
					reqId, err = cl.Create(&s, nil)
				})
				It("returns empty request id", func() {
					Expect(reqId).To(BeEmpty())
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
			})
		})
		Context("Apply", func() {
			When("exist CreateFunc", func() {
				BeforeEach(func() {
					cl.ApplyFunc = func(s api.Spec, body interface{}) (requestId string, err error) {
						ok = true
						return "", nil
					}
					Expect(cl.Apply(&s, nil)).To(Succeed())
				})
				It("run ApplyFunc", func() {
					Expect(ok).To(BeTrue())
				})
			})
			When("not exist ApplyFunc", func() {
				BeforeEach(func() {
					Expect(cl.Apply(&s, nil)).To(Succeed())
					count = httpmock.GetTotalCallCount()
				})
				It("run TestClient.Client.Apply", func() {
					Expect(ok).To(BeFalse())
					Expect(count).To(Equal(1))
				})
			})
			When("not exist ApplyFunc or Client", func() {
				BeforeEach(func() {
					cl.Client = nil
					reqId, err = cl.Apply(&s, nil)
				})
				It("returns empty request id", func() {
					Expect(reqId).To(BeEmpty())
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
			})
		})
		Context("Delete", func() {
			When("exist DeleteFunc", func() {
				BeforeEach(func() {
					cl.DeleteFunc = func(s api.Spec) (requestId string, err error) {
						ok = true
						return "", nil
					}
					Expect(cl.Delete(&s)).To(Succeed())
				})
				It("run DeleteFunc", func() {
					Expect(ok).To(BeTrue())
				})
			})
			When("not exist DeleteFunc", func() {
				BeforeEach(func() {
					Expect(cl.Delete(&s)).To(Succeed())
					count = httpmock.GetTotalCallCount()
				})
				It("run TestClient.Client.Delete", func() {
					Expect(ok).To(BeFalse())
					Expect(count).To(Equal(1))
				})
			})
			When("not exist DeleteFunc or Client", func() {
				BeforeEach(func() {
					cl.Client = nil
					reqId, err = cl.Delete(&s)
				})
				It("returns empty request id", func() {
					Expect(reqId).To(BeEmpty())
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
			})
		})
		Context("Cancel", func() {
			When("exist CancelFunc", func() {
				BeforeEach(func() {
					cl.CancelFunc = func(s api.Spec) (requestId string, err error) {
						ok = true
						return "", nil
					}
					Expect(cl.Cancel(&s)).To(Succeed())
				})
				It("run CancelFunc", func() {
					Expect(ok).To(BeTrue())
				})
			})
			When("not exist CancelFunc", func() {
				BeforeEach(func() {
					Expect(cl.Cancel(&s)).To(Succeed())
					count = httpmock.GetTotalCallCount()
				})
				It("run TestClient.Client.Cancel", func() {
					Expect(ok).To(BeFalse())
					Expect(count).To(Equal(1))
				})
			})
			When("not exist CancelFunc or Client", func() {
				BeforeEach(func() {
					cl.Client = nil
					reqId, err = cl.Cancel(&s)
				})
				It("returns empty request id", func() {
					Expect(reqId).To(BeEmpty())
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
			})
		})
		Context("WatchRead", func() {
			When("exist WatchReadFunc", func() {
				BeforeEach(func() {
					cl.WatchReadFunc = func(ctx context.Context, interval time.Duration, s api.Spec) error {
						ok = true
						return nil
					}
					_ = cl.WatchRead(context.TODO(), time.Second, &s)
				})
				It("run WatchReadFunc", func() {
					Expect(ok).To(BeTrue())
				})
			})
			When("not exist WatchReadFunc", func() {
				BeforeEach(func() {
					_ = cl.WatchRead(context.TODO(), time.Second, &s)
					count = httpmock.GetTotalCallCount()
				})
				It("run TestClient.Client.WatchRead", func() {
					Expect(ok).To(BeFalse())
					Expect(count).To(Equal(1))
				})
			})
			When("not exist WatchReadFunc or Client", func() {
				BeforeEach(func() {
					cl.Client = nil
					err = cl.WatchRead(context.TODO(), time.Second, &s)
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
			})
		})
		Context("WatchList", func() {
			When("exist WatchListFunc", func() {
				BeforeEach(func() {
					cl.WatchListFunc = func(ctx context.Context, interval time.Duration, s api.ListSpec, keyword api.SearchParams) error {
						ok = true
						return nil
					}
					_ = cl.WatchList(context.TODO(), time.Second, &slist, nil)
				})
				It("run WatchListFunc", func() {
					Expect(ok).To(BeTrue())
				})
			})
			When("not exist WatchListFunc", func() {
				BeforeEach(func() {
					_ = cl.WatchList(context.TODO(), time.Second, &slist, nil)
					count = httpmock.GetTotalCallCount()
				})
				It("run TestClient.Client.WatchList", func() {
					Expect(ok).To(BeFalse())
					Expect(count).To(Equal(1))
				})
			})
			When("not exist WatchListFunc or Client", func() {
				BeforeEach(func() {
					cl.Client = nil
					err = cl.WatchList(context.TODO(), time.Second, &slist, nil)
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
			})
		})
		Context("WatchListAll", func() {
			When("exist WatchListFunc", func() {
				BeforeEach(func() {
					cl.WatchListAllFunc = func(ctx context.Context, interval time.Duration, s api.CountableListSpec, keyword api.SearchParams) error {
						ok = true
						return nil
					}
					_ = cl.WatchListAll(context.TODO(), time.Second, &slist, nil)
				})
				It("run WatchListFunc", func() {
					Expect(ok).To(BeTrue())
				})
			})
			When("not exist WatchListAllFunc", func() {
				BeforeEach(func() {
					_ = cl.WatchListAll(context.TODO(), time.Second, &slist, nil)
					count = httpmock.GetTotalCallCount()
				})
				It("run TestClient.Client.WatchListAll", func() {
					Expect(ok).To(BeFalse())
					Expect(count).To(Equal(1))
				})
			})
			When("not exist WatchListAllFunc or Client", func() {
				BeforeEach(func() {
					cl.Client = nil
					err = cl.WatchListAll(context.TODO(), time.Second, &slist, nil)
				})
				It("not returns error", func() {
					Expect(err).To(Succeed())
				})
			})
		})
	})
})
