#!/bin/bash
/usr/bin/psql -q -U postgres -d invitation_development -c "DELETE FROM users WHERE created_at < NOW() - INTERVAL '1 days' AND verified = false;"
