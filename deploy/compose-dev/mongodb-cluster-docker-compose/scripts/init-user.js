use admin;
db.createUser({
    user: "root",
    pwd: "cc",
    roles: [{ role: "dbOwner", db: "admin" }]
});
use cmdb;
db.createUser({
    user: "cc",
    pwd: "cc",
    roles: [{ role: "dbOwner", db: "cmdb" }]
});