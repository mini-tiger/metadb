config=rs.config()
config.members[0].host=$node
rs.reconfig(config, { force: true });
