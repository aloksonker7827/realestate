package utils

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/google/uuid"
	"gitlab.com/umi7/DezHab_user_backend/common/logger"
	"gitlab.com/umi7/DezHab_user_backend/utils/constants"
	"math/rand"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

var UUID = func() string {
	return uuid.New().String()
}

var GoPath = os.Getenv(constants.GoPath)

var Environment = func() string {
	env := os.Getenv(constants.Environment)
	if env == "" {
		env = constants.DefaultEnvironment
	}
	return strings.ToLower(env)
}

var IsProduction = func() bool {
	return Environment() == constants.ProductionEnvironment
}

var GetProviderSubTagFromCtx = func(ctx context.Context) (provider string, subTag string) {
	if ctx.Value(constants.Provider) != nil {
		provider = ctx.Value(constants.Provider).(string)
	}
	if ctx.Value(constants.SubTag) != nil {
		subTag = ctx.Value(constants.SubTag).(string)
	}
	return
}

//This is used to crete the objectId in sake.
var CreateObjectID = func(obj string) string {
	h := sha256.New()
	h.Write([]byte(obj))
	sha256Hash := hex.EncodeToString(h.Sum(nil))
	return sha256Hash[0:24]
}

var RecoverFromPanic = func(ctx context.Context, err interface{}) (interface{}, bool) {
	if err != nil {
		logger.WithAlert(ctx, err, "Panic Recovered!: %v \n %v", err, string(debug.Stack()))
		return err, true
	}
	return nil, false
}

var GetUIntFromContext = func(ctx context.Context, key string) (val uint) {
	if ctx.Value(key) != nil {
		val = ctx.Value(key).(uint)
	}
	return
}

var GetRequestID = func(ctx context.Context) string {
	requestId := ctx.Value(constants.RequestID)
	if requestId != nil {
		return requestId.(string)
	}
	return ""
}

var GetRequestData = func(ctx context.Context) string {
	requestData := ctx.Value(constants.RequestData)
	if requestData != nil {
		return requestData.(string)
	}
	return ""
}

var GetUploadedBy = func(ctx context.Context) string {
	uploadedBy := ctx.Value(constants.UploadedBy)
	if uploadedBy != nil {
		return uploadedBy.(string)
	}
	return ""
}

var GetAction = func(ctx context.Context) string {
	action := ctx.Value(constants.Action)
	if action != nil {
		return action.(string)
	}
	return ""
}

var GetResult = func(ctx context.Context) string {
	result := ctx.Value(constants.Result)
	if result != nil {
		return result.(string)
	}
	return ""
}

var GetUniqueID = func(ctx context.Context) string {
	requestId := ctx.Value(constants.RequestID)
	var fileName string
	if requestId != nil {
		fileName = requestId.(string)
	} else {
		fileName = UUID()
	}
	return fileName
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		rand.Seed(time.Now().UnixNano())
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Find(value string, arr []string) (int, bool) {
	for index, item := range arr {
		if item == value {
			return index, true
		}
	}
	return -1, false
}
