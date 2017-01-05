package concourse_test

import (
	"github.com/18F/concourse-broker/cf"
	"github.com/18F/concourse-broker/concourse"
	"github.com/18F/concourse-broker/config"
	"github.com/concourse/atc"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	"net/http"
)

var _ = Describe("Concourse", func() {
	Describe("NewClient", func() {
		Context("When NewClient is called", func() {
			It("should return a concourseClient", func() {
				env := config.Env{
					AdminUsername: "user",
					AdminPassword: "password",
				}
				client := concourse.NewClient(env)
				_ = client
				//Expect(client).Should(BeAssignableToTypeOf(concourseClient{}))
			})
		})
	})
	Describe("CreateTeam", func() {
		var expectedURL = "/api/v1/teams/team venture"
		var authMethodURL = "/api/v1/teams/team venture/auth/methods"
		var expectedTeam, desiredTeam atc.Team
		var expectedAuthToken = atc.AuthToken{
			Type:  "Bearer",
			Value: "gobbeldigook",
		}

		BeforeEach(func() {
			desiredTeam = atc.Team{
				UAAAuth: &atc.UAAAuth{
					CFSpaces: []string{""},
				},
			}
			expectedTeam = atc.Team{
				ID:   1,
				Name: "team venture",
			}

			team = client.Team("team venture")
		})
		Context("when I create a team successfully", func() {
			BeforeEach(func() {
				atcServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/api/v1/teams/main/auth/token"),
						ghttp.RespondWithJSONEncoded(http.StatusOK, expectedAuthToken),
					),
				)
				atcServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", authMethodURL),
						ghttp.RespondWithJSONEncoded(http.StatusNotFound, nil),
					),
				)
				atcServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("PUT", expectedURL),
						ghttp.VerifyJSONRepresenting(desiredTeam),
						ghttp.RespondWithJSONEncoded(http.StatusCreated, expectedTeam),
					),
				)
			})
			It("returns no error", func() {
				client := concourse.NewClient(env)
				err := client.CreateTeam(cf.Details{OrgName: "team venture"})
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("when I try to create a team with a name that is used for an existing team", func() {
			BeforeEach(func() {
				atcServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/api/v1/teams/main/auth/token"),
						ghttp.RespondWithJSONEncoded(http.StatusOK, expectedAuthToken),
					),
				)
				atcServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", authMethodURL),
						ghttp.RespondWithJSONEncoded(http.StatusOK, []atc.AuthMethod{{}}),
					),
				)
			})
			It("should fail and indicate it could not provision", func() {
				client := concourse.NewClient(env)
				err := client.CreateTeam(cf.Details{OrgName: "team venture"})
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Team team venture already exists"))
			})

		})
		Context("when I create a team and Concourse blows up", func() {
			BeforeEach(func() {
				atcServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/api/v1/teams/main/auth/token"),
						ghttp.RespondWithJSONEncoded(http.StatusOK, expectedAuthToken),
					),
				)
				atcServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", authMethodURL),
						ghttp.RespondWithJSONEncoded(http.StatusNotFound, nil),
					),
				)
				atcServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("PUT", expectedURL),
						ghttp.VerifyJSONRepresenting(desiredTeam),
						ghttp.RespondWithJSONEncoded(http.StatusInternalServerError, nil),
					),
				)
			})
			It("returns an error", func() {
				client := concourse.NewClient(env)
				err := client.CreateTeam(cf.Details{OrgName: "team venture"})
				Expect(err).To(HaveOccurred())
			})
		})

	})
	Describe("DeleteTeam", func() {
		var expectedURL = "/api/v1/teams/team venture"
		var expectedAuthToken = atc.AuthToken{
			Type:  "Bearer",
			Value: "gobbeldigook",
		}

		Context("when I delete a team successfully", func() {
			BeforeEach(func() {
				atcServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/api/v1/teams/main/auth/token"),
						ghttp.RespondWithJSONEncoded(http.StatusOK, expectedAuthToken),
					),
				)
				atcServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("DELETE", expectedURL),
						ghttp.RespondWithJSONEncoded(http.StatusNoContent, nil),
					),
				)
			})
			It("returns no error", func() {
				client := concourse.NewClient(env)
				err := client.DeleteTeam(cf.Details{OrgName: "team venture"})
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("when Concourse blows up when trying to delete an instance", func() {
			BeforeEach(func() {
				atcServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/api/v1/teams/main/auth/token"),
						ghttp.RespondWithJSONEncoded(http.StatusOK, expectedAuthToken),
					),
				)
				atcServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("DELETE", expectedURL),
						ghttp.RespondWithJSONEncoded(http.StatusInternalServerError, nil),
					),
				)
			})
			It("returns an error stating 'couldn't destroy team'", func() {
				client := concourse.NewClient(env)
				err := client.DeleteTeam(cf.Details{OrgName: "team venture"})
				Expect(err).To(HaveOccurred())
			})
		})

	})
})
