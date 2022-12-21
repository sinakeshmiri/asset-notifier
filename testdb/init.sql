USE asset_notif ;
CREATE TABLE dns_records (
  provider VARCHAR(255) NOT NULL ,
  type VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  value VARCHAR(255) NOT NULL
)
