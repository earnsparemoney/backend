module.exports = (sequelize, DataTypes) => {
  const Participation = sequelize.define('Participation', {
    answer: { type: DataTypes.STRING, allowNull: false }
  })

  return Participation
}
