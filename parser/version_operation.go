package parser

import (
	"github.com/hashicorp/go-version"
)

type VersionOperation struct {
	NullOperation
}

func (v *VersionOperation) get(left Operand, right Operand) (*version.Version, *version.Version, error) {
	var leftVer, rightVer *version.Version
	leftVal, ok := left.(string)
	if !ok {
		return nil, nil, newErrInvalidOperand(left, leftVal)
	}
	rightVal, ok := right.(string)
	if !ok {
		return nil, nil, newErrInvalidOperand(right, rightVal)
	}
	var err error
	leftVer, err = version.NewVersion(leftVal)
	if err != nil {
		return nil, nil, err
	}
	rightVer, err = version.NewVersion(rightVal)
	return leftVer, rightVer, err
}

func (v *VersionOperation) EQ(left Operand, right Operand) (bool, error) {
	leftVer, rightVer, err := v.get(left, right)
	if err != nil {
		return false, err
	}
	return leftVer.Equal(rightVer), nil
}

func (v *VersionOperation) NE(left Operand, right Operand) (bool, error) {
	leftVer, rightVer, err := v.get(left, right)
	if err != nil {
		return false, err
	}
	return !leftVer.Equal(rightVer), nil
}

func (v *VersionOperation) LT(left Operand, right Operand) (bool, error) {
	leftVer, rightVer, err := v.get(left, right)
	if err != nil {
		return false, err
	}
	return leftVer.LessThan(rightVer), nil
}

func (v *VersionOperation) GT(left Operand, right Operand) (bool, error) {
	leftVer, rightVer, err := v.get(left, right)
	if err != nil {
		return false, err
	}
	return leftVer.GreaterThan(rightVer), nil
}

func (v *VersionOperation) LE(left Operand, right Operand) (bool, error) {
	leftVer, rightVer, err := v.get(left, right)
	if err != nil {
		return false, err
	}
	return leftVer.LessThanOrEqual(rightVer), nil
}

func (v *VersionOperation) GE(left Operand, right Operand) (bool, error) {
	leftVer, rightVer, err := v.get(left, right)
	if err != nil {
		return false, err
	}
	return leftVer.GreaterThanOrEqual(rightVer), nil
}
