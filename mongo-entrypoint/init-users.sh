if [ "$MONGO_INITDB_ROOT_USERNAME" ] && [ "$MONGO_INITDB_ROOT_PASSWORD" ]; then
  "${mongo[@]}" "$MONGO_INITDB_DATABASE" <<-EOJS
  db.createUser({
     user: $(_js_escape "$MONGO_INITDB_ROOT_USERNAME"),
     pwd: $(_js_escape "$MONGO_INITDB_ROOT_PASSWORD"),
     roles: [ 
       { role: "root", db: "admin" },
       { role: "readWrite", db: $(_js_escape "$MONGO_INITDB_DATABASE") } 
     ]
     })
EOJS
fi

if [ "$MONGO_INITDB_APP_USERNAME" ] && [ "$MONGO_INITDB_APP_PASSWORD" ]; then
  "${mongo[@]}" "$MONGO_INITDB_DATABASE" <<-EOJS
  db.createUser({
     user: $(_js_escape "$MONGO_INITDB_APP_USERNAME"),
     pwd: $(_js_escape "$MONGO_INITDB_APP_PASSWORD"),
     roles: [ 
       { role: "userAdminAnyDatabase", db: "admin" },
       { role: "readWrite", db: $(_js_escape "$MONGO_INITDB_DATABASE") } 
     ]
     })
EOJS
fi

echo ======================================================
echo created $MONGO_INITDB_ROOT_USERNAME in database $MONGO_INITDB_DATABASE
echo created $MONGO_INITDB_APP_USERNAME in database $MONGO_INITDB_DATABASE
echo ======================================================