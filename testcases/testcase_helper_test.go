package testcases_test

import (
	. "github.com/cloudfoundry-incubator/disaster-recovery-acceptance-tests/testcases"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TestcaseHelper", func() {
	Describe("OpenSourceTestCasesWithFocus", func() {
		It("Focuses on a test case", func() {
			tc := OpenSourceTestCasesWithFocus("cf-nfsbroker")

			Expect(tc).To(HaveLen(1))
			Expect(tc[0].Name()).To(Equal("cf-nfsbroker"))
		})

		It("Panics if no suite name is provided", func() {
			Expect(func() {
				OpenSourceTestCasesWithFocus(";djas;klja;ksdljakls")
			}).To(Panic())
		})

		It("Panics if no test cases matching the suite name found", func() {
			Expect(func() {
				OpenSourceTestCasesWithFocus(";djas;klja;ksdljakls")
			}).To(Panic())
		})
	})

	Describe("OpenSourceTestCasesWithRegex", func() {
		It("returns all cases if no skip or focus provided", func() {
			allTc := OpenSourceTestCases()
			tc := OpenSourceTestCasesWithRegexes("", "")

			Expect(tc).To(HaveLen(len(allTc)))
		})

		It("Focusses on a single case", func() {
			tc := OpenSourceTestCasesWithRegexes("", "cf-nfsbroker")

			Expect(tc).To(HaveLen(1))
			Expect(tc[0].Name()).To(Equal("cf-nfsbroker"))
		})

		It("Focusses on multiple cases", func() {
			tc := OpenSourceTestCasesWithRegexes("", "cf-nfsbroker|cf-uaa")

			Expect(tc).To(HaveLen(2))
		})

		It("Excludes a case", func() {
			allTc := OpenSourceTestCases()
			tc := OpenSourceTestCasesWithRegexes("cf-nfsbroker", "")

			Expect(tc).To(HaveLen(len(allTc) - 1))
			Expect(tc).NotTo(ContainElement(NewNFSTestCases()))
		})

		It("Panics if no test cases matching the suite name found", func() {
			Expect(func() {
				OpenSourceTestCasesWithRegexes("", ";djas;klja;ksdljakls")
			}).To(Panic())
		})
	})
})
