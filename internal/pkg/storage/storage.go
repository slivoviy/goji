package storage

import (
	"go.uber.org/zap"
	"strconv"
)

type value struct {
	stringValue string
	intValue    int
	valueType   string
}

type Storage struct {
	inner  map[string]value
	logger *zap.Logger
}

func NewStorage() (Storage, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return Storage{}, err
	}

	defer logger.Sync()
	logger.Info("storage created")

	return Storage{
		inner:  make(map[string]value),
		logger: logger,
	}, nil
}

func (s Storage) Set(k, v string) {
	valueType := checkType(v)

	var val value
	if valueType == "D" {
		intValue, _ := strconv.Atoi(v)
		val = value{
			stringValue: v,
			intValue:    intValue,
			valueType:   valueType,
		}
	} else {
		val = value{
			stringValue: v,
			valueType:   valueType,
		}
	}

	s.inner[k] = val

	s.logger.Info("value set", zap.String("key", k), zap.String("value", val.stringValue))
	s.logger.Sync()
}

func (s Storage) Get(k string) *string {
	result, ok := s.inner[k]
	if !ok {
		return nil
	}

	s.logger.Info("value got", zap.String("key", k), zap.String("value", result.stringValue))

	return &result.stringValue
}

func (s Storage) GetType(k string) string {
	result, ok := s.inner[k]
	if !ok {
		return "No"
	}

	return result.valueType
}

func checkType(v string) string {
	if _, err := strconv.Atoi(v); err == nil {
		return "D"
	} else {
		return "S"
	}
}
