package concourse

import (
	"encoding/json"
	"net/http"

	"code.cloudfoundry.org/lager"
	"github.com/18F/concourse-broker/cf"
	"github.com/18F/concourse-broker/config"
	"github.com/concourse/atc"
	"github.com/concourse/atc/auth/uaa"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Concourse", func() {
	Describe("NewClient", func() {
		Context("When NewClient is called", func() {
			It("should return a concourseClient", func() {
				env := config.Env{
					AdminUsername: "user",
					AdminPassword: "password",
				}
				client := NewClient(env, logger)
				expectedClient := new(concourseClient)
				Expect(client).Should(BeAssignableToTypeOf((expectedClient)))
			})
		})
	})
	Describe("UpdateTeams", func() {
		var expectedURL = "/api/v1/teams/team venture"
		var expectedTeam, desiredTeam atc.Team
		var expectedTeams []atc.Team
		var expectedAuthToken = atc.AuthToken{
			Type:  "Bearer",
			Value: "gobbeldigook",
		}
		BeforeEach(func() {
			env.ClientID = "new-id"
			env.ClientSecret = "new-secret"

			outdatedAuth := uaa.UAAAuthConfig{
				ClientID:     "original-id",
				ClientSecret: "original-secret",
				CFSpaces:     []string{""},
			}
			outdatedAuthData, err := json.Marshal(outdatedAuth)
			Expect(err).NotTo(HaveOccurred())

			currentAuth := uaa.UAAAuthConfig{
				ClientID:     "new-id",
				ClientSecret: "new-secret",
				CFSpaces:     []string{""},
			}
			currentAuthData, err := json.Marshal(currentAuth)
			Expect(err).NotTo(HaveOccurred())

			desiredTeam = atc.Team{
				ID:   1,
				Name: "team venture",
				Auth: map[string]*json.RawMessage{
					uaa.ProviderName: (*json.RawMessage)(&currentAuthData),
				},
			}
			expectedTeams = []atc.Team{
				{
					ID:   1,
					Name: "team venture",
					Auth: map[string]*json.RawMessage{
						uaa.ProviderName: (*json.RawMessage)(&outdatedAuthData),
					},
				},
				{
					ID:   2,
					Name: "team rocket",
					Auth: map[string]*json.RawMessage{
						uaa.ProviderName: (*json.RawMessage)(&currentAuthData),
					},
				},
			}
			expectedTeam = atc.Team{
				ID:   1,
				Name: "team venture",
			}
		})
		Context("when i update teams", func() {
			BeforeEach(func() {
				atcServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/api/v1/teams/main/auth/token"),
						ghttp.RespondWithJSONEncoded(http.StatusOK, expectedAuthToken),
					),
				)
				atcServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/api/v1/teams"),
						ghttp.RespondWithJSONEncoded(http.StatusOK, expectedTeams),
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
			It("updates team uaa credentials", func() {
				client := NewClient(env, logger)
				err := client.UpdateTeams()
				Expect(err).NotTo(HaveOccurred())
				Expect(atcServer.ReceivedRequests()).To(HaveLen(3))
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
			auth := uaa.UAAAuthConfig{CFSpaces: []string{""}}
			authData, err := json.Marshal(auth)
			Expect(err).NotTo(HaveOccurred())

			desiredTeam = atc.Team{
				Auth: map[string]*json.RawMessage{
					uaa.ProviderName: (*json.RawMessage)(&authData),
				},
			}
			expectedTeam = atc.Team{
				ID:   1,
				Name: "team venture",
			}
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
				client := NewClient(env, logger)
				err := client.CreateTeam(cf.Details{OrgName: "team venture"})
				Expect(err).NotTo(HaveOccurred())
				Expect(logger.Logs()).To(HaveLen(0))
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
				client := NewClient(env, logger)
				err := client.CreateTeam(cf.Details{OrgName: "team venture"})
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Team team venture already exists"))
				logs := logger.Logs()
				Expect(logs).To(HaveLen(1))
				Expect(logs[0].LogLevel).To(Equal(lager.ERROR))
				Expect(logs[0].Message).To(ContainSubstring("concourse-client.create-team.existing-team-error"))
				Expect(logs[0].Data["team-name"]).To(Equal("team venture"))
				Expect(logs[0].Data["auth-methods-size"]).To(BeEquivalentTo(1))

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
				client := NewClient(env, logger)
				err := client.CreateTeam(cf.Details{OrgName: "team venture"})
				Expect(err).To(HaveOccurred())
				logs := logger.Logs()
				Expect(logs).To(HaveLen(1))
				Expect(logs[0].LogLevel).To(Equal(lager.ERROR))
				Expect(logs[0].Message).To(ContainSubstring("concourse-client.create-team.unknown-create-error"))
				Expect(logs[0].Data["team-name"]).To(Equal("team venture"))
			})
		})
		Context("when I try to delete a team but I can't auth as an admin", func() {
			BeforeEach(func() {
				atcServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/api/v1/teams/main/auth/token"),
						ghttp.RespondWithJSONEncoded(http.StatusUnauthorized, nil),
					),
				)
			})
			It("returns an error", func() {
				client := NewClient(env, logger)
				err := client.CreateTeam(cf.Details{OrgName: "team venture"})
				Expect(err).To(HaveOccurred())
				logs := logger.Logs()
				Expect(logs).To(HaveLen(1))
				Expect(logs[0].LogLevel).To(Equal(lager.ERROR))
				Expect(logs[0].Message).To(ContainSubstring("concourse-client.create-team.auth-client-error"))
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
				client := NewClient(env, logger)
				err := client.DeleteTeam(cf.Details{OrgName: "team venture"})
				Expect(err).NotTo(HaveOccurred())
				Expect(logger.Logs()).To(HaveLen(0))
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
				client := NewClient(env, logger)
				err := client.DeleteTeam(cf.Details{OrgName: "team venture"})
				Expect(err).To(HaveOccurred())
				logs := logger.Logs()
				Expect(logs).To(HaveLen(1))
				Expect(logs[0].LogLevel).To(Equal(lager.ERROR))
				Expect(logs[0].Message).To(ContainSubstring("concourse-client.delete-team.unknown-delete-error"))
				Expect(logs[0].Data["team-name"]).To(Equal("team venture"))
			})
		})

		Context("when I try to delete a team but I can't auth as an admin", func() {
			BeforeEach(func() {
				atcServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/api/v1/teams/main/auth/token"),
						ghttp.RespondWithJSONEncoded(http.StatusUnauthorized, nil),
					),
				)
			})
			It("returns an error", func() {
				client := NewClient(env, logger)
				err := client.DeleteTeam(cf.Details{OrgName: "team venture"})
				Expect(err).To(HaveOccurred())
				logs := logger.Logs()
				Expect(logs).To(HaveLen(1))
				Expect(logs[0].LogLevel).To(Equal(lager.ERROR))
				Expect(logs[0].Message).To(ContainSubstring("concourse-client.delete-team.auth-client-error"))
			})
		})
	})
})
