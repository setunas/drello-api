package apperr

type AppError struct {
	tags       []Tag // Front tags have higher priorities.
	message    string
	occurredAt string
}

func (e *AppError) Error() string {
	return e.message
}

type Tag string

const (
	RecordNotFound       Tag = "recordNotFound"
	FailedAuthentication Tag = "FailedAuthentication"
	FailedAuthorization  Tag = "FailedAuthorization"
)

func (e *AppError) Tags() []Tag {
	return e.tags
}

func (e *AppError) OccurredAt() string {
	return e.occurredAt
}

func NewAppError(tags []Tag, message string, err error) error {
	appError, ok := err.(*AppError)
	if ok {
		return &AppError{
			tags:       inheritTags(tags, appError.tags),
			message:    combineMessages(message, err),
			occurredAt: appError.occurredAt,
		}
	}

	return &AppError{
		tags:       tags,
		message:    combineMessages(message, err),
		occurredAt: newOccurredAt(),
	}
}

func inheritTags(newTags []Tag, oldTags []Tag) []Tag {
	if len(oldTags) != 0 {
		return oldTags
	} else {
		return append(oldTags, newTags...)
	}
}
