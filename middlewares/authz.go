package middlewares

import "github.com/labstack/echo/v4"

func RoleBasedAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get the student payload from the context (set by JWTAuth middleware)
		// studentPayload := c.Get("studentPayload").(*entity.StudentPayloadJWT)

		// Check if the role is "professor"
		// if studentPayload.Role != "professor" {
		// 	return c.JSON(http.StatusForbidden, echo.Map{
		// 		"error": "access forbidden: only professors can access this resource",
		// 	})
		// }

		// If the role is correct, continue to the next handler
		return next(c)
	}
}
