package nds

import (
	"github.com/qedus/nds/v2/internal/log"
)

func SetLogger(logger log.Logging) {
	log.Logger = logger
}
