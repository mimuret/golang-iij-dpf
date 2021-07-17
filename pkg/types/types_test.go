package types_test

import (
	"github.com/mimuret/golang-iij-dpf/pkg/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("types", func() {
	Context("Boolean", func() {
		Context("String", func() {
			When("Enabled", func() {
				It("return Enabled", func() {
					Expect(types.Enabled.String()).To(Equal("Enabled"))
				})
			})
			When("Disabled", func() {
				It("return Disabled", func() {
					Expect(types.Disabled.String()).To(Equal("Disabled"))
				})
			})
		})
	})
	Context("State", func() {
		Context("String", func() {
			When("Enabled", func() {
				It("return StateBeforeStart", func() {
					Expect(types.StateBeforeStart.String()).To(Equal("BeforeStart"))
				})
			})
			When("Disabled", func() {
				It("return StateRunning", func() {
					Expect(types.StateRunning.String()).To(Equal("Started"))
				})
			})
		})
	})
	Context("Favorite", func() {
		Context("String", func() {
			When("Enabled", func() {
				It("return FavoriteHighPriority", func() {
					Expect(types.FavoriteHighPriority.String()).To(Equal("High"))
				})
			})
			When("Disabled", func() {
				It("return FavoriteLowPriority", func() {
					Expect(types.FavoriteLowPriority.String()).To(Equal("Low"))
				})
			})
		})
	})
})
