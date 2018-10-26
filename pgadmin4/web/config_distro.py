# Path to the online help.
HELP_PATH = '/usr/share/doc/pgadmin4-doc/html'

# Server-side session storage path
if SERVER_MODE:
    SESSION_DB_PATH = '/var/cache/pgadmin/sessions'

# Check for new versions of the application?
UPGRADE_CHECK_ENABLED = False

# Default locations for binary utilities (pg_dump, pg_restore etc)
DEFAULT_BINARY_PATHS = {
    "pg":   "/usr/bin",
    "ppas": "",
    "gpdb": ""
}
