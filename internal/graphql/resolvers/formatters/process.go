package formatters

import (
	"github.com/anti-duhring/autojud/internal/generated/graphql"
	"github.com/anti-duhring/autojud/internal/processes"
)

var courtMap map[processes.Court]graphql.Court = map[processes.Court]graphql.Court{
	processes.COURT_TJPE: graphql.CourtTjpe,
}

func FormatProcess(process *processes.Process) *graphql.Process {
	if process == nil {
		return nil
	}

	court := courtMap[process.Court]

	return &graphql.Process{
		ID:            process.ID.String(),
		ProcessNumber: process.ProcessNumber,
		Court:         court,
		Origin:        process.Origin,
		Judge:         process.Judge,
		ActivePart:    process.ActivePart,
		PassivePart:   process.PassivePart,
		CreatedAt:     process.CreatedAt,
		UpdatedAt:     process.UpdatedAt,
		DeletedAt:     process.DeletedAt,
	}
}
