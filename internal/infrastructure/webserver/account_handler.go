package webserver

import (
	repository "count_on_us/internal/repositories/account"
	usecase "count_on_us/internal/usecases/account"
	dto "count_on_us/internal/usecases/account/dto"
	validator "count_on_us/pkg/validators"
	responses "count_on_us/pkg/webserver/responses"
	"encoding/json"
	"net/http"
)

type WebAccountHandler struct {
	accountRepository repository.AccountRepository
	validator         validator.ValidatorInterface
}

func NewWebAccountHandler(
	repo repository.AccountRepository,
	validator validator.ValidatorInterface) *WebAccountHandler {
	return &WebAccountHandler{
		accountRepository: repo,
		validator:         validator,
	}
}

func (wh *WebAccountHandler) Create(w http.ResponseWriter, r *http.Request) {
	createAccountDTO := dto.CreateAccountInputDTO{}
	_ = json.NewDecoder(r.Body).Decode(&createAccountDTO)
	if errors := wh.validator.Validate(createAccountDTO); len(errors) > 0 {
		responses.MalformedBody(w, errors)
		return
	}
	usecase := usecase.NewCreateAccountUseCase(&wh.accountRepository)
	response, err := usecase.Execute(createAccountDTO)
	if err != nil {
		responses.InternalServerError(w, []string{err.Error()})
		return
	}
	responses.Created(w, &response)
}
