const { User, Task, TaskUser } = require('../models')
const jwt = require('jsonwebtoken')
const config = require('../config/config')

module.exports = {
  async addTask (req, res) {
    try {
      const token = req.header('Authorization')
      if (!token) {
        return res.status(400).send({
          error: 'token should be given!'
        })
      }
      var result = null
      try {
        result = jwt.verify(token, config.authServiceToken.secretKey)
        if (!result) {
          return res.status(400).send({
            error: 'The token is not valid! Please sign in and try again!'
          })
        }
      } catch (err) {
        return res.status(400).send({
          error: 'Token expired, please login again!'
        })
      }
      var task = await Task.create({
        name: req.body.name,
        description: req.body.description,
        content: req.body.content,
        deadline: req.body.deadline,
        adward: req.body.adward,
        publisherId: result.id
      })
      const taskJSON = task.toJSON()

      res.send({
        task: taskJSON
      })
    } catch (err) {
      console.log(err)
      res.status(400).send({
        error: err.errors[0].message
      })
    }
  },
  async deleteTask (req, res) {
    const token = req.header('Authorization')
    if (!token) {
      return res.status(400).send({
        error: 'token should be given!'
      })
    }
    var result = null
    try {
      result = jwt.verify(token, config.authServiceToken.secretKey)
      if (!result) {
        return res.status(400).send({
          error: 'The token is not valid! Please sign in and try again!'
        })
      }
    } catch (err) {
      return res.status(400).send({
        error: 'Token expired, please login again!'
      })
    }
    var task = await Task.findOne({
      where: {
        id: req.params.id,
        publisherId: result.id
      }
    })
    await task.destroy()
    res.send({
      info: 'Delete task successfully!'
    })
  },
  async getAllTasks (req, res) {
    try {
      var tasks = await Task.findAll({
        order: [['deadline', 'DESC']],
        where: {
          status: 0
        },
        include: [{ model: User, as: 'publisher', attributes: ['id', 'username', 'email', 'phone', 'img'] }]
      })
      res.send({ tasks: tasks })
    } catch (err) {
      console.log(err.message)
      res.status(400).send({
        error: 'Some wrong occured when getting data!'
      })
    }
  },
  async getAllPublishedTasks (req, res) {
    try {
      var tasks = await Task.findAll({
        where: {
          publisherId: req.params.id
        },
        include: [{ model: User, as: 'publisher', attributes: ['id', 'username', 'email', 'phone', 'img'] }]
      })
      res.send({ tasks: tasks })
    } catch (err) {
      res.status(400).send({
        error: 'Some wrong occured when getting data!'
      })
    }
  },
  async getAllRunningTasksParticipatesIn (req, res) {
    try {
      var participants = await TaskUser.findAll({
        where: {
          UserId: req.params.id
        }
      }).map(async (participant) => {
        var res = await Task.findOne({
          where: {
            id: participant.EventId
          },
          include: [{ model: User, as: 'publisher', attributes: ['id', 'username', 'email', 'phone', 'img'] }]
        })
        res = res.toJSON()
        return res
      })

      participants = participants.filter(item => item.status == 1)

      res.send({
        tasks: participants
      })
    } catch (err) {
      res.status(400).send({
        error: 'Some wrong occured when getting data!'
      })
    }
  },
  async getAllFinishedTasksParticipatesIn (req, res) {
    try {
      var participants = await TaskUser.findAll({
        where: {
          UserId: req.params.id
        }
      }).map(async (participant) => {
        var res = await Task.findOne({
          where: {
            id: participant.EventId
          },
          include: [{ model: User, as: 'publisher', attributes: ['id', 'username', 'email', 'phone', 'img'] }]
        })
        res = res.toJSON()
        return res
      })

      participants = participants.filter(item => item.status == 2)
      
      res.send({
        tasks: participants
      })
    } catch (err) {
      res.status(400).send({
        error: 'Some wrong occured when getting data!'
      })
    }
  },
  async participateTask (req, res) {
    try {
      const token = req.header('Authorization')
      if (!token) {
        return res.status(400).send({
          error: 'token should be given!'
        })
      }
      var result = null
      try {
        result = jwt.verify(token, config.authServiceToken.secretKey)
        if (!result) {
          return res.status(400).send({
            error: 'The token is not valid! Please sign in and try again!'
          })
        }
      } catch (err) {
        return res.status(400).send({
          error: 'Token expired, please login again!'
        })
      }
      var task = await Task.findOne({
        where: {
          id: req.params.id
        }
      })

      await task.update({
        status: 1,
        UserId: result.id
      })
      res.send({
        info: 'Choose successfully!'
      })
    } catch (err) {
      console.log(err.message)
      res.status(400).send({
        error: 'Some wrong occured when participate in event!!'
      })
    }
  },
  async finishTask (req, res) {
    try {
      const token = req.header('Authorization')
      if (!token) {
        return res.status(400).send({
          error: 'token should be given!'
        })
      }
      var result = null
      try {
        result = jwt.verify(token, config.authServiceToken.secretKey)
        if (!result) {
          return res.status(400).send({
            error: 'The token is not valid! Please sign in and try again!'
          })
        }
      } catch (err) {
        return res.status(400).send({
          error: 'Token expired, please login again!'
        })
      }
      var task = await Task.findOne({
        where: {
          id: req.params.id,
          UserId: result.id
        }
      })

      await task.update({
        status: 2
      })

      res.send({
        info: 'Finish successfully!'
      })
    } catch (err) {
      res.status(400).send({
        error: 'Some wrong occured when finishing!!'
      })
    }
  },
  async getDetail (req, res) {
    try {
      var event = await Task.findOne({
        where: {
          id: req.params.id
        },
        include: [{ model: User, as: 'publisher', attributes: ['id', 'username', 'email', 'phone', 'img'] }]
      })

      res.send({
        event: event
      })
    } catch (err) {
      console.log(err.message)
      res.status(400).send({
        error: 'Some wrong occured when getting detail!!'
      })
    }
  }
}
