package db

import (
	"context"
	"fmt"
	"kisaanSathi/pkg/logger"
	"kisaanSathi/pkg/services/session/models"
	"log"
	"time"

	"go.uber.org/zap"
)

func (g *dbSt) ValidateSession(c context.Context, sqlUsmUsrID string, sqlUsmSssnID string) (bool, error) {
	logger.Log(c).Debug("START")
	logger.Log(c).Debug("END")

	// Retrieve session from database
	session, err := g.GetSession(c, sqlUsmUsrID, sqlUsmSssnID)
	if err != nil {
		logger.Log(c).Error("Error retrieving session:", zap.Error(err))
		return false, err
	}

	// Validate session
	if !isValidSession(c, session) {
		logger.Log(c).Debug("Session is expired or terminated")
		return false, nil
	} else {
		log.Println("Session is valid")
		// Update last access time of session
		g.UpdateSessionLastAccess(c, session)
	}

	return true, nil
}

// Function to retrieve session from the database
func (g *dbSt) GetSession(c context.Context, userID string, sessionID string) (*models.UsmSssnMngr, error) {
	logger.Log(c).Debug("START")
	logger.Log(c).Debug("END")
	var session models.UsmSssnMngr
	query := g.oracle.WithContext(c).
		Table("usm_sssn_mngr usm"). // Specify the table name without alias
		Select(`usm."usm_usr_id", usm."usm_sssn_id", usm."usm_sssn_end_dt_flg", usm."usm_sssn_lst_accs_dt", usm."usm_sup_usr_typ", usm."usm_ip_id", usm."usm_sssn_termntd_flg", usm."usm_login_source"`).
		Where(`usm."usm_usr_id" = ? AND usm."usm_sssn_id" = ? AND usm."usm_sssn_end_dt_flg" = 'N'`, userID, sessionID).
		Order(`"usm_usr_id"`). // Order by column name
		Limit(1).
		Find(&session)

	if query.Error != nil {
		return nil, fmt.Errorf("error retrieving session from database: %w", query.Error)
	}
	return &session, nil
}

// Function to validate session
func isValidSession(c context.Context, session *models.UsmSssnMngr) bool {
	logger.Log(c).Debug("START")
	logger.Log(c).Debug("END")
	// Check if session is expired or terminated
	timeDiff := time.Since(session.UsmSssnLstAccsDt)
	sessionTimeout := getSessionTimeout(c, session)
	return timeDiff.Seconds() < sessionTimeout || session.UsmSssnTermntdFlg == "Y"
}

// Function to get session timeout based on session type
func getSessionTimeout(c context.Context, session *models.UsmSssnMngr) float64 {
	logger.Log(c).Debug("START")
	logger.Log(c).Debug("END")
	// Implement logic to retrieve session timeout based on session type
	if session.UsmSupUsrTyp == "A" || session.UsmSupUsrTyp == "Y" || session.UsmSupUsrTyp == "V" || session.UsmSupUsrTyp == "Z" {
		// Return appropriate session timeout for these session types
		return 3600 // Example session timeout in seconds
	} else {
		// Return default session timeout
		return 1800 // Default session timeout in seconds
	}
}

// Function to update last access time of session
func (g *dbSt) UpdateSessionLastAccess(c context.Context, session *models.UsmSssnMngr) error {
	logger.Log(c).Debug("START")
	logger.Log(c).Debug("END")
	query := g.oracle.WithContext(c).Table("usm_sssn_mngr usm").
		Where(`usm."usm_usr_id" = ? AND usm."usm_sssn_id" = ? AND usm."usm_sssn_end_dt_flg" = 'N'`, session.UsmUsrID, session.UsmSssnID).
		Update(`usm."usm_sssn_lst_accs_dt"`, time.Now())

	if query.Error != nil {
		return fmt.Errorf("error updating session last access time: %w", query.Error)
	}

	return nil
}
