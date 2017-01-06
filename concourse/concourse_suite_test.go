package concourse

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/18F/concourse-broker/config"
	"github.com/concourse/go-concourse/concourse"
	"github.com/onsi/gomega/ghttp"
	"net/http"
	"testing"
)

func TestConcourse(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Concourse Suite")
}

var (
	atcServer *ghttp.Server
	client    concourse.Client
	team      concourse.Team
	env       config.Env
)

var _ = BeforeEach(func() {
	atcServer = ghttp.NewServer()

	client = concourse.NewClient(
		atcServer.URL(),
		&http.Client{},
	)

	team = client.Team("some-team")

	env = config.Env{
		AdminUsername: "user",
		AdminPassword: "password",
		ConcourseURL:  atcServer.URL(),
	}
})

var _ = AfterEach(func() {
	atcServer.Close()
})
