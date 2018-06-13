package checker

import (
	"errors"
)

var (
	ErrInvalidValue  = errors.New("invalid value")
	ErrIsNil  = errors.New("the value can not be nil")

	ErrRuleGt  = errors.New("the value can not match gt rule")
	ErrRuleGte = errors.New("the value can not match gte rule")
	ErrRuleLt  = errors.New("the value can not match lt rule")
	ErrRuleLte = errors.New("the value can not match lte rule")

	ErrRuleLgt  = errors.New("the length can not match lgt rule")
	ErrRuleLgte = errors.New("the length can not match lgte rule")
	ErrRuleLlt  = errors.New("the length can not match llt rule")
	ErrRuleLlte = errors.New("the length can not match llte rule")
)
