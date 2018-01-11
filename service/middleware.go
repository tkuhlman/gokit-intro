package service

import (
	"context"

	"github.com/sirupsen/logrus"
)

// ContentMiddleware defines the function specification used to add service level middleware.
type ContentMiddleware func(Content) Content

// LogMiddleware adds logging middleware to the service.
func LogMiddleware(log *logrus.Logger) ContentMiddleware {
	return func(next Content) Content {
		return logMiddleware{log: log, next: next}
	}
}

// logMiddleware is content service specific middleware that follows the Content interface adding its logic by
// wrapping all of the defined methods.
type logMiddleware struct {
	log  *logrus.Logger
	next Content
}

func (mw logMiddleware) Query(ctx context.Context, query string) (string, error) {
	// Logging at this level is aware of the specific method being called
	mw.log.Debugf("Running query %q", query)
	result, err := mw.next.Query(ctx, query)
	if err != nil {
		mw.log.Errorf("failed query: %v", err)
	} else {
		mw.log.Infof("Processed query: %q", query)
	}
	return result, err
}
