const { Questionnaire, User, Participation, Sequelize } = require('../models')
const jwt = require('jsonwebtoken')
const config = require('../config/config')

module.exports = {
  async getAllQuestionnaires (req, res) {
    try {
      let option = {
        where: {
          endDate: {
            [Sequelize.Op.gte]: Date.now() - 24 * 60 * 60 * 1000
          }
        },
        include: [{ model: User, as: 'publisher', attributes: ['id', 'username', 'email', 'phone', 'img'] }]
      }
      if (req.query.page) {
        let page = parseInt(req.query.page)
        let pageSize = req.query.pageSize ? parseInt(req.query.pageSize) : 5
        option.limit = pageSize
        option.offset = (page - 1) * pageSize
      }
      var questionnaires = await Questionnaire.findAll(option).map(async (questionnaire) => {
        var count = await Participation.findAll({
          where: {
            QuestionnaireId: questionnaire.id
          }
        })
        questionnaire = questionnaire.toJSON()
        questionnaire.numOfFilled = (count.length === undefined) ? 0 : count.length
        return questionnaire
      })
      res.send({
        questionnaires: questionnaires
      })
    } catch (err) {
      res.status(400).send({
        error: 'Some wrong occoured when getting data!'
      })
    }
  },
  async getPublishedQuestionnaires (req, res) {
    try {
      const token = req.header('Authorization')
      if (!token) {
        return res.status(403).send({
          error: 'token should be given!'
        })
      }
      var result = null
      try {
        result = jwt.verify(token, config.authServiceToken.secretKey)
        if (!result) {
          return res.status(403).send({
            error: 'The token is not valid! Please sign in and try again!'
          })
        }
      } catch (err) {
        return res.status(403).send({
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
      var questionnaires = await Questionnaire.findAll(option).map(async (questionnaire) => {
        var count = await Participation.findAll({
          where: {
            QuestionnaireId: questionnaire.id
          }
        })
        questionnaire = questionnaire.toJSON()
        questionnaire.numOfFilled = (count.length === undefined) ? 0 : count.length
        return questionnaire
      })
      res.send({
        questionnaires: questionnaires
      })
    } catch (err) {
      console.log(err)
      res.status(400).send({
        error: 'Fail to get data!'
      })
    }
  },
  async addQuestionnaire (req, res) {
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
        console.log(err)
        return res.status(400).send({
          error: 'Token expired, please login again!'
        })
      }
      
      let user = await User.findOne({
        where: {
          id: result.id
        }
      })

      if (user.balance < req.body.adward * req.body.usernum) {
        return res.status(403).send({
          error: '你的余额不足, 快去赚闲币吧'
        })
      }

      await Questionnaire.create({
        title: req.body.title,
        description: req.body.description,
        questions: req.body.questions,
        endDate: req.body.endDate,
        publisherId: result.id,
        adward: req.body.adward,
        usernum: req.body.usernum
      })

      await user.update({
        balance: user.balance - req.body.adward * req.body.usernum
      })

      var questionnaire = await Questionnaire.findOne({
        where: {
          title: req.body.title
        },
        include: [{ model: User, as: 'publisher', attributes: ['id', 'username', 'email', 'phone', 'img'] }]
      })

      res.send({
        questionnaire: questionnaire.toJSON()
      })
    } catch (err) {
      res.status(400).send({
        error: err.errors !== undefined ? err.errors[0].message : 'Error input please check your input'
      })
    }
  },
  async deleteQuestionnaire (req, res) {
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

      let user = await User.findOne({
        where: {
          id: result.id
        }
      })

      var questionnaire = await Questionnaire.findOne({
        where: {
          id: req.params.id
        }
      })

      if (!questionnaire) {
        res.status(400).send({
          error: 'No questionnaire is found, please check your request!'
        })
      }

      var count = await Participation.findAll({
        where: {
          QuestionnaireId: req.params.id
        }
      })

      count = (count.length === undefined) ? 0 : count.length
      
      console.log(user)

      await user.update({
        balance: user.balance + (questionnaire.usernum - count) * questionnaire.adward
      })

      console.log(user)

      await questionnaire.destroy()
      res.send({
        info: '删除成功, 部分闲币已退回'
      })
    } catch (err) {
      res.status(400).send({
        error: 'Some error occured when deleting questionnaire!'
      })
    }
  },
  async getParticipateQuestionnaire (req, res) {
    try {
      //const token = req.header('Authorization')
      let option = {
        where: {
          UserId: id
        }
      }
      if (req.query.page) {
        let page = parseInt(req.query.page)
        let pageSize = req.query.pageSize ? parseInt(req.query.pageSize) : 5
        option.limit = pageSize
        option.offset = (page - 1) * pageSize
      }

      var result = []
      id = req.params.id
      var paticipations = await Participation.findAll(option)

      if(!paticipations) {
        res.status(400).send({
          error: "user not attend any questionnaire or not such user"
        })
      }

      for (pars in paticipations){
        let qid = participations[pars].id
        let quesOne = await Questionnaire.findOne({
          where: {
            id: qid
          }
        })
        result.push(quesOne)
      }

      res.send({
        questionnaires: JSON.stringify(result)
      })

      
    } catch (error) {
      res.status(400).send({
        error: error
      })
    }
    return null
  },
  async participateQuestionnaire (req, res) {
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

      let uid = result.id
      let qid = req.params.id
      let answer = req.body.answer

      let par = await Participation.findOne({
        where: {
          UserId: uid,
          QuestionnaireId: qid
        }
      })

      if(par){
        return res.status(400).send({
          error: "user already complete this questionnaire"
        })
      }

      await Participation.create({
        UserId: uid,
        QuestionnaireId: qid,
        answer: answer
      })

      var user = await User.findOne({
        where: {
          id: result.id
        }
      })

      var questionnaire = await Questionnaire.findOne({
        where: {
          id: qid
        }
      })

      await user.update({
        balance: user.balance + questionnaire.adward
      })

      res.send({
        info: "success"
      })

    } catch (error) {
      console.log(error)
      res.status(400).send({
        error: err.errors ? err.errors[0].message : 'Error when updating answer!'
      })
    }
    return null
  },
  async getDetil (req, res) {
    try {
      var questionnaire = await Questionnaire.findOne({
        where: {
          id: req.params.id
        },
        include: [{ model: User, as: 'publisher', attributes: ['id', 'username', 'email', 'phone', 'img'] }]
      })
      res.send({
        questionnaire: questionnaire
      })
    } catch (err) {
      res.status(400).send({
        error: 'Some error occured when deleting event!'
      })
    }
  },

  async getAllResults (req, res) {
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

    let uid = result.id
    let qid = req.params.qid

    let quesFind = await Questionnaire.findOne({
      where: {
        id: qid,
        publisherId: uid
      }
    })

    if(!quesFind){
      res.status(400).send({
        error: "you don't own this questionnaire"
      })
    }


    let all = await Participation.findAll({
      where: {
        QuestionnaireId: qid
      }
    }).map((result) => {
      return result.answer
    })

    let allResult = all.join('\n')
    
    res.send({
      result: allResult
    })

  }
}
