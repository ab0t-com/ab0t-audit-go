// ANNOTATION BLOCK
// File: svc_settings.go — the `Settings` designation of the Audit Service.
// Why it exists: typed methods for every Settings endpoint, reached via
// c.Settings.<Method>. Each is context-first, takes typed path/query/body
// params and returns a typed response (or *APIError). Query parameters for
// a method live in its <Method>Params struct in this file.
package auditclient

import (
	"context"
	"encoding/json"
	"net/url"
)

// SettingsGetClassificationSettingsParams holds the query parameters for Settings.GetClassificationSettings.
type SettingsGetClassificationSettingsParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID
}

func (p *SettingsGetClassificationSettingsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// SettingsGetNotificationPrefsParams holds the query parameters for Settings.GetNotificationPrefs.
type SettingsGetNotificationPrefsParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID
}

func (p *SettingsGetNotificationPrefsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// SettingsGetRedactionSettingsParams holds the query parameters for Settings.GetRedactionSettings.
type SettingsGetRedactionSettingsParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID
}

func (p *SettingsGetRedactionSettingsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// SettingsPutClassificationSettingsParams holds the query parameters for Settings.PutClassificationSettings.
type SettingsPutClassificationSettingsParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID
}

func (p *SettingsPutClassificationSettingsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// SettingsPutNotificationPrefsParams holds the query parameters for Settings.PutNotificationPrefs.
type SettingsPutNotificationPrefsParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID
}

func (p *SettingsPutNotificationPrefsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// SettingsPutRedactionSettingsParams holds the query parameters for Settings.PutRedactionSettings.
type SettingsPutRedactionSettingsParams struct {
	OrgID *string `json:"org_id,omitempty"` // Organization ID
}

func (p *SettingsPutRedactionSettingsParams) values() url.Values {
	v := url.Values{}
	if p == nil {
		return v
	}
	if p.OrgID != nil {
		v.Set("org_id", *p.OrgID)
	}
	return v
}

// GetClassificationSettings calls GET /settings/classification.
// Get Classification Settings
//
// Return the org's data-classification rule set. **When to use**:
// classification settings editor page load. **Auth**:
// AuditComplianceReader. **Behavior**: returns the org's saved override if
// one exists, otherwise the hardcoded platform defaults (with
// “is_default=true“).
func (s *SettingsClient) GetClassificationSettings(ctx context.Context, params *SettingsGetClassificationSettingsParams) (*ClassificationSettings, error) {
	b, err := s.c.do(ctx, "GET", "/settings/classification", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out ClassificationSettings
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/settings/classification", Err: err}
	}
	return &out, nil
}

// GetNotificationPrefs calls GET /settings/notifications.
// Get Notification Prefs
//
// Return the caller's notification-feed category preferences. **When to
// use**: notification preferences page load
// (/ui/app/settings/notifications). **Auth**: AuditUIUser (audit.ui.read) —
// personal setting, same tier as the feed. **Behavior**: returns this
// user's saved row for the org if one exists, otherwise the defaults (every
// category enabled, “is_default=true“).
func (s *SettingsClient) GetNotificationPrefs(ctx context.Context, params *SettingsGetNotificationPrefsParams) (*NotificationPrefsSettings, error) {
	b, err := s.c.do(ctx, "GET", "/settings/notifications", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out NotificationPrefsSettings
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/settings/notifications", Err: err}
	}
	return &out, nil
}

// GetRedactionSettings calls GET /settings/redaction.
// Get Redaction Settings
//
// Return the org's effective redaction key list. **When to use**: redaction
// settings editor page load. **Auth**: AuditComplianceReader. **Behavior**:
// returns the org's saved override if one exists, otherwise the hardcoded
// platform defaults (with “is_default=true“).
func (s *SettingsClient) GetRedactionSettings(ctx context.Context, params *SettingsGetRedactionSettingsParams) (*RedactionSettings, error) {
	b, err := s.c.do(ctx, "GET", "/settings/redaction", params.values(), nil)
	if err != nil {
		return nil, err
	}
	var out RedactionSettings
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/settings/redaction", Err: err}
	}
	return &out, nil
}

// PutClassificationSettings calls PUT /settings/classification.
// Put Classification Settings
//
// Persist the org's data-classification rule override. **When to use**:
// classification settings editor "Save changes" button. **Auth**:
// AuditModelAdmin (audit.system.admin, not-suspended). **Behavior**:
// upserts the single per-org row in “settings_classification“. Rules are
// normalized (pattern lower-cased/trimmed, classification validated,
// de-duped by pattern). Returns the persisted rule set.
func (s *SettingsClient) PutClassificationSettings(ctx context.Context, body ClassificationSettings, params *SettingsPutClassificationSettingsParams) (*ClassificationSettings, error) {
	b, err := s.c.do(ctx, "PUT", "/settings/classification", params.values(), body)
	if err != nil {
		return nil, err
	}
	var out ClassificationSettings
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/settings/classification", Err: err}
	}
	return &out, nil
}

// PutNotificationPrefs calls PUT /settings/notifications.
// Put Notification Prefs
//
// Persist the caller's notification-feed category preferences. **When to
// use**: notification preferences page "Save changes" button. **Auth**:
// AuditUIUser (audit.ui.read) — the row is keyed to the caller (org_id +
// user_id), so a user can only ever write their own preferences.
// **Behavior**: upserts the per-user row in “settings_notification_prefs“
// (ReplacingMergeTree(updated_at)). Unknown categories are dropped; missing
// ones default to enabled. Returns the persisted preference set.
func (s *SettingsClient) PutNotificationPrefs(ctx context.Context, body NotificationPrefsSettings, params *SettingsPutNotificationPrefsParams) (*NotificationPrefsSettings, error) {
	b, err := s.c.do(ctx, "PUT", "/settings/notifications", params.values(), body)
	if err != nil {
		return nil, err
	}
	var out NotificationPrefsSettings
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/settings/notifications", Err: err}
	}
	return &out, nil
}

// PutRedactionSettings calls PUT /settings/redaction.
// Put Redaction Settings
//
// Persist the org's redaction key override. **When to use**: redaction
// settings editor "Save changes" button. **Auth**: AuditModelAdmin
// (audit.system.admin, not-suspended). **Behavior**: upserts the single
// per-org row in “settings_redaction“. Keys are normalized (lower-cased,
// de-duped). Returns the persisted list.
func (s *SettingsClient) PutRedactionSettings(ctx context.Context, body RedactionSettings, params *SettingsPutRedactionSettingsParams) (*RedactionSettings, error) {
	b, err := s.c.do(ctx, "PUT", "/settings/redaction", params.values(), body)
	if err != nil {
		return nil, err
	}
	var out RedactionSettings
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, &DecodeError{Endpoint: "/settings/redaction", Err: err}
	}
	return &out, nil
}
