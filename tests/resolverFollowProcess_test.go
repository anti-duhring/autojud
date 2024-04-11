package tests

import (
	"database/sql"
	"time"

	"github.com/99designs/gqlgen/client"
	genGraphql "github.com/anti-duhring/autojud/internal/generated/graphql"
	"github.com/anti-duhring/autojud/internal/processes"
	"github.com/anti-duhring/autojud/tests/mocks"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("resolverFollowProcess", func() {
	type resp struct {
		FollowProcess genGraphql.Process
	}

	makeRequest := func(options ...client.Option) (resp, error) {
		var resp resp
		err := c.Post(`mutation FollowProcess($processNumber: String!) {
			FollowProcess(processNumber: $processNumber) {
        id
        processNumber
        court
        origin 
        judge
        activePart
        passivePart
        createdAt
        updatedAt
        deletedAt
      } 
		}`,
			&resp,
			options...,
		)
		return resp, err
	}

	It("follows a process", func() {
		userID := uuid.New()
		processID := uuid.New()
		origin := "origin"
		judge := "judge"
		activePart := "activePart"
		passivePart := "passivePart"
		processNumber := "123"

		processRepo.(*mocks.MockRepositoryProcesses).Mock.On("GetByProcessNumber", mock.Anything, mock.Anything).Return(&processes.Process{
			ID:            processID,
			ProcessNumber: processNumber,
			Court:         processes.COURT_TJPE,
			Origin:        &origin,
			Judge:         &judge,
			ActivePart:    &activePart,
			PassivePart:   &passivePart,
			CreatedAt:     time.Now().String(),
			UpdatedAt:     time.Now().String(),
			DeletedAt:     nil,
		}, nil)

		processRepo.(*mocks.MockRepositoryProcesses).Mock.On("CreateProcessFollow", mock.Anything, mock.Anything, mock.Anything).Return(&processes.ProcessFollow{
			ID:        uuid.New(),
			UserID:    userID,
			ProcessID: processID,
			CreatedAt: time.Now().String(),
			DeletedAt: nil,
		}, nil)

		resp, err := makeRequest(
			mocks.MockToken(),
			mocks.MockUserID(userID.String()),
			client.Var("processNumber", processNumber),
		)
		Expect(err).ToNot(HaveOccurred())

		Expect(resp.FollowProcess.ID).To(Equal(processID.String()))
		Expect(resp.FollowProcess.ProcessNumber).To(Equal(processNumber))
		Expect(resp.FollowProcess.Court).To(Equal(genGraphql.CourtTjpe))
		Expect(resp.FollowProcess.Origin).To(Equal(&origin))
		Expect(resp.FollowProcess.Judge).To(Equal(&judge))
		Expect(resp.FollowProcess.ActivePart).To(Equal(&activePart))
		Expect(resp.FollowProcess.PassivePart).To(Equal(&passivePart))
		Expect(resp.FollowProcess.CreatedAt).ToNot(BeNil())
		Expect(resp.FollowProcess.UpdatedAt).ToNot(BeNil())
		Expect(resp.FollowProcess.DeletedAt).To(BeNil())
	})
	It("inserts on pending_process table when trying to follow a process that does not exists yet", func() {
		userID := uuid.New()
		processID := uuid.New()
		processNumber := "0000000-00.0000.0.17.0000"

		processRepo.(*mocks.MockRepositoryProcesses).Mock.On("GetByProcessNumber", mock.Anything, mock.Anything).Return(nil, sql.ErrNoRows)

		processRepo.(*mocks.MockRepositoryProcesses).Mock.On("CreateProcess", mock.Anything, mock.Anything).Return(&processes.Process{
			ID:            processID,
			ProcessNumber: processNumber,
			Court:         processes.COURT_TJPE,
		}, nil).Run(func(args mock.Arguments) {
			process := args.Get(1).(*processes.Process)

			Expect(process.ProcessNumber).To(Equal(processNumber))
			Expect(process.Court).To(Equal(processes.COURT_TJPE))
		})

		processRepo.(*mocks.MockRepositoryProcesses).Mock.On("CreatePendingProcess", mock.Anything, mock.Anything).Return(&processes.PendingProcess{
			ID:        uuid.New(),
			ProcessID: processID,
		}, nil).Run(func(args mock.Arguments) {
			createdProcessID := args.Get(1).(string)
			Expect(createdProcessID).To(Equal(processID.String()))
		})

		processRepo.(*mocks.MockRepositoryProcesses).Mock.On("CreateProcessFollow", mock.Anything, mock.Anything, mock.Anything).Return(&processes.ProcessFollow{
			ID:        uuid.New(),
			UserID:    userID,
			ProcessID: processID,
			CreatedAt: time.Now().String(),
			DeletedAt: nil,
		}, nil)

		resp, err := makeRequest(
			mocks.MockToken(),
			mocks.MockUserID(userID.String()),
			client.Var("processNumber", processNumber),
		)
		Expect(err).ToNot(HaveOccurred())

		Expect(resp.FollowProcess.ID).To(Equal(processID.String()))
		Expect(resp.FollowProcess.ProcessNumber).To(Equal(processNumber))
		Expect(resp.FollowProcess.Court).To(Equal(genGraphql.CourtTjpe))
	})
})
