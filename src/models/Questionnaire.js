module.exports = (sequelize, DataTypes) => {
  const Questionnaire = sequelize.define('Questionnaire', {
    title: { type: DataTypes.STRING, unique: true, allowNull: false },
    questions: { type: 'LONGTEXT', allowNull: false },
    startDate: { type: DataTypes.DATE, allowNull: false },
    endDate: { type: DataTypes.DATE, allowNull: false },
    adward: { type: DataTypes.DOUBLE, defaultValue: 0 },
    usernum: { type: DataTypes.INTEGER, defaultValue: 0 }
  })

  return Questionnaire
}
