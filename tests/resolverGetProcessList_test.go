package tests

import (
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
		GetProcessList genGraphql.ProcessList
	}

	makeRequest := func(options ...client.Option) (resp, error) {
		var resp resp
		err := c.Post(`query GetProcessList($limit: Int!, $offset: Int!) {
		    GetProcessList(limit: $limit, offset: $offset) {
          nodes {
            id
            processNumber
          }
          count
          hasNextPage
        }
      }`,
			&resp,
			options...,
		)
		return resp, err
	}

	It("list 10 process from a user", func() {
		var processesList []*processes.Process
		for i := 0; i < 10; i++ {
			processesList = append(processesList, &processes.Process{
				ID:            uuid.New(),
				ProcessNumber: "123",
			})
		}

		processRepo.(*mocks.MockRepositoryProcesses).Mock.On("GetAllByUserID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(processesList, nil)

		processRepo.(*mocks.MockRepositoryProcesses).Mock.On("CountByUserID", mock.Anything, mock.Anything).Return(20, nil)

		resp, err := makeRequest(
			mocks.MockToken(),
			mocks.MockUserID(uuid.NewString()),
			client.Var("limit", 10),
			client.Var("offset", 0),
		)
		Expect(err).ToNot(HaveOccurred())

		Expect(resp.GetProcessList.Count).To(Equal(20))
		Expect(resp.GetProcessList.HasNextPage).To(BeTrue())
		Expect(len(resp.GetProcessList.Nodes)).To(Equal(10))
	})
})
