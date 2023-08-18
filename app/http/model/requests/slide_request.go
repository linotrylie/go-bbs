package requests

type SlideRequest struct {
	Slideid   int    ` json:"slideid"`
	Rank      int    ` json:"rank"`
	Name      string ` json:"name"`
	Url       string ` json:"url"`
	Slidepic  string ` json:"slidepic"`
	Picheight string ` json:"picheight"`
}
