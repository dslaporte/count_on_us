package webserver

import (
	repository "count_on_us/internal/repositories/account"
	usecase "count_on_us/internal/usecases/account"
	usecases "count_on_us/internal/usecases/account"
	dto "count_on_us/internal/usecases/account/dto"
	pkg_strings "count_on_us/pkg/strings"
	validator "count_on_us/pkg/validators"
	responses "count_on_us/pkg/webserver/responses"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

const (
	errRequiredParamId = "param id is required!"
	errAccountNotFound = "account not found!"
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

func (wh *WebAccountHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if pkg_strings.IsEmpty(id) {
		responses.BadRequest(w, []string{errRequiredParamId})
		return
	}
	usecase := usecase.NewFindAccountUseCase(&wh.accountRepository)
	response, err := usecase.Execute(id)
	if err != nil {
		responses.NotFound(w, []string{errAccountNotFound})
		return
	}
	responses.OK(w, response)
}

func (wh *WebAccountHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if pkg_strings.IsEmpty(id) {
		responses.BadRequest(w, []string{errRequiredParamId})
		return
	}
	updateAccountInputDTO := dto.UpdateAccountInputDTO{}
	if errors := wh.validator.Validate(updateAccountInputDTO); len(errors) > 0 {
		responses.MalformedBody(w, errors)
		return
	}
	ucFindAccount := usecases.NewFindAccountUseCase(&wh.accountRepository)
	_, err := ucFindAccount.Execute(id)
	if err != nil {
		responses.NotFound(w, []string{errAccountNotFound})
		return
	}
	ucUpdateAccount := usecases.NewUpdateAccountUseCase(&wh.accountRepository)
	response, err := ucUpdateAccount.Execute(id, updateAccountInputDTO)
	if err != nil {
		responses.InternalServerError(w, []string{err.Error()})
	}
	responses.OK(w, response)
}

func (wh *WebAccountHandler) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("estou no list")
}
