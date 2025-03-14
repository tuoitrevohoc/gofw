package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent"
	usrPkg "github.com/tuoitrevohoc/gofw/backend/gen/go/ent/user"
	graphql1 "github.com/tuoitrevohoc/gofw/backend/gen/go/graphql"
	"github.com/tuoitrevohoc/gofw/backend/gen/go/graphql/model"
	"github.com/tuoitrevohoc/gofw/backend/internal/auth"
	"github.com/tuoitrevohoc/gofw/backend/package/scalars"
	"golang.org/x/crypto/bcrypt"
)

// SignUp is the resolver for the signUp field.
func (r *mutationResolver) SignUp(ctx context.Context, input model.SignUpInput) (*model.AccessToken, error) {
	hasUser, err := r.client.User.Query().Where(
		usrPkg.Email(input.Email),
	).Exist(ctx)

	if err != nil {
		return nil, err
	}

	if hasUser {
		return nil, fmt.Errorf("user already exists")
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user, err := r.client.User.Create().SetEmail(input.Email).SetPassword(string(encryptedPassword)).SetFinishedRegistration(true).Save(ctx)
	if err != nil {
		return nil, err
	}

	token, err := r.authenticationService.Login(ctx, user)
	if err != nil {
		return nil, err
	}

	return &model.AccessToken{
		AccessToken: token.AccessToken,
		Expiry:      int(token.Expiry),
		Viewer: &model.Viewer{
			UserID:          scalars.NewGUID(usrPkg.Table, user.ID),
			Profile:         user,
			IsAuthenticated: true,
		},
	}, nil
}

// SignIn is the resolver for the signIn field.
func (r *mutationResolver) SignIn(ctx context.Context, input model.SignInInput) (*model.AccessToken, error) {
	user, err := r.client.User.Query().Where(usrPkg.Email(input.Email)).Only(ctx)
	if err != nil {
		return nil, err
	}

	passwordMatch := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if passwordMatch != nil {
		return nil, fmt.Errorf("invalid password")
	}

	token, err := r.authenticationService.Login(ctx, user)
	if err != nil {
		return nil, err
	}

	return &model.AccessToken{
		AccessToken: token.AccessToken,
		Expiry:      int(token.Expiry),
		Viewer: &model.Viewer{
			UserID:          scalars.NewGUID(usrPkg.Table, user.ID),
			Profile:         user,
			IsAuthenticated: true,
		},
	}, nil
}

// BeginAuthnRegistration is the resolver for the beginAuthnRegistration field.
func (r *mutationResolver) BeginAuthnRegistration(ctx context.Context, email string) (*model.AuthnRegistrationResponse, error) {
	credientialCreation, err := r.authenticator.BeginRegistration(ctx, email)

	if err != nil {
		return nil, err
	}

	// convert credential to json
	json, err := json.Marshal(credientialCreation)
	if err != nil {
		return nil, fmt.Errorf("something went wrong")
	}

	return &model.AuthnRegistrationResponse{
		CredentialCreation: string(json),
	}, nil
}

// FinishAuthnRegistration is the resolver for the finishAuthnRegistration field.
func (r *mutationResolver) FinishAuthnRegistration(ctx context.Context, email string, response string) (*model.AccessToken, error) {
	parsedResponse, err := protocol.ParseCredentialCreationResponseBytes([]byte(response))

	if err != nil {
		return nil, fmt.Errorf("fail to parse credential creation response")
	}

	user, err := r.authenticator.FinishRegistration(ctx, email, parsedResponse)
	if err != nil {
		return nil, fmt.Errorf("fail to finish passkey registration")
	}

	token, err := r.authenticationService.Login(ctx, user)
	if err != nil {
		return nil, err
	}

	return &model.AccessToken{
		AccessToken: token.AccessToken,
		Expiry:      int(token.Expiry),
		Viewer: &model.Viewer{
			UserID:          scalars.NewGUID(usrPkg.Table, user.ID),
			Profile:         user,
			IsAuthenticated: true,
		},
	}, nil
}

// BeginAuthnLogin is the resolver for the beginAuthnLogin field.
func (r *mutationResolver) BeginAuthnLogin(ctx context.Context) (*model.AuthnLoginResponse, error) {
	credentialRequest, err := r.authenticator.BeginLogin(ctx)
	if err != nil {
		return nil, err
	}

	// convert credential to json
	json, err := json.Marshal(credentialRequest)
	if err != nil {
		return nil, fmt.Errorf("something went wrong")
	}

	return &model.AuthnLoginResponse{
		CredentialRequest: string(json),
	}, nil
}

// FinishAuthnLogin is the resolver for the finishAuthnLogin field.
func (r *mutationResolver) FinishAuthnLogin(ctx context.Context, response string) (*model.AccessToken, error) {
	parsedResponse, err := protocol.ParseCredentialRequestResponseBytes([]byte(response))
	if err != nil {
		return nil, fmt.Errorf("fail to parse credential request response")
	}

	user, err := r.authenticator.FinishLogin(ctx, parsedResponse)
	if err != nil {
		return nil, fmt.Errorf("fail to finish passkey login")
	}

	token, err := r.authenticationService.Login(ctx, user)
	if err != nil {
		return nil, err
	}

	return &model.AccessToken{
		AccessToken: token.AccessToken,
		Expiry:      int(token.Expiry),
		Viewer: &model.Viewer{
			UserID:          scalars.NewGUID(usrPkg.Table, user.ID),
			Profile:         user,
			IsAuthenticated: true,
		},
	}, nil
}

// SignOut is the resolver for the signOut field.
func (r *mutationResolver) SignOut(ctx context.Context) (bool, error) {
	err := r.authenticationService.Logout(ctx)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Viewer is the resolver for the viewer field.
func (r *queryResolver) Viewer(ctx context.Context) (*model.Viewer, error) {
	return auth.ViewerFromContext(ctx), nil
}

// Profile is the resolver for the profile field.
func (r *viewerResolver) Profile(ctx context.Context, obj *model.Viewer) (*ent.User, error) {
	if !obj.IsAuthenticated {
		return nil, fmt.Errorf("user is not authenticated")
	}

	return r.client.User.Get(ctx, obj.UserID.Id())
}

// Viewer returns graphql1.ViewerResolver implementation.
func (r *Resolver) Viewer() graphql1.ViewerResolver { return &viewerResolver{r} }

type viewerResolver struct{ *Resolver }
