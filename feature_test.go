package glow_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/meinto/glow"
	. "github.com/meinto/glow/testutil"
)

var _ = Describe("Feature", func() {
	var features []Feature

	BeforeEach(func() {
		f1, _ := NewFeature("luke", "falcon")
		f2, _ := FeatureFromBranch("feature/luke/falcon")
		features = []Feature{f1, f2}
	})

	It("can be closed", func() {
		ForEachTestSet(features, func(feature interface{}) {
			Expect(feature.(Feature).CanBeClosed()).To(Equal(true))
		})
	})

	It("only closes on develop", func() {
		ForEachTestSet(features, func(feature interface{}) {
			closeBanches := feature.(Feature).CloseBranches([]Branch{})
			Expect(len(closeBanches)).To(Equal(1))
			Expect(closeBanches[0].ShortBranchName()).To(Equal(DEVELOP_BRANCH))
		})
	})

	It("is only allowed to create from develop branch", func() {
		ForEachTestSet(features, func(feature interface{}) {
			f := feature.(Feature)
			for _, testBranch := range MockBranchCollection() {
				testBranchName := testBranch.ShortBranchName()
				if testBranchName == DEVELOP_BRANCH {
					Expect(f.CreationIsAllowedFrom(testBranchName)).To(BeTrue())
				} else {
					Expect(f.CreationIsAllowedFrom(testBranchName)).To(BeFalse())
				}
			}
		})
	})
})