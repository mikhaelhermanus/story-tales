package usecases

import (
	"context"

	"github.com/mfaizfatah/story-tales/app/models"
	"github.com/mfaizfatah/story-tales/app/repository"
)

// all variable const
const (
	// all variable for error
	ErrServer        = "Something wrong with our Server. Please try again later. Thank you"
	ErrNotFound      = "User not found"
	ErrCreated       = "Error when create a new User. Please try again later. Thank you"
	ErrAlreadyEmail  = "Email already created. Please use another e-mail. Thank you"
	ErrAlreadyPhone  = "Phone number already created. Please use another Phone Number. Thank you"
	ErrNotVerified   = "Your e-mail is not Verified"
	ErrBadRequest    = "Your Request is Invalid. Please check the payload"
	ErrNotMatch      = "Email or Password not match"
	ErrInvalidHeader = "Invalid Header"
	ErrTimezones     = "Timezone for Asia/Jakarta not found in our Server. Please try again later. Thank you"
	ErrEncryption    = "Encryption failed"
	ErrCreateToken   = "Error when create a token"

	// layout date
	LayoutDate = "2006-01-02 15:04:05"
)

// uc struct with value interface Repository
type uc struct {
	query repository.Repo
}

// Usecases represent the Usecases contract
type Usecases interface {
	// Authentication for logic auth
	Registration(ctx context.Context, req *models.User) (context.Context, *models.ResponseLogin, string, int, error)
	Login(ctx context.Context, req *models.User) (context.Context, *models.ResponseLogin, string, int, error)
}

/*NewUC will create an object that represent the Usecases interface (Usecases)
 * @parameter
 * r - Repository Interface
 *
 * @represent
 * interface Usecases
 *
 * @return
 * uc struct with value interface Repository
 */
func NewUC(r repository.Repo) Usecases {
	return &uc{query: r}
}
