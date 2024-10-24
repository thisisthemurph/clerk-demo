package main

import (
	"context"
	"github.com/clerk/clerk-sdk-go/v2/user"
	"net/http"
	"os"

	"github.com/clerk/clerk-sdk-go/v2"
	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ContextKey string

const UserContextKey ContextKey = "user"

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	clerkKey := os.Getenv("CLERK_KEY")
	clerk.SetKey(clerkKey)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:*"},
		AllowHeaders: []string{
			echo.HeaderAuthorization,
			echo.HeaderAccept,
			"Host",
			echo.HeaderOrigin,
			"Referer",
			"Sec-Fetch-Dest",
			"User-Agent",
			"X-Forwarded-Host",
			"X-Forwarded-Proto",
			echo.HeaderContentType,
		},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete},
		AllowCredentials: true,
	}))

	e.GET("/public", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Hello, World!",
		})
	})

	e.GET("/private", clerkMiddleware(withClerkUserInContext(func(c echo.Context) error {
		usr := clerkUser(c)
		if !usr.Authenticated {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Hello, authenticated ones!",
			"user":    usr.FirstName,
			"email":   usr.EmailAddresses[0].EmailAddress,
		})
	})))

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

type ClerkUserSession struct {
	clerk.User
	Authenticated bool
}

func NewClerkUserSession(u clerk.User) ClerkUserSession {
	return ClerkUserSession{
		User:          u,
		Authenticated: true,
	}
}

func clerkUser(c echo.Context) ClerkUserSession {
	if u, ok := c.Get(string(UserContextKey)).(clerk.User); ok {
		return NewClerkUserSession(u)
	}
	return ClerkUserSession{}
}

func withClerkUserInContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		claims, ok := clerk.SessionClaimsFromContext(ctx)
		if !ok {
			return next(c)
		}

		usr, err := user.Get(ctx, claims.Subject)
		if err != nil {
			return next(c)
		}
		if usr == nil {
			return next(c)
		}

		c.Set(string(UserContextKey), *usr)
		c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), UserContextKey, *usr)))
		return next(c)
	}
}

func clerkMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		handler := clerkhttp.WithHeaderAuthorization()(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.SetRequest(r)
			if err := next(c); err != nil {
				c.Error(err)
				return
			}
		}))

		handler.ServeHTTP(c.Response(), c.Request())

		if !c.Response().Committed {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Unauthorized or missing session",
			})
		}

		return nil
	}
}
