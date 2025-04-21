package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/nara-ryoya/sqlboiler-mysql-view/db"
	"github.com/nara-ryoya/sqlboiler-mysql-view/repository"
)

func Users(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tenantID, err := strconv.ParseUint(r.Header.Get("X-Tenant-ID"), 10, 64)
	if err != nil || tenantID == 0 {
		http.Error(w, "invalid tenant", http.StatusBadRequest)
		return
	}
	conn, err := db.ConnForTenant(ctx, tenantID)
	if err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	// Start a transaction that implements boil.ContextExecutor
	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		http.Error(w, "transaction error", http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	users, err := repository.ListUsers(ctx, tx)
	if err != nil {
		http.Error(w, "query error", http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(users)
}
