module.exports = {
  port: 4000,
  db: {
    user: process.env.USER || 'EarnMoney',
    password: process.env.PASSWORD || 'EarnMoney',
    database: process.env.DATABASE || 'EarnMoney',
    options: {
      dialect: process.env.DIALECT || 'mysql',
      host: process.env.HOST || 'mysql',
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

