// Code generated by ent, DO NOT EDIT.

package ent

import (
	"errors"
	"fmt"

	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent/predicate"
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent/user"
)

// UserWhereInput represents a where input for filtering User queries.
type UserWhereInput struct {
	Predicates []predicate.User  `json:"-"`
	Not        *UserWhereInput   `json:"not,omitempty"`
	Or         []*UserWhereInput `json:"or,omitempty"`
	And        []*UserWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty"`
	NameNEQ          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGT           *string  `json:"nameGT,omitempty"`
	NameGTE          *string  `json:"nameGTE,omitempty"`
	NameLT           *string  `json:"nameLT,omitempty"`
	NameLTE          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`

	// "email" field predicates.
	Email             *string  `json:"email,omitempty"`
	EmailNEQ          *string  `json:"emailNEQ,omitempty"`
	EmailIn           []string `json:"emailIn,omitempty"`
	EmailNotIn        []string `json:"emailNotIn,omitempty"`
	EmailGT           *string  `json:"emailGT,omitempty"`
	EmailGTE          *string  `json:"emailGTE,omitempty"`
	EmailLT           *string  `json:"emailLT,omitempty"`
	EmailLTE          *string  `json:"emailLTE,omitempty"`
	EmailContains     *string  `json:"emailContains,omitempty"`
	EmailHasPrefix    *string  `json:"emailHasPrefix,omitempty"`
	EmailHasSuffix    *string  `json:"emailHasSuffix,omitempty"`
	EmailEqualFold    *string  `json:"emailEqualFold,omitempty"`
	EmailContainsFold *string  `json:"emailContainsFold,omitempty"`

	// "password" field predicates.
	Password             *string  `json:"password,omitempty"`
	PasswordNEQ          *string  `json:"passwordNEQ,omitempty"`
	PasswordIn           []string `json:"passwordIn,omitempty"`
	PasswordNotIn        []string `json:"passwordNotIn,omitempty"`
	PasswordGT           *string  `json:"passwordGT,omitempty"`
	PasswordGTE          *string  `json:"passwordGTE,omitempty"`
	PasswordLT           *string  `json:"passwordLT,omitempty"`
	PasswordLTE          *string  `json:"passwordLTE,omitempty"`
	PasswordContains     *string  `json:"passwordContains,omitempty"`
	PasswordHasPrefix    *string  `json:"passwordHasPrefix,omitempty"`
	PasswordHasSuffix    *string  `json:"passwordHasSuffix,omitempty"`
	PasswordEqualFold    *string  `json:"passwordEqualFold,omitempty"`
	PasswordContainsFold *string  `json:"passwordContainsFold,omitempty"`

	// "avatar" field predicates.
	Avatar             *string  `json:"avatar,omitempty"`
	AvatarNEQ          *string  `json:"avatarNEQ,omitempty"`
	AvatarIn           []string `json:"avatarIn,omitempty"`
	AvatarNotIn        []string `json:"avatarNotIn,omitempty"`
	AvatarGT           *string  `json:"avatarGT,omitempty"`
	AvatarGTE          *string  `json:"avatarGTE,omitempty"`
	AvatarLT           *string  `json:"avatarLT,omitempty"`
	AvatarLTE          *string  `json:"avatarLTE,omitempty"`
	AvatarContains     *string  `json:"avatarContains,omitempty"`
	AvatarHasPrefix    *string  `json:"avatarHasPrefix,omitempty"`
	AvatarHasSuffix    *string  `json:"avatarHasSuffix,omitempty"`
	AvatarEqualFold    *string  `json:"avatarEqualFold,omitempty"`
	AvatarContainsFold *string  `json:"avatarContainsFold,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *UserWhereInput) AddPredicates(predicates ...predicate.User) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the UserWhereInput filter on the UserQuery builder.
func (i *UserWhereInput) Filter(q *UserQuery) (*UserQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyUserWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyUserWhereInput is returned in case the UserWhereInput is empty.
var ErrEmptyUserWhereInput = errors.New("ent: empty predicate UserWhereInput")

// P returns a predicate for filtering users.
// An error is returned if the input is empty or invalid.
func (i *UserWhereInput) P() (predicate.User, error) {
	var predicates []predicate.User
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, user.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.User, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, user.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.User, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, user.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, user.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, user.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, user.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, user.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, user.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, user.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, user.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, user.IDLTE(*i.IDLTE))
	}
	if i.Name != nil {
		predicates = append(predicates, user.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, user.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, user.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, user.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, user.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, user.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, user.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, user.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, user.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, user.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, user.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, user.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, user.NameContainsFold(*i.NameContainsFold))
	}
	if i.Email != nil {
		predicates = append(predicates, user.EmailEQ(*i.Email))
	}
	if i.EmailNEQ != nil {
		predicates = append(predicates, user.EmailNEQ(*i.EmailNEQ))
	}
	if len(i.EmailIn) > 0 {
		predicates = append(predicates, user.EmailIn(i.EmailIn...))
	}
	if len(i.EmailNotIn) > 0 {
		predicates = append(predicates, user.EmailNotIn(i.EmailNotIn...))
	}
	if i.EmailGT != nil {
		predicates = append(predicates, user.EmailGT(*i.EmailGT))
	}
	if i.EmailGTE != nil {
		predicates = append(predicates, user.EmailGTE(*i.EmailGTE))
	}
	if i.EmailLT != nil {
		predicates = append(predicates, user.EmailLT(*i.EmailLT))
	}
	if i.EmailLTE != nil {
		predicates = append(predicates, user.EmailLTE(*i.EmailLTE))
	}
	if i.EmailContains != nil {
		predicates = append(predicates, user.EmailContains(*i.EmailContains))
	}
	if i.EmailHasPrefix != nil {
		predicates = append(predicates, user.EmailHasPrefix(*i.EmailHasPrefix))
	}
	if i.EmailHasSuffix != nil {
		predicates = append(predicates, user.EmailHasSuffix(*i.EmailHasSuffix))
	}
	if i.EmailEqualFold != nil {
		predicates = append(predicates, user.EmailEqualFold(*i.EmailEqualFold))
	}
	if i.EmailContainsFold != nil {
		predicates = append(predicates, user.EmailContainsFold(*i.EmailContainsFold))
	}
	if i.Password != nil {
		predicates = append(predicates, user.PasswordEQ(*i.Password))
	}
	if i.PasswordNEQ != nil {
		predicates = append(predicates, user.PasswordNEQ(*i.PasswordNEQ))
	}
	if len(i.PasswordIn) > 0 {
		predicates = append(predicates, user.PasswordIn(i.PasswordIn...))
	}
	if len(i.PasswordNotIn) > 0 {
		predicates = append(predicates, user.PasswordNotIn(i.PasswordNotIn...))
	}
	if i.PasswordGT != nil {
		predicates = append(predicates, user.PasswordGT(*i.PasswordGT))
	}
	if i.PasswordGTE != nil {
		predicates = append(predicates, user.PasswordGTE(*i.PasswordGTE))
	}
	if i.PasswordLT != nil {
		predicates = append(predicates, user.PasswordLT(*i.PasswordLT))
	}
	if i.PasswordLTE != nil {
		predicates = append(predicates, user.PasswordLTE(*i.PasswordLTE))
	}
	if i.PasswordContains != nil {
		predicates = append(predicates, user.PasswordContains(*i.PasswordContains))
	}
	if i.PasswordHasPrefix != nil {
		predicates = append(predicates, user.PasswordHasPrefix(*i.PasswordHasPrefix))
	}
	if i.PasswordHasSuffix != nil {
		predicates = append(predicates, user.PasswordHasSuffix(*i.PasswordHasSuffix))
	}
	if i.PasswordEqualFold != nil {
		predicates = append(predicates, user.PasswordEqualFold(*i.PasswordEqualFold))
	}
	if i.PasswordContainsFold != nil {
		predicates = append(predicates, user.PasswordContainsFold(*i.PasswordContainsFold))
	}
	if i.Avatar != nil {
		predicates = append(predicates, user.AvatarEQ(*i.Avatar))
	}
	if i.AvatarNEQ != nil {
		predicates = append(predicates, user.AvatarNEQ(*i.AvatarNEQ))
	}
	if len(i.AvatarIn) > 0 {
		predicates = append(predicates, user.AvatarIn(i.AvatarIn...))
	}
	if len(i.AvatarNotIn) > 0 {
		predicates = append(predicates, user.AvatarNotIn(i.AvatarNotIn...))
	}
	if i.AvatarGT != nil {
		predicates = append(predicates, user.AvatarGT(*i.AvatarGT))
	}
	if i.AvatarGTE != nil {
		predicates = append(predicates, user.AvatarGTE(*i.AvatarGTE))
	}
	if i.AvatarLT != nil {
		predicates = append(predicates, user.AvatarLT(*i.AvatarLT))
	}
	if i.AvatarLTE != nil {
		predicates = append(predicates, user.AvatarLTE(*i.AvatarLTE))
	}
	if i.AvatarContains != nil {
		predicates = append(predicates, user.AvatarContains(*i.AvatarContains))
	}
	if i.AvatarHasPrefix != nil {
		predicates = append(predicates, user.AvatarHasPrefix(*i.AvatarHasPrefix))
	}
	if i.AvatarHasSuffix != nil {
		predicates = append(predicates, user.AvatarHasSuffix(*i.AvatarHasSuffix))
	}
	if i.AvatarEqualFold != nil {
		predicates = append(predicates, user.AvatarEqualFold(*i.AvatarEqualFold))
	}
	if i.AvatarContainsFold != nil {
		predicates = append(predicates, user.AvatarContainsFold(*i.AvatarContainsFold))
	}

	switch len(predicates) {
	case 0:
		return nil, ErrEmptyUserWhereInput
	case 1:
		return predicates[0], nil
	default:
		return user.And(predicates...), nil
	}
}
