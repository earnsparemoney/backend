const { User, Task } = require('../models')
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
      let option = {
        order: [['deadline', 'DESC']],
        where: {
          status: 0
        },
        include: [{ model: User, as: 'publisher', attributes: ['id', 'username', 'email', 'phone', 'img'] }]
      }
      if (req.query.page) {
        let page = parseInt(req.query.page)
        let pageSize = req.query.pageSize ? parseInt(req.query.pageSize) : 5
        option.limit = pageSize
        option.offset = (page - 1) * pageSize
        console.log(option)
      }
      var tasks = await Task.findAll(option)
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
      let option = {
        where: {
          publisherId: result.id
        },
        include: [{ model: User, as: 'publisher', attributes: ['id', 'username', 'email', 'phone', 'img'] }]
      }
      if (req.query.page) {
        let page = parseInt(req.query.page)
        let pageSize = req.query.pageSize ? parseInt(req.query.pageSize) : 5
        option.limit = pageSize
        option.offset = (page - 1) * pageSize
      }
      var tasks = await Task.findAll(option)
      res.send({ tasks: tasks })
    } catch (err) {
      console.log(err)
      res.status(400).send({
        error: 'Some wrong occured when getting data!'
      })
    }
  },
  async getAllRunningTasksParticipatesIn (req, res) {
    try {
      let option = {
        where: {
          UserId: req.params.id,
          status: 1
        },
        include: [{ model: User, as: 'publisher', attributes: ['id', 'username', 'email', 'phone', 'img'] }]
      }
      if (req.query.page) {
        let page = parseInt(req.query.page)
        let pageSize = req.query.pageSize ? parseInt(req.query.pageSize) : 5
        option.limit = pageSize
        option.offset = (page - 1) * pageSize
      }
      var participants = await Task.findAll(option)
      res.send({
        tasks: participants
      })
    } catch (err) {
      console.log(err)
      res.status(400).send({
        error: 'Some wrong occured when getting data!'
      })
    }
  },
  async getAllFinishedTasksParticipatesIn (req, res) {
    try {
      let option = {
        where: {
          UserId: req.params.id,
          status: 2
        },
        include: [{ model: User, as: 'publisher', attributes: ['id', 'username', 'email', 'phone', 'img'] }]
      }
      if (req.query.page) {
        let page = parseInt(req.query.page)
        let pageSize = req.query.pageSize ? parseInt(req.query.pageSize) : 5
        option.limit = pageSize
        option.offset = (page - 1) * pageSize
      }
      var participants = await Task.findAll(option)

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
          status: 0,
          id: req.params.id
        }
      })

      if (!task) {
        return res.status(403).send({
          error: 'Please confirm the status of the task!!'
        })
      }

      if (task.publisherId === result.id) {
        return res.status(403).send({
          error: 'You are the publisher!!'
        })
      }

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
      
      var user = await User.findOne({
        where: {
          id: result.id
        }
      })

      await user.update({
        balance: user.balance + task.adward
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
  async exitTask (req, res) {
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
          status: 1,
          id: req.params.id
        }
      })

      if (!task) {
        return res.status(403).send({
          error: 'Please confirm the status of the task!!'
        })
      }

      if (task.UserId === result.id) {
        return res.status(403).send({
          error: 'You are the participator!!'
        })
      }

      await task.update({
        status: 0,
        UserId: null
      })
      res.send({
        info: 'Exit successfully!'
      })
    } catch (err) {
      console.log(err.message)
      res.status(400).send({
        error: 'Some wrong occured when participate in event!!'
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
