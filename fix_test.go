package glow_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/meinto/glow"
	. "github.com/meinto/glow/testutil"
)

var _ = Describe("Fix", func() {
	var branches []AuthoredBranch

	BeforeEach(func() {
		f1, _ := NewFix("luke", "falcon")
		f2, _ := FixFromBranch("refs/heads/fix/luke/falcon")
		branches = []AuthoredBranch{f1, f2}
	})

	It("can be closed", func() {
		ForEachTestSet(branches, func(branch interface{}) {
			Expect(branch.(AuthoredBranch).CanBeClosed()).To(Equal(true))
		})
	})

	It("only closes on release branch", func() {
		ForEachTestSet(branches, func(branch interface{}) {
			closeBanches := branch.(AuthoredBranch).CloseBranches(MockBranchCollection())
			Expect(len(closeBanches)).To(Equal(1))
			Expect(closeBanches[0].ShortBranchName()).To(Equal(RELEASE_BRANCH))
		})
	})

	It("is only allowed to create from release branch", func() {
		ForEachTestSet(branches, func(branch interface{}) {
			f := branch.(AuthoredBranch)
			for _, testBranch := range MockBranchCollection() {
				testBranchName := testBranch.ShortBranchName()
				if testBranchName == RELEASE_BRANCH {
					Expect(f.CreationIsAllowedFrom(testBranch)).To(BeTrue())
				} else {
					Expect(f.CreationIsAllowedFrom(testBranch)).To(BeFalse())
				}
			}
		})
	})

	// settings like default branch
	// ----------------------------
	It("cannot be published", func() {
		ForEachTestSet(branches, func(branch interface{}) {
			Expect(branch.(AuthoredBranch).CanBePublished()).To(BeFalse())
		})
	})

	It("has no publish branch", func() {
		ForEachTestSet(branches, func(branch interface{}) {
			publishBranch := branch.(AuthoredBranch).PublishBranch()
			Expect(publishBranch).To(BeNil())
		})
	})

	It("has a branch name", func() {
		ForEachTestSet(branches, func(branch interface{}) {
			branchName := branch.(AuthoredBranch).BranchName()
			Expect(branchName).To(Equal("refs/heads/" + FIX_BRANCH))
		})
	})

	It("has a short branch name", func() {
		ForEachTestSet(branches, func(branch interface{}) {
			branchName := branch.(AuthoredBranch).ShortBranchName()
			Expect(branchName).To(Equal(FIX_BRANCH))
		})
	})
})
