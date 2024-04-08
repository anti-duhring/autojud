package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"

	"github.com/anti-duhring/autojud/internal/auth"
	graphql1 "github.com/anti-duhring/autojud/internal/generated/graphql"
	"github.com/anti-duhring/autojud/internal/graphql/resolvers/formatters"
	"github.com/google/uuid"
)

// FollowProcess is the resolver for the FollowProcess field.
func (r *mutationResolver) FollowProcess(ctx context.Context, processNumber string) (*graphql1.Process, error) {
	userID := auth.GetUserID(ctx)
	if userID == "" {
		return nil, fmt.Errorf("Access Denied")
	}

	uUserID, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("Access Denied")
	}

	process, err := r.ProcessService.GetByProcessNumber(processNumber, ctx)
	if err != nil {
		return nil, err
	}

	_, err = r.ProcessService.FollowProcess(process.ID, uUserID, ctx)
	if err != nil {
		return nil, err
	}

	fProcess := formatters.FormatProcess(process)

	return fProcess, nil
}

// GetProcessList is the resolver for the GetProcessList field.
func (r *queryResolver) GetProcessList(ctx context.Context, limit int, offset int) (*graphql1.ProcessList, error) {
	userID := auth.GetUserID(ctx)
	if userID == "" {
		return nil, fmt.Errorf("Access Denied")
	}

	uUserID, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("Access Denied")
	}

	processes, err := r.ProcessService.GetProcessFromUser(uUserID, limit, offset, ctx)
	if err != nil {
		return nil, err
	}

	count, err := r.ProcessService.CountProcessFromUser(uUserID, ctx)
	if err != nil {
		return nil, err
	}

	fProcesses := formatters.FormatProcessList(processes)

	return &graphql1.ProcessList{
		Nodes:       fProcesses,
		Count:       count,
		HasNextPage: count > limit+offset,
	}, nil
}

// Query returns graphql1.QueryResolver implementation.
func (r *Resolver) Query() graphql1.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
