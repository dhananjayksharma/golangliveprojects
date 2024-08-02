#!/bin/bash

# """ deleting database table columns """
TABLE_NAME="stadiums"
DB_NAME="ipldbs"
 
# Update the table
mysql -h $MS_HOST_ENV -P $MS_PORT_ENV -u $MS_USER_ENV -p$MS_PASS_ENV $DB_NAME <<EOF 
# ALTER TABLE $TABLE_NAME DROP stadium_cost;
EOF

echo "Table $TABLE_NAME updated successfully."
