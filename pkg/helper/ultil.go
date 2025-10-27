package helper

import (
	"context"
	gw_response "profile-service/internal/gateway/dto/response"
	"profile-service/pkg/constants"
	"strconv"
	"strings"
)

func CurrentUserFromCtx(ctx context.Context) (*gw_response.CurrentUser, bool) {
	if cu, ok := ctx.Value(constants.CurrentUserKey).(*gw_response.CurrentUser); ok {
		return cu, true
	}
	return nil, false
}

func ParseAppLanguage(header string, defaultVal uint) uint {
	header = strings.TrimSpace(strings.Trim(header, "\""))
	if val, err := strconv.Atoi(header); err == nil {
		return uint(val)
	}
	return defaultVal
}

func GetHeaders(ctx context.Context) map[string]string {
	headers := make(map[string]string)

	if lang, ok := ctx.Value(constants.AppLanguage).(uint); ok {
		headers["X-App-Language"] = strconv.Itoa(int(lang))
	}

	return headers
}

func GetAppLanguage(ctx context.Context, defaultVal uint) uint {
	if lang, ok := ctx.Value(constants.AppLanguage).(uint); ok {
		return lang
	}
	return defaultVal
}
