// go test -v
// go test -run TestDownload/statusok -v
// go test -run TestDownload/statusnotfound -v
// go test -run TestParallelize -v

// Sample test to show how to write a basic sub unit table test.
package example2

import (
	"net/http"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

// TestDownload validates the http Get function can download content and
// handles different status conditions properly.
func TestDownload(t *testing.T) {
	tt := []struct {
		name       string
		url        string
		statusCode int
	}{
		{"statusok", "https://www.ardanlabs.com/blog/index.xml", http.StatusOK},
		{"statusnotfound", "http://rss.cnn.com/rss/cnn_topstorie.rss", http.StatusNotFound},
	}

	t.Log("Given the need to test downloading different content.")
	{
		for i, test := range tt {
			tf := func(t *testing.T) {
				t.Logf("\tTest %d:\tWhen checking %q for status code %d", i, test.url, test.statusCode)
				{
					resp, err := http.Get(test.url)
					if err != nil {
						t.Fatalf("\t%s\tTest %d:\tShould be able to make the Get call : %v", failed, i, err)
					}
					t.Logf("\t%s\tTest %d:\tShould be able to make the Get call.", succeed, i)

					defer resp.Body.Close()

					if resp.StatusCode == test.statusCode {
						t.Logf("\t%s\tTest %d:\tShould receive a %d status code.", succeed, i, test.statusCode)
					} else {
						t.Errorf("\t%s\tTest %d:\tShould receive a %d status code : %v", failed, i, test.statusCode, resp.StatusCode)
					}
				}
			}

			t.Run(test.name, tf)
		}
	}
}

// TestParallelize validates the http Get function can download content and
// handles different status conditions properly but runs the tests in parallel.
func TestParallelize(t *testing.T) {
	type tableTest struct {
		name       string
		url        string
		statusCode int
	}

	tt := []tableTest{
		{"statusok", "https://www.ardanlabs.com/blog/index.xml", http.StatusOK},
		{"statusnotfound", "http://rss.cnn.com/rss/cnn_topstorie.rss", http.StatusNotFound},
	}

	t.Log("Given the need to test downloading different content.")
	{
		for testID, test := range tt {
			tf := func(testID int, test tableTest) func(t *testing.T) {
				return func(t *testing.T) {
					t.Parallel()

					t.Logf("\tTest %d:\tWhen checking %q for status code %d", testID, test.url, test.statusCode)
					{
						resp, err := http.Get(test.url)
						if err != nil {
							t.Fatalf("\t%s\tTest %d:\tShould be able to make the Get call : %v", failed, testID, err)
						}
						t.Logf("\t%s\tTest %d:\tShould be able to make the Get call.", succeed, testID)

						defer resp.Body.Close()

						if resp.StatusCode == test.statusCode {
							t.Logf("\t%s\tTest %d:\tShould receive a %d status code.", succeed, testID, test.statusCode)
						} else {
							t.Errorf("\t%s\tTest %d:\tShould receive a %d status code : %v", failed, testID, test.statusCode, resp.StatusCode)
						}
					}
				}
			}

			t.Run(test.name, tf(testID, test))
		}
	}
}