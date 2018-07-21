package extract

import (
	"testing"
	"io/ioutil"
	"path/filepath"
)

var frontPageUrls = []string{
	"https://itunes.apple.com/gb/app/mondo/id1052238659",
	"/-play-store-redirect",
	"https://www.theguardian.com/technology/2017/dec/17/monzo-facebook-of-banking",
	"https://www.telegraph.co.uk/personal-banking/current-accounts/monzo-atom-revolut-starling-everything-need-know-digital-banks/",
	"https://www.thetimes.co.uk/article/tom-blomfield-the-man-who-made-monzo-g8z59dr8n",
	"https://www.standard.co.uk/tech/monzo-prepaid-card-current-accounts-challenger-bank-a3805761.html",
	"/features/apple-pay",
	"/features/travel",
	"https://www.fscs.org.uk/",
	"/features/switch",
	"/features/overdrafts",
	"https://itunes.apple.com/gb/app/mondo/id1052238659",
	"/-play-store-redirect",
	"mailto:help@monzo.com",
	"https://monzo.com/community",
	"https://itunes.apple.com/gb/app/mondo/id1052238659",
	"/-play-store-redirect",
	"/about",
	"/blog",
	"/press",
	"/careers",
	"https://web.monzo.com",
	"/community",
	"/transparency",
	"/blog/how-money-works",
	"/tone-of-voice",
	"/faq",
	"/legal/terms-and-conditions",
	"/legal/fscs-information",
	"/legal/privacy-policy",
	"/legal/cookie-policy",
	"https://itunes.apple.com/gb/app/mondo/id1052238659",
	"/-play-store-redirect",
	"https://twitter.com/monzo",
	"https://www.facebook.com/monzobank",
	"https://www.linkedin.com/company/monzo-bank",
	"https://www.youtube.com/monzobank",
	"mailto:help@monzo.com",
	"tel:08008021281",
	"tel:+442038720620",
}

func TestExtractUrlsFromHtml(t *testing.T) {

	path :=filepath.Join("testData", "main_page_test_file.html")
	t.Log(path)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	resultList := ExtractUrlsFromHtml(data)

	if len(frontPageUrls) != len(resultList) {
		t.Fatalf("Error expected result size %d but got %d", len(frontPageUrls), len(resultList))
	}
	for index, r := range resultList {
		if r != frontPageUrls[index] {
			t.Fatalf("Error expected %q but got %q at index %d", frontPageUrls[index], r, index)
		}
	}
}
