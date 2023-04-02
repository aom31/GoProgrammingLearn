package main

import "net/http"

func main() {
	ms := NewMicroservice()

	ms.POST("/citizen", func(ctx IContext) error {
		ctx.log("POST: /citizen")
		status := map[string]interface{}{
			"status": "success",
		}
		ctx.Response(http.StatusOK, status)
		return nil
	})
}
