module.exports = {
  port: 4000,
  db: {
    user: 'root',
    password: 'sactestdatabase',
    database: 'EarnMyMoney',
    options: {
      dialect: process.env.DIALECT || 'mysql',
      host: '222.200.190.31',
      port: 33336,
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
