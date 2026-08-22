package db

// import (
// 	"context"
// 	"errors"
// 	"kisaanSathi/pkg/logger"
// 	"kisaanSathi/pkg/services/user/models"
// 	"time"

// 	"github.com/golang-jwt/jwt/v5"
// 	"go.uber.org/zap"
// 	"gorm.io/gorm"
// )

// var jwtSecret = []byte("abcde") // TODO: store in env

// // Login authenticates a user and returns JWT token
// func (g *registerStore) Login(c context.Context, email string, password string) (*models.LoginResponse, error) {
// 	logger.Log(c).Debug("START Login")
// 	defer logger.Log(c).Debug("END Login")

// 	// Check if database is available
// 	if g.store == nil {
// 		logger.Log(c).Warn("Database unavailable - returning mock login data")
// 		return &models.LoginResponse{
// 			Token: "mock_token_" + email,
// 			User: &models.UserDetails{
// 				ID:    1,
// 				Name:  "Mock User",
// 				Email: email,
// 				Role:  "user",
// 			},
// 		}, nil
// 	}

// 	var user models.UserDetails
// 	err := g.store.WithContext(c).Table("kisan.users").
// 		Where("email = ?", email).
// 		First(&user).Error
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			logger.Log(c).Error("User not found", zap.String("email", email))
// 			return nil, errors.New("invalid email or password")
// 		}
// 		logger.Log(c).Error("error while fetching user data", zap.Error(err))
// 		return nil, errors.New("invalid email or password")
// 	}

// 	// Compare hashed password
// 	// if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
// 	// 	return nil, errors.New("invalid email or password")
// 	// }

// 	// Generate JWT token
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"user_id": user.ID,
// 		"email":   user.Email,
// 		"role":    user.Role,
// 		"exp":     time.Now().Add(24 * time.Hour).Unix(),
// 	})

// 	tokenString, err := token.SignedString(jwtSecret)
// 	if err != nil {
// 		logger.Log(c).Error("JWT generation failed", zap.Error(err))
// 		return nil, err
// 	}

// 	resp := &models.LoginResponse{
// 		Token: tokenString,
// 		Name:  user.Name,
// 		Role:  user.Role,
// 		Email: user.Email,
// 	}

// 	return resp, nil
// }

// func (g *registerStore) Register(c context.Context, phone, email, password string) error {
// 	logger.Log(c).Debug("START Register")
// 	defer logger.Log(c).Debug("END Register")

// 	//hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	var count int64
// 	err := g.store.WithContext(c).Select("count * ").Table("kisan.users").Where("email = ?", email).Count(&count).Error
// 	if err != nil {
// 		logger.Log(c).Error("Error checking existing user", zap.Error(err))
// 		return nil
// 	}
// 	if count > 0 {
// 		return errors.New("user already exists with this email")
// 	}

// 	err = g.store.WithContext(c).Exec(`
// 		INSERT INTO kisan.users (name, phone, role, language, soil_type, district, lat, lng, created_at, password, email)
// 		VALUES (?, ?, ?, ?, ?, ?, 0, 0, now(), ?, ?)`,
// 		"jitendra", phone, "farmer", "hindi", "moisture", "dewas", password, email,
// 	).Error

// 	if err != nil {
// 		logger.Log(c).Error("Error inserting user", zap.Error(err))
// 		return err
// 	}
// 	return nil
// }

// func (g *registerStore) Logout(c context.Context, logoutFlag string) error {
// 	logger.Log(c).Debug("START Logout")
// 	defer logger.Log(c).Debug("END Logout")
// 	// You can blacklist JWT here if using Redis, else client just deletes token.
// 	return nil
// }

// // db/register_store.go
// func (g *registerStore) GetUserDetails(c context.Context, email string) (*models.UserDetails, error) {
// 	// strip "Bearer " prefix
// 	// token = strings.TrimPrefix(token, "Bearer ")

// 	// parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
// 	// 	return jwtSecret, nil
// 	// })
// 	// if err != nil || !parsed.Valid {
// 	// 	return nil, errors.New("invalid token")
// 	// }

// 	// claims := parsed.Claims.(jwt.MapClaims)
// 	// email := claims["email"].(string)

// 	var user models.UserDetails
// 	err := g.store.WithContext(c).Table("kisan.users").Where("email = ?", email).First(&user).Error
// 	if err != nil {
// 		logger.Log(c).Error("Error fetching user details", zap.Error(err))
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func (g *registerStore) UserupdateDetails(ctx context.Context, Email, Bio string) (*models.UserDetails, error) {
// 	logger.Log(ctx).Debug("START UserupdateDetails")
// 	defer logger.Log(ctx).Debug("END UserupdateDetails")

// 	err := g.store.WithContext(ctx).Exec(`
// 		UPDATE kisan.users
// 		SET
// 			bio = COALESCE(NULLIF(?, ''), bio),
// 		WHERE email = ?`,
// 		Bio,
// 		Email,
// 	).Error
// 	if err != nil {
// 		logger.Log(ctx).Error("Error updating user details", zap.Error(err))
// 		return nil, err
// 	}

// 	var updatedUser models.UserDetails
// 	err = g.store.WithContext(ctx).Table("kisan.users").Where("email = ?", Email).First(&updatedUser).Error
// 	if err != nil {
// 		logger.Log(ctx).Error("Error fetching updated user details", zap.Error(err))
// 		return nil, err
// 	}

// 	return &updatedUser, nil
// }

// func (g *registerStore) UpdateAvatar(ctx context.Context, email string, avatarURL string) error {
// 	logger.Log(ctx).Debug("START UpdateAvatar")
// 	defer logger.Log(ctx).Debug("END UpdateAvatar")
// 	err := g.store.WithContext(ctx).Exec(`
// 		UPDATE kisan.users
// 		SET
// 			avatar = ?
// 		WHERE email = ?`,
// 		avatarURL,
// 		email,
// 	).Error
// 	if err != nil {
// 		logger.Log(ctx).Error("Error updating avatar", zap.Error(err))
// 		return err
// 	}
// 	return nil
// }

// // func (g *registerStore) RefreshToken(c context.Context, request *models.DtlsRequest) ([]*models.DtlsResponse, error) {
// // 	// Implementation for refreshing token goes here
// // 	return nil, nil
// // }
