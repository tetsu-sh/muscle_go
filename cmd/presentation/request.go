package presentation

import "github.com/labstack/echo/v4"

type trainCreateRequest struct {
	Train struct {
		TrainName string `json:"trainname"`
		Volume    int    `json:"volume"`
		Set       int    `json:"set"`
		Rep       int    `json:"rep"`
	} `json:"user"`
}

func newTrainCreateRequest() *trainCreateRequest {
	return new(trainCreateRequest)
}

func (r *trainCreateRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil

}
