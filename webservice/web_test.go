package webservice

import (
	"net/http"
	"log"
//	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"testing"
	"os"
)

func DoHttpAndExpect(c *http.Client, method, url string) ([]byte, error){
	log.SetOutput(os.Stdout)

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Printf("Making request for %s %s falied with %v", method, url, err)
		return nil, err
	}
	req.Close = true

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil{
		log.Printf("Reading response for %s %s failed with %v", method, url, err)
		return nil, err
	}

	if res.StatusCode != 1 {
			log.Printf("Expected 1 but got %d with %s for %s %s", res.StatusCode, resBody, method, url)
			return nil, err
		}

	return resBody, nil
}

type WebTestSuite struct{
	suite.Suite
	c *http.Client
}

func (suite *WebTestSuite) SetupTest() {
	suite.c = &http.Client{}
}

func (suite *WebTestSuite) TestWebService() {
	res, err := DoHttpAndExpect(suite.c, "GET", "http://localhost:8080/hello")
	if err != nil {
		log.Printf("error: %s", err)
	} else {
		log.Printf("success: %s", res)
	}
}

func TestWebTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}