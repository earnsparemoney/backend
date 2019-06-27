module.exports = {
  port: 4000,
  db: {
    user: process.env.USER || 'root',
    password: process.env.PASSWORD || 'abc25834934F',
    database: process.env.DATABASE || 'EarnMyMoney',
    options: {
      dialect: process.env.DIALECT || 'mysql',
      host: process.env.HOST || 'localhost',
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

