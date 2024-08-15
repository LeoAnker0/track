package responses

import (

)

/*
    errorMessages
        1 = token no longer valid, relogin


*/

type JSONMessage struct {
    Message string `json:"message"`
    HTTPCode int    `json:"httpCode"`
    ErrorCode string `json:"errorCode,omitempty"`
    ErrorMessage string `json:"errorMessage,omitempty"`
    Data     interface{} `json:"data,omitempty"`
}



