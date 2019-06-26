module.exports = (sequelize, DataTypes) => {
  const Task = sequelize.define('Task', {
    name: { type: DataTypes.STRING, unique: true, allowNull: false },
    description: { type: DataTypes.STRING, allowNull: false },
    content: { type: DataTypes.STRING, allowNull: false },
    deadline: { type: DataTypes.DATE, allowNull: false },
    adward: { type: DataTypes.DOUBLE, allowNull: false },
    status: { type: DataTypes.INTEGER, defaultValue: 0 }
  })

  return Task
}
