package model

type NewsLetterNewsItem struct {
	ArticleURL        string `xml:"ArticleURL"`
	NewsArticleID     string `xml:"NewsArticleID"`
	PublishDate       string `xml:"PublishDate"`
	Taxonomies        string `xml:"Taxonomies"`
	TeaserText        string `xml:"TeaserText"`
	ThumbnailImageURL string `xml:"ThumbnailImageURL"`
	Title             string `xml:"Title"`
	OptaMatchId       string `xml:"OptaMatchId"`
	LastUpdateDate    string `xml:"LastUpdateDate"`
	IsPublished       string `xml:"IsPublished"`
}

type NewListInformation struct {
	ClubName            string `xml:"ClubName"`
	ClubWebsiteURL      string `xml:"ClubWebsiteURL"`
	NewsletterNewsItems struct {
		News []NewsLetterNewsItem `xml:"NewsletterNewsItem"`
	} `xml:"NewsletterNewsItems"`
}
