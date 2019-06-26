const Sequelize = require('sequelize')
const config = require('../config/config')

const sequelize = new Sequelize(
  config.db.database,
  config.db.user,
  config.db.password,
  config.db.options
)

console.log('Connecting to database', config.db.options.host, '...')

var User = sequelize.import('./User.js')
var Questionnaire = sequelize.import('./Questionnaire.js')
var Participation = sequelize.import('./Participation.js')
var Task = sequelize.import('./Task.js')

Questionnaire.belongsTo(User, { as: 'publisher' })
User.belongsToMany(Questionnaire, { through: Participation })
Task.belongsTo(User, { as: 'publisher' })
User.hasMany(Task, { as: 'tasks' })

module.exports = {
  User: User,
  Questionnaire: Questionnaire,
  Participation: Participation,
  Task: Task,
  sequelize: sequelize,
  Sequelize: Sequelize
}
