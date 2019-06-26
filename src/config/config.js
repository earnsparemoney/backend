module.exports = {
  port: 4000,
  db: {
    user: 'root',
    password: 'abc25834934F',
    database: 'EarnMyMoney',
    options: {
      dialect: process.env.DIALECT || 'mysql',
      host: 'localhost',
      port: 3306,
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
