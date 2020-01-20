package glow_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/meinto/glow"
	. "github.com/meinto/glow/testutil"
)

var _ = Describe("Hotfix", func() {
	var branches []Branch

	BeforeEach(func() {
		f1, _ := NewHotfix("0.0.1")
		f2, _ := HotfixFromBranch("refs/heads/hotfix/v0.0.1")
		branches = []Branch{f1, f2}
	})

	It("can be closed", func() {
		ForEachTestSet(branches, func(branch interface{}) {
			Expect(branch.(Branch).CanBeClosed()).To(Equal(true))
		})
	})

	It("closes on release branches & develop branch", func() {
		ForEachTestSet(branches, func(branch interface{}) {
			closeBanches := branch.(Branch).CloseBranches(MockBranchCollection())
			Expect(len(closeBanches)).To(Equal(2))
			Expect(closeBanches[0].ShortBranchName()).To(Equal(RELEASE_BRANCH))
			Expect(closeBanches[1].ShortBranchName()).To(Equal(DEVELOP_BRANCH))
		})
	})

	It("is only allowed to create from master branch", func() {
		ForEachTestSet(branches, func(branch interface{}) {
			f := branch.(Branch)
			for _, testBranch := range MockBranchCollection() {
				testBranchName := testBranch.ShortBranchName()
				if testBranchName == MASTER_BRANCH {
					Expect(f.CreationIsAllowedFrom(testBranchName)).To(BeTrue())
				} else {
					Expect(f.CreationIsAllowedFrom(testBranchName)).To(BeFalse())
				}
			}
		})
	})

	It("can be published", func() {
		ForEachTestSet(branches, func(branch interface{}) {
			Expect(branch.(Branch).CanBePublished()).To(BeTrue())
		})
	})

	It("can be published on master", func() {
		ForEachTestSet(branches, func(branch interface{}) {
			publishBranch := branch.(Branch).PublishBranch()
			Expect(publishBranch.ShortBranchName()).To(Equal(MASTER_BRANCH))
		})
	})

	// settings like default branch
	// ----------------------------
	It("has a branch name", func() {
		ForEachTestSet(branches, func(branch interface{}) {
			branchName := branch.(Branch).BranchName()
			Expect(branchName).To(Equal("refs/heads/" + HOTFIX_BRANCH))
		})
	})

	It("has a short branch name", func() {
		ForEachTestSet(branches, func(branch interface{}) {
			branchName := branch.(Branch).ShortBranchName()
			Expect(branchName).To(Equal(HOTFIX_BRANCH))
		})
	})
})