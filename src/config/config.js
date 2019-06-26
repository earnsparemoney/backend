module.exports = {
  port: 4000,
  db: {
    user: process.env.DB_USER || 'root',
    password: process.env.DB_PASSWD || 'sactestdatabase',
    database: process.env.DB_NAME || 'EarnMoney',
    options: {
      dialect: process.env.DIALECT || 'mysql',
      host: process.env.HOST || '222.200.190.31',
      port: process.env.PORT || 33336,
      define: {
        charset: 'utf8mb4',
        collate: 'utf8mb4_unicode_ci',
        supportBigNumbers: true,
        bigNumberStrings: true
      }
    }
  },
  authServiceToken: {
    secretKey: process.env.SECRET || 'secret'
  }
}

