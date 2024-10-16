package storage

import (
	"go.uber.org/zap"
	"strconv"
)

type value struct {
	stringValue string
	intValue    int
	valueType   ValueType
}

type ValueType string

const (
	ValueTypeInt       ValueType = "D"
	ValueTypeString    ValueType = "S"
	ValueTypeUndefined ValueType = "U"
)

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
	valueType := evaluateType(v)

	var val value
	switch valueType {
	case ValueTypeInt:
		intValue, _ := strconv.Atoi(v)
		val = value{
			stringValue: v,
			intValue:    intValue,
			valueType:   valueType,
		}
	case ValueTypeString:
		val = value{
			stringValue: v,
			valueType:   valueType,
		}
	case ValueTypeUndefined:
		s.logger.Error(
			"trying to set value of unknown type",
			zap.String("type", string(valueType)),
			zap.String("value", v),
			zap.String("key", k),
		)
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

func (s Storage) GetType(k string) ValueType {
	result, ok := s.inner[k]
	if !ok {
		return ValueTypeUndefined
	}

	return result.valueType
}

func evaluateType(v string) ValueType {
	_, err := strconv.Atoi(v)
	if err != nil {
		return ValueTypeString
	}

	return ValueTypeInt
}
