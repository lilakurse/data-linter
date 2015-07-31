package mock

import "github.com/GabbyyLS/data-linter/models"

var (
	EmptyProblemList = []*models.Problem{}
	Problems         = []*models.Problem{{Original: "A file with a problem"}, {Original: "Another file with a problem"}}
)
