package http

import (
	"net/http"

	"hexagonal/template/internal/src/iteractor"
	"hexagonal/template/internal/src/transport_layer/http/model"

	"github.com/joomcode/errorx"
	"github.com/labstack/echo/v4"
)

type DummyController struct {
	useCase iteractor.IDummyIteractor
}

func ConstructorDummyController() *DummyController {
	return &DummyController{
		useCase: iteractor.ConstructorDummyIteractor(),
	}
}

func (c *DummyController) GetDummy(echoContext echo.Context) error {
	dummyRequest := &model.DummyRequest{}

	requestError := validateDummyRequest(echoContext, dummyRequest)
	if requestError != nil {
		return model.ResponseError(echoContext, requestError)
	}

	response, count, iteractorError := c.useCase.GetDummy(dummyRequest.ToDummy(), dummyRequest.Limit, dummyRequest.Offset)
	if iteractorError != nil {
		return model.ResponseError(echoContext, errorx.InternalError.Wrap(iteractorError, "Iteractor error exception"))
	}

	return echoContext.JSON(http.StatusOK, model.ToPageable(response, len(response), count, int(dummyRequest.Offset)))
}

func validateDummyRequest(echoContext echo.Context, request *model.DummyRequest) *errorx.Error {
	if error := echoContext.Bind(request); error != nil {
		return errorx.IllegalArgument.Wrap(error, "error on request bind")
	}

	if error := request.IsValid(); error != nil {
		return errorx.IllegalArgument.Wrap(error, "error on request validation")
	}

	if request.Limit == 0 {
		request.Limit = model.MAX_LIMIT
	}

	return nil
}
