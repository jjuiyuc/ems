package utils

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("UtilsTime", func() {
	Describe("AddDate", func() {
		Context("leap year: add one month to 2020-01-31T16:00:00Z", func() {
			It("should return 2020-02-29T16:00:00Z", func() {
				seedUtDate := time.Date(2020, 1, 31, 16, 0, 0, 0, time.UTC)
				expectedResponseDate := time.Date(2020, 2, 29, 16, 0, 0, 0, time.UTC)
				date := AddDate(seedUtDate, 0, 1, 0)
				Expect(date).To(Equal(expectedResponseDate))
			})
		})

		Context("leap year: subtract one month to 2020-03-31T16:00:00Z", func() {
			It("should return 2020-02-29T16:00:00Z", func() {
				seedUtDate := time.Date(2020, 3, 31, 16, 0, 0, 0, time.UTC)
				expectedResponseDate := time.Date(2020, 2, 29, 16, 0, 0, 0, time.UTC)
				date := AddDate(seedUtDate, 0, -1, 0)
				Expect(date).To(Equal(expectedResponseDate))
			})
		})

		Context("normal year: add one month to 2021-01-31T16:00:00Z", func() {
			It("should return 2021-02-28T16:00:00Z", func() {
				seedUtDate := time.Date(2021, 1, 31, 16, 0, 0, 0, time.UTC)
				expectedResponseDate := time.Date(2021, 2, 28, 16, 0, 0, 0, time.UTC)
				date := AddDate(seedUtDate, 0, 1, 0)
				Expect(date).To(Equal(expectedResponseDate))
			})
		})

		Context("normal year: add one month to 2021-03-31T16:00:00Z", func() {
			It("should return 2021-04-30T16:00:00Z", func() {
				seedUtDate := time.Date(2021, 3, 31, 16, 0, 0, 0, time.UTC)
				expectedResponseDate := time.Date(2021, 4, 30, 16, 0, 0, 0, time.UTC)
				date := AddDate(seedUtDate, 0, 1, 0)
				Expect(date).To(Equal(expectedResponseDate))
			})
		})

		Context("normal year: subtract one month to 2021-03-31T16:00:00Z", func() {
			It("should return 2021-02-28T16:00:00Z", func() {
				seedUtDate := time.Date(2021, 3, 31, 16, 0, 0, 0, time.UTC)
				expectedResponseDate := time.Date(2021, 2, 28, 16, 0, 0, 0, time.UTC)
				date := AddDate(seedUtDate, 0, -1, 0)
				Expect(date).To(Equal(expectedResponseDate))
			})
		})

		Context("normal year: subtract one month to 2021-04-30T16:00:00Z", func() {
			It("should return 2021-03-31T16:00:00Z", func() {
				seedUtDate := time.Date(2021, 4, 30, 16, 0, 0, 0, time.UTC)
				expectedResponseDate := time.Date(2021, 3, 31, 16, 0, 0, 0, time.UTC)
				date := AddDate(seedUtDate, 0, -1, 0)
				Expect(date).To(Equal(expectedResponseDate))
			})
		})

		Context("normal year: subtract ten month to 2021-04-30T16:00:00Z", func() {
			It("should return 2020-06-30T16:00:00Z", func() {
				seedUtDate := time.Date(2021, 4, 30, 16, 0, 0, 0, time.UTC)
				expectedResponseDate := time.Date(2020, 6, 30, 16, 0, 0, 0, time.UTC)
				date := AddDate(seedUtDate, 0, -10, 0)
				Expect(date).To(Equal(expectedResponseDate))
			})
		})

		Context("leap year: add one year to 2019-02-28T16:00:00Z", func() {
			It("should return 2020-02-29T16:00:00Z", func() {
				seedUtDate := time.Date(2019, 2, 28, 16, 0, 0, 0, time.UTC)
				expectedResponseDate := time.Date(2020, 2, 29, 16, 0, 0, 0, time.UTC)
				date := AddDate(seedUtDate, 1, 0, 0)
				Expect(date).To(Equal(expectedResponseDate))
			})
		})

		Context("leap year: subtract one year to 2020-02-29T16:00:00Z", func() {
			It("should return 2019-02-28T16:00:00Z", func() {
				seedUtDate := time.Date(2020, 2, 29, 16, 0, 0, 0, time.UTC)
				expectedResponseDate := time.Date(2019, 2, 28, 16, 0, 0, 0, time.UTC)
				date := AddDate(seedUtDate, -1, 0, 0)
				Expect(date).To(Equal(expectedResponseDate))
			})
		})

		Context("normal year: add one year to 2022-06-30T16:00:00Z", func() {
			It("should return 2023-06-30T16:00:00Z", func() {
				seedUtDate := time.Date(2022, 6, 30, 16, 0, 0, 0, time.UTC)
				expectedResponseDate := time.Date(2023, 6, 30, 16, 0, 0, 0, time.UTC)
				date := AddDate(seedUtDate, 1, 0, 0)
				Expect(date).To(Equal(expectedResponseDate))
			})
		})

		Context("normal year: subtract one year to 2022-07-31T16:00:00Z", func() {
			It("should return 2021-07-31T16:00:00Z", func() {
				seedUtDate := time.Date(2022, 7, 31, 16, 0, 0, 0, time.UTC)
				expectedResponseDate := time.Date(2021, 7, 31, 16, 0, 0, 0, time.UTC)
				date := AddDate(seedUtDate, -1, 0, 0)
				Expect(date).To(Equal(expectedResponseDate))
			})
		})
	})
})
