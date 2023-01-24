package util

type LocalMdMetadata struct {
	ModifiedOn         string `json:"modifiedOn"`
	CreatedOn          string `json:"createdOn"`
	AuthorName         string `json:"authorName"`
	AuthorEmail        string `json:"authorEmail"`
	AuthorBanner       string `json:"authorBanner"`
	AuthorImage        string `json:"authorImage"`
	Title              string `json:"title"`
	Description        string `json:"description"`
	Tags               string `json:"tags"`
	TagsArr            string `json:"tagsArr"`
	OgType             string `json:"ogType"`
	OgURL              string `json:"ogURL"`
	OgDescription      string `json:"ogDescription"`
	OgImage            string `json:"ogImage"`
	TwitterCard        string `json:"twitterCard"`
	TwitterDescription string `json:"twitterDescription"`
	TwitterTitle       string `json:"twitterTitle"`
	TwitterURL         string `json:"twitterURL"`
	TwitterIamge       string `json:"twitterIamge"`
	DirPath            string `json:"dirPath"`
	ContentFilePath    string `json:"contentFilePath"`
}
